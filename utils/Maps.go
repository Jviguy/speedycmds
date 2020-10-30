package utils

import (
	"github.com/Jviguy/GoingCommando/command"
	"github.com/Jviguy/GoingCommando/command/commandGroup"
)

func GetAllKeysCommands(m map[string]command.Command) []string {
	keys := make([]string, 0, len(m))
	for k,_ := range m {
		keys = append(keys, k)
	}
	return keys
}

func GetAllKeysGroups(m map[string]commandGroup.Group) []string {
	keys := make([]string, 0, len(m))
	for k,_ := range m {
		keys = append(keys, k)
	}
	return keys
}
