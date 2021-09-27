package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Dolore duis velit magna eu sunt et excepteur cupidatat ullamco non ex Aliquip amet consequat enim enim occaecat quis fugiat officia quis aliquip pariatur Consequat cillum eiusmod proident culpa duis dolor incididunt id occaecat ex consectetur id dolor Mollit fugiat sit irure do cupidatat officia deserunt laboris ametElit ut dolore incididunt irure elit consectetur Quis exercitation reprehenderit nostrud occaecat nisi ipsum nulla cillum quis labore tempor minim magna ullamco Et nostrud nostrud enim sunt esse excepteur exercitation ad officia sint aliqua exercitation doCillum Lorem reprehenderit minim minim Voluptate magna ipsum incididunt voluptate cillum enim Aute cupidatat fugiat tempor sint labore dolore dolore consectetur id anim ea voluptate Occaecat dolore do aute nulla reprehenderit Minim magna aliquip magna commodo dolore culpa sint nostrud sit non fugiat Eu sunt incididunt deserunt enim adipisicing nulla quis nostrud culpa dolore Mollit pariatur id velit qui duis irureId ea minim commodo labore ullamco proident laborum ad adipisicing Elit ullamco officia culpa magna amet voluptate dolor Occaecat pariatur ea consectetur sunt anim et Lorem in tempor labore velit pariatur velit Minim qui aliquip ad aliquip id nulla labore dolore nulla culpa culpa Irure laborum mollit veniam dolore Lorem Commodo aliqua esse incididunt sit in aliquip incididunt deserunt"

	//find out the length of the words that "occurs the most by length" and print the most occuring "word length" and the number of occurrences

	//use strings.Split() to split the string into words
	words := strings.Split(str, " ")
	wordsCountBySize := getWordsCountBySize(words)
	maxOccurance, maxOccuranceSize := getMaxOccuranceBySize(wordsCountBySize)
	fmt.Println(maxOccurance, maxOccuranceSize)
}

func getWordsCountBySize(words []string) map[int]int {
	//create a map to store the word count by size
	wordCountBySize := make(map[int]int)

	//iterate over the words
	for _, word := range words {
		//get the length of the word
		wordLength := len(word)

		//increment the word count by size
		wordCountBySize[wordLength]++
	}

	//return the word count by size
	return wordCountBySize
}

func getMaxOccuranceBySize(wordsCountBySize map[int]int) (int, int) {
	//create a variable to store the max occurance
	var maxOccurance int

	//create a variable to store the max occurance size
	var maxOccuranceSize int

	//iterate over the word count by size
	for size, count := range wordsCountBySize {
		//if the current occurance is greater than the max occurance
		if count > maxOccurance {
			//set the max occurance to the current occurance
			maxOccurance = count

			//set the max occurance size to the current occurance size
			maxOccuranceSize = size
		}
	}

	//return the max occurance and the max occurance size
	return maxOccurance, maxOccuranceSize
}
