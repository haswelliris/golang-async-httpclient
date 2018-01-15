package httpclient

import (
    "fmt"
    "time"
  
	"github.com/haswelliris/golang-async-httpclient/service"
)
func Asynchronous() [10]httpclient.Infomation {
    start := time.Now()    
    customer := httpclient.GetCustomers()     
    destinations := httpclient.GetDestinations(customer)     
    var infos [10]httpclient.Infomation     
    quotes := [10]chan httpclient.Quoting{}     
    weathers := [10]chan httpclient.Weather{}   
  
    for i, _ := range quotes {     
        quotes[i] = make(chan httpclient.Quoting)     
    }      
  
    for i, _ := range weathers {     
        weathers[i] = make(chan httpclient.Weather)     
    }   
  
    for index, dest := range destinations {      
        i := index      
        d := dest      
        go func() {       
            quotes[i] <- httpclient.GetQuoting(d)      
        }()
      
        go func() {       
            weathers[i] <- httpclient.GetWeather(d)      
        }()     
    }      
    for index, dest := range destinations {   
        infos[index] = httpclient.Infomation{W:<-weathers[index],D:dest, Q:<-quotes[index]}  
    } 
  
    fmt.Println(time.Since(start))   
    return infos
} 
