package main

/*
This problem describes two processes, the producer and the consumer who share a commmon fixed sized buffer
-> Producer generates data and put it into the buffer, and repeats
-> Consumer consumes from the buffer, one piece at a time

Producer can't add stuff to the bufferif its full, and that the consumer won't try to remove data from an empty buffer
-> Producer must either sleep, or discard data if the buffer is full

*/
func main(){

}