package fibo

type F struct {
	Seed   int
	Result int
}

func (f *F) Calc() int {
	r := 0
	for i := 1; i <= f.Seed; i++ {
		r += i
	}
	f.Result = r
	return f.Result
}
