package ch8

import "testing"

func TestExec(t *testing.T) {

}

func TestIsEven(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "even", args: args{2}, want: true},
		{name: "odd", args: args{1}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.i); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
