package service

import "time"

type Weather struct{}
type Destinations struct{}
type Quoting struct{}
type Customers struct{}
type Infomation struct {
	W Weather
	D Destinations
	Q Quoting
}

func GetCustomers() Customers {
	time.Sleep(150 * time.Millisecond)
	return Customers{}
}

func GetDestinations(c Customers) [10]Destinations {
	time.Sleep(250 * time.Millisecond)
	return [10]Destinations{}
}

func GetQuoting(d Destinations) Quoting {
	time.Sleep(170 * time.Millisecond)
	return Quoting{}
}

func GetWeather(d Destinations) Weather {
	time.Sleep(330 * time.Millisecond)
	return Weather{}
}
