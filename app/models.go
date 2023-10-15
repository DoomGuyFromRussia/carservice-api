package main

type Car struct {
	id       int
	producer string
	model    string
	year     int
	vin      string
}

type Client struct {
	id      int
	name    string
	surname string
	adress  string
	phone   string
}

type Order struct {
	id          int
	car         Car
	client      Client
	data        string
	description string
	status      string
}
