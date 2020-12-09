package main

import (
	"fmt"
	"github.com/Jviguy/SpeedyCmds"
	"github.com/Jviguy/SpeedyCmds/command/commandGroup"
	"github.com/Jviguy/SpeedyCmds/command/commandMap"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	dg, err := discordgo.New("Bot " + "TOKEN HERE")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	//i might remove the making of a command.Map but atm its in for init command injection
	handler := SpeedyCmds.New(dg,commandMap.New(),true,"!")
	//init a new base command group the name in the struct doesnt matter btw
	pg := commandGroup.New("api")
	//add a command to the said command group
	pg.AddCommand("ping",Ping{"Ping"})
	//add the said command group to the commandHandler HERE THE NAME MATTERS
	handler.GetCommandHandler().RegisterCommandGroup("api",pg)
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
