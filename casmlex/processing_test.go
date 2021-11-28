package casmlex


import "testing"


func TestIsInSlice(t *testing.T) {
	mainSlice := []string{"a", "b", "cd", "hello", "world"}
	falseSlice := []string{"aa", "bb", "c", "d", "helloworld", "bro"}

	_ = falseSlice

	for _, key := range mainSlice {
		if !isInSlice(key, mainSlice) {
			t.Error("for", key, "expected true, got false")
		}
	}

	for _, key := range falseSlice {
		if isInSlice(key, mainSlice) {
			t.Error("for", key, "expected false, got true")
		}
	}

} 


func TestIsByteInSlice(t *testing.T) {
	mainSlice := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'}
	falseSlice := []byte{'k', 'l', 'y', 'z', 'x'}

	for _, key := range mainSlice {
		if !isByteInSlice(key, mainSlice) {
			t.Error("for", string(key), "expected true, got false")
		}
	}

	for _, key := range falseSlice {
		if isByteInSlice(key, mainSlice) {
			t.Error("for", string(key), "expected false, got true")
		}
	}
}


func TestGetSourceCode (t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("panic() expected")
        }
    }()

	
	GetSourceCode("file.txt")
}


func TestIsOpenBracket(t *testing.T) {
	openBrackets := []byte{'(', '[', '{'}
	closeBrackets := []byte{')', ']', '}'}

	for _, bracket := range openBrackets {
		if !isOpenBracket(bracket) {
			t.Error("for", string(bracket), "expected true, got false")
		}
	} 

	for _, bracket := range closeBrackets {
		if isOpenBracket(bracket) {
			t.Error("for", string(bracket), "expected false, got true")
		}
	} 
}