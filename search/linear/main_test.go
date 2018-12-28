package main

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	var refArr []int = []int{1,6,2,4,9,10}

	res := linearSearch(refArr, 3)

	if (res == -1) {
		t.Log("ok, 测试通过")
	} else {
		t.Error("测试错误")
	}
}
