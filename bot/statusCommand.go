package bot

import "github.com/bwmarrin/discordgo"

func statusCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
	return "Currently 0 players online."
}
