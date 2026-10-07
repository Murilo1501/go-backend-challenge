CREATE TABLE Wallet (
  id int PRIMARY KEY,
  player_id varchar(255) NOT NULL,
  currency varchar(255),
  balance BIGINT NOT NULL,
  version BIGINT NOT NULL, 
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
);