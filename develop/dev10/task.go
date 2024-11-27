package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/ziutek/telnet"
)

func main() {
	var err error
	t := flag.Int("timeout", 10, "timeout")
	host := flag.String("host", "localhost", "host")
	port := flag.String("port", "8080", "port")
	auto := flag.Bool("auto", false, "")
	flag.Parse()
	timeout := time.Duration(*t) * time.Second
	address := fmt.Sprint(*host, ":", *port)
	var conn *telnet.Conn

	conn, err = telnet.DialTimeout("tcp", address, timeout)
	if err != nil {
		log.Fatalln(err)
	}
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGALRM)
	go func() {
		<-ctx.Done()
		log.Println("Gracefull shoutdown")
		os.Exit(1)

	}()

	go func() {
		if *auto {
			i := 0
			for {
				msg := fmt.Sprintln(strconv.Itoa(i))
				_, err = conn.Write([]byte(msg))
				if err != nil {
					log.Fatalln(err)
				}
				i++
				time.Sleep(100 * time.Millisecond)
			}
		} else {
			reader := bufio.NewScanner(os.Stdin)
			for reader.Scan() {
				input := reader.Text()
				if input != "" {
					if _, err = conn.Write([]byte(input)); err != nil {
						fmt.Fprintln(os.Stderr, err)
					}
				}
			}
			ctx.Done()
		}

	}()

	for {
		reply := make([]byte, 1024)
		_, err = conn.Read(reply)
		if err != nil {
			log.Fatalln(err)
		}
		str := string(reply)
		fmt.Fprint(os.Stdout, str)
	}
}
