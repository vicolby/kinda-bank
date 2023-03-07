package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountById(int) (*Account, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres password=12345 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db}, nil
}

func (s *PostgresStorage) Init() error {
	err := s.CreateAccountTable()
	return err
}

func (s *PostgresStorage) CreateAccountTable() error {
	_, err := s.db.Exec(`
	CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		number SERIAL NOT NULL,
		balance FLOAT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`)

	return err
}

func (s *PostgresStorage) CreateAccount(a *Account) error {
	query := `INSERT INTO accounts (first_name, last_name, number, balance, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.Query(query, a.FirstName, a.LastName, a.Number, a.Balance, a.Created_at)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStorage) UpdateAccount(a *Account) error {
	return nil
}

func (s *PostgresStorage) GetAccountById(id int) (*Account, error) {
	return nil, nil
}
