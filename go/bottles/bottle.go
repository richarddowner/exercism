package bottles

import (
// "fmt"
// "strconv"
// "strings"
)

// number > 1: {8} {bottles} of beer on the wall, {8} {bottles} of beer.\nTake one down and pass it around, {7} bottles of beer on the wall.\n
// number = 1: {1} {bottle} of beer on the wall, {1} {bottle} of beer.\nTake it down and pass it around, {no more} bottles of beer on the wall.\n
// number = 0: {No more} {bottles} of beer on the wall, {no more} {bottles} of beer.\nGo to the store and buy some more, {99} bottles of beer on the wall.\n

// const words = "{n} {bottles} of beer on the wall, {n} {bottles} of beer.\n{task}, {n-1} bottles of beer on the wall.\n"

var bottles int = 99

func Verse(number int) string {

	if number > 1 {
		fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", bottles, bottles)
		bottles--
		fmt.Println("Take one down and pass it around, %d bottles of beer on the wall.\n", bottles)
		Verse(bottles)
	}
	if number == 1 {
		fmt.Println("1 bottle of beer on the wall, 1 bottle of beer.")
		fmt.Println("Take one down and pass it around, no more bottles of beer on the wall.\n")
		Verse(bottles)
	} else {
		fmt.Println("No more bottles of beer on the wall, no more bottles of beer.")
		fmt.Println("Go to the store and buy some more, 99 bottles of beer on the wall.")
		return
	}
}
