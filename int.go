package optional

type Int struct {
	set   bool
	value int
}

func (p *Int) IsSet() bool {
	return p.set
}

func (p *Int) Value() int {
	return p.value
}

func (p *Int) Set(v int) {
	p.set = true
	p.value = v
}
