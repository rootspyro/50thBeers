package utils

import "strings"

func NameToId( name string ) string {

  name = strings.ToLower(name)
  name = strings.ReplaceAll(name, " ", "_")

  return name
}
