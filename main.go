package main

import (
	"Snek/bot"
)

func main() {
	go startServer()
	bot.StartBot()
}
