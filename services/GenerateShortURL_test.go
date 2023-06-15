package services

import "testing"

func TestHashShortURL(t *testing.T) {
	type args struct {
		URL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HashShortURL(tt.args.URL); got != tt.want {
				t.Errorf("HashShortURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
