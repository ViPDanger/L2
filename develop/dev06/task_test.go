package main

import (
	"io"
	"log"
	"os/exec"
	"testing"
	"time"
)

type CustomOut struct {
	data []byte
}

func (c *CustomOut) Write(data []byte) (n int, err error) {
	c.data = append(c.data, data...)
	return len(data), nil
}

func TestCut01(t *testing.T) {
	expected := "Введите строки\nsecond\nfive\neight\n"
	cmd := exec.Command(".\\task.exe", "-f", "2")
	sendtoWriter := []string{"first	second	third\n", "four	five	six\n", "seven	eight	nine\n", "/p"}
	defaultTest(t, expected, cmd, sendtoWriter)
}

func TestCut02(t *testing.T) {
	expected := "Введите строки\nfirst\nsecond	third\nfour\nfive\nsix\n"
	cmd := exec.Command(".\\task.exe", "-d", ".", "-s")
	sendtoWriter := []string{"first.second	third\n", "four.five.six\n", "seven	eight	nine\n", "/p"}
	defaultTest(t, expected, cmd, sendtoWriter)
}
func TestCut03(t *testing.T) {
	expected := "Введите строки\nthird\nnine\n"
	cmd := exec.Command(".\\task.exe", "-f", "3", "-d", " ", "-s")
	sendtoWriter := []string{"first second third\n", "four five\n", "six,seven eight nine\n", "/p"}
	defaultTest(t, expected, cmd, sendtoWriter)
}
func defaultTest(t *testing.T, expected string, cmd *exec.Cmd, sendtoWriter []string) {
	reader, writer := io.Pipe()
	out := &CustomOut{}
	var err error
	if err != nil {
		log.Fatal(err)
	}
	cmd.Stdout = out
	cmd.Stdin = reader
	go func() {
		defer writer.Close()
		for _, str := range sendtoWriter {
			if _, err := writer.Write([]byte(str)); err != nil {
				log.Fatalln(err)
			}
			time.Sleep(time.Millisecond * 100)
		}
	}()
	if err = cmd.Run(); err != nil {
		log.Fatal(err)
	}
	result := string(out.data)
	if expected != result {
		t.Fatal("Incorrect result:\n", result, "\n--Expected\n", expected)
	}
}
