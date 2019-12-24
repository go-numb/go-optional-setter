package optional

type Bool struct {
	set   bool
	value bool
}

func (p *Bool) IsSet() bool {
	return p.set
}

func (p *Bool) Value() bool {
	return p.value
}

func (p *Bool) Set(v bool) {
	p.set = true
	p.value = v
}
