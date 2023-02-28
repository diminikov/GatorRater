DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts (
  id         INT AUTO_INCREMENT NOT NULL,
  username      VARCHAR(128) NOT NULL,
  password     VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO accounts
  (username, password)
VALUES
  ('Administrator', 'Password');