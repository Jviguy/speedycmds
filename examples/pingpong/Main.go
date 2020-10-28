package main

import (
	"fmt"
	"github.com/Jviguy/GoingCommando"
	"github.com/Jviguy/GoingCommando/command/commandmap"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
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
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = dg.Close()
}
