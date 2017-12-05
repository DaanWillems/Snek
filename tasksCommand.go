package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	_ "github.com/go-sql-driver/mysql"

	"github.com/knq/dburl"

	"test/models"
)

func tasksCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) string {
	db, err := dburl.Open("[redacted]")
	if err != nil {
		log.Fatal(err)
	}

	tasks, err := models.TaskAll(db)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%v", tasks)
}
