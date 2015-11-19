package magnalanguage

import (
	"fmt"
	"old_magna/magna/magnauser"
	"strings"

	"github.com/magna/magnagraph"
	"regexp"
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

}

func Categorize(aWord string) int {
	return 0
}

func IsImportant(aWord string) bool {
	return false
}

func TokenizeQuery(queryObject MagnaQueryObject) []magnagraph.Node {
	var tokenArray []magnagraph.Node

	regexSpaces := regexp.MustCompile("[^A-z|']") //Find spaces
	words := regexSpaces.ReplaceAllString(queryObject.QueryString, "\n")
	//fmt.Println(words)
	splitQueryString := strings.Split(words, "\n")
	fmt.Println(splitQueryString)
	for index, aWord := range splitQueryString {
		newNode := magnagraph.Node{index, aWord, 1, nil, -1}
		tokenArray = append(tokenArray, newNode)
	}
	return tokenArray
}

func BuildQuery(queryObject MagnaQueryObject) {

}

func ExecuteQuery(){
	
}
