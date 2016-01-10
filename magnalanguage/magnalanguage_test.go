package magnalanguage_test

import(
  "fmt"
  "testing"
  . "github.com/magna/magnalanguage"
   "github.com/magna/magnagraph"
   "github.com/magna/magnauser"
  "strconv"
)

func TestTokenize(t *testing.T){
  //Tokenize test takes in an array of sentences and then places them in to queryObjects and then tokenizes each query object.
  sentences := []string{"the world is a crazy place we live in.", "What is the fascination with all this stuff out here.", "Making the world a better place will help us all in the end.", "I want to go to far away places to see new things.", "To be trusting is a very daring thing.", "What is this crazy golang thing about?", "Blast it to the max.", "The light side is strong in this one.", "Do you like hockey?", "I will fly away one day to a distant place.", "We have found an ally in you.", "To play these games you will need more controllers", "It is fun to roll around in the grass."}
  nodes := []MagnaQueryObject{}
  for _, aSentence := range sentences{
    aQueryObject := MagnaQueryObject{magnauser.User{"Wallace Mathers", 98239842773}, aSentence}
    nodes = append(nodes, aQueryObject)
  }
  tokenArray := []magnagraph.Node{}
  for _, anode := range nodes {
    tempNode := TokenizeQuery(anode)
    for _, tNode := range tempNode{
      tokenArray = append(tokenArray, tNode)
    }
  }
  fmt.Println(tokenArray)
}
func TestProcessing(t *testing.T){
  //Processing test does the beginning of the tokenize test and then processes all of the nodes.
  sentences := []string{"the world is a crazy place we live in.", "What is the fascination with all this stuff out here.", "Making the world a better place will help us all in the end.", "I want to go to far away places to see new things.", "To be trusting is a very daring thing.", "What is this crazy golang thing about?", "Blast it to the max.", "The light side is strong in this one.", "Do you like hockey?", "I will fly away one day to a distant place.", "We have found an ally in you.", "To play these games you will need more controllers", "It is fun to roll around in the grass."}
  nodes := []MagnaQueryObject{}
  for _, aSentence := range sentences {
    aQueryObject := MagnaQueryObject{magnauser.User{"Wallace Mathers", 98239842773}, aSentence}
    nodes = append(nodes, aQueryObject)
  }
  tokenArray := []magnagraph.Node{}
  for _, anode := range nodes {
    tempNode := TokenizeQuery(anode)
    for _, tNode := range tempNode{
      tokenArray = append(tokenArray, tNode)
    }
  }
  for _, pNode := range tokenArray {
    ProcessNode(&pNode)//Does not return the correct node processed even though it goes through all the steps. Need to fix.
  }
  fmt.Println(tokenArray)
}
