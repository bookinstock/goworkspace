package sys

import (
	"fmt"
	"os"
)

func RunExit() {

	defer fmt.Println("!")

	os.Exit(3)
}
