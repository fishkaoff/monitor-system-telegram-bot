package bot

import (
	"fmt"

	"github.com/fishkaoff/telegram-client/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update)
			continue
		}
		b.handleMessage(update)
	}
}

func (b *Bot) handleCommand(update tgbotapi.Update) {
	var response string

	switch update.Message.Command() {
	case messages.GETMETRICCOMMAND:
		response = b.checkUrls(update.Message.Chat.ID)
		b.sendMessage(response, update)
		break

	case messages.ADDSITECOMMAND:
		response = messages.SENDDATA
		b.userStatus[update.Message.Chat.ID] = 1
		b.sendMessage(response, update)

	case messages.DELETESITECOMMAND:
		response = messages.SENDDATA
		b.userStatus[update.Message.Chat.ID] = 2
		b.sendMessage(response, update)
		break

	case messages.HELPCOMMAND:
		b.sendMessage(messages.HELP, update)

	case messages.REGISTERCOMMAND:
		b.sendMessage(b.SaveUser(update.Message.Chat.ID), update)
	default:
		b.sendMessage(messages.HELP, update)
		b.userStatus[update.Message.Chat.ID] = 0
	}
}

func (b *Bot) handleMessage(update tgbotapi.Update) {
	var response string
	// if userStatus == 1 we need to add url
	if b.userStatus[update.Message.Chat.ID] == 1 {

		// request to usecase for add url
		response = b.addUrl(update.Message.Chat.ID, update.Message.Text)
		b.userStatus[update.Message.Chat.ID] = 0
		b.sendMessage(response, update)

		return
	}

	// if userStatus == 1 we need to delete url
	if b.userStatus[update.Message.Chat.ID] == 2 {
		response = b.deleteUrl(update.Message.Chat.ID, update.Message.Text)
		b.userStatus[update.Message.Chat.ID] = 0
		b.sendMessage(response, update)
		return
	}

	b.sendMessage(messages.UNKMOWNCOMMAND, update)
}

func (b *Bot) sendMessage(text string, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := b.bot.Send(msg); err != nil {
		errorMessage := fmt.Sprintf("Cannot send message to user: %v", update.Message.Chat.ID)
		b.sugar.LogError(errorMessage)
	}
}
