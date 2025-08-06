package component

type ComponentStack struct {
	components []*Component
}

func NewComponentStack() *ComponentStack {
	return &ComponentStack{components: []*Component{}}
}

func (cs *ComponentStack) Push(c *Component) {
	cs.components = append(cs.components, c)
}

func (cs *ComponentStack) Pop() *Component {
	if len(cs.components) == 0 {
		return nil
	}
	c := cs.components[len(cs.components)-1]
	cs.components = cs.components[:len(cs.components)-1]
	return c
}
