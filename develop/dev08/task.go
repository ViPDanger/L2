package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	ps "github.com/mitchellh/go-ps"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	executeDir, _ := os.Executable()
	executeDir = filepath.Dir(executeDir)
	for {

		fmt.Print("KOTSHELL: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if input != "" {
			if err = kotShell(input, executeDir); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}

	}
}

func kotShell(input string, executeDir string) error {
	errPath := errors.New("path is required")
	errArgs := errors.New("args is required")
	input = strings.TrimFunc(input, func(r rune) bool {
		switch r {
		case '\n':
			return true
		case '\r':
			return true
		}
		return false
	})

	args := strings.Split(input, " ")
	switch args[0] {
	case "cd":
		if len(args) < 2 {
			return errPath
		}
		return os.Chdir(args[1])
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		fmt.Println(dir)
	case "echo":
		if len(args) < 2 {
			return errArgs
		}
		for i := 1; i < len(args); i++ {
			fmt.Fprintln(os.Stdout, args[i])
		}
	case "kill":
		if len(args) != 2 {
			return errArgs
		}
		pid, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			return nil
		}
		err = proc.Kill()
		if err != nil {
			return err
		}
	case "fork":
	case "ps":
		processList, err := ps.Processes()
		if err != nil {
			log.Println("ps.Processes() Failed, are you using windows?")
			return nil
		}

		for _, process := range processList {
			log.Printf("%d\t%s\n", process.Pid(), process.Executable())
		}
	case "exit":
		os.Exit(0)
	case "exec":
		cmd := exec.Command(args[1], args[2:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	case "nc":
		cmd := exec.Command(executeDir+"\\nc\\nc.exe", args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	default:
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	return nil
}
