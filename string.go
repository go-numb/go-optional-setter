package optional

type String struct {
	set   bool
	value string
}

func (p *String) IsSet() bool {
	return p.set
}

func (p *String) Value() string {
	return p.value
}

func (p *String) Set(v string) {
	p.set = true
	p.value = v
}
