package main

import (
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
)

func main() {
	token := ""

	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	timers := &Timers{}

	err = run(api, timers)
	if err != nil {
		panic(err)
	}
}

func run(api *tgbotapi.BotAPI, timers *Timers) error {
	update := tgbotapi.NewUpdate(0)

	update.Timeout = 60

	updatesChan, err := api.GetUpdatesChan(update)
	if err != nil {
		return err
	}

	for {
		for received := range updatesChan {
			switch received.Message.Command() {
			case "ping":
				go handlePing(api, received.Message)

			case "loop":
				go handleLoop(api, timers, received.Message)

			case "timer":

			case "list":

			case "stop":
				go handleStop(timers, received.Message)

			case "help":

			default:

			}
		}
	}

	return nil
}

func handlePing(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	fmt.Printf("%d: ping\n", message.From.ID)
	msg := tgbotapi.NewMessage(message.Chat.ID, "pong")
	bot.Send(msg)
}

func handleLoop(bot *tgbotapi.BotAPI, timers *Timers, message *tgbotapi.Message) {
	payload := strings.Split(message.Text, " ")

	if len(payload) == 1 {
		return
	}

	sleepFor, err := time.ParseDuration(payload[1])
	if err != nil {
		return
	}

	tickCallback := func(iteration int) {
		timePassed := sleepFor * time.Duration(iteration)

		notification := timePassed.String()

		if strings.Contains(notification, "m0s") {
			notification = strings.TrimSuffix(notification, "0s")
		}

		if strings.Contains(notification, "h0m") {
			notification = strings.TrimSuffix(notification, "0m")
		}

		msg := tgbotapi.NewMessage(message.Chat.ID, notification)

		bot.Send(msg)
	}

	timers.AddLoopedTimer(message.From.ID, sleepFor, tickCallback)
}

func handleStop(timers *Timers, message *tgbotapi.Message) {
	timers.DisableTimer(message.From.ID)
}
