package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Tgbot() {

	//подключаемся к боту
	bot, err := tgbotapi.NewBotAPI("5633405489:AAEuNchpVfMot9oK0x2aG4rMfJjr2cv9p_8")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Auth on account %v", &bot.Self.UserName)

	//иниц канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates := bot.GetUpdatesChan(ucfg)

	for update := range updates {

		//user который написал боту
		UserName := update.Message.From.FirstName
		Loc := update.Message.Location

		if Loc != nil {
			lat := Loc.Latitude
			lon := Loc.Longitude
			data := GetWather(lat, lon)
			posiotion := fmt.Sprintf("%v, %v, %v", data.Geo_object.Province.Name, data.Geo_object.Locality.Name, data.Geo_object.District.Name)
			coord := fmt.Sprintf("Ваши координаты %v :: %v\n %v", lat, lon, posiotion)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, coord)

			bot.Send(msg)
			fmt.Printf("%v", data)
			temp := data.Fact["temp"]
			feels_like := data.Fact["feels_like"]
			wind_speed := data.Fact["wind_speed"]
			humidity := data.Fact["humidity"]

			fmt.Printf("%v , %v ", temp, feels_like)
			str := fmt.Sprintf("Сейчас %v °C, ощущается как %v °C\n Ветер %v м/с, влажность воздуха %v% ", temp, feels_like, wind_speed, humidity)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, str)
			bot.Send(msg)
		}

		// str := fmt.Sprintf("lat: %v lon: %v", lat, lon)

		// ID чата/диалога.
		// Может быть идентификатором как чата с пользователем
		// (тогда он равен UserID) так и публичного чата/канала
		ChatID := update.Message.Chat.ID

		//Text сообщения

		//Ответим userу
		if Loc == nil {
			reply := "Привет, " + UserName
			//подготовим сообщение для отправки
			msg := tgbotapi.NewMessage(ChatID, reply)
			msg2 := tgbotapi.NewMessage(ChatID, "Отправь свою геопозицию, чтобы узнать погоду")
			//отправим сообщение
			bot.Send(msg)
			bot.Send(msg2)
		}
	}
}

func connDB() (*sql.DB, error) {

	connDB, err := sql.Open("mysql", "user:user@tcp(127.0.0.1)/mybase")

	return connDB, err
}
