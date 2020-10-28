package main

import (
	"github.com/Jviguy/GoingCommando/command/ctx"
	"github.com/bwmarrin/discordgo"
)

type Ping struct {
	name string
}

func (p Ping) GetName() string {
	return p.name
}

func (p Ping) Execute(ctx ctx.Ctx, session *discordgo.Session) error {

}

