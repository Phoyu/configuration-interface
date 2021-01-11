package test

import (
  "os"
  "testing"
)

import (
  "github.com/Phoyu/configuration-interface/env"
)

func TestCheckRequiredProperties(t *testing.T) {
  testConfiguration := env.NewConfiguration()

  errProperties := testConfiguration.CheckRequiredProperties()
  if errProperties != nil {
    t.Fail()
  }

  testConfiguration.AddRequiredProperty("TEST_PROPERTY")
  errProperties = testConfiguration.CheckRequiredProperties()
  if errProperties == nil {
    t.Fail()
  }

  os.Setenv("TEST_PROPERTY","test")
  os.Setenv("TEST_PROPERTY2","test2")
  errProperties = testConfiguration.CheckRequiredProperties()
  if errProperties != nil {
    t.Fail()
  }

  testConfiguration.AddRequiredProperty("TEST_PROPERTY2")
  errProperties = testConfiguration.CheckRequiredProperties()
  if errProperties != nil {
    t.Fail()
  }

  requiredProperties := []string{"TEST_PROPERTY","TEST_PROPERTY2"}
  for i, value := range testConfiguration.GetRequiredProperties() {
    if requiredProperties[i] != value {
      t.Fail()
    }
  }

  os.Unsetenv("TEST_PROPERTY")
  os.Unsetenv("TEST_PROPERTY2")
}

func TestGetProperty(t *testing.T) {
  testConfiguration := env.NewConfiguration()

  none, err := testConfiguration.GetProperty("TEST_PROPERTY")
  if none != "" || err == nil {
    t.Fatal(none,err)
  }

  os.Setenv("TEST_PROPERTY","test")
  os.Setenv("TEST_PROPERTY2","test2")
  none, err = testConfiguration.GetProperty("TEST_PROPERTY2")

  if none != "test2" || err != nil {
    t.Fatal(none,err)
  }
}
