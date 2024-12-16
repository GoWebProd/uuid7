package uuid7

import (
	"strings"
	"testing"
)

func TestScan(t *testing.T) {
	stringTest := "0193cf94-e889-78da-a958-bca85957636d"
	badTypeTest := 6
	invalidTest := "f47ac10b-58cc-0372-8567-0e02b2c3d4"

	byteTest := make([]byte, 16)
	byteTestUUID := MustParse(stringTest)
	copy(byteTest, byteTestUUID[:])

	// sunny day tests

	var uuid UUID

	if err := uuid.Scan(stringTest); err != nil {
		t.Fatal(err)
	}

	if err := uuid.Scan([]byte(stringTest)); err != nil {
		t.Fatal(err)
	}

	if err := uuid.Scan(byteTest); err != nil {
		t.Fatal(err)
	}

	// bad type tests

	err := uuid.Scan(badTypeTest)
	if err == nil {
		t.Error("int correctly parsed and shouldn't have")
	}

	if !strings.Contains(err.Error(), "unable to scan type") {
		t.Error("attempting to parse an int returned an incorrect error message")
	}

	// invalid/incomplete uuids

	err = uuid.Scan(invalidTest)
	if err == nil {
		t.Error("invalid uuid was parsed without error")
	}

	if !strings.Contains(err.Error(), "scan uuid error: bad UUID") {
		t.Errorf("attempting to parse an invalid UUID returned an incorrect error message: %s", err)
	}

	err = uuid.Scan(byteTest[:len(byteTest)-2])
	if err == nil {
		t.Error("invalid byte uuid was parsed without error")
	}

	if !strings.Contains(err.Error(), "unable to scan type []byte with length 14 into UUID") {
		t.Errorf("attempting to parse an invalid byte UUID returned an incorrect error message: %s", err)
	}

	// empty tests

	uuid = UUID{}

	var emptySlice []byte

	err = uuid.Scan(emptySlice)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range uuid {
		if v != 0 {
			t.Error("UUID was not nil after scanning empty byte slice")
		}
	}

	if !uuid.Empty() {
		t.Error("UUID was not empty after scanning empty byte slice")
	}

	uuid = UUID{}

	var emptyString string

	err = uuid.Scan(emptyString)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range uuid {
		if v != 0 {
			t.Error("UUID was not nil after scanning empty string")
		}
	}

	if !uuid.Empty() {
		t.Error("UUID was not empty after scanning empty string")
	}

	uuid = UUID{}

	err = uuid.Scan(nil)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range uuid {
		if v != 0 {
			t.Error("UUID was not nil after scanning nil")
		}
	}
}

func TestValue(t *testing.T) {
	stringTest := "0193cf97-66af-7308-9504-44040772cff2"
	uuid := MustParse(stringTest)

	val, _ := uuid.Value()
	if val != stringTest {
		t.Error("Value() did not return expected string")
	}
}
