package lib

import "fmt"

func Foo(a int, b string) string {
	return fmt.Sprintf("I see %d %s.", a, b)
}
