package app

import "testing"

func TestParseLink(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "valid link",
			args: args{
				link: "https://www.google.com",
			},
			want: "google.com",
		},
		{
			name: "valid link with query params",
			args: args{
				link: "https://www.google.com/search?q=go+lang&oq=go+lang#hash",
			},
			want: "google.com",
		},
		{
			name: "valid link with path",
			args: args{
				link: "https://www.google.com/calendar/render?tab=wc",
			},
			want: "google.com",
		},
		{
			name: "valid link with path and query params",
			args: args{
				link: "https://www.google.com/calendar/render?tab=wc&q=go+lang&oq=go+lang#hash",
			},
			want: "google.com",
		},
		{
			name: "invalid link",
			args: args{
				link: "not a link",
			},
			want: "",
		},
		{
			name: "empty link",
			args: args{
				link: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLink(tt.args.link); got != tt.want {
				t.Errorf("ParseLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
