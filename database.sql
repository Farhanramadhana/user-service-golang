CREATE TABLE users (
	id serial PRIMARY KEY,
  full_name VARCHAR(60) NOT NULL,
  phone_number VARCHAR(13) NOT NULL UNIQUE,
  password VARCHAR(60) NOT NULL,
  created_at timestamptz,
  updated_at timestamptz
);
