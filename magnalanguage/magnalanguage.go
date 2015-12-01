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
	stem = phaseOne(stem)

	edRegexp := regexp.MustCompile(".*[aeiou].*ed$")
	if len(edRegexp.FindAllString(stem, 1)) > 0 {
		stem := edRegexp.ReplaceAllString(stem, stem[0:len(stem)-2])
		fmt.Println(stem)
	}
	ingRegexp := regexp.MustCompile(".*[aeiou].*ing$")
	if len(ingRegexp.FindAllString(stem, 1)) > 0 {
		stem := ingRegexp.ReplaceAllString(stem, aWord[0:len(stem)-3])
		fmt.Println(stem)
	}
	atRegexp := regexp.MustCompile(".*[aeiou].*at$")
	if len(atRegexp.FindAllString(stem, 1)) > 0 {
		stem := atRegexp.ReplaceAllString(stem, stem+"e")
		fmt.Println(stem)
	}
	blRegexp := regexp.MustCompile(".*[aeiou].*bl$")
	if len(blRegexp.FindAllString(stem, 1)) > 0 {
		stem := blRegexp.ReplaceAllString(stem, stem+"e")
		fmt.Println(stem)
	}
	izRegexp := regexp.MustCompile(".*[aeiou].*iz$")
	if len(izRegexp.FindAllString(stem, 1)) > 0 {
		stem := izRegexp.ReplaceAllString(stem, stem+"e")
		fmt.Println(stem)
	}
	yRegexp := regexp.MustCompile(".*[aeiou].y$")
	if len(yRegexp.FindAllString(stem, 1)) > 0 {
		stem := yRegexp.ReplaceAllString(stem, stem[0:len(stem)-1]+"i")
		fmt.Println(stem)
	}
	ationalRegexp := regexp.MustCompile(".*[aeiou].ational$")
	if len(ationalRegexp.FindAllString(stem, 1)) > 0 {
		stem := ationalRegexp.ReplaceAllString(stem, stem[0:len(stem)-7]+"ate")
		fmt.Println(stem)
	}
	tionalRegexp := regexp.MustCompile(".*[aeiou].tional$")
	if len(tionalRegexp.FindAllString(stem, 1)) > 0 {
		stem := tionalRegexp.ReplaceAllString(stem, stem[0:len(stem)-6]+"tion")
		fmt.Println(stem)
	}

	return stem
}

func phaseOne(aStem string) string {
	ssesRegexp := regexp.MustCompile(".*sses$")
	if len(ssesRegexp.FindAllString(aStem, 1)) > 0 {
		aStem = ssesRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"ss")
		fmt.Println(aStem)
	}
	iesRegexp := regexp.MustCompile(".*ies$")
	if len(iesRegexp.FindAllString(aStem, 1)) > 0 {
		aStem = iesRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3]+"i")
		fmt.Println(aStem)
	}
	ssRegexp := regexp.MustCompile(".*ss$")
	if len(ssRegexp.FindAllString(aStem, 1)) > 1 {
		aStem = ssRegexp.ReplaceAllString(aStem, aStem)
		fmt.Println(aStem)
	}
	sRegexp := regexp.MustCompile(".*s$")
	if len(sRegexp.FindAllString(aStem, 1)) > 1 {
		aStem = sRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-1])
		fmt.Println(aStem)
	}
	return aStem
}
func phaseTwo(aStem string) string {
	return aStem
}
func phaseThree(aStem string) string {
	return aStem
}
func phaseFour(aStem string) string {
	return aStem
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
