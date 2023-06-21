CREATE USER 'fredy'@'localhost' IDENTIFIED BY '12345';
CREATE DATABASE wininbrowser_calendar_db;
GRANT ALL PRIVILEGES ON wininbrowser_calendar_db.* TO 'fredy'@'localhost';
FLUSH PRIVILEGES;

USE wininbrowser_calendar_db;

CREATE TABLE users (
  id INT NOT NULL PRIMARY KEY
);
CREATE TABLE events (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  start DATETIME NOT NULL,
  end DATETIME NOT NULL,
  allDay BOOLEAN NOT NULL,
  location VARCHAR(255)
);
CREATE TABLE participants (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  eventId INT NOT NULL,
  userId INT NOT NULL,
  FOREIGN KEY (eventId) REFERENCES events(id),
  FOREIGN KEY (userId) REFERENCES users(id)
);
