package main

import (
	"github.com/Jviguy/SpeedyCmds/command/ctx"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

type Ping struct {

}

func (p Ping) Execute(ctx ctx.Ctx, session *discordgo.Session) error {
	_, err := session.ChannelMessageSend(ctx.GetChannel().ID, "Pong at " + strconv.Itoa(int(session.HeartbeatLatency().Milliseconds())) + "ms !")
	return err
}
