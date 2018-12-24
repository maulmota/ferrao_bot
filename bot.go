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

var kwsTimCook = []string{
	"Apple",
	"apple",
	"tim cook",
	"Tim Cook",
}

var timCook = []string{
	"Apple ta uma merda",
	"Apple ta um lixo",
}

var kwsFakeNews = []string{
	"http",
	"https",
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

var kwsNewAppleDevices = []string{
	"iphone",
	"iPhone",
	"MacBook",
	"macbook",
}

var newAppleDevices = []string{
	"Apple ta uma merda",
	"Vou pegar no muambeiro",
	"Ta valendo a pena",
	"Comprei o novo e dei o antigo pro paulo",
}

var replyImage = []string{
	"kkkkkkkkkkkk",
	"eHAOIUHEAOIHEAIUOHEAUEA",
	"L O L",
	"perdi mto forte",
	"ehaiuoheaiuheiauhea",
	"\U0001F602\U0001F602\U0001F602\U0001F602",
	"Caralho",
	"A O V I V O",
}

var kwsAmir = []string{
	"amir",
	"Amir",
}

var amir = []string{
	"Ta muito caro la",
	"Joao saiu de la, nao vale mais a pena",
}

var kwsOcupado = []string{
	"outback",
	"Outback",
	"bora",
	"Bora",
}

var ocupado = []string{
	"To ocupado",
	"To em Sp",
	"tenho pf de calc iv amanha",
	"to com enxaqueca",
	"bora no donaldo pra economizar",
	"vou avisar para a Natalia",
}

var kwsCanada = []string{
	"Canada",
	"canada",
	"Canadá",
	"canada",
	"quebec",
	"Quebec",
}

var canada = []string{
	"Mota se achando o foda aí porque foi pro Canada",
	"tu ja está no caminho sem volta nessa bosta de canada",
	"lombo canadense é a unica coisa boa do canada",
	"Enjoy Canada without friends, @maulmota",
	"normie é ir pro canada cara, discurso de patricinha metida",
	"foi mal se no canada é sempre domingo",
}

var kwsAndroid = []string{
	"android",
	"Android",
	"samsung",
	"Samsung",
}

var android = []string{
	"outback",
	"nao vou nem entrar nessa discussao",
	"affe mano",
	"ah nao mano",
	"Google",
	"google",
}

var kwsChrome = []string{
	"Chrome",
	"chrome",
	"gmail",
	"Gmail",
}

var chrome = []string{
	"nao sei oq tu escuta ai no canada mas tu ta se achando para caralho vai tomar no cu",
	"tu ta cego com sei la q merda teu coworker falou para vc",
	"mota eh vira casaca",
	"q lixo cara",
}

var kwsCo7 = []string{
	"co7",
	"Co7",
	"Cossetti",
	"cossetti",
}

var co7 = []string{
	"Co7 é o pedreiro so faz serviço de carga",
	"Botei ele para arrumar os armários do César",
}

func ferraoReply(text *string, keywords []string, possibleAnswers []string, probability int) bool {

	match := false
	rand.Seed(time.Now().UTC().UnixNano())
	r := rand.Intn(100 - 0)

	for _, keyword := range keywords {
		if strings.Contains(*text, keyword) && (r < probability) {
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
		log.Printf("[%s] %s \n", update.Message.From.UserName, update.Message.Text)

		var message string = ""

		switch {
		case ferraoReply(&update.Message.Text, kwsChrome, chrome, 80):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsAndroid, android, 80):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsNewAppleDevices, newAppleDevices, 80):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsTimCook, timCook, 80):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsFakeNews, fakeNews, 80):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsAmir, amir, 80):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsOcupado, ocupado, 80):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsCanada, canada, 80):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsCo7, co7, 80):
			message = update.Message.Text
		case strings.Contains(update.Message.Text, "@viniciusferrao") && (r < 90) && (update.Message.From.UserName != "viniciusferrao"):
			message = mentionFerrao[rand.Intn(len(mentionFerrao))]
		case update.Message.Photo != nil && (r < 30) && (update.Message.From.UserName != "viniciusferrao"):
			message = replyImage[rand.Intn(len(replyImage))]
		case (update.Message.From.UserName == "maulmota") && (r < 5) && (update.Message.From.UserName != "viniciusferrao"):
			message = nossaMota[rand.Intn(len(nossaMota))]
		}

		if message == "" {
			continue
		}

		// send the message
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
		msg.ReplyToMessageID = update.Message.MessageID

		time.Sleep(time.Millisecond * 5000)
		bot.Send(msg)
	}
}
