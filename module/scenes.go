package module

import (
	"encoding/json"
	"fmt"
)

type scene struct {
	SceneId   string `json:"sceneId"`
	SceneName string `json:"sceneName"`
}

// Includes a list of all the scenes you have registered.
type Scenes struct {
	SceneList []scene `json:"body"`
}

// Converts the API response body to a list of scenes.
func convertToScenes(data interface{}) (Scenes, error) {
	dataSlice, ok := data.([]interface{})
	if !ok {
		return Scenes{}, fmt.Errorf("data is not a slice")
	}

	var scenes []scene

	for _, s := range dataSlice {
		jsonData, error := json.Marshal(s)
		if error != nil {
			return Scenes{}, error
		}

		var scene scene
		error = json.Unmarshal(jsonData, &scene)
		if error != nil {
			return Scenes{}, error
		}
		scenes = append(scenes, scene)
	}

	return Scenes{SceneList: scenes}, nil
}

// Retrieve all scenes.
func (scenes Scenes) GetScenes() (Scenes, error) {
	c := NewSwitchbotAPIClient()
	resp, err := c.SendAPIRequest("https://api.switch-bot.com/v1.0/scenes", "GET", nil)

	if err != nil {
		return Scenes{}, err
	}

	var r apiResponse
	json.Unmarshal(resp, &r)

	if r.StatusCode == 190 {
		return Scenes{}, fmt.Errorf("system error: Device internal error due to device states not synchronized with server")
	}

	scenes, error := convertToScenes(r.Body)
	if error != nil {
		panic(error)
	}
	return scenes, nil

}
