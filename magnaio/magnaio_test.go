package main

import (
	"fmt"
	"testing"
)

func TestReadInput(t *testing.T) {
	inputValues := []string{}
  i := 0
  for i<10 {
	   temp := ReadInput()
     inputValues = append(inputValues, temp)
  }
  fmt.Println(inputValues)
}
