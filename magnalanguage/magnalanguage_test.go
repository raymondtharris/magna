package magnalanguage_test

import(
  "fmt"
  "testing"
  . "github.com/magna/magnalanguage"
  "strconv"
)

func TestPorters(t *testing.T){

  //Load file to read from
  //Read file
  // split words
  //run words through Porters

}
func TestMeasure(t *testing.T){
  words := []string{"the", "world", "farther", "places", "trusting", "higher", "max", "light"}
  for _, aWord := range words {
    meas := FindMeasure(aWord)
    fmt.Println(aWord + " measure: "+ strconv.Itoa(meas))
  }
}
func TestTokenize(t *testing.T){

}
