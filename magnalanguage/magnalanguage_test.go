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
  words := []string{"the", "world", "farther", "places", "trusting", "higher", "max", "light", "hockey", "fly", "ally", "controll", "roll"}
  for _, aWord := range words {
    meas := FindMeasure(aWord)
    stem := Porters(aWord, meas)
    fmt.Println("word " + aWord + " stem " + stem + " measure " + strconv.Itoa(meas))
  }
}
func TestMeasure(t *testing.T){
  words := []string{"the", "world", "farther", "places", "trusting", "higher", "max", "light"}
  for _, aWord := range words {
    meas := FindMeasure(aWord)
    fmt.Println(aWord + " measure: "+ strconv.Itoa(meas))
  }
}
func TestTokenize(t *testing.T){
  sentences := []string{"the world is a crazy place we live in.", "What is the fascination with all this stuff out here.", "Making the world a better place will help us all in the end.", "I want to go to far away places to see new things.", "To be trusting is a very daring thing.", "What is this crazy golang thing about?", "Blast it to the max.", "The light side is strong in this one.", "Do you like hockey?", "I will fly away one day to a distant place.", "We have found an ally in you.", "To play these games you will need more controllers", "It is fun to roll around in the grass."}
  nodes := []Node
  for _, aSentence := range words {
    aQueryObject := MagnaQueryObject{magnauser.User{"Wallace Mathers", 98239842773}, aSentence}
    nodes = append(nodes, aQueryObject)
  }
  fmt.Println(nodes)
  tokenArray := []magnagraph.Node{}
  for _, anode := range nodes {
    tempNode := TokenizeQuery(anode)
    tokenArray = append(tokenArray, tempNode)
  }
  fmt.Println(tokenArray)
}
func TestProcessing(t *testing.T){
  sentences := []string{"the world is a crazy place we live in.", "What is the fascination with all this stuff out here.", "Making the world a better place will help us all in the end.", "I want to go to far away places to see new things.", "To be trusting is a very daring thing.", "What is this crazy golang thing about?", "Blast it to the max.", "The light side is strong in this one.", "Do you like hockey?", "I will fly away one day to a distant place.", "We have found an ally in you.", "To play these games you will need more controllers", "It is fun to roll around in the grass."}
  nodes := []Node
  for _, aSentence := range words {
    aQueryObject := MagnaQueryObject{magnauser.User{"Wallace Mathers", 98239842773}, aSentence}
    nodes = append(nodes, aQueryObject)
  }
  fmt.Println(nodes)
  tokenArray := []magnagraph.Node{}
  for _, anode := range nodes {
    tempNode := TokenizeQuery(anode)
    tokenArray = append(tokenArray, tempNode)
  }
  fmt.Println(tokenArray)
  for _, pNode := range tokenArray {
    ProcessNode(ProcessNode)
  }
  fmt.Println(tokenArray)
}
