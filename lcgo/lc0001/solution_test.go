package lc0001_test

import (
	"lcgo/lc0001"
	"testing"
	"fmt"
)
func Test_LC_0001(t *testing.T) {
	nums := []int{2, 7, 11, 15}

	res := lc0001.TwoSum(nums, 9)
	fmt.Printf("%v\n", res)
	//println(res)
}