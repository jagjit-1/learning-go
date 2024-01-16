package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2;

	fmt.Print("Write ", i, " as ");
	switch i {
	case 1 :
		fmt.Println("one");
	
	case 2 : 
		fmt.Println("two");
	
	case 3 :
		fmt.Println("three");
	}

	t := time.Now();
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon");
	default:
		fmt.Println("It's after noon");
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Im a bool");
		case int:
			fmt.Println("im an int");
		default:
			fmt.Printf("dont know the type :( %T\n", t);
		}
	}

	whatAmI(true);
	whatAmI("sfesf");
	whatAmI(12);



}