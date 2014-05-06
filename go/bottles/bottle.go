package bottles

import (
	"fmt"
	"strconv"
	"strings"
)

// number > 1: {8} {bottles} of beer on the wall, {8} {bottles} of beer.\nTake one down and pass it around, {7} bottles of beer on the wall.\n
// number = 1: {1} {bottle} of beer on the wall, {1} {bottle} of beer.\nTake it down and pass it around, {no more} bottles of beer on the wall.\n
// number = 0: {No more} {bottles} of beer on the wall, {no more} {bottles} of beer.\nGo to the store and buy some more, {99} bottles of beer on the wall.\n

// const words = "{n} {bottles} of beer on the wall, {n} {bottles} of beer.\n{task}, {n-1} bottles of beer on the wall.\n"

func Verse(number int) string {

	n := ""
	bottles := ""
	task := ""

	if number > 1 {
		n = strconv.Itoa(number)
		bottles = "bottles"
		task = "Take one down and pass it around,"
	}
	if number == 1 {
		n = strconv.Itoa(number)
		bottles = "bottle"
		task = "Take it down and pass it around,"
	}
	if number == 0 {
		n = "no more"
		bottles = "bottles"
		task = "Go to the store and buy some more,"
	}

	words := n + " " + bottles + " of beer on the wall, " + n + " " + bottles + " of beer.\n" + task + " "

	number = number - 1

	if number > 1 {
		n = strconv.Itoa(number)
		bottles = "bottles"
		task = "Take one down and pass it around,"
	}
	if number == 1 {
		n = strconv.Itoa(number)
		bottles = "bottle"
		task = "Take it down and pass it around,"
	}
	if number == 0 {
		n = "no more"
		bottles = "bottles"
		task = "Go to the store and buy some more,"
	}

	words = words + n + " " + bottles + " of beer on the wall.\n"

	return capitalize(words)
}

func capitalize(words string) (output string) {
	output = strings.ToUpper(string(words[0]))
	output = output + words[1:]
	return
}
