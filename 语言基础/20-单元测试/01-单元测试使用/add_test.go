package main

import "testing"

func TestAdd(t *testing.T) {
	res := add(1, 2)

	if res != 3 {
		t.Fatalf("add(1,2) 执行错误，期望值%v,实际值%v", 3, res)
	}
	t.Logf("add(1,2)执行成功")
}
