package bottles

import (
	//"fmt"
	"strconv"

// "strings"
)

// number > 1: {8} {bottles} of beer on the wall, {8} {bottles} of beer.\nTake one down and pass it around, {7} bottles of beer on the wall.\n
// number = 1: {1} {bottle} of beer on the wall, {1} {bottle} of beer.\nTake it down and pass it around, {no more} bottles of beer on the wall.\n
// number = 0: {No more} {bottles} of beer on the wall, {no more} {bottles} of beer.\nGo to the store and buy some more, {99} bottles of beer on the wall.\n

// const words = "{n} {bottles} of beer on the wall, {n} {bottles} of beer.\n{task}, {n-1} bottles of beer on the wall.\n"

var bottles int = 99

func Verse(number int) (output string) {

	if number > 1 {
		// fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", bottles, bottles)
		output += strconv.Itoa(bottles) + " bottles of beer on the wall, " + strconv.Itoa(bottles) + " bottles of beer.\n"
		bottles--
		//fmt.Printf("Take one down and pass it around, %d bottles of beer on the wall.\n", bottles)
		output += "Take one down and pass it around, " + strconv.Itoa(bottles) + " bottles of beer on the wall.\n"
		Verse(bottles)
	}
	if number == 1 {
		// fmt.Println("1 bottle of beer on the wall, 1 bottle of beer.")
		output += "1 bottle of beer on the wall, 1 bottle of beer.\n"
		// fmt.Println("Take one down and pass it around, no more bottles of beer on the wall.\n")
		output += "Take one down and pass it around, no more bottles of beer on the wall.\n"
		bottles--
		Verse(bottles)
	} else {
		// fmt.Println("No more bottles of beer on the wall, no more bottles of beer.")
		output += "No more bottles of beer on the wall, no more bottles of beer.\n"
		// fmt.Println("Go to the store and buy some more, 99 bottles of beer on the wall.")
		output += "Go to the store and buy some more, 99 bottles of beer on the wall.\n"
		return
	}
	return
}
