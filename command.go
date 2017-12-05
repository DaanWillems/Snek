package main

import (
	"github.com/bwmarrin/discordgo"
)

var commandMap map[string]func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string

func loadCommands() {
	commandMap = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string{
		"!status": statusCommand,
		"!id":     getIdCommand,
		"!task":   tasksCommand,
		"!help":   statusCommand,
	}
}

func evaluateCommand(s *discordgo.Session, m *discordgo.MessageCreate) string {
	args, cmd := parseCommand(m.Content)

	if _, ok := commandMap[cmd]; !ok {
		return commandMap["!help"](s, m, args)
	}

	return commandMap[cmd](s, m, args)
}
