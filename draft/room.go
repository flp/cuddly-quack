package draft

import (
	uuid "github.com/satori/go.uuid"
)

type Room struct {
	// TODO Implement
	Name string
	UUID uuid.UUID
	// - authentication?
	//   - basic: password
	//   - invite only
	//   - etc
	// - list of authenticated players
	// - state & state machine:
	//   - open
	//   - started
	//   - playing
	//   - closed
}
