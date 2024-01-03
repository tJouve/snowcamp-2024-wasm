package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func binaryRepresentationOf(number interface{}) string {

	switch number.(type) {
	case uint32:
		return strconv.FormatInt(int64(number.(uint32)), 2)
	case uint64:
		return strconv.FormatInt(int64(number.(uint64)), 2)
	default:
		return "🤭 oups"
	}

}

func putTwoValuesIntoOnlyOne(position, length uint32) uint64 {
	fmt.Println("------------------------------------------------------------------")
	fmt.Println(" 🤓 How to put", position, "(Position) and", length, "(Length) into only one value:")
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("")

	fmt.Println("Binary representation of Position and Length:")
	fmt.Println("  position:", position, "->", binaryRepresentationOf(position))
	fmt.Println("  length  :", length, "->", binaryRepresentationOf(length))
	fmt.Println("")

	/*
		## Put the `position` of the left of 32 zero bits

		**Shifts left** by pushing zeros in from the right
		and let the leftmost bits fall off

		position << 32n
	*/
	positionShiftLeft := uint64(position) << uint64(32)
	
	fmt.Println("Position (", binaryRepresentationOf(position), ") + 32 0 bits: put the position on the left")
	fmt.Println("  position:       ", positionShiftLeft, "->", binaryRepresentationOf(positionShiftLeft))
	fmt.Println("")

	/*
		## Put the `length` at right (replace the last 32 zero bits of the right)

		Sets each bit to 1 if one of two bits is 1 with a **Or**

		position << 32n | length

	*/
	// (position, length) = 42949672980 = 101000000000000000000000000000010100 🤔
	positionShiftLeftOrLength := uint64(position) << uint64(32) | uint64(length)


	fmt.Println("Position (", binaryRepresentationOf(position), ") on the left and Length(", binaryRepresentationOf(length), ") on the right into the same place")
	fmt.Println("  position_length:", positionShiftLeftOrLength, "->", binaryRepresentationOf(positionShiftLeftOrLength))
	fmt.Println("")

	return positionShiftLeftOrLength

}

func extractTwoValuesFromOnlyOne(positionAndLength uint64) {
	fmt.Println("------------------------------------------------------------------")
	fmt.Println(" 🤓 Position and Length are embedded in this value: ", positionAndLength)
	fmt.Println(" 🤓 How to extract Position and Length from", positionAndLength, "?")
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("")

	fmt.Println("This is the binary representation of", positionAndLength, ":", binaryRepresentationOf(positionAndLength))
	fmt.Println("")

	/*
		Extract the position
		Operation: Shift Right
		positionAndLength >> 32b

	*/

	positionAndLengthShiftRight := uint32(positionAndLength >> uint64(32))
	fmt.Println("1- Extract the Position by removing the 32 bits on the right from", binaryRepresentationOf(positionAndLength), ":")

	fmt.Println("    Apply Shift Right(positionAndLength >> 32b):")
	fmt.Println("      Binary Position=", binaryRepresentationOf(positionAndLengthShiftRight))
	fmt.Println("      Decimal Position=", positionAndLengthShiftRight, "🎉")
	fmt.Println("")

	/*
		Extract the length
		Operation: And Mask

		positionAndLength & mask

		Extract the lowest 32 bits, using the mask 11111111111111111111111111111111 (32 "1")
		& (AND) sets each bit to 1 if both bits are 1

	*/
	stringMask := "11111111111111111111111111111111"
	var bigInt big.Int
	mask, _ := bigInt.SetString(stringMask, 2)

	fmt.Println("2- Extract the Length by applying a mask on", positionAndLength, "(positionAndLength):")
	fmt.Println("    Binary representation of positionAndLength:", binaryRepresentationOf(positionAndLength))
	fmt.Println("    Binary representation of the Mask         :    ", stringMask)
	fmt.Println("")
	fmt.Println("    Operation: ", binaryRepresentationOf(positionAndLength), "&", stringMask)

	fmt.Println("    Starting from the right,")
	fmt.Println("      If a bit of Mask == 1 and a bit of Value == 1")
	fmt.Println("      Then the result is == 1")
	fmt.Println("      Else the result is == the value of the bit of Value")
	fmt.Println("")

	positionAndLengthAndMask := uint32(positionAndLength & mask.Uint64())

	fmt.Println("    Apply And Mask:")
	fmt.Println("      Binary Length=", binaryRepresentationOf(positionAndLengthAndMask))
	fmt.Println("      Decimal Length=", positionAndLengthAndMask, "🎉")
	fmt.Println("")

}

func main() {
	kindOfTupleWithPositionAndLength := putTwoValuesIntoOnlyOne(10, 20)

	extractTwoValuesFromOnlyOne(kindOfTupleWithPositionAndLength)

}
