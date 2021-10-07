



Each connection consumes around 20KB which is the sum from the goroutine, buffer in net/http and buffers from gorilla/ws
So, to optimize the memory consumption of the application we must apply those above.

1 - goroutine
Using epoll. 
2, 3 - net/http buffer and gorilla/ws buffer by using gobwas/ws