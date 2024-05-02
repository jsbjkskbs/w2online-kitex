package db

type _MysqlConfig struct {
	MysqlDSN string
}

var (
	Conf _MysqlConfig
)