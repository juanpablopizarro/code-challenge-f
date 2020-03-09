package parser

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	// LengthOffset represents the amount of characteres of Length field.
	LengthOffset = 2

	// TypeOffset represents the amount of characteres of Type field.
	TypeOffset = 1

	// TypeNumOffset represents the amount of characters of Type Number field.
	TypeNumOffset = 2

	// MinOffset represents the amount of characters used to define the tlv into the string
	MinOffset = TypeNumOffset + TypeOffset + LengthOffset

	//Alphanumeric type descriptor
	Alphanumeric = 'A'

	//Number type descriptor
	Number = 'N'
)

// Tlv encapsulates the .. I don't know what TLV represents :)
type Tlv struct {
	Length  int
	Value   interface{}
	Type    rune
	TypeNum int
}

// Unmarshal parses a given byte array using tlv format
func Unmarshal(input []byte) (map[string]string, error) {
	if input == nil || len(input) < MinOffset {
		return nil, errors.New("invalid unmarshal's input")
	}

	result := make(map[string]string)

	for i := 0; len(input) > 5; i++ {
		// read value's length
		length, err := readLength(input)
		if err != nil {
			return nil, err
		}
		// remove the length from input
		input = input[LengthOffset:]

		// read the value
		strValue, err := readValueStr(input, length)
		if err != nil {
			return nil, err
		}
		// removing the value from input
		input = input[length:]

		// reading the type
		t, err := readType(input)
		if err != nil {
			return nil, err
		}
		// removes the type from input
		input = input[TypeOffset:]

		ok := checkValueType(strValue, t)
		if !ok {
			return nil, errors.New("invalid value")
		}

		// reading the type number field
		tNum, err := readTypeNum(input)
		if err != nil {
			return nil, err
		}
		// removing the type number from input
		input = input[TypeNumOffset:]

		tlv := Tlv{
			Length:  length,
			Value:   strValue,
			Type:    t,
			TypeNum: tNum,
		}

		// I convert the values to string because the challenge wants a map[string]string as an output
		result[strconv.Itoa(i)] = tlvToStr(tlv)
	}
	return result, nil
}

// checkValueType compares the type and check if the value is valid
func checkValueType(s string, t rune) bool {
	if t == Number {
		_, err := strconv.Atoi(s)
		return err == nil
	}
	return true
}

// tlvToStr just format the struct to a string because the challenge enunciate.
func tlvToStr(tlv Tlv) string {
	return fmt.Sprintf("{%v %v %c %v}", tlv.Length, tlv.Value, tlv.Type, tlv.TypeNum)
}

// readLength reads the length value
func readLength(input []byte) (int, error) {
	if input == nil || len(input) < LengthOffset {
		return 0, errors.New("error reading length: invalid length")
	}

	v, err := strconv.Atoi(string(input[:LengthOffset]))
	// convert the first two bytes to string and then to int.
	return v, err
}

// readValue read's the given length from input
func readValueStr(input []byte, length int) (string, error) {
	if input == nil || len(input) < length {
		return "", errors.New("error reading the value of given input")
	}

	return string(input[:length]), nil
}

// readType return the type as a rune
func readType(input []byte) (rune, error) {
	if input == nil || len(input) == 0 {
		return 0, errors.New("error reading the type of given input")
	}

	v := input[0]
	if v == byte(Alphanumeric) {
		return Alphanumeric, nil
	} else if v == byte(Number) {
		return Number, nil
	} else {
		fmt.Printf("type=%c\n", v)
		fmt.Printf("v == byte(Alphanumeric): %v\n", v == byte(Alphanumeric))
		fmt.Printf("v == byte(Number): %v\n", v == byte(Number))
		return 0, errors.New("invalid type")
	}
}

// readTypeNum reads and parse the typeNum
func readTypeNum(input []byte) (int, error) {
	if input == nil || len(input) < TypeNumOffset {
		return 0, errors.New("error reading the type number of given input")
	}

	return strconv.Atoi(string(input[:TypeNumOffset]))
}
