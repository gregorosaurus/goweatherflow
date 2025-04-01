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
				log.Printf("Received observation from %s. \n%+v", v.SerialNumber, v)
				//process tempest observation
				if temperature, err := v.AirTemperature(); err == nil {
					log.Printf("Current Temperature:%fºC", temperature)
				} else {
					log.Printf("Error occurred reading the temperature: %s", err)
					t.Fail()
				}
			case *RapidWindObservation:
				//process rapid wind
				log.Printf("Received rapid wind from %s", v.SerialNumber)
			case *RainStartEvent:
				//process Rain start event
				log.Printf("Received rain start from %s", v.SerialNumber)
			case *DeviceStatus:
				//process device status
				log.Printf("Received device status from %s", v.SerialNumber)
			case *HubStatus:
				//process hub status
				log.Printf("Received hub status from %s", v.SerialNumber)
			case *LightningStrikeEvent:
				//process lightning strike
				log.Printf("Received lightning strike from %s", v.SerialNumber)
			default:
				log.Println("Unknown message type")
			}
		}
	}()

	go func() {
		time.Sleep(time.Second * 60)
		log.Println("Stopping server")
		server.Stop()
	}()

	server.Start() //Blocks until server.Stop() is called.

	if messageCount == 0 {
		t.Error("No messages received")
		t.Fail()
	}
}
