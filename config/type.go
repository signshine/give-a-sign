package config

type Config struct {
	DB     DBConfig     `json:"db"`
	Server ServerConfig `json:"server"`
	SMTP   SMTPConfig   `json:"smtp"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ServerConfig struct {
	HttpPort          uint   `json:"httpPort"`
	Secret            string `json:"secret"`
	AuthExpMinute     uint   `json:"authExpMin"`
	AuthRefreshMinute uint   `json:"authExpRefreshMin"`
}

type SMTPConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	From     string `json:"from"`
	Username string `json:"username"`
	Password string `json:"password"`
}
