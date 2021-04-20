package goweatherflow

import "time"

//RainStartEvent is the message sent when rain has been detected.
type RainStartEvent struct {
	SerialNumber    string    `json:"serial_number"`
	Type            string    `json:"type"`
	HubSerialNumber string    `json:"hub_sn"`
	EVT             []float64 `json:"evt"`
}

//Time returns the time of the rain start message
func (msg RainStartEvent) Time() time.Time {
	if len(msg.EVT) == 0 {
		return time.Now().UTC()
	}

	return time.Unix(int64(msg.EVT[0]), 0).UTC()
}
