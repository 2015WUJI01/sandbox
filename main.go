package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("host", "engine10.uptimerobot.com")
	s, _ := cmd.Output()
	fmt.Println(string(s))
}
