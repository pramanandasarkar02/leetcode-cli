package config


type Config struct{
	OutputFileName string 
}


func NewConfig(OutputFileName string) * Config {
	return &Config{
		OutputFileName: OutputFileName,
	}
}