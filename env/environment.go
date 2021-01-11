package env

import (
  "os"
  "fmt"
)

import (
  "github.com/Phoyu/configuration-interface"
)

type EnvironmentConfiguration struct{
  requiredProperties []string
}

func NewConfiguration() configuration.Configuration {
  return new(EnvironmentConfiguration)
}

func (ec *EnvironmentConfiguration) GetProperty(propertyId string) (string,error) {
  prop, exists := os.LookupEnv(propertyId)
  if !exists {
    return "",configuration.ErrPropertyNotFound(propertyId)
  }
  return prop,nil
}

func (ec *EnvironmentConfiguration) SetProperty(propertyId string,property string) error {
  return fmt.Errorf("EnvironmentConfiguration: %w",configuration.ErrPropertyCouldNotBeSet())
}

func (ec *EnvironmentConfiguration) AddRequiredProperty(property string) {
  ec.requiredProperties = append(ec.requiredProperties,property)
}

func (ec *EnvironmentConfiguration) CheckRequiredProperties() error {
  for _, property := range ec.requiredProperties {
    _, exists := os.LookupEnv(property)
    if !exists {
      return fmt.Errorf("Missing required property: %w",configuration.ErrPropertyNotFound(property))
    }
  }
  return nil
}
