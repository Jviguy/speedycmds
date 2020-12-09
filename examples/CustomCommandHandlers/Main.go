package main

import (
	"CustomCommandHandlers/commandhandler"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Jviguy/SpeedyCmds/command/commandMap"
	"github.com/bwmarrin/discordgo"
)

func main() {
	//my token was there for a old bot again so fuck me lmao
	dg, err := discordgo.New("Bot " + "TOKEN HERE")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	var prefixes = []string{"!", "go ", "bruh ", "-"}
	//use of a custom implemented Handler that allows for multiple prefixes
	handler := commandhandler.New(dg, commandMap.New(), prefixes)
	handler.GetCommandHandler().RegisterCommand("ping", Ping{"ping"}, true)
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = dg.Close()
}
