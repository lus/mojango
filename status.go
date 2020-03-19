package mojango

import "encoding/json"

// These constants represent the possible states of the Mojang services
const (
	StatusGreen = "green"
	StatusYellow = "yellow"
	StatusRed = "red"
)

// Status contains all states of the Mojang services
type Status struct {
	MinecraftWebsite string
	MojangWebsite string
	Session string
	SessionServer string
	AuthServer string
	Account string
	Textures string
	API string
}

// parseStatusFromBody parses a status object from the response body of the API
func parseStatusFromBody(body []byte) (*Status, error) {
	// Parse multiple single states out of the response body
	var rawStates []struct{
		Key string
		Status string
	}
	err := json.Unmarshal(body, &rawStates); if err != nil {
		return nil, err
	}

	// Create the status object and put the corresponding values in it
	status := new(Status)
	for _, state := range rawStates {
		switch state.Key {
		case "minecraft.net":
			status.MinecraftWebsite = state.Status
			break
		case "mojang.com":
			status.MojangWebsite = state.Status
			break
		case "session.minecraft.net":
			status.Session = state.Status
			break
		case "sessionserver.mojang.com":
			status.SessionServer = state.Status
			break
		case "authserver.mojang.com":
			status.AuthServer = state.Status
			break
		case "account.mojang.com":
			status.Account = state.Status
			break
		case "textures.minecraft.net":
			status.Textures = state.Status
			break
		case "api.mojang.com":
			status.API = state.Status
			break
		}
	}

	// Return the status object
	return status, nil
}
