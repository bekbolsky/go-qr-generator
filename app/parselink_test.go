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
		{"Test - 1", args{"https://www.google.com/"}, "www.google.com-www.google.com"},
		{"Test - 2", args{"https://e.mail.ru/tomyself/"}, "e.mail.ru-tomyself"},
		{"Test - 3", args{"https://sozdik.kz/ru/"}, "sozdik.kz-ru"},
		{"Test - 4", args{"https://eduindex.kz/pedagogical-work/1451-aza-til-bilimini-salalaryn-satylaj-keshendi-taldau-tehnologijasy-aryly-megertu.html"}, "eduindex.kz-1451-aza-t"},
		{"Test - 5", args{"http://edunews.kz/lentnews/2427-bgn-halyaraly-aza-tl-oamy-oamdy-brlestgn-trala-otyrysy-tt.html"}, "edunews.kz-2427-bgn-h"},
		{"Test - 6", args{"https://www.youtube.com/watch?v=dQw4w9WgXcQ"}, "www.youtube.com-watchvdQw4"},
		{"Test - 7", args{"https://automatetheboringstuff.com/2e/chapter12/"}, "automatetheboringstuff.com-chapter12"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseLink(tt.args.link); got != tt.want {
				t.Errorf("ParseLink() = %v, want %v", got, tt.want)
			}
		})
	}
}
