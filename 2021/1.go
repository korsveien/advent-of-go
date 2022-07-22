package main

import "fmt"

var depths = [10]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func main() {
	var lastDepth = -1
	var count = 0

	for _, curDepth := range depths {
		if lastDepth == -1 {
			fmt.Printf("%d (N/A. No previous measurement)\n", curDepth)
		} else if curDepth > lastDepth {
			count += 1
			fmt.Printf("%d (increased)\n", curDepth)
		} else {
			fmt.Printf("%d (decreased)\n", curDepth)
		}

		lastDepth = curDepth
	}

	fmt.Printf("\n%d measurements are larger than the previous measurement.\n", count)
}
