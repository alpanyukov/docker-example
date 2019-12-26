package handlers

import "strconv"

type CounterMessaging struct {
	count int
}

func (m *CounterMessaging) getMessage() interface{} {
	m.count++
	return "Counter: " + strconv.Itoa(m.count)
}

func NewCounterMessaging() *CounterMessaging {
	return &CounterMessaging{}
}
