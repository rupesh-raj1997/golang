package main

import (
	"errors"
	"fmt"
)

type email struct {
	status    string
	recipient *user
}

type user struct {
	email  string
	status string
}

type analytics struct {
	totalBounces int
}

func (u *user) updateStatus(status string) error {
	if status != "email_bounced" {
		err := fmt.Errorf("error updating user status: %w: %s", errors.New("invalid status"), status)
		return err
	}
	u.status = status
	return nil
}

func (a *analytics) track(event string) error {
	if event != "email_bounced" {
		err := fmt.Errorf("error tracking user bounce: %w: %s", errors.New("invalid status"), event)

		return err
	}
	a.totalBounces++
	return nil
}
