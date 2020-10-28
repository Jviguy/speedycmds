package command

import (
	"github.com/Jviguy/GoingCommando/command/ctx"
	"github.com/bwmarrin/discordgo"
)
//the simple command Interface for executing code when its called.
type Command interface {
	GetName() string
	Setname(newname string)
	Execute(ctx ctx.Ctx,session *discordgo.Session) error
}

