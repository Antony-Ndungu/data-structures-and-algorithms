package sum_test

import (
	"testing"

	"github.com/Antony-Ndungu/data-structures-and-algorithms/go/sum"
)

const success = "\u2713"
const failed = "\u2717"

func TestInts(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"one to five", args{nums: []int{1, 2, 3, 4, 5}}, 15},
		{"nothing", args{nil}, 0},
		{"one and minus one", args{[]int{-1, 1}}, 0},
	}

	t.Log("Given the need to test adding a list of integers")
	{

		for i, tt := range tests {
			t.Logf("\tTest %d:\tWhen checking the sum of %v", i, tt.args.nums)
			{
				t.Run(tt.name, func(t *testing.T) {
					if got := sum.Ints(tt.args.nums...); got != tt.want {
						t.Errorf("\t%s Ints() = %v, want %v", failed, got, tt.want)
					} else {
						t.Logf("\t%s Ints() = %v, want %v", success, got, tt.want)
					}
				})
			}
		}

	}

}
