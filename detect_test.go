package icu4go

import "testing"
import "io/ioutil"

func TestDetectChinese(t *testing.T) {
	data, err := ioutil.ReadFile("input_cn.html")
	if err != nil {
		t.Errorf("error: %s", err.String())
	}
	detector, err := NewCharsetDetector()
	if err != nil {
		t.Errorf("error: %s", err.String())
	}
	charset := detector.GuessCharset(data)
	if charset != "GB18030" {
		t.Errorf("the charset guess was wrong")
	}
	detector.Free()
}

func TestDetectLatin1(t *testing.T) {
	data, err := ioutil.ReadFile("input_latin1.html")
	if err != nil {
		t.Errorf("error: %s", err.String())
	}
	detector, err := NewCharsetDetector()
	if err != nil {
		t.Errorf("error: %s", err.String())
	}
	charset := detector.GuessCharset(data)
	if charset != "ISO-8859-1" {
		t.Errorf("the charset guess was wrong")
	}
	detector.Free()
}