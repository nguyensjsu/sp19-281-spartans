package functions

type test struct {
	S string
	E string
}

func GetValue1(a, b string) test {
	return test{
		a, b,
	}
}
