package coding_interviews

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReversePrint(t *testing.T) {
	var list *ListNode = nil
	fmt.Println(reversePrint(list))
}

//func Test09(t *testing.T) {
//	obj := Constructor()
//	obj.AppendTail(2)
//	v := obj.DeleteHead()
//	fmt.Println(v)
//}

func Test_minArray(t *testing.T) {
	fmt.Println(minArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(minArray([]int{7, 8, 9, 0, 1, 2, 3, 4, 5, 6}))
	fmt.Println(minArray([]int{2, 2, 2, 0, 1}))
	fmt.Println(minArray([]int{}))
	fmt.Println(minArray([]int{3}))
}

func Test_fib(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(fib(i))
	}
}

func Test_multiMatrix(t *testing.T) {
	type args struct {
		a [2][2]int
		b [2][2]int
	}
	tests := []struct {
		name string
		args args
		want [2][2]int
	}{
		// TODO: Add test cases.
		{
			"1",
			args{
				a: [2][2]int{
					{1, 1},
					{1, 0},
				},
				b: [2][2]int{
					{1, 1},
					{1, 0},
				},
			},
			[2][2]int{
				{2, 1},
				{1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiMatrix(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("multiMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
