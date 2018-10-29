package _map

import (
	"fmt"
)

func lengthOfNoRepeatingStrsub(str string) int {
	lastOccurred := make(map[byte]int) //empty map
	start := 0
	maxLength := 0
	fmt.Println(lastOccurred)
	//byte与string是ASCII码的装换关系
	for i, ch := range []byte(str) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	fmt.Println(lastOccurred)
	return maxLength
}

func main() {
	fmt.Println(lengthOfNoRepeatingStrsub("bbb"))
	fmt.Println(lengthOfNoRepeatingStrsub("abcdefg"))
	fmt.Println(lengthOfNoRepeatingStrsub("abdeeeedfdfdfd"))

}
