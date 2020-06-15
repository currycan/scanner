package core

type Config struct {
	Name string
}

type Service struct {
	Name    string
	Project string
	Host    string
	Port    int
}

type Services struct {
	Services []Service
}

var (
	Name    string
	Project string
	Host    string
	Port    int
)
