package configModel

type Config struct {
	Postgresql  Postgresql
	Redis   Redis
	Runtime Runtime
	Server  Server
}
