package utils

import (
	"io/ioutil"
	"log"
	"os/user"
)

func WriteToFile(data string, file string) {
	ioutil.WriteFile(file, []byte(data), 777)
}

func WriteBytesToFile(data []byte, file string) {
	ioutil.WriteFile(file, data, 777)
}

func ReadBytesFromFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	return data, err
}

func ReadFromFile(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	return string(data), err
}

func GetHomeDir() (homeDir string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
