package main

import (
	"fmt"
	"time"

	"github.com/Warashi/magicalclock"
)

func main() {
	h, m, _ := magicalclock.New(dayLength).Apply(time.Now().Local().Clock())
	fmt.Printf("%d:%d\n", h, m)
}
