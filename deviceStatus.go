package goweatherflow

import "time"

type DeviceStatus struct {
	SerialNumber     string  `json:"serial_number"`
	Type             string  `json:"type"`
	HubSerialNumber  string  `json:"hub_sn"`
	TimeStamp        int64   `json:"timestamp"`
	UptimeSeconds    int64   `json:"uptime"`
	Voltage          float64 `json:"voltage"`
	FirmwareRevision int     `json:"firmware_revision"`
	RSSI             float64 `json:"rssi"`
	HubRSSI          float64 `json:"hub_rssi"`
	SensorStatus     uint32  `json:"sensor_status"`
	Debug            int     `json:"debug"`
}

//Time returns the time of the observation
func (msg DeviceStatus) Time() time.Time {
	return time.Unix(msg.TimeStamp, 0).UTC()
}
