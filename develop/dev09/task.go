package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	url := flag.String("url", "", "")
	timeout := flag.Duration("timeout", 7*time.Second, "")
	output_path := flag.String("output", "task.html", "")

	flag.Parse()
	content, err := wGet(*url, *timeout)
	if err != nil {
		log.Fatalln("wGet error: ", err)
	}
	err = os.WriteFile(*output_path, content, 0666)
	if err != nil {
		log.Fatalln("os.WriteFile error: ", err)
	}
}
func wGet(url string, timeout time.Duration) (content []byte, err error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	ctx, cancel_func := context.WithTimeout(context.Background(), timeout)
	request = request.WithContext(ctx)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		cancel_func()
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		cancel_func()
		return nil, fmt.Errorf("INVALID RESPONSE; status: %s", response.Status)
	}
	res, err := io.ReadAll(response.Body)
	cancel_func()
	return res, err
}
