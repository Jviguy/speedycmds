package utils

import "github.com/Jviguy/GoingCommando/command"

func GetAllKeys(m map[string]command.Command) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
