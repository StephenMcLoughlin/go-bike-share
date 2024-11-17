package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(host string, port int, user string, password string, dbName string) (*Postgres, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to db")

	return &Postgres{
		db: db,
	}, nil
}

func (p *Postgres) Close() {
	p.db.Close()
}

func (p *Postgres) ExecQuery(query string, args ...interface{}) (sql.Result, error) {
	result, err := p.db.Exec(query, args...)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	return result, nil
}
