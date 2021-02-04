package bloomfilter

import "testing"

func TestMurmurHash2(t *testing.T) {
	type args struct {
		b    []byte
		seed uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "case1", args: args{b: []byte("abc"), seed: 7}, want: 957085255},
		{name: "case2", args: args{b: []byte("aaaa"), seed: 17}, want: 455770523},
		{name: "case3", args: args{b: []byte("hello"), seed: 23}, want: 3811527086},
		{name: "case4", args: args{b: []byte("python"), seed: 29}, want: 151976208},
		{name: "case5", args: args{b: []byte("aaaaaaaaa aaaaaaaaaa"), seed: 42}, want: 2860058555},
		{name: "case6", args: args{b: []byte(""), seed: 7}, want: 2193385065},
		{name: "case7", args: args{b: []byte(" "), seed: 7}, want: 3009246799},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MurmurHash2(tt.args.b, tt.args.seed); got != tt.want {
				t.Errorf("MurmurHash2() = %v, want %v", got, tt.want)
			}
		})
	}
}