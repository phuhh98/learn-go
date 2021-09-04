package main

func foo(sum *int, s ...int) int {
	if len(s) != 1 {
		foo(sum, s[1:]...)
	}
	*sum = *sum + s[0]
	return *sum
}

//Bar function return sum of input []int
func Bar(s []int) int {
	var sum int
	foo(&sum, s...)

	return sum
}
