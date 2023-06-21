CREATE USER 'auth_ms'@'localhost' IDENTIFIED BY 'Auth_ms12#$';
CREATE DATABASE auth_db;
GRANT ALL PRIVILEGES ON auth_db.* TO 'auth_ms'@'localhost';
FLUSH PRIVILEGES;

use auth_db;
