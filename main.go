package main

import (
	"errors"
)

const (
	planFree = "free"
	planPro  = "pro"
)

var errUnsupportedPlan = errors.New("unsupported plan")

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	// ?
	switch plan {
	case planFree:
		return messages[:2], nil
	case planPro:
		return messages[:], nil
	default:
		return nil, errUnsupportedPlan
	}
}
