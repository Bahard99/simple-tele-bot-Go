package main

import (
	"fmt"
	"os"

	bt "github.com/SakoDroid/telego"
	cfg "github.com/SakoDroid/telego/configs"
	objs "github.com/SakoDroid/telego/objects"
)

const token string = "[your-tele-bot-token-here]"

var bot *bt.Bot

func main()  {
	
	up := cfg.DefaultUpdateConfigs()

	//Bot Configs
	cf := cfg.BotConfigs {
		BotAPI: cfg.DefaultBotAPI,
		APIKey: token, 
		UpdateConfigs: up,
		Webhook: false,
		LogFileAddress: cfg.DefaultLogFile,
	}

	var err error

	//Creating bot using created configs
	bot, err = bt.NewBot(&cf)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = bot.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	start()

}

func start()  {
	
	bot.AddHandler("/start", func(u *objs.Update)  {

		//Create the custom keyboard
		kb := bot.CreateKeyboard(false, false, false, "type ...")

		//Add button, first is text and the second is row num
		kb.AddButton("/hi", 1)
		kb.AddButton("/start", 1)
		kb.AddButton("/pic", 1)
		kb.AddButton("/inline-keyboard", 1)

		//Reply message and pass the keyboard to send method (to be a reply set 0 with "u.Message.MessageId")
		_, err := bot.AdvancedMode().ASendMessage(u.Message.Chat.Id, "Hi im telegram bot!", "", 0, false, false, nil, false, false, kb)
		if err != nil {
			fmt.Println(err)
			return
		}

	}, "private", "group")

	bot.AddHandler("/hi", func(u *objs.Update)  {

		// Sends the message to a message (to be a reply set 0 with "u.Message.MessageId")
		_, err := bot.SendMessage(u.Message.Chat.Id, "Hi "+u.Message.From.FirstName+", Im a telegram bot! \n/start : for any info\npic : for picture", "", 0, false, false)
		if err != nil {
			fmt.Println(err)
			return
		}

	}, "private")

	bot.AddHandler("/inline-keyboard", func(u *objs.Update)  {
		
		//Creates inline keyboard
		kb := bot.CreateInlineKeyboard()

		kb.AddURLButton("url", "https://google.com", 1)

		_, err := bot.AdvancedMode().ASendMessage(u.Message.Chat.Id, "hi, this is inline keyboard", "", 0, false, false, nil, false, false, kb)
		if err != nil {
			fmt.Println(err)
			return
		}

	}, "private")

	bot.AddHandler("/pic", func(u *objs.Update)  {
		
		//msgId for reply type-message
		chatId := u.Message.Chat.Id
		// msgId  := u.Message.MessageId
		
		//Create media sender for sendir a URL (became reply if filled with msgId)
		mediaSender1 := bot.SendPhoto(chatId, 0, "Kucing 1", "")

		_, err := mediaSender1.SendByFileIdOrUrl("https://i.insider.com/61d1c0e2aa741500193b2d18?width=1000&format=jpeg&auto=webp", false, false)
		if err != nil {
			fmt.Println(err)
			return
		}

		//Create media sender for sending file (became reply if filled with msgId)
		mediaSender2 := bot.SendPhoto(chatId, 0, "Kucing 2", "")

		file, err := os.Open("funny_cat.jpg")
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = mediaSender2.SendByFile(file, false, false)
		if err != nil {
			fmt.Println(err)
			return
		}
	}, "private")

	//Custom handler for reply with text after "panggil aku "
	bot.AddHandler("panggil aku *", func(u *objs.Update)  {

		textAll		 := u.Message.Text
		lenIndexText := len(textAll)
		nameFromText := string(textAll[12:lenIndexText])

		// Sends the message to a message (to be a reply set 0 with "u.Message.MessageId")
		_, err := bot.SendMessage(u.Message.Chat.Id, "Hi "+nameFromText+", Semoga sehat selalu!", "", 0, false, false)
		if err != nil {
			fmt.Println(err)
			return
		}

	}, "private")

	//Register the channel
	messageChannel, _ := bot.AdvancedMode().RegisterChannel("", "message")

	for {
		//Wait for updates
		up := <- *messageChannel

		//Print the text
		fmt.Println(up.Message.Text)
	}
}