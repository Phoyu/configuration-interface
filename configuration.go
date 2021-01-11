package configuration

type Configuration interface {
  SetProperty(string,string) error
  GetProperty(string) (string,error)

  AddRequiredProperty(string)
  GetRequiredProperties() []string
  CheckRequiredProperties() error
}

type AppConfig struct {
  properties          map[string]string
  originConfiguration Configuration
}

func NewAppConfig(config Configuration) *AppConfig {
  return &AppConfig {
    properties: make(map[string]string),
    originConfiguration: config,
  }
}

func (ac *AppConfig) AddRequiredProperty(property string) {
  ac.originConfiguration.AddRequiredProperty(property)
}

func (ac *AppConfig) Load() error {
  requiredConfigErr := ac.originConfiguration.CheckRequiredProperties()
  if requiredConfigErr != nil {
    return requiredConfigErr
  }

  for _, property := range ac.originConfiguration.GetRequiredProperties() {
    ac.properties[property], _ = ac.originConfiguration.GetProperty(property)
  }

  return nil
}

func (ac *AppConfig) GetProperty(property string) string {
  return ac.properties[property]
}
