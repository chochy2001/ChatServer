package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var (
	portNetcat = flag.Int("port", 8080, "Port to connect to")
	hostNetcat = flag.String("host", "localhost", "Host to connect to")
)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", *hostNetcat, *portNetcat))
	if err != nil {
		log.Fatalln(err)
	}
	done := make(chan struct{}) //Control channel
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			return
		}
		done <- struct{}{}
	}()
	CopyContent(conn, os.Stdin)
	errClose := conn.Close()
	if errClose != nil {
		return
	}
	<-done
}

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatalln(err)
	}
}
