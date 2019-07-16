package domain

type Greeting struct {
	Name string
}

func (g *Greeting) SetName(name string) error {
	if name == "" {
		return EmptyNameError
	}
	g.Name = name

	return nil
}
