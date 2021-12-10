package sort

import (
	"fmt"
	"sort"
)

func Run() {
	strs := []string{"b", "a", "c"}
	sort.Strings(strs)
	fmt.Println("strs=", strs, sort.StringsAreSorted(strs))

	ints := []int{2, 1, 3}
	sort.Ints(ints)
	fmt.Println("ints=", ints, sort.IntsAreSorted(ints))
}
