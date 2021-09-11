package main

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		x, y int
	}
	tests := []struct {
		testName string
		args     args
		want     int
	}{
		{
			testName: "plus",
			args:     args{x: 1, y: 2},
			want:     3,
		},
		{
			testName: "minus",
			args:     args{x: 1, y: -2},
			want:     -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add = %v, want %v", got, tt.want)
			}
		})
	}
}
