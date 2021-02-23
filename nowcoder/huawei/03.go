package main

import (
	"fmt"
	"io"
	"sort"
)

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func HJ03() {
	for {
		var total int
		n, err := fmt.Scan(&total)
		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
		if n == 0 || total == 0 {
			return
		}
		intSlice := make(IntSlice, 0)

		intMap := map[int]interface{}{}

		for i := 0; i < total; i++ {
			var tmp int
			_, err := fmt.Scan(&tmp)
			if err != nil {
				panic(err)
			}
			if _, existed := intMap[tmp]; !existed {
				intMap[tmp] = nil
				intSlice = append(intSlice, tmp)
			}
		}
		sort.Sort(intSlice)

		for _, item := range intSlice {
			fmt.Println(item)
		}
	}
}
