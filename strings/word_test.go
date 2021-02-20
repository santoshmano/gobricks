package word

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	if IsPalindrome("detartrated") == false {
		t.Error(`IsPalindrome("detartrated") = false, should be true`)
	}

	if IsPalindrome("kayak") == false {
		t.Error(`IsPalindrome("kayak") = false, should be true`)
	}

	if IsPalindrome("") == false {
		t.Error(`IsPalindrome("") = false, should be true`)
	}

	if IsPalindrome("adsf") == true {
		t.Error(`IsPalindrome("asdf") = true, should be false`)
	}
}

func TestIsPalindromeCanal(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if IsPalindrome(input) == false {
		t.Errorf(`isPalindrom(%v) = false, should be true`, input)
	}
}

func TestIsPalindromeFrench(t *testing.T) {
	if IsPalindrome("वलव") == false {
		t.Error(`isPalindrome("वलव") = false, should be true`)
	}
}
