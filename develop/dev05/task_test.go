package main

import (
	"log"
	"os/exec"
	"testing"
)

func TestGrep01(t *testing.T) {
	testFileName := "sample.txt"
	expected := "Hotel California\nThe Road To Hell. Part 2\nSultans Of Swing\nMy Fathers Son\n"
	cmd := exec.Command(".\\task.exe", "a", testFileName)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	result := string(out)
	if expected != result {
		t.Fatal("Incorrect result: \n", expected, "\n --Expected\n", result)
	}
}

func TestGrep02(t *testing.T) {
	testFileName := "sample.txt"
	expected := "bring me to life\nsultans of swing\nin the army now\nthe temple of the king\nTotal rows count: 4\n"
	cmd := exec.Command(".\\task.exe", "-c", "-i", "in", testFileName)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	result := string(out)
	if expected != result {
		t.Fatal("Incorrect result: \n", expected, "\n --Expected\n", result)
	}

}

func TestGrep03(t *testing.T) {
	testFileName := "sample.txt"
	expected := "The Temple Of The King\nMy Fathers Son\nRiders on the Storm\n"
	cmd := exec.Command(".\\task.exe", "-C", "1", "-F", "Fathers", testFileName)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	result := string(out)
	if expected != result {
		t.Fatal("Incorrect result: \n", result, "\n --Expected\n", expected)
	}

}
