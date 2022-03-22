package ids

import (
	"github.com/google/uuid"
)

type uuidInterface interface {
	uuid() uuid.UUID
}

type TagID struct {
	id uuid.UUID
}

func (id TagID) uuid() uuid.UUID {
	return id.id
}

type TaskID struct {
	id uuid.UUID
}

func (id TaskID) uuid() uuid.UUID {
	return id.id
}
