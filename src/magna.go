package main

import (
  "fmt"
  "../../docker_testing/src/magnauser"
  "strings"
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
  var newUser = magnauser.MagnaUser{"Tim", 113232444443}
  greeting(newUser)
}

func greeting(mUser magnauser.MagnaUser){
  if mUser.Name == "" {
    fmt.Println("Hello, I am Magna. Who might you be?")
    var user string
    fmt.Scanln(&user)
    fmt.Println("Nice to meet you "+user+".")
  } else {
    fmt.Println("Hello, "+ mUser.Name + ". Welcome back.")
  }
}

func whatToSearch(){
  fmt.Println("What can I do for you?")
  var queryInput string
  fmt.Scanf("%s",&queryInput)
  fmt.Println(queryInput)
  splitQuery  := strings.Split(queryInput, " ")
  fmt.Println(splitQuery)
  for i:= 0; i< len(splitQuery); i++ {

    fmt.Println(splitQuery[i])
  }
}

func main() {
  introduction()
  //for {
    whatToSearch()
  //}
}
