package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

type Robot struct {
	name string
}

func (r *Robot) Name() (name string) {
	name += string(randInt(65, 90))
	name += string(randInt(65, 90))
	name += strconv.Itoa(randInt(0, 9))
	name += strconv.Itoa(randInt(0, 9))
	name += strconv.Itoa(randInt(0, 9))
	return name
}

func (r *Robot) Reset() {

}

func randInt(min int, max int) (num int) {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}
