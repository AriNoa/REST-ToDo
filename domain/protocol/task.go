package protocol

type TaskID string

type Task struct {
	title       string
	description string
	tags        []TaskID
	completed   bool
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

func (t Task) Tags() []TaskID {
	return append([]TaskID{}, t.tags...)
}

func (t Task) Completed() bool {
	return t.completed
}

func NewTask(
	title string,
	description string,
	tags []TaskID,
	completed bool,
) Task {
	return Task{
		title:       title,
		description: description,
		tags:        tags,
		completed:   completed,
	}
}
