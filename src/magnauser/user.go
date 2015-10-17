package magnauser

import (
  "fmt"
)

type MagnaUser struct{
  Name string
  IdTag int
}

func (mUser MagnaUser) String() string {
  return fmt.Sprintf("username: %v, Tag: %v", mUser.Name, mUser.IdTag)
}
