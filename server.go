package goweatherflow

import (
	"encoding/json"
	"fmt"
	"net"
	"regexp"
	"sync"
	"time"
)

type TempestServer struct {
	conn    *net.UDPConn
	running bool

	serverWaitGroup sync.WaitGroup

	ReceivedMessage chan WeatherFlowMessage
}

//NewServer creates a new server that will start listening
//for new Tempest UDP events
func NewServer() (*TempestServer, error) {
	serverAddr, err := net.ResolveUDPAddr("udp4", ":50222")
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp4", serverAddr)
	if err != nil {
		return nil, err
	}

	server := new(TempestServer)
	server.conn = conn

	server.ReceivedMessage = make(chan WeatherFlowMessage, 50)

	return server, nil
}

//Start starts the server by spinning up a goroutine.
//The function returns only after the server is stopped.
func (server *TempestServer) Start() {
	if server.running {
		return
	}

	server.running = true

	server.serverWaitGroup.Add(1)
	go server.startServer()
	server.serverWaitGroup.Wait()
}

func (server *TempestServer) startServer() {
	buf := make([]byte, 1024)
	for server.running {
		server.conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		n, _, _, addr, err := server.conn.ReadMsgUDP(buf, nil)
		if err != nil {
			//err was likely a timeout, we'll continue.
			continue
		}

		sliceCopy := make([]byte, n)
		copy(sliceCopy, buf[:n])
		go server.processData(addr, sliceCopy)
	}
	server.serverWaitGroup.Done()
}

func (server *TempestServer) processData(addr net.Addr, data []byte) error {
	stringData := string(data)
	// fmt.Println(stringData)
	typeRegex := regexp.MustCompile("\"type\":\\s*\"(.*?)\"")
	typeMatches := typeRegex.FindStringSubmatch(stringData)
	if len(typeMatches) != 2 {
		return fmt.Errorf("Unable to find type in message")
	}

	messageType := typeMatches[1]

	switch messageType {
	case "evt_precip":
		msg := new(RainStartEvent)
		err := json.Unmarshal(data, msg)
		if err != nil {
			return fmt.Errorf("Unable to deserialize json message message: %s", err)
		}
		server.ReceivedMessage <- msg
	case "evt_strike":
		msg := new(LightningStrikeEvent)
		err := json.Unmarshal(data, msg)
		if err != nil {
			return fmt.Errorf("Unable to deserialize json message message: %s", err)
		}
		server.ReceivedMessage <- msg
	case "rapid_wind":
		msg := new(RapidWindObservation)
		err := json.Unmarshal(data, msg)
		if err != nil {
			return fmt.Errorf("Unable to deserialize json message message: %s", err)
		}
		server.ReceivedMessage <- msg
	case "obs_st":
		msg := new(TempestObservation)
		err := json.Unmarshal(data, msg)
		if err != nil {
			return fmt.Errorf("Unable to deserialize json message message: %s", err)
		}
		server.ReceivedMessage <- msg
	case "device_status":
		msg := new(DeviceStatus)
		err := json.Unmarshal(data, msg)
		if err != nil {
			return fmt.Errorf("Unable to deserialize json message message: %s", err)
		}
		server.ReceivedMessage <- msg
	case "hub_status":
		msg := new(HubStatus)
		err := json.Unmarshal(data, msg)
		if err != nil {
			return fmt.Errorf("Unable to deserialize json message message: %s", err)
		}
		server.ReceivedMessage <- msg
	default:
		return fmt.Errorf("Unknown message received: %s", messageType)
	}

	return nil
}

//Stop initiates the server stop.  This will wait until the server shuts down and
//stops listening for messages.
func (server *TempestServer) Stop() {
	server.running = false
	server.serverWaitGroup.Wait()
}

//Close closes the server.  Once close is called, the server can not
//be reopened.
func (server *TempestServer) Close() error {
	return server.conn.Close()
}
