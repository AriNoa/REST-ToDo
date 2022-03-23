package tag

import "github.com/AriNoa/REST-ToDo/domain/protocol"

type Name struct {
	raw string
}

func (n Name) String() string {
	return n.raw
}

func (n Name) Equal(other Name) bool {
	return n.raw == other.raw
}

func NewName(name string) (Name, error) {
	if len(name) < 1 {
		return Name{}, protocol.ErrNoTagName
	}

	return Name{
		raw: name,
	}, nil
}
