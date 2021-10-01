



Each connection consumes around 20KB which is the sum from the goroutine, buffer in net/http e buffers from gorilla/ws
So, to optimize the memory consumption of the application we must apply those above.

1 - goroutine
Using epoll. 


2 - net/http buffer

3- gorilla/ws buffer