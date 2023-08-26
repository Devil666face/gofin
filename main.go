package main

import (
	. "github.com/Devil666face/gofinabot/cmd"
)

func main() {
	switch Cli() {
	case MIGRATE:
		err := Migrate()
		if err != nil {
			panic(err)
		}

	case START:
		bot, err := Bot()
		if err != nil {
			panic(err)
		}
		bot.Start()
	}
}
