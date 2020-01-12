package utils

func HandlePanic(err *error) {
	if r := recover(); r != nil {
		*err = r.(error)
	}
}

func Check(err error)  {
	if err != nil {
		panic(err)
	}
}
