
-- show all users
SELECT User FROM mysql.user;
-- create user
CREATE USER 'go'@'localhost' IDENTIFIED BY 'password';
-- delete user
DROP USER IF EXISTS 'user'@'host'; -- host == localhost
-- create database
CREATE DATABASE go_library_api;
-- grant privileges
GRANT ALL PRIVILEGES ON go_library_api.* to go@localhost;


-- db and table operations
SHOW DATABASES;
USE <database_name>;
SHOW TABLES;
