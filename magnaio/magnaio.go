package magnaio
import (
  //"fmt"
  "bufio"
  "os"
)

func ReadInput() string {
	consoleReader := bufio.NewReader(os.Stdin)
	rawInput, _ := consoleReader.ReadString('\n')
	inputValue := rawInput[0 : len(rawInput)-1]
	return inputValue
}

func ErrorDisplay() string {
  return ""
}
