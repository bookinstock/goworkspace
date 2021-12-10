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

	strs = []string{"aaa", "bb", "cccc"}
	sort.Sort(byLength(strs))
	fmt.Println("strs=", strs)
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
