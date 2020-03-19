package mojango

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
)

// Represents an API client
type Client struct {
	client *fasthttp.Client
}

// Creates a new fasthttp client and wraps it into an API client
func New() *Client {
	return &Client{
		client: &fasthttp.Client{
			Name: "mojango",
		},
	}
}

// Fetches the states of all Mojang services and wraps them into a single object
func (client *Client) FetchStatus() (*Status, error) {
	// Call the Mojang status endpoint
	code, body, err := client.client.Get(nil, "https://status.mojang.com/check"); if err != nil {
		return nil, err
	}

	// Handle possible errors
	if code != fasthttp.StatusOK {
		return nil, errorFromCode(code)
	}

	// Parse the result into a status object
	return parseStatusFromBody(body)
}

// Fetches the current UUID of the given username
func (client *Client) FetchUUID(username string) (string, error) {
	return client.FetchUUIDAtTime(username, -1)
}

// Fetches the UUID of the given username at a given timestamp
func (client *Client) FetchUUIDAtTime(username string, timestamp int64) (string, error) {
	// Call the Mojang profile endpoint
	atExtension := ""
	if timestamp >= 0 {
		atExtension = "?at=" + strconv.FormatInt(timestamp, 10)
	}
	code, body, err := client.client.Get(nil, "https://api.mojang.com/users/profiles/minecraft/" + username + atExtension); if err != nil {
		return "", err
	}

	// Handle possible errors
	if code != fasthttp.StatusOK {
		return "", errorFromCode(code)
	}

	// Parse the result into a map containing the profile data
	var result map[string]interface{}
	err = json.Unmarshal(body, &result); if err != nil {
		return "", err
	}

	// Return the UUID of the requested profile
	return result["id"].(string), nil
}

func (client *Client) FetchMultipleUUIDs(usernames []string) {
	// TODO: Add multiple UUID fetching
}

func (client *Client) FetchNameHistory(uuid string) {
	// TODO: Add name history fetching
}

func (client *Client) FetchProfile(uuid string) {
	// TODO: Add profile fetching
}
