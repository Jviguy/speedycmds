package GoingCommando

import (
	"github.com/Jviguy/GoingCommando/command"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Handler interface {
	Handle(s *discordgo.Session, msg *discordgo.MessageCreate)
}
//this is the basic premade handler with support for one prefix u can change it up how u want because of the whole interfaces
type PremadeHandler struct {
	dg *discordgo.Session
	cmds command.Map
	Prefix string
}
//initalizes a basic premade handler for you.
func New(dg *discordgo.Session,cmds command.Map,prefix string) PremadeHandler {
	h := PremadeHandler{dg,cmds,prefix}
	dg.AddHandler(h.Handle)
	return h
}
//called when a message is sent checks if its a bot and if it has the right prefix
func (h PremadeHandler) Handle(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.Bot {
		return
	}
	if strings.HasPrefix(msg.Content,h.Prefix){
		args := strings.Split(strings.TrimPrefix(msg.Content,h.Prefix)," ")
		args , cmd := Shift(args,0)
		err := h.cmds.Execute(cmd,args,s)
		if err != nil{
			_,err = s.ChannelMessageSend(msg.ChannelID,"An Error Occurred while executing that command" +
				"\nDEBUG: " + err.Error())
		}
	}
}

func Shift(a []string,i int) ([]string,string) {
	b := a[i]
	copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
	a[len(a)-1] = ""     // Erase last element (write zero value).
	a = a[:len(a)-1]     // Truncate slice.
	return a,b
}

