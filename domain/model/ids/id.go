package ids

import "github.com/google/uuid"

type ID interface {
	TagID | TaskID
}

func Parse[T ID](raw string) (T, error) {
	id, err := uuid.Parse(raw)

	return T{id: id}, err
}

func New[T ID]() T {
	return T{
		id: uuid.New(),
	}
}
