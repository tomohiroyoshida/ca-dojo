DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS characters;
DROP TABLE IF EXISTS user_characlers;

DROP DATABASE IF EXISTS ca_dojo;
CREATE DATABASE IF NOT EXISTS ca_dojo;

USE ca_dojo;

CREATE TABLE IF NOT EXISTS users (
  id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
  name varchar(128) NOT NULL UNIQUE,
  token varchar(255) NOT NULL UNIQUE,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS characters (
  id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
  name varchar(128) NOT NULL UNIQUE,
  weight INT NOT NULL,
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS user_characters (
  id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
  user_id INT NOT NULL REFERENCES users(id),
  character_id INT NOT NULL REFERENCES characters(id),
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO users(name, token) VALUES("山田1", 'token1');
INSERT INTO users(name, token) VALUES("山田2", 'token2');
INSERT INTO users(name, token) VALUES("山田3", 'token3');

INSERT INTO characters(name, weight) VALUES("ドラゴンA", 1);
INSERT INTO characters(name, weight) VALUES("ドラゴンB", 1);
INSERT INTO characters(name, weight) VALUES("ドラゴンC", 1);
INSERT INTO characters(name, weight) VALUES("ドラゴンD", 1);
INSERT INTO characters(name, weight) VALUES("ドラゴンE", 2);
INSERT INTO characters(name, weight) VALUES("ドラゴンF", 2);
INSERT INTO characters(name, weight) VALUES("ドラゴンG", 3);
INSERT INTO characters(name, weight) VALUES("ドラゴンH", 3);
INSERT INTO characters(name, weight) VALUES("ドラゴンI", 4);
INSERT INTO characters(name, weight) VALUES("ドラゴンJ", 4);
INSERT INTO characters(name, weight) VALUES("ドラゴンK", 5);
