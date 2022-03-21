package protocol

type Task struct {
	id          string
	title       string
	description string
	tags        []string
	completed   bool
}

func (t Task) Id() string {
	return t.id
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

func (t Task) Tags() []string {
	return append([]string{}, t.tags...)
}

func (t Task) Completed() bool {
	return t.completed
}

func NewTask(
	id string,
	title string,
	description string,
	tags []string,
	completed bool,
) Task {
	return Task{
		id:          id,
		title:       title,
		description: description,
		tags:        tags,
		completed:   completed,
	}
}
