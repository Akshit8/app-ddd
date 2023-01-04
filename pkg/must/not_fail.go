package must

func NotFail(err error) {
	if err != nil {
		panic(err)
	}
}

func NotFailf(f func() error) {
	NotFail(f())
}
