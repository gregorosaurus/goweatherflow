package goweatherflow

import "time"

type RapidWindObservation struct {
	SerialNumber    string    `json:"serial_number"`
	Type            string    `json:"type"`
	HubSerialNumber string    `json:"hub_sn"`
	ObservationData []float64 `json:"ob"`
}

//Time returns the time of the observation
func (msg RapidWindObservation) Time() time.Time {
	if len(msg.ObservationData) == 0 {
		return time.Now()
	}

	return time.Unix(int64(msg.ObservationData[0]), 0).UTC()
}

//WindSpeed returns the wind speed in metres per second
func (msg RapidWindObservation) WindSpeed() (float64, error) {
	if len(msg.ObservationData) != 3 {
		return 0, WeatherFlowMessageLengthError
	}
	return msg.ObservationData[1], nil
}

//WindDirection is the wind direction in degrees true.  The wind direction
//is the degrees that wind is originating from.  IE 180º is wind blowing
//from the south.
func (msg RapidWindObservation) WindDirection() (float64, error) {
	if len(msg.ObservationData) != 3 {
		return 0, WeatherFlowMessageLengthError
	}
	return msg.ObservationData[2], nil
}
