package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"runtime"
)

// https://stackoverflow.com/questions/7580508/getting-chrome-to-accept-self-signed-localhost-certificate

// https://gist.github.com/denji/12b3a568f092ab951456

func main() {
	var err error
	var addrs []net.Addr

	var ifaces []net.Interface

	ifaces, err = net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, err = i.Addrs()
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			log.Printf("addr: %v", ip)
		}
	}

	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)

	log.Println("Listening https on :443... ")
	//err = http.ListenAndServe(":3000", nil)
	err = http.ListenAndServeTLS(":443", "./examples/server/server.crt", "./examples/server/server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
