package fizzbuzz

import (
	"fmt"
	"github.com/sekou-diarra/fiz-buzz-server/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	FIZZ     = "fizz"
	BUZZ     = "buzz"
	FIZZBUZZ = "fizzbuzz"
)

func Test_fizzBuzzService_ComputeFizzBuzz(t *testing.T) {

	tests := []struct {
		name string
		args model.FizzbuzzParam
		want []string
	}{
		{
			"return_the_correct_sequence",
			model.FizzbuzzParam{FstMultiple: 2, SecMultiple: 5, Limit: 10, Label1: FIZZ, Label2: BUZZ},
			[]string{"1", FIZZ, "3", FIZZ, BUZZ, FIZZ, "7", FIZZ, "9", FIZZBUZZ},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFizzBuzzService()
			got := f.ComputeFizzBuzz(tt.args).Sequences
			assert.Equal(t, tt.want, got, "sequence should be the same")
		})
	}
}

func Test_fizzBuzzService_isMultipleOf(t *testing.T) {

	type args struct {
		numberToTest int
		multiple1    int
		multiple2    int
	}

	type want struct {
		res1 bool
		res2 bool
	}

	tests := []struct {
		name string
		args args
		want want
	}{

		{
			"is_multiple_of_first_param",
			args{18, 9, 5},

			want{true, false},
		},

		{
			"is_multiple_of_second_param",
			args{10, 9, 5},
			want{
				false, true,
			},
		},

		{
			"is_multiple_of_both_param",
			args{10, 2, 5},
			want{
				true, true,
			},
		},

		{
			"is_not_multiple_of_both_param",
			args{3, 2, 5},
			want{
				false, false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := isMultipleOf(tt.args.numberToTest, tt.args.multiple1, tt.args.multiple2)

			assert.Equal(t, tt.want.res1, got1, fmt.Sprintf("%s failed", tt.name))
			assert.Equal(t, tt.want.res2, got2, fmt.Sprintf("%s failed", tt.name))
		})
	}
}
