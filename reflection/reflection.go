package reflection

func walk(x interface{}, fn func(input string)) {
	fn("this is a nonsensical string")
}
