# Telegram Bot for monitor system

## About
	Project сreated for education purposes.
	It divided for microservices:
		1. Telegram BOT
		2. Database microservice
		3. Monitor microservice
	In the future im planning to create web site which will allow to configure user servers and view statisctics
	Architecture of app:
	https://miro.com/app/board/uXjVM7W6hhI=/?share_link_id=600100255279
## Technologies
	TelegramBOT: Golang, gRPC, Rest
	Monitor Microservice: Golang, Rest
	Database Microservice: Golang, gRPC, Postgresql
## Architecture of microservices
	All layers of microservices implemented by interfaces.
## Start Up on local machine
	Clone all repos to local machine
### Database start
#### Edit settings in docker-compose.yaml and .env
	$ cd <database microservice>
	$ docker-compose up
	$ go run main.go
### Monitor microservice start
	$ go run main.go
### Telegram Bot start
#### Change settings in .env
	$ cd <Telegram Bot dir>
	$ go run cmd/main.go
