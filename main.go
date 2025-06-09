package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// Write a getMessageText function.
// It should accept a pointer to an Analytics struct and a Message struct (not a pointer).
// It should look at the Success field of the Message struct and, based on that,
// increment the MessagesTotal, MessagesSucceeded, or MessagesFailed fields of
// the Analytics struct as appropriate.

func getMessageText(anal *Analytics, msg Message) {
	if msg.Success {
		(*anal).MessagesSucceeded++
	} else {
		(*anal).MessagesFailed++
	}
	(*anal).MessagesTotal++
}
