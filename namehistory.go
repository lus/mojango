package mojango

// NameHistoryEntry represents an entry of the name history of an account
type NameHistoryEntry struct {
	Name string `json:"name"`
	ChangedToAt int64 `json:"changedToAt"`
}
