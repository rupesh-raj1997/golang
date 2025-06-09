package main

import (
	"strings"
)

func removeProfanity(message *string) {
	msg := *message
	msg = strings.ReplaceAll(msg, "fubb", "****")
	msg = strings.ReplaceAll(msg, "shiz", "****")
	msg = strings.ReplaceAll(msg, "witch", "*****")
	*message = msg
}
