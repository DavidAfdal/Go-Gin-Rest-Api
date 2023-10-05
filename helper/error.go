package helper

func ErrorPanic(e error) {
	if e != nil {
		panic(e)
	}
}