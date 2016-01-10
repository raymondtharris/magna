package  magnalanguage

import (
	"fmt"
	"github.com/magna/magnauser"
	"strings"

	"regexp"

	"github.com/magna/magnagraph"
	"github.com/magna/magnaporters"
)

type MagnaQueryObject struct {
	User        magnauser.User
	QueryString string
	StemString string
	TokenArray []magnagraph.Node
	ProcessStage int
}

const vowelRegexp = "[aeiou|y]"
const consanantRegexp = "^" + vowelRegexp

func (mqo MagnaQueryObject) String() string {
	return fmt.Sprintf("User: %v, Query: %v\n", mqo.User, mqo.QueryString)
}

//const CommonDict = {"in", "a", "the", "of", "an"}
func ProcessNode(aNode *magnagraph.Node) {
	//fmt.Println(aNode.Value)
	aNode.Measure = FindMeasure(aNode.Value)
	aNode.Stem = Porters(aNode.Value, aNode.Measure)
	fmt.Println(aNode.Stem + "  " )
}

func AppendToStemmedSring(stemString string, aNode magnagraph.Node) {
	stemString = stemString + aNode.Stem + " "
}

func Categorize(aWord string) int {
	return 0
}

func IsImportant(aWord string) bool {
	return false
}

func CreateStemString(aQueryObject MagnaQueryObject) {
	if aQueryObject.ProcessStage > 2 {
		for _, aStem := range aQueryObject.TokenArray {
			aQueryObject.StemString = aQueryObject.StemString + " " + aStem
		}
	}
}

func PartOfSpeachTag(sentenceTokens []magnagraph.Node) bool {
	for _, aWord := range sentenceTokens{
		TagWord(aWord)
	}
	fmt.Println(sentenceTokens)
}

func TagWord(aNode magnagraph.Node) {

}


//TokenizeQuery functions makes an array of Nodes that can be used with Magna's
// graph to build a functional query. The function takes in a QueryObject and then
// splits each word into a token. Those are then pushed into an array to be later
// processed by the AI system.
func TokenizeQuery(queryObject MagnaQueryObject) []magnagraph.Node {
	var tokenArray []magnagraph.Node

	regexSpaces := regexp.MustCompile("[^A-z|']") //Find spaces
	words := regexSpaces.ReplaceAllString(queryObject.QueryString, "\n")
	//fmt.Println(words)
	splitQueryString := strings.Split(words, "\n")
	fmt.Println(splitQueryString)
	for index, aWord := range splitQueryString {
		newNode := magnagraph.Node{index, aWord, 1, nil, -1, ""}
		tokenArray = append(tokenArray, newNode)
	}
	return tokenArray
}

func ProcessQueryObject(queryObject MagnaQueryObject) {
	StemQueryObject(queryObject.TokenArray)
}

func StemQueryObject(tokenArray []magnagraph.Node) {
	for _, aNode := range tokenArray {
		aNode.Stem , aNode.Measure = magnaporters.PortersAlgorithm(aNode.Value)
	}
}

func ExecuteQuery() {

}
