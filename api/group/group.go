package group

type Group struct {
	name    string
	groupID uint16
}

func (g *Group) GetName() string {
	return g.name
}

func NewGroup(name string, groupID uint16) *Group {
	return &Group{name: name, groupID: groupID}
}

func (g *Group) UpdateName(name string) {
	g.name = name
}
