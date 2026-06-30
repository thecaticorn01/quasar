package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"github.com/psanford/wormhole-william/wormhole"
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

	ctx := context.Background()

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

func send(text, filePath string) {
	var client wormhole.Client

	if text != "" {
		fmt.Println("Initializing encrypted transfer channel...")
		code, status, err := client.SendText(ctx, text)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Initialization failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\nChannel established! Secure connection code:\n%s\n\n", code)
		fmt.Println("Awaiting peer connection...")

		s := <-status
		if s.OK {
			fmt.Println("Payload successfully delivered.")
		} else {
			fmt.Fprintf(os.Stderr, "Transmission block failed: %v\n", s.Error)
		}
		return
	}

	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read target file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		fileInfo, err = file.Stat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to inspect file stat: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Staging file for transfer: %s (%d bytes)\n", fileInfo.Name(), fileInfo.Size())
		code, status, err := client.SendFile(ctx, fileInfo.Name(), file, wormhole.WithProgress(func(sent, total int64) {
			fmt.Printf("\rTransmitting: %d/%d bytes", sent, total)
		}))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Channel initialization failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("\nChannel established! Secure connection code:\n%s\n\n", code)
		fmt.Println("Awaiting peer connection...")

		s := <-status
		if s.OK {
			fmt.Println("\nPayload successfully delivered.")
		} else {
			fmt.Fprintf(os.Stderr, "Transmission block failed: %v\n", s.Error)
		}
		return

		fmt.Println("You must specify either --text or --file to transmit.")
		
	}

}
func recv(code string) {
	// Implement receiving logic here
}
