CREATE DATABASE IF NOT EXISTS projectmanager DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE projectmanager;

CREATE TABLE IF NOT EXISTS users (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		email VARCHAR(100) NOT NULL,
		password VARCHAR(100) NOT NULL,
		firstName VARCHAR(100) NOT NULL,
		lastName VARCHAR(100) NOT NULL,
		createAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

		PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;

  CREATE TABLE IF NOT EXISTS tasks (
		id INT UNSIGNED NOT NULL AUTO_INCREMENT,
		name VARCHAR(100) NOT NULL,
		status ENUM('TODO', 'IN_PROGRESS', 'DONE') NOT NULL DEFAULT 'TODO',
		assignedToId INT UNSIGNED NOT NULL,
		createAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

		PRIMARY KEY (id),
		FOREIGN KEY (assignedToId) REFERENCES users(id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO users (email, password, firstName, lastName) VALUES ('example@example.com', 'password123', 'John', 'Doe');
INSERT INTO users (email, password, firstName, lastName) VALUES ('another@example.com', 'pass123', 'Jane', 'Smith');
INSERT INTO users (email, password, firstName, lastName) VALUES ('test@example.com', 'test123', 'Alex', 'Johnson');

INSERT INTO tasks (name, status, assignedToId) VALUES ('Task 1', 'TODO', 1);
INSERT INTO tasks (name, status, assignedToId) VALUES ('Task 2', 'IN_PROGRESS', 2);
INSERT INTO tasks (name, status, assignedToId) VALUES ('Task 3', 'DONE', 3);

