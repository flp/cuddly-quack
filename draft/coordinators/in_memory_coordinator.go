package coordinators

import (
	"errors"
	"sync"

	"github.com/flp/cuddly-quack/draft"

	uuid "github.com/satori/go.uuid"
)

// InMemoryCoordinator is a very basic implementation
// of a draft room coordinator.  It will house draft rooms
// in RAM.  Note that if the node dies, so will all the current draft rooms.
// This is intended to be a preliminary implementation for testing purposes.
// In the future, we'll want more persistent draft rooms.
type InMemoryCoordinator struct {
	Lock       sync.Mutex
	DraftRooms map[string]*draft.Room
}

var DefaultInMemoryCoordinator *InMemoryCoordinator

func NewInMemoryCoordinator() *InMemoryCoordinator {
	return &InMemoryCoordinator{
		DraftRooms: make(map[string]*draft.Room),
	}
}

func (i *InMemoryCoordinator) GetDraftRoom(id string) (*draft.Room, error) {
	var room *draft.Room
	var ok bool

	i.Lock.Lock()
	room, ok = i.DraftRooms[id]
	i.Lock.Unlock()

	if !ok {
		return nil, errors.New("Couldn't find draft room")
	}

	return room, nil
}

func (i *InMemoryCoordinator) CreateDraftRoom(name string) (*draft.Room, error) {
	var room *draft.Room

	i.Lock.Lock()

	id := uuid.NewV4()
	room = &draft.Room{
		Name: name,
		UUID: id,
	}

	i.DraftRooms[id.String()] = room
	i.Lock.Unlock()

	return room, nil
}

func init() {
	DefaultInMemoryCoordinator = NewInMemoryCoordinator()
}
