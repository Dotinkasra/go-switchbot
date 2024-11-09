package module

import (
	"encoding/json"
)

// TrunOn: Turn on the infrared device.
func (d infraredRemoteDevice) TrunOn() string {
	c := NewSwitchbotAPIClient()
	uri := "https://api.switch-bot.com/v1.0/devices/" + d.DeviceId + "/commands"
	body := map[string]string{
		"commandType": "command",
		"command":     "turnOn",
		"parameter":   "default",
	}

	resp, error := c.SendAPIRequest(uri, "POST", body)
	if error != nil {
		panic(error)
	}

	var response apiResponse
	json.Unmarshal(resp, &response)

	return response.Message
}

// TrunOff: Turn off the infrared device.
func (d infraredRemoteDevice) TrunOff() string {
	c := NewSwitchbotAPIClient()
	uri := "https://api.switch-bot.com/v1.0/devices/" + d.DeviceId + "/commands"
	body := map[string]string{
		"commandType": "command",
		"command":     "turnOff",
		"parameter":   "default",
	}

	resp, error := c.SendAPIRequest(uri, "POST", body)
	if error != nil {
		panic(error)
	}

	var response apiResponse
	json.Unmarshal(resp, &response)

	return response.Message
}

// CustomCommand: Execute custom commands added by the app.
func (d infraredRemoteDevice) CustomCommand(command string) string {
	c := NewSwitchbotAPIClient()
	uri := "https://api.switch-bot.com/v1.0/devices/" + d.DeviceId + "/commands"
	body := map[string]string{
		"commandType": "customize",
		"command":     command,
		"parameter":   "default",
	}

	resp, error := c.SendAPIRequest(uri, "POST", body)
	if error != nil {
		panic(error)
	}

	var response apiResponse
	json.Unmarshal(resp, &response)

	return response.Message
}

// Execute: Execute the scene.
func (s scene) Execute() string {
	c := NewSwitchbotAPIClient()
	uri := "https://api.switch-bot.com/v1.0/scenes/" + s.SceneId + "/execute"
	resp, error := c.SendAPIRequest(uri, "POST", nil)
	if error != nil {
		panic(error)
	}
	var response apiResponse
	json.Unmarshal(resp, &response)

	return response.Message
}
