package mnstring

func MyString(str chan<- string) {
	str <- "test"
	close(str)
}
