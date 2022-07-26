package File

import (
	"log"
	"os"
)

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}

func WriteToFile(name string, msg string) error {
	if !FileExists(name) {
		CreateFile(name)
	}

	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(msg)

	return nil
}

func ClearFile(name string) error {
	if err := os.Truncate(name, 0); err != nil {
		return err
	}

	return nil
}
