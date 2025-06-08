package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		err := errors.New("invalid sizes")
		return nil, err
	}

	userMap := make(map[string]user)

	for i, name := range names {
		userMap[name] = user{name: name, phoneNumber: phoneNumbers[i]}
	}

	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}
