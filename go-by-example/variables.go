package main

import "fmt"

func main() {
	// inferred type
	var a = "intial";
	fmt.Println(a);

	// explicit type
	var b, c int = 1, 2;
	fmt.Println(b, c);

	var e int;
	fmt.Println(e);

	f := "apple";
	fmt.Println(f);

}