package main

import (
	"fmt"
)

func frequencySort(s string) string {
	mapCharacter := make(map[string]int)

	for _, character := range s {
		mapCharacter[string(character)] += 1
	}

	push := func(heap *[]string, character string) {
		*heap = append(*heap, character)

		i := len(*heap) - 1
		for {
			if i <= 0 {
				break
			}

			parent := (i - 1) / 2

			curChar := (*heap)[i]
			curParentChar := (*heap)[parent]

			// fmt.Println(mapCharacter)
			// fmt.Println(curChar, curParentChar)
			// fmt.Println(i, parent)

			if mapCharacter[curParentChar] < mapCharacter[curChar] {
				(*heap)[parent], (*heap)[i] = (*heap)[i], (*heap)[parent]
			}
			i = parent
		}
	}

	pop := func(heap *[]string) string {
		(*heap)[len(*heap)-1], (*heap)[0] = (*heap)[0], (*heap)[len(*heap)-1]
		max := (*heap)[len(*heap)-1]

		(*heap) = (*heap)[:len(*heap)-1]

		heapLen := len(*heap)

		i := 0
		for {
			curLeft := 2*i + 1
			curRight := 2*i + 2
			biggest := i

			if biggest >= heapLen {
				break
			}

			curCharacter := (*heap)[biggest]

			if heapLen > curLeft {
				leftCharacter := (*heap)[curLeft]

				if mapCharacter[leftCharacter] > mapCharacter[curCharacter] {
					biggest = curLeft
					curCharacter = (*heap)[curLeft]
				}
			}

			if heapLen > curRight {
				rightCharacter := (*heap)[curRight]
				if mapCharacter[rightCharacter] > mapCharacter[curCharacter] {
					biggest = curRight
				}
			}

			if biggest == i {
				break
			}

			(*heap)[i], (*heap)[biggest] = (*heap)[biggest], (*heap)[i]

			i = biggest
		}

		return max
	}

	heap := []string{}
	for character, _ := range mapCharacter {
		push(&heap, character)
	}

	solution := ""
	for {
		if len(heap) <= 0 {
			break
		}

		result := pop(&heap)

		for _ = range mapCharacter[result] {
			solution += result
		}
	}

	return solution
}

func main() {
	fmt.Println(frequencySort("Aabb"))
}
