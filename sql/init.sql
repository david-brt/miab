-- underscore necessary as user is a reserved keyword
CREATE TABLE
  User_ (
    id serial PRIMARY KEY,
    username varchar(34) UNIQUE NOT NULL,
    password_hash_salted varchar(72) NOT NULL,
    latest_login_attempt timestamp NOT NULL,
    failed_login_attempts int NOT NULL
  );

CREATE TABLE
  Message (
    id serial PRIMARY KEY,
    content text NOT NULL,
    sender int,
    sender_name varchar(34),
    timestamp timestamp NOT NULL,
    has_been_read boolean DEFAULT false,
    FOREIGN KEY (sender) REFERENCES User_ (id)
  );

INSERT INTO
  User_ (username, password_hash_salted, latest_login_attempt, failed_login_attempts)
VALUES
-- default password: 'password'
  ('miab', '$2a$12$HyunzAGcR4j0gZ4W3t/DVe2jW8eO/VhBCu4KPJsROSiIQ.6OI8zMK', CURRENT_TIMESTAMP, 0);

INSERT INTO
  Message (content, sender, sender_name, timestamp,has_been_read)
VALUES
  (
    'init',
    (
      SELECT
        id
      FROM
        User_
      WHERE
        username = 'glumanda'
    ),
    'glumanda',
    CURRENT_TIMESTAMP,
    false
  );
