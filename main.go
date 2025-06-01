package main

func (e email) cost() int {
	if e.isSubscribed {
		return 2 * len(e.body)
	}
	return 5 * len(e.body)
}

func (e email) format() string {
	str := "'" + e.body + "'" + " | "
	if e.isSubscribed {
		return str + "Subscribed"
	}
	return str + "Not Subscribed"
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
