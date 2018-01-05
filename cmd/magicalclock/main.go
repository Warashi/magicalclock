package main

import (
	"fmt"
	"time"

	"github.com/Warashi/magicalclock"
)

func main() {
	h, m, _ := magicalclock.New(dayLength).Apply(time.Now().Local().Clock())
	fmt.Printf("%02d:%02d\n", h, m)
}
