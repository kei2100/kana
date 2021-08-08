package kana

import (
	"strings"
)

// NormalizePSM は、文字列 s に含まれる長音記号風の文字を、
// 全角の長音記号 (U+30FC: Katakana-Hiragana Prolonged Sound Mark) にノーマライズします。
func NormalizePSM(s string) string {
	b := &strings.Builder{}
	for _, r := range s {
		if _, ok := likePSMTable[r]; ok {
			b.WriteRune(0x30FC)
		} else {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// likePSMTable は、長音記号風の文字のテーブルです。
var likePSMTable = map[rune]struct{}{
	0x002D: {}, // Hyphen-Minus
	//0x00AD: {}, // Soft Hyphen (必要？)
	0x02D7: {}, // Modifier Letter Minus Sign
	0x2010: {}, // Hyphen
	0x2011: {}, // Non-Breaking Hyphen
	0x2012: {}, // Figure Dash
	0x2013: {}, // En Dash
	0x2014: {}, // Em Dash
	0x2015: {}, // Horizontal Bar
	0x2043: {}, // Hyphen Bullet
	0x2212: {}, // Minus Sign
	//0x2796: {}, // Two-Em Dash (必要？)
	//0x2E3A: {}, // Heavy Minus Sign (必要？)
	//0x2E3B: {}, // Three-Em Dash (必要？)
	0x30FC: {}, // Katakana-Hiragana Prolonged Sound Mark
	0xFE58: {}, // Small Em Dash
	0xFE63: {}, // Small Hyphen-Minus
	0xFF0D: {}, // Fullwidth Hyphen-Minus
	0xFF70: {}, // Halfwidth Katakana-Hiragana Prolonged Sound Mark
}
