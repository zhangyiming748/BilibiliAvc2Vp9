package constant

type Param struct {
	Root string
}

func (p *Param) GetRoot() string {
	return p.Root
}
func (p *Param) SetRoot(s string) {
	p.Root = s
}
