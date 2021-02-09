package main

import (
    "fmt"
    "bytes"
)

func main() {
    //fmt.Println(CombAll([]rune("abc")))
    fmt.Println(Combinations([]string{"abc"}))
}

func CombAll(set []rune) (subsets []string) {
    length := uint(len(set))

    var buf bytes.Buffer
    for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
        for obj := uint(0); obj < length; obj++ {
            if (subsetBits >> obj)&1 == 1 {
                buf.WriteString(string(set[obj]))
            }
        }

        subsets = append(subsets, buf.String())
        buf.Reset()
    }

    return subsets
}

func Combinations(set []string, n int) (subsets [][]string) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []string

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}
