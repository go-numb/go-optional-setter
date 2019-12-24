package optional

import "time"

type Time struct {
	set   bool
	value time.Time
}

func (p *Time) IsSet() bool {
	return p.set
}

func (p *Time) Value() time.Time {
	return p.value
}

func (p *Time) Set(v time.Time) {
	p.set = true
}
