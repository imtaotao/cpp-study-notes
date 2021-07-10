package main

import "net/http"

// 错误的
func do() error {
	res, err := http.Get("http://xxxxxxxxxx")
	defer res.Body.Close()
	if err != nil {
		return err
	}

	// ..code...

	return nil
}

// 正确的
func do1() error {
	res, err := http.Get("http://xxxxxxxxxx")
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return err
	}

	// ..code...

	return nil
}

// 在错误的位置使用 defer
func main() {
	// do()
	do1()
}
