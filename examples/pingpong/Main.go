package main

import (
	"fmt"
	"github.com/Jviguy/GoingCommando"
	"github.com/Jviguy/GoingCommando/command/commandmap"
	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("Bot " + "")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	//i might remove the making of a command.Map but atm its in for init command injection
	handler := GoingCommando.New(dg,commandmap.New(),"!")
	handler.GetCommandHandler().RegisterCommand("Ping",Ping{"Ping"},true)
}
