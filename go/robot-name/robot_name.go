package robotname

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Robot struct {
	name string
}

func (r *Robot) Name() (name string) {
	rand.Seed(time.Now().UTC().UnixNano())
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
	num = rand.Intn(max-min) + min
	fmt.Println(num)
	return
}
