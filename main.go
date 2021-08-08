package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron/v3"
	"github.com/slack-go/slack"
)

func main() {
	token := "your_token"
	client := slack.New(token)

	c := cron.New()

	c.AddFunc("@every 1m", func() {
		gameId := "gameId"
		url := "https://baseball.yahoo.co.jp/live/mlb/game/" + gameId + "/live"
		doc, err := goquery.NewDocument(url)
		if err != nil {
			_, _, err := client.PostMessage("#error_channel", slack.MsgOptionText(err.Error(), true))
			panic(err)
		}

		batterName := doc.Find(".batter em").Text()

		if batterName != "大谷" {
			return
		}

		_, _, e := client.PostMessage("#your_channel", slack.MsgOptionText("OHTANISAN!!", true))
		if e != nil {
			_, _, e := client.PostMessage("#error_channel", slack.MsgOptionText(e.Error(), true))
			if e != nil {
				panic(err)
			}
		}
	})

	c.Start()

	for {
	}
}
