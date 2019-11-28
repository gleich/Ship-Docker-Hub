package main

import (
	log "github.com/sirupsen/logrus"
	logErr
	"os"
	"fmt"
)

func main() {
	fmt.Println("Logging into docker hub as user:")
	file, err := os.Open("dockerUsername.txt")
	if err != nil {
		fmt.Println("Make sure that the dockerUsername.txt file in the volume")
		log.Error("dockerUsername.txt not found. Make sure it is in volume")
	}
	fmt.Println(file)
}