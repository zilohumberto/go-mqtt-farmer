package handler

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// getOrCreate checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func getOrCreate(filename string) *os.File {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			log.Panic(err)
		}
		return file
	}
	if info.IsDir() {
		log.Panic(errors.New("is file is already as a folder"))
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Panic(err)
	}
	return file
}

func SaveInformation(uuid string, action string, topic string, payload []byte) {
	filename := "/tmp/" + uuid + ".in"
	file := getOrCreate(filename)
	defer file.Close()
	s := fmt.Sprintf("%s|%s|%s\n", action, topic, payload)
	_, err := file.WriteString(s)
	if err != nil {
		log.Panic(err)
	}
}
