package goweatherflow

import (
	"log"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	log.Println("Starting test goweatherflow server")
	server, err := NewServer()
	if err != nil {
		log.Fatalln(err)
	}

	var messageCount int = 0

	//Process channel data on a go routine.
	go func() {
		for {
			msg := <-server.ReceivedMessage
			messageCount += 1
			log.Println("recieved message new message!")
			switch v := msg.(type) {
			case *TempestObservation:
				log.Printf("Received observation from %s", v.SerialNumber)
				//process tempest observation
			case *RapidWindObservation:
				//process rapid wind
			case *RainStartEvent:
				//process Rain start event
			case *DeviceStatus:
				//process device status
			case *HubStatus:
				//process hub status
			case *LightningStrikeEvent:
				//process lightning strike
			default:
				log.Println("Unknown message type")
			}
		}
	}()

	go func() {
		time.Sleep(time.Second * 10)
		log.Println("Stopping server")
		server.Stop()
	}()

	server.Start() //Blocks until server.Stop() is called.

	if messageCount == 0 {
		t.Error("No messages received")
		t.Fail()
	}
}
