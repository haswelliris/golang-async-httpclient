package httpclient

import (
	"fmt"
	"time"

	"github.com/haswelliris/golang-async-httpclient/service"
)

func Asynchronous() [10]service.Infomation {
	start := time.Now()
	customer := service.GetCustomers()
	destinations := service.GetDestinations(customer)
	var infos [10]service.Infomation
	quotes := [10]chan service.Quoting{}
	weathers := [10]chan service.Weather{}

	for i, _ := range quotes {
		quotes[i] = make(chan service.Quoting)
	}

	for i, _ := range weathers {
		weathers[i] = make(chan service.Weather)
	}

	for index, dest := range destinations {
		i := index
		d := dest
		go func() {
			quotes[i] <- service.GetQuoting(d)
		}()

		go func() {
			weathers[i] <- service.GetWeather(d)
		}()
	}
	for index, dest := range destinations {
		infos[index] = service.Infomation{W: <-weathers[index], D: dest, Q: <-quotes[index]}
	}
	fmt.Println(time.Since(start))
	return infos
}
