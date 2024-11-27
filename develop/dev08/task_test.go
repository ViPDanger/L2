package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestKotShell01(t *testing.T) {
	cmd := exec.Command(".\\task.exe")
	executeDir, _ := os.Getwd()
	expected :=
		`KOTSHELL: pwd
` + executeDir + `
KOTSHELL: cd nc
KOTSHELL: echo HelloWorld!
HelloWorld!
KOTSHELL: pwd
` + executeDir + `\nc
KOTSHELL: exit`
	sendtoWriter := []string{
		"pwd",
		"cd nc",
		"echo HelloWorld!",
		"pwd",
		"exit"}
	defaultTest(t, expected, cmd, sendtoWriter)
}

// Прошу прощения, я не представляю как тестировать ps и nc
func TestKotShell02(t *testing.T) {
	cmd := exec.Command(".\\task.exe")
	executeDir, _ := os.Getwd()
	expected := executeDir + ":" + executeDir + "\\nc:HelloWorld!" + executeDir + "\\nc:"
	sendtoWriter := []string{
		"ps",
		"exit"}
	defaultTest(t, expected, cmd, sendtoWriter)
}

type CustomOut struct {
	data []byte
}

func (c *CustomOut) Write(data []byte) (n int, err error) {
	c.data = append(c.data, data...)
	return len(data), nil
}

func defaultTest(t *testing.T, expected string, cmd *exec.Cmd, sendtoWriter []string) {

	reader, writer := io.Pipe()
	out := &CustomOut{}
	var err error
	cmd.Stdout = out
	cmd.Stdin = reader
	cmd.Stderr = os.Stderr
	go func() {
		defer writer.Close()
		for _, str := range sendtoWriter {
			time.Sleep(time.Millisecond * 100)
			if _, err := writer.Write([]byte(str + "\n")); err != nil {
				log.Fatalln(err)
			}
			if _, err := out.Write([]byte(str + "\n")); err != nil {
				log.Fatalln(err)
			}

		}
	}()
	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}
	result := strings.ReplaceAll(string(out.data), "", "")
	result = strings.Trim(result, "\n \r")
	expected = strings.Trim(expected, "\n \r")
	//result = strings.ReplaceAll(result, " ", "")
	if strings.Compare(result, expected) != 0 {
		t.Fatal("Incorrect result\n--Result--\n", result, "\n--Expected--\n", expected)
	}
}
