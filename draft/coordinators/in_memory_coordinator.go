package coordinators

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"sync"

	"github.com/flp/cuddly-quack/draft"
	"github.com/flp/cuddly-quack/wire"

	uuid "github.com/satori/go.uuid"
)

// InMemoryCoordinator is a very basic implementation
// of a draft room coordinator.  It will house draft rooms
// in RAM.  Note that if the node dies, so will all the current draft rooms.
// This is intended to be a preliminary implementation for testing purposes.
// In the future, we'll want more persistent draft rooms.
type InMemoryCoordinator struct {
	Lock         sync.Mutex
	DraftRooms   map[string]*draft.Room
	UserRegistry map[string]bool
}

var DefaultInMemoryCoordinator *InMemoryCoordinator

func NewInMemoryCoordinator() *InMemoryCoordinator {
	return &InMemoryCoordinator{
		DraftRooms:   make(map[string]*draft.Room),
		UserRegistry: make(map[string]bool),
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

func (i *InMemoryCoordinator) CreateDraftRoom(req *wire.CreateDraftRequest) (*draft.Room, error) {
	var room *draft.Room

	i.Lock.Lock()

	id := uuid.NewV4()
	room = &draft.Room{
		Name: req.Name,
		UUID: id,
	}

	i.DraftRooms[id.String()] = room
	i.Lock.Unlock()

	return room, nil
}

func (i *InMemoryCoordinator) RegisterUser(userID string, roomID string) error {
	hash := sha512.Sum512([]byte(userID + roomID))
	key := string(hash[:len(hash)-1])
	var err error

	i.Lock.Lock()

	registered, ok := i.UserRegistry[key]
	if !registered || !ok {
		i.UserRegistry[key] = true
		err = nil
	} else {
		err = errors.New(fmt.Sprintf("User %s Already Registered for draft room: %s", userID, roomID))
	}

	i.Lock.Unlock()

	return err
}

func init() {
	DefaultInMemoryCoordinator = NewInMemoryCoordinator()
}
