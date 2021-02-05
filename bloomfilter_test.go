package bloomfilter

import "testing"

func TestBloomFilter_Check(t *testing.T) {
	b, _ := NewBloomFilter(20000, 0.05)
	b.Set([]byte("Hello"))
	b.Set([]byte("Python"))
	b.Set([]byte(" "))
	b.Set([]byte(""))
	b.Set([]byte("#$%^&*("))
	b.Set([]byte("Today is a big day!"))
	b.Set([]byte("Bottle112345"))
	b.Set([]byte("123456!@#$%^&*"))
	b.Set([]byte(" "))
	b.Set([]byte(""))
	b.Set([]byte("	Monday	"))
	b.Set([]byte("a"))

	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case1", args: args{data: []byte(" ")}, want: true},
		{name: "case2", args: args{data: []byte("a")}, want: true},
		{name: "case3", args: args{data: []byte("Bottle112345")}, want: true},
		{name: "case4", args: args{data: []byte("Today is a big day!")}, want: true},
		{name: "case5", args: args{data: []byte("c")}, want: false},
		{name: "case6", args: args{data: []byte("Today is a great day!")}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := b.Check(tt.args.data); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
