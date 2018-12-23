package main

import (
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var nossaMota = []string{
	"Nossa Mota",
	"Mota é muito invejoso",
	"Vira latismo over 9000",
}

var timCook = []string{
	"Apple ta uma merda",
	"Apple ta um lixo",
}

var fakeNews = []string{
	"F E I Q U E",
	"fake nils",
	"fake news",
}

var mentionFerrao = []string{
	"oi",
	"fala",
}

var newAppleDevices = []string{
	"Apple ta uma merda",
	"Vou pegar no muambeiro",
	"Ta valendo a pena",
	"Comprei o novo e dei o antigo pro paulo",
}

var kwsNewAppleDevices = []string{
	"iphone",
	"iPhone",
	"MacBook",
	"macbook",
}

var replyImage = []string{
	"kkkkkkkkkkkk",
	"eHAOIUHEAOIHEAIUOHEAUEA",
	"L O L",
	"perdi mto forte",
	"L O L",
	"\U0001F602\U0001F602\U0001F602\U0001F602",
	"Caralho",
	"A O V I V O",
}

var amir = []string{
	"Ta muito caro la",
	"Joao saiu de la, nao vale mais a pena",
}

var ocupado = []string{
	"To ocupado",
	"To em Sp",
	"tenho pf de calc iv amanha",
	"to com enxaqueca",
	"bora no donaldo pra economizar",
	"vou avisar para a Natalia",
}

func ferraoReply(text *string, keywords []string, possibleAnswers []string, probability int) bool {

	match := false
	rand.Seed(time.Now().UTC().UnixNano())
	r := rand.Intn(100 - 0)

	for _, keyword := range keywords{
		if (strings.Contains(*text, keyword) && (r < probability))	{
			*text = possibleAnswers[rand.Intn(len(possibleAnswers))]
			match = true
			break
		}
	}

	return match
}

func main() {
	bot, err := tgbotapi.NewBotAPI("602146626:AAH7fAjSsxSqV3NcaEwbaYeHWVi3dHpRJow")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)


	// To avoid dealing with old messages
	time.Sleep(time.Millisecond * 5000)
	updates.Clear()

	for update := range updates {

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		rand.Seed(time.Now().UTC().UnixNano())
		r := rand.Intn(100 - 0)

		// Logging messages that arrived
		log.Printf("[%s] %s \n", update.Message.From.UserName, update.Message.Text )


		var message string = ""

		switch {

		//case (strings.Contains(update.Message.Text, "macbook") || strings.Contains(update.Message.Text, "iphone")|| strings.Contains(update.Message.Text, "apple watch")) && (r < 80) && (update.Message.From.UserName != "viniciusferrao"):
			//	message = newAppleDevices[rand.Intn(len(newAppleDevices))]


			case ferraoReply(&update.Message.Text, kwsNewAppleDevices, newAppleDevices, 100):
				message = update.Message.Text

			case (strings.Contains(update.Message.Text, "apple") && (r < 80) && (update.Message.From.UserName != "viniciusferrao")):
				message = timCook[rand.Intn(len(timCook))]

			case (strings.Contains(update.Message.Text, "https") || strings.Contains(update.Message.Text, "http")) && (r < 60) && (update.Message.From.UserName != "viniciusferrao"):
				message = fakeNews[rand.Intn(len(fakeNews))]

			case strings.Contains(update.Message.Text, "amir") && (r < 80) && (update.Message.From.UserName != "viniciusferrao"):
				message = amir[rand.Intn(len(amir))]

			case strings.Contains(update.Message.Text, "bora") && (r < 90) && (update.Message.From.UserName != "viniciusferrao"):
				message = ocupado[rand.Intn(len(ocupado))]

			case strings.Contains(update.Message.Text, "@viniciusferrao") && (r < 90) && (update.Message.From.UserName != "viniciusferrao"):
				message = mentionFerrao[rand.Intn(len(mentionFerrao))]

			case update.Message.Photo != nil && (r < 30) && (update.Message.From.UserName != "viniciusferrao"):
				message = replyImage[rand.Intn(len(replyImage))]

			case (update.Message.From.UserName == "maulmota") && (r < 5) && (update.Message.From.UserName != "viniciusferrao"):
				message = nossaMota[rand.Intn(len(nossaMota))]
		}

		if message == ""{
			continue
		}

		// send the message
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
		msg.ReplyToMessageID = update.Message.MessageID

		time.Sleep(time.Millisecond * 5000)
		bot.Send(msg)
	}
}
