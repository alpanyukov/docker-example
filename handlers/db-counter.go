package handlers

import (
	"log"

	"github.com/gomodule/redigo/redis"
)

type DBCounterMessaging struct {
	redis.Conn
	field string
}

func (m *DBCounterMessaging) getMessage() interface{} {
	_, err := m.Do("INCR", m.field)
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	newValue, _ := redis.Int(m.Do("GET", m.field))
	return newValue
}

func NewDBCounterMessaging(host string) *DBCounterMessaging {
	conn, err := redis.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	field := "counter"
	exists, _ := redis.Bool(conn.Do("EXISTS", field))
	if !exists {
		conn.Do("SET", field, 0)
	}

	return &DBCounterMessaging{
		conn,
		field,
	}
}
