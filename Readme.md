# goweatherflow 

goweatherflow is a library used to receive UDP messages from WeatherFlow devices such as the WeatherFlow Tempest weather station. 

## Installation
```
go get github.com/gregorosaurus/goweatherflow
```

## Supported messages
 - [x] Rain Start Event (evt_precip)
 - [x] Lightning Strike Event (evt_strike)
 - [x] Rapid Wind (rapid_wind)
 - [ ] AIR Observation (obs_air)
 - [ ] Sky Observation (obs_sky)
 - [x] Tempest Observation (obs_st)
 - [x] Device Status (device_status)
 - [x] Hub Status (hub_status) 
 - [ ] Radio stats

## Basic Usage

To quickly get started, see the example code below:
```
    server, err := goweatherflow.NewServer()
	if err != nil {
		log.Fatalln(err)
	}

    //Process channel data on a go routine. 
    go func() {
		for {
			msg := <-server.ReceivedMessage
			log.Println("recieved message new message!")
			switch v := msg.(type) {
			case *goweatherflow.TempestObservation:
				log.Printf("Received observation from %s", v.SerialNumber)
				//process tempest observation
			case *goweatherflow.RapidWindObservation:
				//process rapid wind
			case *goweatherflow.RainStartEvent:
				//process Rain start event
			case *goweatherflow.DeviceStatus:
				//process device status
			case *goweatherflow.HubStatus:
				//process hub status
			case *goweatherflow.LightningStrikeEvent:
				//process lightning strike
			default:
				log.Println("Unknown message type")
			}
		}
	}()

	server.Start()  //Blocks until server.Stop() is called. 
```

## Architecture

The library continually receives UDP messages on port 50222, and places them on a channel to be received by the caller. 
It is the responisbility of the caller to receive these messages.  

