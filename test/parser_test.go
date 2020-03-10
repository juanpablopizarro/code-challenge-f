package parser

import (
	"fmt"
	"testing"

	"github.com/juanpablopizarro/code-challenge-f/parser"
)

func TestParseSuccessfully(t *testing.T) {
	testData := []byte("11AB398765UJ1A050200N23")
	expected := map[string]string{
		"0": "{11 AB398765UJ1 A 5}",
		"1": "{2 00 N 23}",
	}

	actual, err := parser.Unmarshal(testData)

	if err != nil {
		t.Errorf("expected err=nil but error '%v' found", err)
	}

	for eK, eV := range expected {
		aV, ok := actual[eK]
		if !ok {
			t.Errorf("expected key '%v' was not found in actual", eK)
		}

		if eV != aV {
			t.Errorf("expected value is not correct in actual for key=%v. Expected value: %v, Actual value: %v", eK, eV, aV)
		}
	}
}

func TestParseEmptyInput(t *testing.T) {
	testData := []byte("")
	_, err := parser.Unmarshal(testData)

	if err.Error() != "invalid unmarshal's input" {
		t.Errorf("expected err 'invalid unmarshal's input'. Found: %v", err)
	}
}

func TestParseInvalidLength(t *testing.T) {
	testData := []byte("1XAB398765UJ1A050200N23")
	_, err := parser.Unmarshal(testData)

	if err == nil {
		t.Error("error not found with an invalid length")
	}
}

func TestParseInvalidType(t *testing.T) {
	testData := []byte("11AB398765UJ1R05")
	_, err := parser.Unmarshal(testData)

	fmt.Printf("err=%v\n", err.Error())
	fmt.Println("err.Error() != 'invalid type': "+err.Error() != "invalid type")
	if err == nil || err.Error() != "invalid type" {
		t.Errorf("error not found. We want 'invalid type'. We got: %v", err)
	}
}

func TestParseInvalidValue(t *testing.T) {
	testData := []byte("11AB398765UJ1N05")
	_, err := parser.Unmarshal(testData)

	if err == nil || err.Error() != "invalid value" {
		t.Errorf("error not found. We want 'invalid value'. err=%v", err)
	}

}
