package mojango

import "github.com/valyala/fasthttp"

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

func (client *Client) FetchUUID(username string) {
	// TODO: Add UUID fetching
}

func (client *Client) FetchUUIDAtTime(username string, timestamp int64) {
	// TODO: Add UUID at timestamp fetching
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
