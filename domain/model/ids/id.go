package ids

import "github.com/google/uuid"

type idConstraints interface {
	~struct {
		id uuid.UUID
	}
}

func Parse[T idConstraints](raw string) (T, error) {
	id, err := uuid.Parse(raw)

	return T{id: id}, err
}

func New[T idConstraints]() T {
	return T{
		id: uuid.New(),
	}
}

func String(id uuidInterface) string {
	return id.uuid().String()
}
