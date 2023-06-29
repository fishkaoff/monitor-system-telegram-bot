package bot

import (
	"bytes"
	"encoding/json"
	"fmt"

	"strings"

	"github.com/fishkaoff/telegram-client/internal/models"
	"github.com/fishkaoff/telegram-client/pkg/messages"
)

func (b *Bot) addUrl(chatID int64, url string) string {

	// validate url
	url = strings.TrimSpace(url)
	if !b.mw.CheckUrl(url) {
		return messages.NOTURL
	}

	// check for similar url
	userUrls := b.storageMicro.Get(chatID)

	if b.mw.CheckMatches(userUrls, url) {
		return messages.URLALREADYADDED
	}

	// add url to db
	resp := b.storageMicro.Save(chatID, url)

	return resp
}

func (b *Bot) deleteUrl(chatID int64, url string) string {
	url = strings.TrimSpace(url)

	// validate url
	if !b.mw.CheckUrl(url) {
		return messages.NOTURL
	}

	// query to database for delete url
	resp := b.storageMicro.Delete(chatID, url)

	return resp
}

func (b *Bot) checkUrls(chatID int64) string {

	// get urls
	userUrls := b.storageMicro.Get(chatID)

	// put urls to struct
	structedUrls := models.RequestRows{Urls: userUrls}

	// encode struct to json
	request, err := json.Marshal(structedUrls)
	if err != nil {
		b.sugar.LogError(err.Error())
		return messages.SERVERERROR
	}

	response := b.checkUrlsMicro.SendUrls(request)

	parsedResponse, err := b.parseResponse(response)
	if err != nil {
		b.sugar.LogError(err.Error())
		return messages.SERVERERROR
	}

	return b.renderResponse(parsedResponse)
}

func (b *Bot) parseResponse(response []byte) ([]models.ResponseRow, error) {
	var parsedJSON models.ResponseRows

	r := bytes.NewReader(response)
	decoder := json.NewDecoder(r)

	err := decoder.Decode(&parsedJSON)
	if err != nil {
		b.sugar.LogMessage(string(response))
		return []models.ResponseRow{}, err
	}

	return parsedJSON.Urls, nil
}

func (b *Bot) renderResponse(unRenderedResponse []models.ResponseRow) string {
	var renderedResponse string
	var status string

	for _, checkedUrl := range unRenderedResponse {
		if checkedUrl.StatusCode != 200 {
			status = messages.URLUNAWAILABLE
		} else {
			status = messages.URLAWAILABLE
		}

		renderedResponse += fmt.Sprintf("🗄️%s :\n ➖status: %s\n ➖code: %v\n ➖response time: %v ms\n", checkedUrl.Url, status, checkedUrl.StatusCode, checkedUrl.ResponseTime)
	}

	return renderedResponse
}
