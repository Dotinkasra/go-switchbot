
# Go SwitchBot API Client

A Go package for interacting with the SwitchBot API. This package provides functionality to list devices, execute commands, and run scenes through the SwitchBot cloud API.

## Features
- Retrieve a list of registered devices and infrared remote devices.
- Turn devices on/off and execute custom commands.
- Run predefined scenes.

## Installation
To install this package, use `go get`:

```bash
go get github.com/Dotinkasra/go-switchbot
```

## Setup
Ensure that the `SWITCHBOT_TOKEN` environment variable is set. This token is required for authorization when making API requests.

```bash
export SWITCHBOT_TOKEN=your_switchbot_token
```

## Usage
### Import the Package
```go
import (
    "fmt"
    "log"

    "github.com/Dotinkasra/go-switchbot/module"
)
```

### Example Code
Here's a simple example of how to use this package to get devices and scenes.

#### `main.go`
```go
package main

import (
    "fmt"
    "log"

    "github.com/Dotinkasra/go-switchbot/module"
)

func main() {
    // Get list of devices
    d := new(module.Devices)
    devices, err := d.GetDevices()
    if err != nil {
        log.Fatalf("Error getting devices: %v", err)
    }
    fmt.Printf("Devices: %+v\n", devices)

    // Get list of scenes and execute one
    s := new(module.Scenes)
    scenes, err := s.GetScenes()
    if err != nil {
        log.Fatalf("Error getting scenes: %v", err)
    }
    fmt.Printf("Scenes: %+v\n", scenes)

    // Execute the first scene (example)
    if len(scenes.SceneList) > 0 {
        result := scenes.SceneList[0].Execute()
        fmt.Printf("Executed scene result: %s\n", result)
    }
}
```

## API Overview
### `Devices` Struct
- **`GetDevices() (Devices, error)`**: Retrieves all registered devices.

### `infraredRemoteDevice` Methods
- **`TrunOn() string`**: Turns on the infrared device.
- **`TrunOff() string`**: Turns off the infrared device.
- **`CustomCommand(command string) string`**: Sends a custom command to the infrared device.

### `Scenes` Struct
- **`GetScenes() ([]Scene, error)`**: Retrieves all available scenes.
- **`Execute()`**: Executes the scene.

## License
MIT