#!/bin/bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL

CREATE TABLE product_category(
	id VARCHAR PRIMARY KEY, 
	category_name VARCHAR(30) NOT NULL UNIQUE
);

CREATE TABLE product( 
	id VARCHAR PRIMARY KEY NOT NULL, 
	product_name VARCHAR(50) NOT NULL, 
	amount	INT NOT NULL, 
	price INT NOT NULL, 
	expire VARCHAR(6) NOT NULL, 
	category_id VARCHAR REFERENCES product_category(id)
);

EOSQL

