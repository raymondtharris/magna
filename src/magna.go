package main

import (
  "fmt"

)

type MagnaNode struct{
  Index int
  Value string
  Weight int
}


func displayInput(){
  fmt.Println("What would you like to explore?")
  //Take in input from the user.
}



func introduction(){
  //Determine if this is a new user
  greeting("")
}

func greeting(MagnaUser string){
  if MagnaUser == "" {
    fmt.Println("Hello, I am Magna. Who might you be?")
    fmt.Scanln(&MagnaUser)
    fmt.Println("Nice to meet you "+MagnaUser+".")
  } else {
    fmt.Println("Hello, "+ MagnaUser + ". Welcome back.")
  }
}

func whatToSearch(){
  fmt.Println("What can I do for you?")
  var queryInput string
  fmt.Scanln(&queryInput)
}

func main() {
  introduction()
  for {
    whatToSearch()
  }
}
