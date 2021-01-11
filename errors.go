package configuration

import (
  "fmt"
)

func ErrPropertyNotFound(property string) error {
  return fmt.Errorf("Could not find %v property",property)
}

func ErrPropertyCouldNotBeSet() error {
  return fmt.Errorf("Property could not be set")
}
