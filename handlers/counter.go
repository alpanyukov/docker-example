package handlers

import "strconv"

type counterMessaging struct {
	count int
}

func (m *counterMessaging) getMessage() interface{} {
	m.count++
	return "Counter: " + strconv.Itoa(m.count)
}

func NewCounterMessaging() *counterMessaging {
	return &counterMessaging{}
}
