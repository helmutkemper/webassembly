package main

import (
	"io"
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

	//fs := http.FileServer(http.Dir("./"))
	//http.Handle("/", fs)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCORS(&w)

		// Verifica se a requisição é do tipo OPTIONS
		if r.Method == "OPTIONS" {
			return
		}

		// Serve os arquivos estáticos
		http.ServeFile(w, r, r.URL.Path[1:])
	})

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

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	//(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET")
	//(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// Construir o URL do OpenStreetMap usando o caminho da requisição
	url := "https://tile.openstreetmap.org" + r.URL.Path

	// Fazer a requisição para o OpenStreetMap
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Erro ao buscar o recurso do OpenStreetMap", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copiar os cabeçalhos da resposta do OpenStreetMap para o cliente
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Copiar o status code da resposta do OpenStreetMap para o cliente
	w.WriteHeader(resp.StatusCode)

	// Copiar o corpo da resposta do OpenStreetMap para o cliente
	io.Copy(w, resp.Body)
}

/*
func main() {
	// Define a função que lida com as requisições, incluindo o cabeçalho CORS
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Adiciona o cabeçalho CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Verifica se a requisição é do tipo OPTIONS
		if r.Method == "OPTIONS" {
			return
		}

		// Serve os arquivos estáticos
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	log.Println("Servidor iniciado na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
*/
