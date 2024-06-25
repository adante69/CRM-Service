-- +goose Up
-- SQL в этой секции будет применен при запуске команды `goose up`.

CREATE TABLE accounts (
                          id SERIAL PRIMARY KEY,
                          email VARCHAR(255) NOT NULL UNIQUE,
                          password VARCHAR(255) NOT NULL
);

CREATE TABLE contacts (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          phone VARCHAR(255) NOT NULL,
                          description TEXT
);

CREATE TABLE partners (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          description TEXT
);

CREATE TABLE partner_contacts (
                                  partner_id INT NOT NULL REFERENCES partners(id) ON DELETE CASCADE,
                                  contact_id INT NOT NULL REFERENCES contacts(id) ON DELETE CASCADE,
                                  PRIMARY KEY (partner_id, contact_id)
);

CREATE TABLE bids (
                      id SERIAL PRIMARY KEY,
                      description TEXT,
                      amount INT,
                      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose Down
-- SQL в этой секции будет применен при запуске команды `goose down`.

DROP TABLE IF EXISTS bids;
DROP TABLE IF EXISTS partner_contacts;
DROP TABLE IF EXISTS partners;
DROP TABLE IF EXISTS contacts;
DROP TABLE IF EXISTS accounts;