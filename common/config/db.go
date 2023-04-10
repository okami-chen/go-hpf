package config

type Mysql struct {
	Host         string
	Config       string
	Dbname       string
	Username     string
	Password     string
	MaxIdleConns int
	MaxOpenConns int
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ")/" + m.Dbname + "?" + m.Config
}
