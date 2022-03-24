package tag

import (
	"github.com/AriNoa/REST-ToDo/domain/model/ids"
	model "github.com/AriNoa/REST-ToDo/domain/model/tag"
	"github.com/AriNoa/REST-ToDo/domain/protocol"
)

type Repository struct {
	tags map[ids.TagID]model.Tag
}

func (r *Repository) Create(name model.Name) (model.Tag, error) {
	id := ids.New[ids.TagID]()
	tag, err := model.New(id, name)

	if err != nil {
		return model.Tag{}, err
	}

	r.tags[id] = tag

	return tag.Clone(), nil
}

func (r *Repository) FindByID(id ids.TagID) (model.Tag, bool) {
	tag, ok := r.tags[id]

	return tag.Clone(), ok
}

func (r *Repository) FindByName(name model.Name) (model.Tag, bool) {
	for _, tag := range r.tags {
		if tag.Name().Equal(name) {
			return tag.Clone(), true
		}
	}

	return model.Tag{}, false
}

func (r *Repository) Update(tag model.Tag) error {
	r.tags[tag.Id()] = tag.Clone()

	return nil
}

func (r *Repository) Delete(id ids.TagID) error {
	if _, ok := r.tags[id]; !ok {
		return protocol.ErrTagNotFound
	}

	delete(r.tags, id)

	return nil
}
