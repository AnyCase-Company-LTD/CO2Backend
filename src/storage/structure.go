package storage

import (
	"time"
)

type Event struct {
	Date        time.Time `json:"date"`
	DeviceId    string    `json:"deviceId"`
	Co2         int       `json:"cO2"`
	Temperature float32   `json:"temperature"`
	Humidity    float32   `json:"humidity"`
}

type Sensor struct {
	Name     string `json:"name"`
	DeviceId string `json:"device_id"`
}

type SensorList struct {
	Data []Sensor `json:"data"`
}
