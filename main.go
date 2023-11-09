package main

import (
	"fmt"
	"time"
)

type Message struct {
	From     string
	Playload string
}

type Server struct {
	msgChan  chan Message
	quitchan chan struct{}
}

func (s *Server) StartAndListen() {

	for {
		select {
		case msg := <-s.msgChan:
			fmt.Printf("received message from: %s payload: %s\n", msg.From, msg.Playload)
		case <-s.quitchan:
			fmt.Println("Server is gracefully shut down")
			return
		}
	}
}

func SendMessageToServer(msgCh chan Message, payload string) {
	fmt.Println("Sending the message")
	msg := Message{
		From:     "Keshav",
		Playload: payload,
	}
	msgCh <- msg
}

func main() {
	s := &Server{
		msgChan:  make(chan Message),
		quitchan: make(chan struct{}),
	}
	go s.StartAndListen()

	for i := 1; i < 10; i++ {
		go func() {
			time.Sleep(10 * time.Microsecond)
			SendMessageToServer(s.msgChan, "Hello server")
		}()
	}

	time.Sleep(1 * time.Second) // Wait for a while
	// select {}

	close(s.quitchan)                  // Gracefully shut down the server by closing the quitchan
	time.Sleep(100 * time.Millisecond) // Give the server some time to process the quit signal
}
