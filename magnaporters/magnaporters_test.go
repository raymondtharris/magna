package magnaporters_test

import (
  "log"
  . "github.com/magna/magnaporters"
  "testing"
)

func TestMeasure(t *testing.T){
  //Measure Test uses an array of words and calculates the measure of each word.
  words := []string{"the", "world", "farther", "places", "trusting", "higher", "max", "light"}
  for _, aWord := range words {
    meas := FindMeasure(aWord)
    log.Println(aWord + " measure: "+ strconv.Itoa(meas))
  }
}

func TestPortersAlgorithm(t *testing.T){
    words := []string{"the", "world", "farther", "places", "trusting", "higher", "max", "light"}
    for _, aWord := range words {
      aStem, aMeasure := PortersAlgorithm(aWord)
      log.Println(aWord + " measure: "+ strconv.Itoa(aMeasure) + " stem: " +aStem)
    }
}
