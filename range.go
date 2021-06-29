package main

func main() {
	var s = "abc"

	for i := 0; i < 3; i++ {
		println(s[i])
	}

	println("----")

	for i := range s {
		println(s[i])
	}

	println("----")

	for _, c := range s {
		println(c)
	}
}
