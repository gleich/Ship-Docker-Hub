package main

import (
	"fmt"
	"github.com/Matt-Gleich/logErr"
	"os"
	"io/ioutil"
)

func main() {
	fmt.Println("Logging into docker hub as user:")
	file, err := os.Open("dockerUsername.txt")
	logErr.Log(err, "Make sure that the dockerUsername.txt file is in the volume")
	defer file.Close()
	b, err := ioutil.ReadAll(file)
	logErr.Log(err, "")
	fmt.Println(string(b))
	loginCmdexec.Command("docker", "--version")
}
