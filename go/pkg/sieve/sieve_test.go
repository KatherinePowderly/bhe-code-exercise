package sieve_test

import (
	"errors"
	"math/big"
	"ssse-exercise-sieve/pkg/sieve"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNthPrime(t *testing.T) {
	type args struct {
		number int64
	}
	type want struct {
		res int64
		err error
	}
	tests := []struct {
		name string
		args *args
		want *want
	}{
		{
			name: "Error: Negative number is invalid",
			args: &args{number: -1},
			want: &want{res: 0, err: errors.New("invalid number: -1, negative numbers not allowed")},
		},
		{
			name: "Success Case 1",
			args: &args{number: 0},
			want: &want{res: 2, err: nil},
		},
		{
			name: "Success Case 2",
			args: &args{number: 19},
			want: &want{res: 71, err: nil},
		},
		{
			name: "Success Case 3",
			args: &args{number: 99},
			want: &want{res: 541, err: nil},
		},
		{
			name: "Success Case 4",
			args: &args{number: 500},
			want: &want{res: 3581, err: nil},
		},
		{
			name: "Success Case 5",
			args: &args{number: 986},
			want: &want{res: 7793, err: nil},
		},
		{
			name: "Success Case 6",
			args: &args{number: 2000},
			want: &want{res: 17393, err: nil},
		},
		{
			name: "Success Case 7",
			args: &args{number: 1000000},
			want: &want{res: 15485867, err: nil},
		},
		{
			name: "Success Case 8",
			args: &args{number: 10000000},
			want: &want{res: 179424691, err: nil},
		},
		{
			name: "Success Case 9",
			args: &args{number: 100000000},
			want: &want{res: 2038074751, err: nil},
		},
	}
	for _, tt := range tests {
		tt := tt // pin
		t.Run(tt.name, func(t *testing.T) {
			sieve := sieve.NewSieve()
			got, err := sieve.GetNthPrime(tt.args.number)
			require.Equal(t, tt.want.err, err)
			require.Equal(t, tt.want.res, got)
		})
	}
}

func FuzzNthPrime(f *testing.F) {
	sieve := sieve.NewSieve()

	f.Fuzz(func(t *testing.T, n int64) {
		sieve, err := sieve.GetNthPrime(n)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !big.NewInt(sieve).ProbablyPrime(0) {
			t.Errorf("the sieve produced a non-prime number at index %d", n)
		}
	})
}
