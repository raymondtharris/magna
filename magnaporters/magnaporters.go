package magnaporters

import (
  "fmt"
)

//FindMeasure function will calculate the measure of a word defined by Porters algorithm.
//The function takes in a string and will return the integer value of the meaure.
func FindMeasure(aWord string) int {
	measureRegexp := regexp.MustCompile("[aeiou][^aeiou]")
	aMeasure := len(measureRegexp.FindAllString(aWord, -1))
	return aMeasure
}
