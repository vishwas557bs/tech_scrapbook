package main

import (
	"log"
	"net"
	"time"
)

func main() {

	listner, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go do(conn)
	}

}

func do(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(10 * time.Second)
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nhello world!!\r\n"))
	conn.Close()
}

//imporvements
//using thread pools to make sure large number of threads are not created.
//time out mechanism
