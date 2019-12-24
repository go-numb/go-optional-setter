package optional

type Float64 struct {
	set   bool
	value float64
}

func (p *Float64) IsSet() bool {
	return p.set
}

func (p *Float64) Value() float64 {
	return p.value
}

func (p *Float64) Set(v float64) {
	p.set = true
	p.value = v
}
