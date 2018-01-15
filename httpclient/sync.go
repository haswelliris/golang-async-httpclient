package httpclient

import (
	"fmt"
	"time"

	"github.com/haswelliris/golang-async-httpclient/service"
)

func Synchronous() [10]service.Infomation {
	start := time.Now()
	customer := service.GetCustomers()
	destinations := service.GetDestinations(customer)
	var infos [10]service.Infomation
	for index, dest := range destinations {
		q := service.GetQuoting(dest)
		w := service.GetWeather(dest)
		infos[index] = service.Infomation{W: w, D: dest, Q: q}
	}
	fmt.Println(time.Since(start))
	return infos
}
