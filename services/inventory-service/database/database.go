package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBInventory struct {
	Drive    string
	User     string
	Password string
	Port     int
	Host     string
	Conn     *sql.DB
	DataBase string
}

func NewDB(drive string, user string, password string, port int, host string, db string) *DBInventory {

	return &DBInventory{Drive: drive, User: user, Password: password, Port: port, Host: host, DataBase: db}
}

func (db *DBInventory) Conection() error {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.DataBase,
	)
	fmt.Println("string de conexao: ", url)
	var err error
	if db.Conn, err = sql.Open(db.Drive, url); err != nil {
		return err
	}
	return nil

}
