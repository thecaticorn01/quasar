package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	recvCmd := flag.NewFlagSet("recv", flag.ExitOnError)

	sendTextFlag := sendCmd.String("text", "", "Text to send")
	sendFileFlag := sendCmd.String("file", "", "File to send")

	if len(os.Args) < 2 {
		fmt.Println("Usage: quasar [send|recv] [options]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "send":
		sendCmd.Parse(os.Args[2:])
		if *sendTextFlag == "" && *sendFileFlag == "" {
			fmt.Println("Please provide either -text or -file option")
			os.Exit(1)
		}
		send(*sendTextFlag, *sendFileFlag)
	case "recv":
		recvCmd.Parse(os.Args[2:])
		if len(recvCmd.Args()) < 1 {
			fmt.Println("Please provide a share code")
			os.Exit(1)
		}
		recv(recvCmd.Arg(0))
	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println("Usage: quasar [send|recv] [options]")
		os.Exit(1)
	}
}

func send(text, file string) {
	// Implement sending logic here
}
func recv(code string) {
	// Implement receiving logic here
}
