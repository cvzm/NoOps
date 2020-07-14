package main

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func Test4(t *testing.T) {
	cmd := exec.Command("ansible", "all", "-m", "ping")
	cmd.Stdout = new(outstream)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmd.Wait())
}

type outstream struct{}

func (out outstream) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}
