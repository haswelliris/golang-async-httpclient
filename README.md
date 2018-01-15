# golang-async-httpclient
golang async httpclient
## Reactive动机  
Reactive动机的作用就是将服务端的需要被客户端调用内部服务操作集中到对方开放层，从而避免所有内部服务暴露在公网。此外，Reactive动机还使得流量开销减少以及延迟降低。并且降低一个主机承载所有服务的压力，起到负载均衡和分布式计算的作用。  
![image](https://github.com/haswelliris/golang-async-httpclient/asserts/1.png)  
## 使用同步方法解决问题
![image](https://github.com/haswelliris/golang-async-httpclient/asserts/2.png)  
## 搭建异步机制解决问题
![image](https://github.com/haswelliris/golang-async-httpclient/asserts/3.png)  
## 测试结果
![image](https://github.com/haswelliris/golang-async-httpclient/asserts/4.png)  
异步机制的速度会比使用同步的速度快很多。  
## 小结  
因为熟悉nodejs异步服务端开发，所用用异步实现轻车熟路。  
