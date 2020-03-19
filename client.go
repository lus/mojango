package mojango

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
)

// Mojango represents an API client
type Mojango interface {
	FetchStatus() (*Status, error)
	FetchUUID(username string) (string, error)
	FetchUUIDAtTime(username string, timestamp int64) (string, error)
	FetchMultipleUUIDs(usernames []string) (map[string]string, error)
	FetchNameHistory(uuid string) ([]NameHistoryEntry, error)
	FetchProfile(uuid string, unsigned bool) (*Profile, error)
}

// client represents a internal wrapper around the fasthttp client
type client struct {
	*fasthttp.Client
}

// New creates a new fasthttp client and wraps it into an API client
func New() Mojango {
	return &client{
		&fasthttp.Client{
			Name: "mojango",
		},
	}
}

// FetchStatus fetches the states of all Mojang services and wraps them into a single object
func (client *client) FetchStatus() (*Status, error) {
	// Call the Mojang status endpoint
	code, body, err := client.Get(nil, "https://status.mojang.com/check"); if err != nil {
		return nil, err
	}

	// Handle possible errors
	if code != fasthttp.StatusOK {
		return nil, errorFromCode(code)
	}

	// Parse the result into a status object and return it
	return parseStatusFromBody(body)
}

// FetchUUID fetches the current UUID of the given username
func (client *client) FetchUUID(username string) (string, error) {
	return client.FetchUUIDAtTime(username, -1)
}

// FetchUUIDAtTime fetches the UUID of the given username at a given timestamp
func (client *client) FetchUUIDAtTime(username string, timestamp int64) (string, error) {
	// Call the Mojang profile endpoint
	atExtension := ""
	if timestamp >= 0 {
		atExtension = "?at=" + strconv.FormatInt(timestamp, 10)
	}
	code, body, err := client.Get(nil, "https://api.mojang.com/users/profiles/minecraft/" + username + atExtension); if err != nil {
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

// FetchMultipleUUIDs fetches the UUIDs of the given usernames
func (client *client) FetchMultipleUUIDs(usernames []string) (map[string]string, error) {
	// Define the request object
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)
	request.SetRequestURI("https://api.mojang.com/profiles/minecraft")
	reqBody, err := json.Marshal(usernames); if err != nil {
		return nil, err
	}
	request.SetBody(reqBody)

	// Define the response object
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	// Call the Mojang profile endpoint
	err = client.Do(request, response); if err != nil {
		return nil, err
	}

	// Define the important response values
	code := response.StatusCode()
	body := response.Body()

	// Handle possible errors
	if code != fasthttp.StatusOK {
		return nil, errorFromCode(code)
	}

	// Parse the response body into a list of results
	var rawResults []struct{
		UUID string `json:"id"`
		Name string `json:"name"`
	}
	err = json.Unmarshal(body, &rawResults); if err != nil {
		return nil, err
	}

	// Parse the list of results into a map and return it
	result := make(map[string]string)
	for _, rawResult := range rawResults {
		result[rawResult.Name] = rawResult.UUID
	}
	return result, nil
}

// FetchNameHistory fetches all names of the given UUID and their corresponding changing timestamps
func (client *client) FetchNameHistory(uuid string) ([]NameHistoryEntry, error) {
	// Call the Mojang profile endpoint
	code, body, err := client.Get(nil, "https://api.mojang.com/user/profiles/" + uuid + "/names"); if err != nil {
		return nil, err
	}

	// Handle possible errors
	if code != fasthttp.StatusOK {
		return nil, errorFromCode(code)
	}

	// Parse the response body into a list of name history entries and return it
	var entries []NameHistoryEntry
	err = json.Unmarshal(body, &entries); if err != nil {
		return nil, err
	}
	return entries, nil
}

// FetchProfile fetches the profile of the given UUID
func (client *client) FetchProfile(uuid string, unsigned bool) (*Profile, error) {
	// Call the Mojang profile endpoint
	code, body, err := client.Get(nil, "https://sessionserver.mojang.com/session/minecraft/profile/" + uuid + "?unsigned=" + strconv.FormatBool(unsigned)); if err != nil {
		return nil, err
	}

	// Handle possible errors
	if code != fasthttp.StatusOK {
		return nil, errorFromCode(code)
	}

	// Parse the response body into a profile and return it
	profile := new(Profile)
	err = json.Unmarshal(body, profile); if err != nil {
		return nil, err
	}
	return profile, nil
}
