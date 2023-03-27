package config

type Config struct {
	Port  string `env:"PORT,default=8080" short:"p" long:"port" description:"HTTP Port"`
	DbURL string `env:"DBURL,default=postgres://postgres:kamalesh@localhost:5432/calculator" long:"dbUrl" description:"Database connection URL"`
}
