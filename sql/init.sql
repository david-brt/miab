-- underscore necessary as user is a reserved keyword
CREATE TABLE
  User_ (
    id serial PRIMARY KEY,
    username varchar(100) UNIQUE NOT NULL
  );

CREATE TABLE
  Message (
    id serial PRIMARY KEY,
    content text NOT NULL,
    sender int,
    timestamp timestamp NOT NULL,
    FOREIGN KEY (sender) REFERENCES User_ (id)
  );

INSERT INTO
  User_ (username)
VALUES
  ('admin');

INSERT INTO
  Message (content, sender, timestamp)
VALUES
  (
    'init',
    (
      SELECT
        id
      FROM
        User_
      WHERE
        username = 'admin'
    ),
    CURRENT_TIMESTAMP
  );
