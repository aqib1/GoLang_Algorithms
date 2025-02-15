package main

import "fmt"

func main() {
	fmt.Println(reverse("aqib"))
}

func reverse(str string) string {
	if len(str) == 0 {
		return str
	}

	return reverse(str[1:]) + str[0:1]
}
