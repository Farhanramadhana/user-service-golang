CREATE TABLE users (
	id serial PRIMARY KEY,
  full_name VARCHAR(60) NOT NULL,
  phone_number VARCHAR(13) NOT NULL UNIQUE,
  password VARCHAR(60) NOT NULL,
  created_at timestamptz,
  updated_at timestamptz
);

CREATE TABLE login_logs(
  id serial PRIMARY KEY,
  user_id int REFERENCES users(id),
  total_login_success int,
  last_login_at timestamptz,
  created_at timestamptz,
  updated_at timestamptz
);