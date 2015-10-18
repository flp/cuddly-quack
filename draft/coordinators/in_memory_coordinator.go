package coordinators

import (
	"errors"
	
	"github.com/flp/cuddly-quack/draft"
)

// InMemoryCoordinator is a very basic implementation
// of a draft room coordinator.  It will house draft rooms
// in RAM.  Note that if the node dies, so will all the current draft rooms.
// This is intended to be a preliminary implementation for testing purposes.
// In the future, we'll want more persistent draft rooms.
type InMemoryCoordinator struct {
	DraftRooms map[string]*draft.Room
}

func NewInMemoryCoordinator() (*InMemoryCoordinator) {
	return &InMemoryCoordinator{
		DraftRooms: make(map[string]*draft.Room),
	}
}

func(i *InMemoryCoordinator) DraftRoom(id string, userID string) (*draft.Room, error) {
	// TODO Implement
	room, ok := i.DraftRooms(id)
	if !ok {
		return nil, errors.New("Couldn't find draft room")
	}
	
	return room, nil
}

func (i *InMemoryCoordinator) CreateDraftRoom(id string) (error) {
	i[id] = &draft.Room{}
}
