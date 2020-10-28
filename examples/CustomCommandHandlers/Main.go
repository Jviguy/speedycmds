package main

import (
	"CustomCommandHandlers/commandhandler"
	"fmt"
	"github.com/Jviguy/GoingCommando/command/commandmap"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	dg, err := discordgo.New("Bot " + "NzUyNjkxMDU2ODgxMTcyNjIy.X1bULA.CrYPU434skqnFOpzbr-FkovwAfE")
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	var prefixes = []string {"!","go ","bruh ","-"}
	//use of a custom implemented Handler that allows for multiple prefixes
	handler := commandhandler.New(dg,commandmap.New(),prefixes)
	handler.GetCommandHandler().RegisterCommand("ping" , Ping{"ping"},true)
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = dg.Close()
}
