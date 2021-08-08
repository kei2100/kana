package kana

import (
	"fmt"
	"testing"
)

func TestNormalizePSM(t *testing.T) {
	tt := []struct {
		s    string
		want string
	}{
		{
			s:    "日\u002d本\uFF70語",
			want: "日\u30FC本\u30FC語",
		},
		{
			s: "\u002D" +
				"\u02D7" +
				"\u2010" +
				"\u2011" +
				"\u2012" +
				"\u2013" +
				"\u2014" +
				"\u2015" +
				"\u2043" +
				"\u2212" +
				"\u30FC" +
				"\uFE58" +
				"\uFE63" +
				"\uFF0D" +
				"\uFF70",
			want: "\u30FC" +
				"\u30FC" +
				"\u30FC" +
				"\u30FC" +
				"\u30FC" +
				"\u30FC" +
				"\u30FC" +
				"\u30FC" +
				"\u30FC" +
				"\u30FC" +
				"\u30FC" +
				"\u30FC" +
				"\u30FC" +
				"\u30FC" +
				"\u30FC",
		},
	}
	for i, te := range tt {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			got := NormalizePSM(te.s)
			if got != te.want {
				t.Errorf("\ngot :%U\nwant:%U", []rune(got), []rune(te.want))
			}
		})
	}
}
