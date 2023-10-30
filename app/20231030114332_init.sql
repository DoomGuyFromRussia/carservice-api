-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

create table cars (
	id serial primary key,
	producer TEXT NOT NULL,
	model TEXT NOT NULL,
	year INTEGER NOT NULL,
	vin	TEXT NOT NULL

);

CREATE TABLE clients (
	id	serial PRIMARY KEY NOT NULL,
	name	TEXT NOT NULL,
	surname	TEXT NOT NULL,
	address	TEXT NOT NULL,
	phone	TEXT NOT NULL
);

CREATE TABLE orders (
	id	serial NOT NULL primary key,
	carId	INTEGER NOT NULL,
	clientId	INTEGER NOT NULL,
	date	TEXT NOT NULL,
	description	TEXT NOT NULL,
	status	TEXT NOT NULL
);

create table clientsCars (
	clientId integer,
	carId integer,
	FOREIGN KEY(clientId) REFERENCES clients(id),
	FOREIGN KEY(carId) REFERENCES cars(id)
);
CREATE TABLE clientsOrders (
	clientId	INTEGER,
	orderId	INTEGER,
	FOREIGN KEY(clientId) REFERENCES clients(id),
	FOREIGN KEY(orderId) REFERENCES orders(id)
);




-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
drop table cars
drop table clients
drop table orders
drop table clientsCars
drop table clientsOrders

