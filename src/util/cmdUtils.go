package util

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

// RunCmd takes a command and args and runs it, streaming output to stdout
func RunCmd(cmdName string, cmdArgs ...string) error {
	cmd := exec.Command(cmdName, cmdArgs...)
	log.Printf(" Run: %s \n", cmd.String())
	value, err := scannerScan(cmd.StdoutPipe())
	if err != nil {
		return err
	}
	valueErr, err := scannerScan(cmd.StderrPipe())
	if err != nil {
		return err
	}
	if err := cmd.Run(); err != nil {
		return err
	}
	<-value
	<-valueErr
	return nil
}

//管道输出处理
func scannerScan(r io.Reader, err error) (chan string, error) {
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	var value = make(chan string, 2)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		value <- scanner.Text()
	}()
	return value, nil
}
