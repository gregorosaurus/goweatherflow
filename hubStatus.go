package goweatherflow

import "time"

type HubStatus struct {
	SerialNumber     string  `json:"serial_number"`
	Type             string  `json:"type"`
	TimeStamp        int64   `json:"timestamp"`
	UptimeSeconds    uint64  `json:"uptime"`
	FirmwareRevision int     `json:"firmware_revision"`
	RSSI             float64 `json:"rssi"`
	ResetFlags       string  `json:"reset_flags"`
	Sequence         int     `json:"seq"`
	FS               []int   `json:"fs"`
	RadioStats       []int   `json:"radio_stats"`
	MQTTStats        []int   `json:"mqtt_stats"`
}

//Time returns the time of the observation
func (msg HubStatus) Time() time.Time {
	return time.Unix(msg.TimeStamp, 0).UTC()
}
