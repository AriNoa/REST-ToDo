package protocol

type Tag struct {
	id   string
	name string
}

func (t Tag) Id() string {
	return t.name
}

func (t Tag) Tag() string {
	return t.name
}

func NewTag(id string, name string) Tag {
	return Tag{
		id:   id,
		name: name,
	}
}
