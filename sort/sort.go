package sorter

import "sort"

func BubbleSort(els []int) {
	for keepRunning := true; keepRunning; {
		keepRunning = false
		for i := 0; i < len(els)-1; i++ {
			if els[i] > els[i+1] {
				els[i], els[i+1] = els[i+1], els[i]
				keepRunning = true
			}
		}
	}
}

// Following the Benchmark results, we can optimize our solution like so:
func Sort(els []int) {
	if len(els) <= 1000 {
		BubbleSort(els)
		return
	}
	sort.Ints(els)
}
