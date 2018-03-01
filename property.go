package schemaorg

type Property interface {
	Name() string
}

func NewProperty(name string) Property {
	return property{
		name: name,
	}
}

type property struct {
	name string
}

func (p property) Name() string {
	return p.name
}
