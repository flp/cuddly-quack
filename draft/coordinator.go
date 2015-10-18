package draft

// Coordinator is an interface which
// keeps track of active draft rooms and who is authenticated to join them.
type Coordinator interface{} {
	DraftRoom(string id, string userID) (*Room, error)
}
