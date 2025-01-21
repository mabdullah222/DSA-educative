package questions

import "fmt"

func towerOfHanoi(num int, src byte, dst byte, temp byte) {
	if num < 1 {
		return
	}
	towerOfHanoi(num-1, src, temp, dst)
	//DONOT change print statement
	fmt.Printf("Move %d disk from peg %c to peg %c ", num, src, dst)
	// write some code here
	towerOfHanoi(num-1, temp, dst, src)
}
