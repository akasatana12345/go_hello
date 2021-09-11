package main

import "testing"

func TestStringAdd(t *testing.T) {
	type args struct {
		x, y string
	}
	tests := []struct {
		testName string
		args     args
		want     string
	}{
		{
			testName: "lower",
			args:     args{x: "a", y: "b"},
			want:     "ab",
		},
		{
			testName: "upper",
			args:     args{x: "A", y: "B"},
			want:     "AB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := stringAdd(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("add = %v, want %v", got, tt.want)
			}
		})
	}
}
