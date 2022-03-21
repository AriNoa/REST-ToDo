package tag

import (
	"github.com/AriNoa/REST-ToDo/domain/model/ids"
)

type Tag struct {
	id   ids.TagID
	name Name
}

func (t Tag) Id() ids.TagID {
	return t.id
}

func (t *Tag) Name() *Name {
	return &t.name
}

func (t Tag) Clone() Tag {
	return Tag{
		id:   t.id,
		name: t.name,
	}
}

func New(
	id ids.TagID,
	name Name,
) (Tag, error) {
	return Tag{
		id:   id,
		name: name,
	}, nil
}
