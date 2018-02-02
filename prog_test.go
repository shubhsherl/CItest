package mypkg

import "testing"

func TestProg_checkIP(t *testing.T) {
	actualResult := checkIP("00172.45.232.0")
	var expectedResult = false

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got %s", expectedResult, actualResult)
	}
}

func TestProg_Rmspecial(t *testing.T) {
	actualResult1 := Rmspecial("Iam .@#ee LoCkEd")
	var expectedResult1 = "IameeLoCkEd"

	if actualResult1 != expectedResult1 {
		t.Fatalf("Expected %s but got %s", expectedResult1, actualResult1)
	}
}
