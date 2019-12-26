package handlers

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

type FileCounterMessaging struct {
	cnt int
	*os.File
	*bufio.ReadWriter
}

func (m *FileCounterMessaging) getMessage() interface{} {
	m.cnt++
	m.syncFileWithTemp()
	return "Counter: " + strconv.Itoa(m.cnt)
}

func (m *FileCounterMessaging) syncTempWithFile() {
	countRaw, err := m.ReadWriter.ReadString('\n')
	if err != io.EOF {
		log.Println(err.Error())
	}
	log.Println("readed", countRaw)
	if countRaw == "" {
		m.cnt = 0
		return
	}
	count, _ := strconv.Atoi(countRaw)
	m.cnt = count
}
func (m *FileCounterMessaging) syncFileWithTemp() {
	log.Println("save number", m.cnt)
	m.File.Truncate(0)
	m.File.Seek(0, 0)
	_, err := m.ReadWriter.WriteString(strconv.Itoa(m.cnt))
	m.ReadWriter.Flush()
	if err != nil {
		log.Println("cant write to file", err.Error())
	}
}

func NewFileCounterMessaging() *FileCounterMessaging {
	path := "data/counter"
	f, err := openOrCreate(path)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReadWriter(bufio.NewReader(f), bufio.NewWriter(f))
	fcm := &FileCounterMessaging{0, f, r}
	fcm.syncTempWithFile()
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
		f, err = os.Create("data/counter")
		if err != nil {
			return nil, err
		}
	} else {
		f, err = os.OpenFile(path, os.O_RDWR, os.ModePerm)
	}
	return f, err
}
