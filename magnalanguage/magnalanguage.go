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

const vowelRegexp = "[aeiou|y]"
const consanantRegexp = "^" + vowelRegexp

func (mqo MagnaQueryObject) String() string {
	return fmt.Sprintf("User: %v, Query: %v", mqo.User, mqo.QueryString)
}

//const CommonDict = {"in", "a", "the", "of", "an"}
func ProcessNode(aNode magnagraph.Node) {
	//fmt.Println(aNode.Value)
	aNode.Measure = FindMeasure(aNode.Value)
	aNode.Stem = Porters(aNode.Value)
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

func Porters(aWord string) string {
	stem := aWord
	stem = phaseOne(stem)
	stem = phaseTwo(stem)
	stem = phaseThree(stem)
	stem = phaseFour(stem)

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
	if len(ssRegexp.FindAllString(aStem, 1)) > 0 {
		aStem = ssRegexp.ReplaceAllString(aStem, aStem)
		fmt.Println(aStem)
	}
	sRegexp := regexp.MustCompile(".*s$")
	if len(sRegexp.FindAllString(aStem, 1)) > 0 {
		aStem = sRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-1])
		fmt.Println(aStem)
	}
	p1aRegexp := regexp.MustCompile(consanantRegexp + ".*" + vowelRegexp + "eed$")
	if len(p1aRegexp.FindAllString(aStem, 1)) > 0 {
		aStem = p1aRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3]+"ee")
		fmt.Println(aStem)
	}
	p1bRegexp := regexp.MustCompile(".*" + vowelRegexp + "(ed$|ing$)")
	if len(p1bRegexp.FindAllString(aStem, 1)) > 0 {
		edRegexp := regexp.MustCompile(".*[aeiou].*ed$")
		if len(edRegexp.FindAllString(aStem, 1)) > 0 {
			aStem := edRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-2])
			fmt.Println(aStem)
		}
		ingRegexp := regexp.MustCompile(".*[aeiou].*ing$")
		if len(ingRegexp.FindAllString(aStem, 1)) > 0 {
			aStem := ingRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
			fmt.Println(aStem)
		}

		atRegexp := regexp.MustCompile(".*[aeiou].*at$")
		if len(atRegexp.FindAllString(aStem, 1)) > 0 {
			aStem := atRegexp.ReplaceAllString(aStem, aStem+"e")
			fmt.Println(aStem)
		}
		blRegexp := regexp.MustCompile(".*[aeiou].*bl$")
		if len(blRegexp.FindAllString(aStem, 1)) > 0 {
			aStem := blRegexp.ReplaceAllString(aStem, aStem+"e")
			fmt.Println(aStem)
		}
		izRegexp := regexp.MustCompile(".*[aeiou].*iz$")
		if len(izRegexp.FindAllString(aStem, 1)) > 0 {
			aStem := izRegexp.ReplaceAllString(aStem, aStem+"e")
			fmt.Println(aStem)
		}

	}

	yRegexp := regexp.MustCompile(".*[aeiou].y$")
	if len(yRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := yRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-1]+"i")
		fmt.Println(aStem)
	}

	return aStem
}
func phaseTwo(aStem string) string {
	ationalRegexp := regexp.MustCompile(".*[aeiou].ational$")
	if len(ationalRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ationalRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-7]+"ate")
		fmt.Println(aStem)
	}
	tionalRegexp := regexp.MustCompile(".*[aeiou].tional$")
	if len(tionalRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := tionalRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-6]+"tion")
		fmt.Println(aStem)
	}
	enciRegexp := regexp.MustCompile(".*[aeiou].enci$")
	if len(enciRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := enciRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"ence")
		fmt.Println(aStem)
	}
	anciRegexp := regexp.MustCompile(".*[aeiou].anci$")
	if len(anciRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := anciRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"ance")
		fmt.Println(aStem)
	}
	izerRegexp := regexp.MustCompile(".*[aeiou].izer$")
	if len(izerRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := izerRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"ize")
		fmt.Println(aStem)
	}
	abliRegexp := regexp.MustCompile(".*[aeiou].abli$")
	if len(abliRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := abliRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"able")
		fmt.Println(aStem)
	}
	alliRegexp := regexp.MustCompile(".*[aeiou].alli$")
	if len(alliRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := alliRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"al")
		fmt.Println(aStem)
	}
	entliRegexp := regexp.MustCompile(".*[aeiou].entli$")
	if len(entliRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := entliRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"ent")
		fmt.Println(aStem)
	}
	eliRegexp := regexp.MustCompile(".*[aeiou].eli$")
	if len(eliRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := eliRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3]+"e")
		fmt.Println(aStem)
	}
	ousliRegexp := regexp.MustCompile(".*[aeiou].ousli$")
	if len(ousliRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ousliRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"ous")
		fmt.Println(aStem)
	}
	izationRegexp := regexp.MustCompile(".*[aeiou].ization$")
	if len(izationRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := izationRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-7]+"ize")
		fmt.Println(aStem)
	}
	ationRegexp := regexp.MustCompile(".*[aeiou].ation$")
	if len(ationRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ationRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"ate")
		fmt.Println(aStem)
	}
	atorRegexp := regexp.MustCompile(".*[aeiou].ator$")
	if len(atorRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := atorRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"ate")
		fmt.Println(aStem)
	}
	alismRegexp := regexp.MustCompile(".*[aeiou].alism$")
	if len(alismRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := alismRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"al")
		fmt.Println(aStem)
	}
	ivenessRegexp := regexp.MustCompile(".*[aeiou].iveness$")
	if len(ivenessRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ivenessRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-7]+"ive")
		fmt.Println(aStem)
	}
	fulnessRegexp := regexp.MustCompile(".*[aeiou].fulness$")
	if len(fulnessRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := fulnessRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-7]+"ful")
		fmt.Println(aStem)
	}
	ousnessRegexp := regexp.MustCompile(".*[aeiou].ousness$")
	if len(ousnessRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ousnessRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-7]+"ous")
		fmt.Println(aStem)
	}
	alitiRegexp := regexp.MustCompile(".*[aeiou].aliti$")
	if len(alitiRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := alitiRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"al")
		fmt.Println(aStem)
	}
	ivitiRegexp := regexp.MustCompile(".*[aeiou].iviti$")
	if len(ivitiRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ivitiRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"ive")
		fmt.Println(aStem)
	}
	bilitiRegexp := regexp.MustCompile(".*[aeiou].biliti$")
	if len(bilitiRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := bilitiRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-6]+"ble")
		fmt.Println(aStem)
	}

	return aStem
}
func phaseThree(aStem string) string {
	icateRegexp := regexp.MustCompile(".*[aeiou].icate$")
	if len(icateRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := icateRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"ic")
		fmt.Println(aStem)
	}
	ativeRegexp := regexp.MustCompile(".*[aeiou].ative$")
	if len(ativeRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ativeRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5])
		fmt.Println(aStem)
	}
	alizeRegexp := regexp.MustCompile(".*[aeiou].alize$")
	if len(alizeRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := alizeRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"al")
		fmt.Println(aStem)
	}
	icitiRegexp := regexp.MustCompile(".*[aeiou].iciti$")
	if len(icitiRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := icitiRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5]+"ic")
		fmt.Println(aStem)
	}
	icalRegexp := regexp.MustCompile(".*[aeiou].ical$")
	if len(icalRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := icalRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4]+"ic")
		fmt.Println(aStem)
	}
	fulRegexp := regexp.MustCompile(".*[aeiou].ful$")
	if len(fulRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := fulRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		fmt.Println(aStem)
	}
	nessRegexp := regexp.MustCompile(".*[aeiou].ness$")
	if len(nessRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := nessRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		fmt.Println(aStem)
	}

	return aStem
}
func phaseFour(aStem string) string {
	alRegexp := regexp.MustCompile(".*[aeiou].al$")
	if len(alRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := alRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-2])
		fmt.Println(aStem)
	}
	anceRegexp := regexp.MustCompile(".*[aeiou].ance$")
	if len(anceRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := anceRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		fmt.Println(aStem)
	}
	enceRegexp := regexp.MustCompile(".*[aeiou].ence$")
	if len(enceRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := enceRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		fmt.Println(aStem)
	}
	erRegexp := regexp.MustCompile(".*[aeiou].er$")
	if len(erRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := erRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-2])
		fmt.Println(aStem)
	}
	icRegexp := regexp.MustCompile(".*[aeiou].ic$")
	if len(icRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := icRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-2])
		fmt.Println(aStem)
	}
	ableRegexp := regexp.MustCompile(".*[aeiou].able$")
	if len(ableRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ableRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		fmt.Println(aStem)
	}
	ibleRegexp := regexp.MustCompile(".*[aeiou].ible$")
	if len(ibleRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ibleRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		fmt.Println(aStem)
	}
	antRegexp := regexp.MustCompile(".*[aeiou].ant$")
	if len(antRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := antRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		fmt.Println(aStem)
	}
	ementRegexp := regexp.MustCompile(".*[aeiou].ement$")
	if len(ementRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ementRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-5])
		fmt.Println(aStem)
	}
	mentRegexp := regexp.MustCompile(".*[aeiou].ment$")
	if len(mentRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := mentRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		fmt.Println(aStem)
	}
	entRegexp := regexp.MustCompile(".*[aeiou].ent$")
	if len(entRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := entRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		fmt.Println(aStem)
	}

	s_tionRegexp := regexp.MustCompile(".*[aeiou].[s|t]ion$")
	if len(s_tionRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := s_tionRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-4])
		fmt.Println(aStem)
	}
	ouRegexp := regexp.MustCompile(".*[aeiou].ou$")
	if len(ouRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ouRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-2])
		fmt.Println(aStem)
	}
	ismRegexp := regexp.MustCompile(".*[aeiou].ism$")
	if len(ismRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ismRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		fmt.Println(aStem)
	}
	ateRegexp := regexp.MustCompile(".*[aeiou].ate$")
	if len(ateRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ateRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		fmt.Println(aStem)
	}
	itiRegexp := regexp.MustCompile(".*[aeiou].iti$")
	if len(itiRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := itiRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		fmt.Println(aStem)
	}
	ousRegexp := regexp.MustCompile(".*[aeiou].ous$")
	if len(ousRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := ousRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		fmt.Println(aStem)
	}
	iveRegexp := regexp.MustCompile(".*[aeiou].ive$")
	if len(iveRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := iveRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		fmt.Println(aStem)
	}
	izeRegexp := regexp.MustCompile(".*[aeiou].ize$")
	if len(izeRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := izeRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-3])
		fmt.Println(aStem)
	}

	return aStem
}
func phaseFive(aStem string) string {
	eRegexp := regexp.MustCompile(".*[aeiou].e$")
	if len(eRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := eRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-1])
		fmt.Println(aStem)
	}
	lRegexp := regexp.MustCompile(".*[aeiou].l$")
	if len(lRegexp.FindAllString(aStem, 1)) > 0 {
		aStem := lRegexp.ReplaceAllString(aStem, aStem[0:len(aStem)-1])
		fmt.Println(aStem)
	}

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
