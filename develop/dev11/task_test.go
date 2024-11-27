package main

import (
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestDev11_01(t *testing.T) {
	r := strings.NewReader("user_id=3&date=2019-09-09")
	res, err := http.Post("http://localhost:8080/create_event", "", r)
	expected := `{"result":{"user_id":"3","date":"2019-09-09"},"action":"Create Event!"}`
	if err != nil {
		log.Fatalln(err)
	}
	data := make([]byte, 100)
	_, err = res.Body.Read(data)
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}
	result := string(data)
	result = result[strings.IndexRune(result, '{') : strings.LastIndex(result, "}")+1]
	res.Body.Close()
	if result != expected {
		t.Fatal("Incorrect result\n--Result--\n", result, "\n--Expected--\n", expected)
	}

}

func TestDev11_02(t *testing.T) {
	res, err := http.Get("http://localhost:8080/events_for_week?user_id=ИванПавлович&date=2024-08-07")
	expected := `{"result":{"user_id":"ИванПавлович","date":"2024-08-07"},"action":"Events for Week!"}`
	if err != nil {
		log.Fatalln(err)
	}
	data := make([]byte, 100)
	_, err = res.Body.Read(data)
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}
	result := string(data)
	result = result[max(strings.IndexRune(result, '{'), 0):min(strings.LastIndex(result, "}")+1, len(result)-1)]
	res.Body.Close()
	if result != expected {
		t.Fatal("Incorrect result\n--Result--\n", result, "\n--Expected--\n", expected)
	}

}
func TestDev11_03(t *testing.T) {
	res, err := http.Get("http://localhost:8080/events_for_week")
	if err != nil {
		log.Fatalln(err)
	}
	if res.StatusCode != 400 {
		log.Fatalln("respound status code =", res.StatusCode, " must been 400")
	}
}

func TestDev11_04(t *testing.T) {
	res, err := http.Get("http://localhost:8080/boop")
	if err != nil {
		log.Fatalln(err)
	}
	if res.StatusCode != 404 {
		log.Fatalln("respound status code =", res.StatusCode, " must been 404")
	}
}
