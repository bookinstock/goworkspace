package ctrlflow

import "fmt"

func loop() {

	// for
	for i := 0; i < 5; i++ {
		fmt.Println("loop1 i=", i)
	}

	// while
	i := 0
	for i < 5 {
		fmt.Println("loop2 i=", i)
		i += 1
	}

	// break
	i = 0
	for {
		if i == 3 {
			break
		}
		fmt.Println("loop3 i=", i)
		i += 1
	}

	// continue
	i = 0
	for i < 5 {
		if i == 3 {
			i += 1
			continue
		}
		fmt.Println("loop4 i=", i)
		i += 1
	}
}
