package main

import (
	"fmt"
	"time"

	"github.com/Warashi/magicalclock"
)

const dayLength = 26

//go:generate stringer -type=JapaneseWeekDay
type JapaneseWeekDay int

const (
	日 JapaneseWeekDay = iota
	月
	火
	水
	木
	金
	土
)

func main() {
	t := time.Now().Local()
	_, M, D := t.Date()
	W := JapaneseWeekDay(t.Weekday())
	h, m, _ := magicalclock.New(dayLength).Apply(t.Clock())
	fmt.Printf("%d月%d日(%s) %02d:%02d\n", M, D, W, h, m)
}
