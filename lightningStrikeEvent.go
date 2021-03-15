package goweatherflow

import "time"

//LightningStrikeEvent is the message sent when rain has been detected.
type LightningStrikeEvent struct {
	SerialNumber    string    `json:"serial_number"`
	Type            string    `json:"type"`
	HubSerialNumber string    `json:"hub_sn"`
	EVT             []float64 `json:"evt"`
}

//Time returns the time of the lightning strike event
func (msg LightningStrikeEvent) Time() time.Time {
	if len(msg.EVT) == 0 {
		return time.Now()
	}

	return time.Unix(int64(msg.EVT[0]), 0)
}

//Distance returns the distance away the lightning strike was in km
func (msg LightningStrikeEvent) Distance() (int, error) {
	if len(msg.EVT) != 3 {
		return 0, WeatherFlowMessageLengthError
	}

	return int(msg.EVT[1]), nil
}

//Energy returns the energy value for the lightning strike
func (msg LightningStrikeEvent) Energy() (float64, error) {
	if len(msg.EVT) != 3 {
		return 0, WeatherFlowMessageLengthError
	}

	return msg.EVT[2], nil
}
