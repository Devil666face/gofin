package main

import (
	. "github.com/Devil666face/gofinabot/cmd"
)

func main() {
	bot, err := Bot()
	if err != nil {
		panic(err)
	}
	bot.Start()
}
