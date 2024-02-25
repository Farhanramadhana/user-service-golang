CREATE TABLE user (
	id serial PRIMARY KEY,
  full_name VARCHAR(60) NOT NULL,
  phone_number VARCHAR(13) NOT NULL,
  password VARCHAR(60) NOT NULL,
  created_at timestampttz,
  updated_at timestampttz
);
