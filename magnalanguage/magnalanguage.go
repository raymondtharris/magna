package magnalanguage

import (
	"fmt"
	"old_magna/magna/magnauser"
	"strings"

	"regexp"

	"github.com/magna/magnagraph"
)

type MagnaQueryObject struct {
	User        magnauser.User
	QueryString string
}

func (mqo MagnaQueryObject) String() string {
	return fmt.Sprintf("User: %v, Query: %v", mqo.User, mqo.QueryString)
}

//const CommonDict = {"in", "a", "the", "of", "an"}
func ProcessNode(aNode magnagraph.Node) {
	//fmt.Println(aNode.Value)
	aNode.Measure = FindMeasure(aNode.Value)
	aNode.Stem = FindStem(aNode.Value)
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

func FindMeasure(aWord string) int {
	return 0
}

func FindStem(aWord string) string {
	//Find ing stem
	ingRegexp := regexp.MustCompile(".*[aeiou].*ing$")
	if len(ingRegexp.FindAllString(aWord, 1)) > 0 {
		stem := ingRegexp.ReplaceAllString(aWord, aWord[0:len(aWord)-3])
		fmt.Println(stem)
		FindStem(stem)
	}
	edRegexp := regexp.MustCompile(".*[aeiou].*ed$")
	if len(edRegexp.FindAllString(aWord, 1)) > 0 {
		stem := edRegexp.ReplaceAllString(aWord, aWord[0:len(aWord)-2])
		fmt.Println(stem)
		FindStem(stem)
	}
	lyRegexp := regexp.MustCompile(".*[aeiou].*ly$")
	if len(lyRegexp.FindAllString(aWord, 1)) > 0 {
		stem := lyRegexp.ReplaceAllString(aWord, aWord[0:len(aWord)-2])
		fmt.Println(stem)
		FindStem(stem)
	}
	ssesRegexp := regexp.MustCompile(".*sses$")
	if len(ssesRegexp.FindAllString(aWord, 1)) > 0 {
		stem := ssesRegexp.ReplaceAllString(aWord, aWord[0:len(aWord)-4]+"ss")
		fmt.Println(stem)
		FindStem(stem)
	}
	ssRegexp := regexp.MustCompile(".*ss$")
	if len(ssRegexp.FindAllString(aWord, 1)) > 1 {
		stem := ssRegexp.ReplaceAllString(aWord, aWord[0:len(aWord)-2])
		fmt.Println(stem)
		FindStem(stem)
	}

	sRegexp := regexp.MustCompile(".*s$")
	_ = sRegexp
	esRegexp := regexp.MustCompile(".*es$")
	_ = esRegexp
	return aWord
}

func Porters(aWord string) string {
	stem := aWord
	ssesRegexp := regexp.MustCompile(".*sses$")
	if len(ssesRegexp.FindAllString(stem, 1)) > 0 {
		stem = ssesRegexp.ReplaceAllString(stem, stem[0:len(stem)-4]+"ss")
		fmt.Println(stem)
	}
	iesRegexp := regexp.MustCompile(".*ies$")
	if len(iesRegexp.FindAllString(stem, 1)) > 0 {
		stem = iesRegexp.ReplaceAllString(stem, stem[0:len(stem)-3]+"i")
		fmt.Println(stem)
	}
	ssRegexp := regexp.MustCompile(".*ss$")
	if len(ssRegexp.FindAllString(stem, 1)) > 1 {
		stem = ssRegexp.ReplaceAllString(stem, stem)
		fmt.Println(stem)
	}
	sRegexp := regexp.MustCompile(".*s$")
	if len(sRegexp.FindAllString(stem, 1)) > 1 {
		stem = sRegexp.ReplaceAllString(stem, stem[0:len(stem)-1])
		fmt.Println(stem)
	}

	return stem
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

func BuildQuery(queryObject MagnaQueryObject) {

}

func ExecuteQuery() {

}
