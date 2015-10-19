package magnauser

import (
	"fmt"
)

//User is the object to store info about the user
type User struct {
	Name  string //Name of the User
	IDTag int    //IDTag of the User
}

func (mUser User) String() string {
	return fmt.Sprintf("username: %v, Tag: %v", mUser.Name, mUser.IDTag)
}

//Interest is an object to store something the user likes
type Interest struct {
	Name string //Name of the interest
}

func (mInterest Interest) String() string {
	return fmt.Sprintf("%v", mInterest.Name)
}
