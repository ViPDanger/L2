package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"testing"
)

// task.exe -k 2 test01.txt
func TestSort01(t *testing.T) {
	testFileName := "test01.txt"
	CreateTestFiles("sample.txt", testFileName)
	cmd := exec.Command(".\\task.exe", "-k", "2", testFileName)
	expected := []string{
		"2 Вуду чайлд 5",
		"5 Кукла колдуна 12",
		"6 Сомбади вот ви юзед ту ноy 9",
		"2 Стрэнджер тхингс 3",
		"1 Хэванс дор 4",
		"12 Смелс лайк э тин спирит 6",
		"14 Число вселенной 42"}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	CheckFile(t, testFileName, expected)
}

func TestSort02(t *testing.T) {
	testFileName := "test02.txt"
	CreateTestFiles("sample.txt", testFileName)
	cmd := exec.Command(".\\task.exe", "-r", testFileName)
	expected := []string{
		"6 Сомбади вот ви юзед ту ноy 9",
		"5 Кукла колдуна 12",
		"2 Стрэнджер тхингс 3",
		"2 Вуду чайлд 5",
		"14 Число вселенной 42",
		"12 Смелс лайк э тин спирит 6",
		"1 Хэванс дор 4"}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	CheckFile(t, testFileName, expected)
}

func TestSort03(t *testing.T) {
	testFileName := "test03.txt"
	CreateTestFiles("sample.txt", testFileName)
	cmd := exec.Command(".\\task.exe", "-n", testFileName)
	expected := []string{
		"1 Хэванс дор 4",
		"2 Стрэнджер тхингс 3",
		"2 Вуду чайлд 5",
		"5 Кукла колдуна 12",
		"6 Сомбади вот ви юзед ту ноy 9",
		"12 Смелс лайк э тин спирит 6",
		"14 Число вселенной 42"}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	CheckFile(t, testFileName, expected)
}

func TestSort04(t *testing.T) {
	testFileName := "test04.txt"
	CreateTestFiles("sample.txt", testFileName)
	cmd := exec.Command(".\\task.exe", "-n", "-r", "-k", "2", testFileName)
	expected := []string{
		"14 Число вселенной 42",
		"5 Кукла колдуна 12",
		"6 Сомбади вот ви юзед ту ноy 9",
		"12 Смелс лайк э тин спирит 6",
		"2 Вуду чайлд 5",
		"1 Хэванс дор 4",
		"2 Стрэнджер тхингс 3"}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	CheckFile(t, testFileName, expected)
}

func CheckFile(t *testing.T, testName string, expected []string) {
	file, err := os.Open(testName)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	i := 1
	result := make([]string, 0)
	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
		i++
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	for i, _ := range result {
		if expected[i] != result[i] {
			t.Fatal("Incorrect result in row", i+1, ". Expected", expected[i], "; Result", result[i])
		}
	}
}

func CreateTestFiles(sampleName string, testName string) {
	_, err := os.Stat(sampleName)
	if err != nil {
		log.Fatalf("sampleFile not found: %s", err)
	}
	sampleFile, _ := os.Open(sampleName)
	testFile, err := os.Create(testName)
	if err != nil {
		log.Fatal(err)
	}
	defer testFile.Close()
	_, err = io.Copy(testFile, sampleFile)
	if err != nil {
		log.Fatal(err)
	}
}
