package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func getIdCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
	return fmt.Sprintf("Channel Id is %v", m.ChannelID)
}
