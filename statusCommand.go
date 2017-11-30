package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func statusCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
	return fmt.Sprintf("Currently %v players online.", server.PlayerAmount)
}
