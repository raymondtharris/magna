package magnaio
import (
  //"fmt"
  "bufio"
  "os"
)

const (
  None = 0
  Light = 1
  Normal = 2
  Verbose = 3
)

type Logger struct{
  Verbose int
  LoggingOn bool
}


func ReadInput() string {
	consoleReader := bufio.NewReader(os.Stdin)
  //Add code for Logger
	rawInput, _ := consoleReader.ReadString('\n')
	inputValue := rawInput[0 : len(rawInput)-1]
	return inputValue
}

func ErrorDisplay() string {
  return ""
}
