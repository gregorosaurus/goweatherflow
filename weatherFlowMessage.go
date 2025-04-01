package goweatherflow

import (
	"fmt"
	"time"
)

var weatherFlowMessageLengthError error = fmt.Errorf("invalid weather flow data length")

type WeatherFlowPrecipType int

const WeatherFlowPrecipTypeNone WeatherFlowPrecipType = 0
const WeatherFlowPrecipTypeRain WeatherFlowPrecipType = 1
const WeatherFlowPrecipTypeHail WeatherFlowPrecipType = 2

type WeatherFlowMessage interface {
	Time() time.Time
}
