package handlers

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type fileCounterMessaging struct {
	filepath string
}

func (m *fileCounterMessaging) getMessage() interface{} {
	cnt, _ := m.readFromFile()
	cnt++
	m.writeToFile(cnt)
	return "Counter: " + strconv.Itoa(cnt)
}

func (m *fileCounterMessaging) readFromFile() (int, error) {
	cntRaw, err := ioutil.ReadFile(m.filepath)
	if err != nil {
		return 0, err
	}

	cnt, err := strconv.Atoi(string(cntRaw))
	if err != nil {
		return 0, err
	}

	return cnt, nil
}

func (m fileCounterMessaging) writeToFile(cnt int) error {
	return ioutil.WriteFile(m.filepath, []byte(strconv.Itoa(cnt)), 0644)
}

func NewFileCounterMessaging() *fileCounterMessaging {
	path := "data/counter"
	_, err := openOrCreate(path)
	if err != nil {
		log.Fatal(err)
	}
	fcm := &fileCounterMessaging{filepath: path}
	return fcm
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func openOrCreate(path string) (*os.File, error) {
	var f *os.File
	var err error
	if !fileExists(path) {
		// Папка должна быть создана! Это не рекурсивное создание
		_, err = os.Create(path)
		if err != nil {
			return nil, err
		}
	}
	f, err = os.OpenFile(path, os.O_RDWR, 0644)
	return f, err
}
