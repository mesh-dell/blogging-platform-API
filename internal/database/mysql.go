package database

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/mesh-dell/blogging-platform-API/config"
)

type MySQLDB struct {
	Client *sql.DB
}

func NewMySQLDB(config config.Config) (*MySQLDB, error) {
	cfg := mysql.NewConfig()
	cfg.User = config.DBUser
	cfg.Passwd = config.DBPassword
	cfg.Net = "tcp"
	cfg.Addr = config.DBAddr
	cfg.DBName = config.DBName

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("db open: %v", err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		return nil, fmt.Errorf("db ping: %v", pingErr)
	}
	return &MySQLDB{Client: db}, nil
}
