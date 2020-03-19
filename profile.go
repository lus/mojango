package mojango

// Profile represents a whole player profile
type Profile struct {
	UUID string `json:"id"`
	Name string `json:"name"`
	Properties []ProfileProperty `json:"properties"`
}

// ProfileProperty represents a property of a player profile
type ProfileProperty struct {
	Name string `json:"name"`
	Value string `json:"value"`
	Signature string `json:"signature"`
}
