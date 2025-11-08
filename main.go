package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	const filepath = "/home/pires/.bashrc"

	file, err := os.OpenFile(filepath, os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	command := "export PATH=$PATH:" + path

	file.WriteString(command)

	file.Sync()

	cmd := exec.Command(command)
	defer cmd.Run()

	fmt.Print("ok\n")

}
