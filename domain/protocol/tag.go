package protocol

type TagID string

type Tag struct {
	name string
}

func (t Tag) Tag() string {
	return t.name
}

func NewTag(name string) Tag {
	return Tag{
		name: name,
	}
}
