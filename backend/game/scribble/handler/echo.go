package echo

const (
	name string = "hello"
)

type Echo struct{}

func (e *Echo) HandleInteraction() {
}

func (e *Echo) Name() string {
	return name
}
