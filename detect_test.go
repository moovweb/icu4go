package icu4go

import "testing"
import "io/ioutil"

func TestDetectChinese(t *testing.T) {
	data, err := ioutil.ReadFile("input_cn.html")
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	detector, err := NewCharsetDetector()
	if err != nil {
		t.Errorf("error: %s", err.Error())
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
		t.Errorf("error: %s", err.Error())
	}
	detector, err := NewCharsetDetector()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	charset := detector.GuessCharset(data)
	if charset != "ISO-8859-1" {
		t.Errorf("the charset guess was wrong")
	}
	detector.Free()
}

func _TestDetectCp1252(t *testing.T) {
	data, err := ioutil.ReadFile("input_cp1252.html")
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	detector, err := NewCharsetDetector()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	charset := detector.GuessCharset(data)
	if charset != "windows-1252" {
		println(charset)
		t.Errorf("the charset guess was wrong")
	}
	detector.Free()
}

func TestDetectUtf8(t *testing.T) {
	data, err := ioutil.ReadFile("input_utf8.html")
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	detector, err := NewCharsetDetector()
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}
	charset := detector.GuessCharset(data)
	if charset != "UTF-8" {
		println(charset)
		t.Errorf("the charset guess was wrong")
	}
	detector.Free()
}
