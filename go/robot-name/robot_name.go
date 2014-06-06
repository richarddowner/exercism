package robotname

import (
	//"fmt"
	"math/rand"
	"strconv"
	"time"
)

var identity string = ""

type Robot struct {
	name string
}

func (r *Robot) Name() (name string) {
	if identity == "" {
		rand.Seed(time.Now().UTC().UnixNano()) // this needs to be lazy init
		name += string(randInt(65, 90))
		name += string(randInt(65, 90))
		name += strconv.Itoa(randInt(0, 9))
		name += strconv.Itoa(randInt(0, 9))
		name += strconv.Itoa(randInt(0, 9))
		identity = name
	}
	return identity
}

func (r *Robot) Reset() {
	identity = ""
}

func randInt(min int, max int) (num int) {
	num = rand.Intn(max-min) + min
	return
}
