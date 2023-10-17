package main

import (
	"fmt"
	"testing"
)

func Test_updateMessage(t *testing.T){

	msg = "Hello Wrold"

	wg.Add(2)

	go updateMeassage("abc")
	go updateMeassage("hi, there")

	wg.Wait()

	if msg != "hi, there"{
		fmt.Println(msg)
		t.Error("incorrect value in msg")
	}

	// run the command go test -race .
}