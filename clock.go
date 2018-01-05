package magicalclock

type clock int

type Clock interface {
	Apply(hour, minute, second int) (h, m, s int)
}

func New(length int) Clock {
	return clock(length)
}

func (c clock) Apply(h, m, s int) (hour, minute, second int) {
	u := 3600*h + 60*m + s
	u *= int(c)
	u /= 24
	s = u % 60
	u /= 60
	m = u % 60
	u /= 60
	h = u
	return h, m, s
}
