package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	s := "A man, a plan, a canal: Panama"
	for idx := 0; idx < b.N; idx++ {
		isPalindrome(s)
	}
}
func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "A man, a plan, a canal: Panama",
			args: args{s: "A man, a plan, a canal: Panama"},
			want: true,
		},
		{
			name: "race a car",
			args: args{s: "race a car"},
			want: false,
		},
		{
			name: " ",
			args: args{s: " "},
			want: true,
		},
		{
			name: "0P",
			args: args{s: "0P"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
