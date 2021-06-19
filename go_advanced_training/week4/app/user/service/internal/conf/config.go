package conf

type MongoDB struct {
	Uri string `yaml:"uri"`
	Db  string `yaml:"db"`
}

type Data struct {
	MongoDB MongoDB `yaml:"mongodb"`
}

type Http struct {
	Addr string `yaml:"addr"`
}

type Server struct {
	Http Http `yaml:"http"`
}

type Config struct {
	Server *Server `yaml:"server"`
	Data   *Data   `yaml:"data"`
}
