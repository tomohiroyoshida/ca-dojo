DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS Characters;
DROP TABLE IF EXISTS UserCharaclers;

CREATE DATABASE IF NOT EXISTS ca_dojo;

USE ca_dojo;

CREATE TABLE IF NOT EXISTS Users (
  id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
  name varchar(128) NOT NULL UNIQUE,
  token varchar(255) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Characters (
  id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
  name varchar(128) NOT NULL UNIQUE,
  rarity INT NOT NULL
);

CREATE TABLE IF NOT EXISTS UserCharaclers (
  id INT PRIMARY KEY NOT NULL UNIQUE AUTO_INCREMENT,
  user_id INT NOT NULL REFERENCES Users(id),
  character_id id INT NOT NULL REFERENCES Characters(id)
);

INSERT INTO Users(name, rarity) VALUES("ドラゴンA", 1);
INSERT INTO Users(name, rarity) VALUES("ドラゴンA", 1);
INSERT INTO Users(name, rarity) VALUES("ドラゴンA", 1);
INSERT INTO Users(name, rarity) VALUES("ドラゴンA", 1);
INSERT INTO Users(name, rarity) VALUES("ドラゴンB", 2);
INSERT INTO Users(name, rarity) VALUES("ドラゴンB", 2);
INSERT INTO Users(name, rarity) VALUES("ドラゴンC", 3);
INSERT INTO Users(name, rarity) VALUES("ドラゴンC", 3);
INSERT INTO Users(name, rarity) VALUES("ドラゴンD", 4);
INSERT INTO Users(name, rarity) VALUES("ドラゴンD", 4);
INSERT INTO Users(name, rarity) VALUES("ドラゴンE", 5);