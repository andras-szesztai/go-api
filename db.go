package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MySQL")

	return &MySQLStorage{db: db}
}

func (storage *MySQLStorage) Init() (*sql.DB, error) {
	if err := storage.createProjectsTable(); err != nil {
		return nil, err
	}
	if err := storage.createUsersTable(); err != nil {
		return nil, err
	}
	if err := storage.createTasksTable(); err != nil {
		return nil, err
	}
	return storage.db, nil
}

func (storage *MySQLStorage) createProjectsTable() error {
	_, err := storage.db.Exec(`CREATE TABLE IF NOT EXISTS projects (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		email VARCHAR(100) NOT NULL,
		firstName VARCHAR(100) NOT NULL,
		lastName VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		createAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

		PRIMARY KEY (id),
		UNIQUE KEY (email)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8`)
	if err != nil {
		return err
	}
	return nil
}

func (storage *MySQLStorage) createUsersTable() error {
	_, err := storage.db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		email VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		firstName VARCHAR(100) NOT NULL,
		lastName VARCHAR(100) NOT NULL,
		createAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

		PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8`)
	if err != nil {
		return err
	}
	return nil
}

func (storage *MySQLStorage) createTasksTable() error {
	_, err := storage.db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(100) NOT NULL,
		status ENUM('TODO', 'IN_PROGRESS', 'DONE') NOT NULL DEFAULT 'TODO',
		projectId INT UNSIGNED NOT NULL,
		assignedToId INT UNSIGNED NOT NULL,
		createAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

		PRIMARY KEY (id),
		FOREIGN KEY (projectId) REFERENCES projects(id),
		FOREIGN KEY (assignedToId) REFERENCES users(id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8`)
	if err != nil {
		return err
	}
	return nil
}
