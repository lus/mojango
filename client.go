package mojango

import "github.com/valyala/fasthttp"

type Client struct {
	// TODO: Add client fields
}

func (client *Client) FetchStatus() {
	// TODO: Add server status fetching and parsing
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
