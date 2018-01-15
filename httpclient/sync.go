package httpclient

import (
	"fmt"
	"time"

	"github.com/haswelliris/golang-async-httpclient/httpclient"
)

func Synchronous() [10]httpclient.Infomation {
	start := time.Now()
	customer := httpclient.GetCustomers()
	destinations := httpclient.GetDestinations(customer)
	var infos [10]httpclient.Infomation
	for index, dest := range destinations {
		q := httpclient.GetQuoting(dest)
		w := httpclient.GetWeather(dest)
		infos[index] = httpclient.Infomation{W: w, D: dest, Q: q}
	}
	fmt.Println(time.Since(start))
	return infos
}
