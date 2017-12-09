package main

import "./line"

func main() {
	// accessToken := "Y5GIIbj2672wnp7nR0JtSXJm5RwphO9WsVtyzDNitSI"
	msg := "hello"

	err := line.SendMessage(msg)
	if err != nil {
		panic(err)
	}

}
