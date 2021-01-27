package main

import (
	"io/ioutil"

	"github.com/mutsuki333/log"
)

func main() {
	log.Println("This is for debug")
	log.Info("This is info")

	log.SetOutput(ioutil.Discard)
	log.Println("No more debug logs")
}
