CREATE TABLE IF NOT EXISTS Users (
  id UUID PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(36) NOT NULL,
  role VARCHAR(20) NOT NULL,
  verification_code INT NOT NULL,
  verified BOOL NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS Sessions (
  id UUID PRIMARY KEY,
  session_id UUID NOT NULL,
  token VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  expires_at TIMESTAMP NOT NULL
)
-- Add more table creation statements as needed
CREATE TABLE IF NOT EXISTS ActiveSessions (
  id UUID PRIMARY KEY,
  session_id UUID NOT NULL,
  token VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  expires_at TIMESTAMP NOT NULL
)