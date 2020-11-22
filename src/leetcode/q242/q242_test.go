package q242

import "testing"

func TestIsAnagram(t *testing.T)  {
	s1 := "a"
	s2 := "b"
	if !isAnagram(s1,s2) {
		t.FailNow()
	}
}