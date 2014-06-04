package robotname

type Robot struct {
	name string
}

func (r *Robot) Name() string {
	return "Test"
}
