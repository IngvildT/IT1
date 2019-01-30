package strings_test
import (
  "strings"
  "testing")

func Test_Index(t *testing.T) {
c​ onst s, sep, want = "chicken", "ken", 4 got := string.Index(s, sep) ifgot!=want {
t​ .Errorf("Index(%q,%q) = %v; want %v", s, sep, want, got) }​
}

