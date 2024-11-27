package main

import (
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestWGet01(t *testing.T) {
	FileName := "test01.html"
	cmd := exec.Command(".\\task.exe", "-url", "https://github.com", "-output", FileName)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	if fileSize, err := FileSize(FileName); err != nil {
		t.Fatal(err)
	} else if fileSize < 2000 {
		t.Fatal(errors.New("Размер скачанного менее 2х тысяч символов"))
	}
}
func TestWGet02(t *testing.T) {
	FileName := "test02.html"
	cmd := exec.Command(".\\task.exe", "-url", "https://ya.ru", "-output", FileName, "-timeout", "1")
	if err := cmd.Run(); err != nil {
		if err.Error() == "exit status 2" {
			return
		}
		log.Fatal(err)
	}

}

func TestWGet03(t *testing.T) {
	FileName := "test03.html"
	cmd := exec.Command(".\\task.exe", "-url", "https://ya.ru", "-output", FileName)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	if fileSize, err := FileSize(FileName); err != nil {
		t.Fatal(err)
	} else if fileSize < 2000 {
		t.Fatal(errors.New("Размер скачанного менее 2х тысяч символов"))
	}
}
func FileSize(FileName string) (int, error) {
	File, err := os.Open(FileName)
	if err != nil {
		return 0, err
	}
	defer File.Close()
	FileData, err := io.ReadAll(File)
	if err != nil {
		return 0, err
	}
	return len(FileData), nil

}
