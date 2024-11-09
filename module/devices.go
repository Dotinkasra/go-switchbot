package module

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type device struct {
	DeviceId           string `json:"deviceId"`
	DeviceName         string `json:"deviceName"`
	DeviceType         string `json:"deviceType"`
	EnableCloudService bool   `json:"enableCloudService"`
	HubDeviceId        string `json:"hubDeviceId"`
}

type infraredRemoteDevice struct {
	DeviceId    string `json:"deviceId"`
	DeviceName  string `json:"deviceName"`
	RemoteType  string `json:"remoteType"`
	HubDeviceId string `json:"hubDeviceId"`
}

// Devices: Includes a list of physical devices and a list of infrared devices.
type Devices struct {
	DeviceList         []device               `json:"deviceList"`         //List of physical devices.
	InfraredRemoteList []infraredRemoteDevice `json:"infraredRemoteList"` //List of infrared devices.
}

func convertToDevices(data interface{}) (Devices, error) {
	var devices Devices
	err := mapstructure.Decode(data, &devices)
	if err != nil {
		return Devices{}, fmt.Errorf("failed to convert")
	}
	return devices, nil
}

// GetDevices: Retrieve all devices.
func (devices Devices) GetDevices() (Devices, error) {
	c := NewSwitchbotAPIClient()
	resp, err := c.SendAPIRequest("https://api.switch-bot.com/v1.0/devices", "GET", nil)

	if err != nil {
		return Devices{}, err
	}

	var r apiResponse
	json.Unmarshal(resp, &r)

	if r.StatusCode == 190 {
		return Devices{}, fmt.Errorf("system error: Device internal error due to device states not synchronized with server")
	}

	devices, error := convertToDevices(r.Body)
	if error != nil {
		panic(error)
	}
	return devices, nil
}
