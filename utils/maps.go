package utils

import (
	"github.com/Jviguy/SpeedyCmds/command"
)

// CommandMapToKeys converts a map from key to command to a string slice of keys.
func CommandMapToKeys(m map[string]command.Command) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}
