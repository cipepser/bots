package main

import "./line"

func main() {
	msg := "hello"

	err := line.SendMessage(msg)
	if err != nil {
		panic(err)
	}

}
