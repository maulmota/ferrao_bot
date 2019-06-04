package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"math/rand"
	"strings"
	"time"
)

type videonote struct {
	fileId   string
	duration int
}

var kwsTimMaia = []string{
	"tim maia",
	"que delicia",
	"que delícia",
}

var timMaiaAudio = []string{
	"CQADAQADVAADMdEYRhqWi07QgrMIAg",
	"CQADAQADIQADxYFoRo6kFYP0RxFaAg",
}

var felizNatal = []string{
	"Feliz Páscoa",
}

var kwsFeliznatal = []string{
	"feliz natal",
}

var nossaMota = []string{
	"eu tenho q concordar o mota",
	"Nossa Mota",
	"Mota é muito invejoso",
	"Vira latismo over 9000",
}

var voiceCanadaNossaMota = []string{
	"AwADAQADCwAD0lE5R4-K8R6LN3zPAg",
	"AwADAQADJgAD5b2JRS6uR4M1Es73Ag",
	"AwADAQADGQADLRu5RVkzYdXBrtkRAg",
}

var voiceHugo = []string{
	"AwADAQADIgADojd5RhYfwFxbaOYCAg",
}

var kwsTimCook = []string{
	"Apple",
	"tim cook",
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

var kwsFatoOuFake = []string{
	"é fake",
	"eh fake",
	"fato ou fake",
	"é verdade",
	"eh verdade",
	"fake news",
}

var fatoOuFake = []string{
	"Muito bom, mas provavelmente fake",
	"Só pode ser sacanagem",
	"Se fosse fake news eu saberia",
	"falaram q eh fake",
	"Claramente Fake",
	"Sim é fake.",
	"Mais fake news impossível",
	"pior q não é feique",
	"Feique",
	"NAO TEM MAIS FEIQUE DO BOLARO",
	"Fake the Faking Fakers.",
	"Na real esse fake news poderia ser real",
	"não é fake",
	"Não entendi absolutamente porra nenhuma. Não sei se é fake ou real.",
	"Não sei se é real ou não",
}
var mentionFerrao = []string{
	"oi",
	"fala",
}

var videoMentionFerrao = []videonote{
	videonote{
		fileId:   "DQADAQADOgADWuihRvpcibjr3LhHAg",
		duration: 6,
	},
}

var kwsLetsPlay = []string{
	"vamos jogar",
	"vai jogar",
	"jogar",
	"d3",
}

var LetsPlay = []string{
	"vou tomar banho",
	"Tão jogando?",
	"to ligando o DESQUETOPE",
	"daqui a pouco vou ter que ir jantar e ai miou",
	"to la",
	"ta baixando o patch",
	"vou entrar",
	"who mais gonna play",
}

var kwsNewAppleDevices = []string{
	"iPhone",
	"MacBook",
	"Apple Watch",
	"Mac Pro",
	"fagOS",
	"macos",
}

var newAppleDevices = []string{
	"Apple ta uma merda",
	"Vou pegar no muambeiro",
	"Ta valendo a pena",
	"Comprei o novo e dei o antigo pro paulo",
	"a apple deu rollback",
	"RALADOR DE QUEIJO",
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
	"Amir",
}

var kwsAmirPhoto = []string{
	"kafta",
	"shwarma",
	"homus",
	"joao",
}

var amirPhoto = []string{
	"AgADAQAD5akxG-iEGUUdzJm7rz5FxXdlDDAABPRBJK8_scwxwDEDAAEC", // Amir
	"AgADAQAD5qkxG-iEGUWKtUfvkuleKzx1DDAABO3chmZZ_T93FzYDAAEC", // Amir
}

var amir = []string{
	"Ta muito caro la",
	"Joao saiu de la, nao vale mais a pena",
}

var kwsOcupado = []string{
	"Bora",
	"Vamos lá",
	"vamos la",
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
	"Canadá",
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

var voiceLol = []string{
	"AwADAQADLAADnLEwRfUP3N9QawoSAg",
	"AwADAQADLgADnLEwRTSrLDQWhre-Ag",
	"AwADAQADNQADnLEwRYfiKgua6PnsAg",
	"AwADAQADQgADGBS5RYa21YUzsybTAg",
	"AwADAQADKwADRazwRNfpMzv6ESpHAg",
}

var kwsAndroid = []string{
	"Android",
	"Samsung",
}

var android = []string{
	"nao vou nem entrar nessa discussao",
	"affe mano",
	"ah nao mano",
}

var kwsChrome = []string{
	"Chrome",
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
	"cossetti",
}

var co7 = []string{
	"Co7 é o pedreiro so faz serviço de carga",
	"Botei ele para arrumar os armários do César",
}

var kwsSumido = []string{
	"sumido",
	"ferrao morreu",
	"rip ferrao",
	"rip ferrão",
	"sumiu",
}

var kwsTomarNoCu = []string{
	"me manda tomar no cu",
	"tomar no cu",
	"toma no cu",
	"me xinga",
}

var voiceTomarNoCu = []string{
	"AwADAQADUQADZW4gRd2fV8q-IinjAg",
	"AwADAQADgQADpILYRjKn9boi0189Ag",
	"AwADAQADYQADgMzIRRNI_JOwIk6aAg",
}

var macacoVideo = []videonote{
	videonote{
		fileId:   "DQADAQADbwADn0BIRsIBf-nusV6kAg",
		duration: 10,
	},
}

var kwsMacacoVideo = []string{
	"macaco",
	"gorila",
	"preto",
	"chipanze",
	"crioulo",
	"criolo",
	"harambe",
}

var randomPhotoId = []string{
	"AgADAQADwqcxG9JRMUfjPcPDRQrOa20aAzAABMQi7nMGvR8iR3oAAgI",
	"AgADAQADvacxG-s_IUf67VIkprcF5zck9y8ABGd47xxniCxx-6wAAgI",
	"AgADAQADt6cxGzo2cEc8iXrbS_cMhqUt9y8ABKkQmalKNUieGHAAAgI",
	"AgADAQAD7acxG1Y4MEbuOeeF8bgnQmjpCjAABIOXWXaWtaVLoZUBAAEC",
	"AgADAQAD46kxG-iEGUW9hIwYVMyy64FoAzAABHZhdnUaT5rcNTACAAEC",
	"AgADAQADBqgxGxqbGEXmWBfSIBkuZCeuCjAABC0ovq3C-_tKwoUCAAEC",
	"AgADAQAD5KkxG-iEGUWkSgd_XoEhexc2FDAABEKtMRz6sJWt-3wAAgI",
	"AgADAQADB6gxGxqbGEUsmN6uJ66MtSFtDDAABC7YkzIas8W3xTkDAAEC",
	"AgADAQADCKgxGxqbGEWxlnrf6STq3zC5CjAABBMLDhI5tXf6WnUCAAEC",
	"AgADAQADCagxGxqbGEXV9XVf52bYblr6CjAABLY_hXcF4kvp-f8BAAEC",
	"AgADAQADCqgxGxqbGEWzX82CRB0ekMN1DDAABP_oaz12olmzfjoDAAEC",
	"AgADAQADC6gxGxqbGEXC2OJ0rDP2PyLtCjAABMR38Dvit2EwBBcCAAEC", // matsuda
	"AgADAQADDKgxGxqbGEVTWYjfl_d9U7W3CjAABDcyKSHwgM06GHoCAAEC", // matsuda
	"AgADAQADDagxGxqbGEW21D3-wkXZey0CCzAABDp9sX3ZHeOQDhkCAAEC",
	"AgADAQADDqgxGxqbGEWjhiob0_va5Pt4DDAABFabgRs98CaNPjQDAAEC",
	"AgADAQAD56kxG-iEGUXWW9P6PH3NEo95DDAABCUNEhMEM_1YUDMDAAEC",
	"AgADAQADD6gxGxqbGEUb5ZY90we8dQnrCjAABIfw9PVPqV_EjRUCAAEC",
	"AgADAQADEKgxGxqbGEVSdXk4jooFZQ00FDAABKp7oUi7NmHWaXoAAgI",
	"AgADAQADEagxGxqbGEWW_S4E4-5O5bBlAzAABAnXSUSYyzXq_ysCAAEC",
	"AgADAQADEqgxGxqbGEXbdwLtNWaGTuJrDDAABOcwq9xqG3RtXjgDAAEC",
}

func ferraoReply(text *string, keywords []string, possibleAnswers []string, probability int) bool {

	match := false
	rand.Seed(time.Now().UTC().UnixNano())
	r := rand.Intn(100 - 0)

	for _, keyword := range keywords {
		if strings.Contains(strings.ToLower(*text), strings.ToLower(keyword)) && (r < probability) {
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
		audioId := ""
		voiceId := ""
		photoId := ""
		videoId := videonote{fileId: "", duration: 0}

		switch {
		case strings.Contains(update.Message.Text, "@viniciusferrao") && ferraoReply(&update.Message.Text, kwsFatoOuFake, fatoOuFake, 90):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsFeliznatal, felizNatal, 80):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsChrome, chrome, 80):
			message = update.Message.Text
		case strings.Contains(update.Message.Text, "@viniciusferrao") && ferraoReply(&update.Message.Text, kwsTomarNoCu, voiceTomarNoCu, 90) && (update.Message.From.ID == 390998014):
			voiceId = update.Message.Text
		case strings.Contains(update.Message.Text, "@viniciusferrao") && (update.Message.From.ID == 444973217):
			voiceId = "AwADAQADHQADdVsoRd3xzhPNeufBAg" // Porra Erick
		case strings.Contains(update.Message.Text, "@viniciusferrao") && r < 5 && (update.Message.From.UserName == "hsantanna88"):
			voiceId = voiceHugo[rand.Intn(len(voiceHugo))]
		case ferraoReply(&update.Message.Text, kwsAndroid, android, 80):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsNewAppleDevices, newAppleDevices, 70):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsLetsPlay, LetsPlay, 90):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsTimCook, timCook, 70):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsFakeNews, fakeNews, 25):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsAmir, amir, 70):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsOcupado, ocupado, 70):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsCanada, canada, 70):
			if r > 50 {
				message = update.Message.Text
			} else {
				voiceId = voiceCanadaNossaMota[rand.Intn(len(voiceCanadaNossaMota))]
			}
		case ferraoReply(&update.Message.Text, kwsCo7, co7, 80):
			message = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsSumido, randomPhotoId, 90):
			photoId = update.Message.Text
		case ferraoReply(&update.Message.Text, kwsAmirPhoto, amirPhoto, 90):
			photoId = update.Message.Text
		case strings.Contains(update.Message.Text, "macaco") && (r < 90):
			videoId.fileId = macacoVideo[rand.Intn(len(macacoVideo))].fileId
			videoId.duration = macacoVideo[rand.Intn(len(macacoVideo))].duration
		case strings.Contains(update.Message.Text, "@viniciusferrao") && (r < 90) && (update.Message.From.UserName != "viniciusferrao"):
			if r < 35 {
				videoId.fileId = videoMentionFerrao[rand.Intn(len(videoMentionFerrao))].fileId
				videoId.duration = videoMentionFerrao[rand.Intn(len(videoMentionFerrao))].duration
			} else {
				message = mentionFerrao[rand.Intn(len(mentionFerrao))]
			}
		case update.Message.Photo != nil && (r < 40) && (update.Message.From.UserName != "viniciusferrao"):
			if r < 20 {
				voiceId = voiceLol[rand.Intn(len(voiceLol))]
			} else {
				message = replyImage[rand.Intn(len(replyImage))]
			}
		case (update.Message.From.UserName == "maulmota") && (r < 5) && (update.Message.From.UserName != "viniciusferrao"):
			message = nossaMota[rand.Intn(len(nossaMota))]
		case update.Message.Sticker != nil:
			if update.Message.Sticker.FileID == "CAADAQADjgADZ_LqCT89ryHtPlS0Ag" {
				message = "Isso de ficar de RIP Ferrão já deu no saco, depois ficam de babaquice. E sim to puto"
			}
		case ferraoReply(&update.Message.Text, kwsTimMaia, timMaiaAudio, 100):
			audioId = update.Message.Text // Geraldão
		}

		if message == "" && audioId == "" && voiceId == "" && photoId == "" && videoId.fileId == "" {
			continue
		}
		time.Sleep(time.Millisecond * 3000)

		switch {
		case audioId != "":
			audio := tgbotapi.NewAudioShare(update.Message.Chat.ID, audioId)
			audio.ReplyToMessageID = update.Message.MessageID
			bot.Send(audio)
		case voiceId != "":
			voice := tgbotapi.NewVoiceShare(update.Message.Chat.ID, voiceId)
			voice.ReplyToMessageID = update.Message.MessageID
			bot.Send(voice)
		case photoId != "":
			photo := tgbotapi.NewPhotoShare(update.Message.Chat.ID, photoId)
			bot.Send(photo)
		case videoId.fileId != "" && videoId.duration != 0:
			video := tgbotapi.NewVideoNoteShare(update.Message.Chat.ID, videoId.duration, videoId.fileId)
			bot.Send(video)
		default:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
