DROP TABLE IF EXISTS user;
CREATE TABLE user (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  name TEXT NOT NULL,
  username TEXT NOT NULL,
  email TEXT NOT NULL,
  password TEXT NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO user
    (name, username, email, password)
VALUES
    ('Mimmo Baluyut', 'mimmo', 'mimmo.baluyut@icloud.com', 'Rockyou'),
    ('Oona Waters', 'oona', 'oona.waters@fakeemail.com', 'oonasplaceholder'),
    ('Jaya Spencer', 'jaya', 'jaya.spencer@fakeemail.com', 'jayaspaceholder');

-- default testing data