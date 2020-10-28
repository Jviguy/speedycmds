package command

import "github.com/bwmarrin/discordgo"
//the simple command Interface for executing code when its called.
type Command interface {
	GetName() string
	Execute(args []string,session *discordgo.Session) error
}
