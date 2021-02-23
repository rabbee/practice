package main

import (
	"bufio"
	"fmt"
	"os"
)

func HJ02() {
	reader := bufio.NewReader(os.Stdin)
	inStr, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	cntMap := make(map[byte]int)

	for i := 0; i < len(inStr); i++ {
		if 'A' <= inStr[i] && inStr[i] <= 'Z' {
			cntMap[inStr[i]-'A'+'a']++
		} else if 'a' <= inStr[i] && inStr[i] <= 'z' {
			cntMap[inStr[i]]++
		}
	}

	target, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	if 'A' <= target[0] && target[0] <= 'Z' {
		fmt.Println(cntMap[target[0]-'A'+'a'])
	} else {
		fmt.Println(cntMap[target[0]])
	}
}
