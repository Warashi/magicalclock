// Package magicalclock は一日が何時間かを指定して、それに併せて1秒の長さを調整するパッケージ
package magicalclock

type clock int

// Clock は時刻を改変するApplyを持つinterface
type Clock interface {
	Apply(hour, minute, second int) (h, m, s int)
}

func New(length int) Clock {
	return clock(length)
// New の引数には一日を何時間に設定するかを指定
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
