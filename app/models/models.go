package models

type Car struct {
	Id       int
	Producer string
	Model    string
	Year     string
	Vin      string
}

func (c *Car) Describe() string {
	return c.Producer + c.Model + c.Year + c.Vin
}

type Client struct {
	id      int
	name    string
	surname string
	address string
	phone   string
}

type Order struct {
	id          int
	carId       int
	clientId    int
	date        string
	description string
	status      string
}
