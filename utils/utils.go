package utils

//PanicOnError panics if the provided error is not nil
func PanicOnError(e error) {
	if e != nil {
		panic(e)
	}
}
