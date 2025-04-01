package goweatherflow

import (
	"time"
)

type TempestObservation struct {
	SerialNumber     string      `json:"serial_number"`
	Type             string      `json:"type"`
	HubSerialNumber  string      `json:"hub_sn"`
	ObservationData  [][]float64 `json:"obs"`
	FirmwareRevision int         `json:"firmware_revision"`
}

const TemperatureObservationMessageParameterCount = 18

// Time returns the time of the observation
func (msg TempestObservation) Time() time.Time {
	if len(msg.ObservationData) == 0 {
		return time.Now().UTC()
	}

	return time.Unix(int64(msg.ObservationData[0][0]), 0).UTC()
}

// WindLull returns the wind lull for a mimum 3 second sample in metres per second.
func (msg TempestObservation) WindLull() (float64, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return msg.ObservationData[0][1], nil
}

// WindAverage returns the average wind over report interval in metres per second
func (msg TempestObservation) WindAverage() (float64, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return msg.ObservationData[0][2], nil
}

// WindGust returns the maximum wind over a 3 second sample.
func (msg TempestObservation) WindGust() (float64, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return msg.ObservationData[0][3], nil
}

// WindDirection returns the origin of the wind in degrees true.
func (msg TempestObservation) WindDirection() (float64, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return msg.ObservationData[0][4], nil
}

// WindSampleInterval returns the seconds that was used to sample the wind.
func (msg TempestObservation) WindSampleInterval() (int, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return int(msg.ObservationData[0][5]), nil
}

// StationPressure returns the pressure of the station in millibars
func (msg TempestObservation) StationPressure() (float64, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return msg.ObservationData[0][6], nil
}

// AirTemperature returns the tempearture of the outside air in degrees c
func (msg TempestObservation) AirTemperature() (float64, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return msg.ObservationData[0][7], nil
}

// RelativeHumidity returns the relative humidity in percentage
func (msg TempestObservation) RelativeHumidity() (float64, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return msg.ObservationData[0][8], nil
}

// Illuminance returns the Illuminance in lux
func (msg TempestObservation) Illuminance() (float64, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return msg.ObservationData[0][9], nil
}

// UV returns the UV index
func (msg TempestObservation) UV() (int, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return int(msg.ObservationData[0][10]), nil
}

// SolarRadiation returns the amount of solar radiation in watts per square metre
func (msg TempestObservation) SolarRadiation() (float64, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return msg.ObservationData[0][11], nil
}

// PrecipitationAccumulated returns the amount of precipitation in mm in the reporting interval
func (msg TempestObservation) PrecipitationAccumulated() (float64, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return msg.ObservationData[0][12], nil
}

// PrecipitationType returns the type of preciptation observed in the reporting interval
func (msg TempestObservation) PrecipitationType() (WeatherFlowPrecipType, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return WeatherFlowPrecipType(msg.ObservationData[0][13]), nil
}

// LightningStrikeAverageDistance returns the average distance away lightning strikes are in km in the
// reporting interval
func (msg TempestObservation) LightningStrikeAverageDistance() (float64, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return msg.ObservationData[0][14], nil
}

// LightningStrikeCount returns the number of lightning strikes in the reporting interval
func (msg TempestObservation) LightningStrikeCount() (int, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return int(msg.ObservationData[0][15]), nil
}

// BatteryVoltage returns the current battery voltage.
func (msg TempestObservation) BatteryVoltage() (float64, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	return msg.ObservationData[0][16], nil
}

// ReportInterval is the interval in minutes of this observation
func (msg TempestObservation) ReportInterval() (time.Duration, error) {
	if len(msg.ObservationData) >= 1 && len(msg.ObservationData[0]) != TemperatureObservationMessageParameterCount {
		return 0, weatherFlowMessageLengthError
	}

	//we convert this to seconds and then a duration since a duration is an integer,
	//we dont want to lose any precision in case weather flow sends a decimal minute
	//over.
	return time.Second * time.Duration(60*msg.ObservationData[0][17]), nil
}
