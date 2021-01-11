package configuration

type Configuration interface {
  SetProperty(string,string) error
  GetProperty(string) (string,error)

  AddRequiredProperty(string)
  CheckRequiredProperties() error
}
