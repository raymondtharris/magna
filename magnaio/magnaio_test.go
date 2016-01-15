package magnaio_test

import (
	"log"
	"testing"
  . "github.com/magna/magnaio"
)

func TestReadInput(t *testing.T) {
	inputValues := []string{"Jimmy", "what", "calendar", "maxxie", "who", "amazing things are happening", "fire", "The Walking Dead", "Making cool things", "Lioness"}
  i := 0
  for i < 10 {
    temp := ReadInput()
    _ = temp
     //inputValues = append(inputValues, temp)
  }
  log.Println(inputValues)
}
