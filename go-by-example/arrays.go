package main

import "fmt"

func main() {
	var a [5]int;
	fmt.Println(a);

	a[4] = 100;
	fmt.Println(a[4]);

	b := [4]int {1,2,3,4};
	fmt.Println(b);

	var twoD [3][4]int;
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			twoD[i][j] = i + j;
		}
	}
	
	fmt.Println(twoD);
}