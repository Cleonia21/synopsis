package main

func main() {
	{ // new
		type customType struct {
			a int
		}

		var ptr *customType
		// ptr.a = 1 // panic

		// The NEW built-in function allocates memory. The first argument is a type,
		// not a value, and the value returned is a pointer to a newly
		// allocated zero value of that type.
		ptr = new(customType)
		ptr.a = 1
	}
}
