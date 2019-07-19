package utils

func HandlePanic(err *error) {
	if r := recover(); r != nil {
		*err = r.(error)
	}
}

func HandleErr(err error)  {
	if err != nil {
		panic(err)
	}
}
