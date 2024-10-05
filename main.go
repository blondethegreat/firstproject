package main

import (
	"log"
	"net"
	"net/http"
	"GoBruseWayne/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.Hello)

	http.HandleFunc("/shelbywalk", handlers.ShelbyWalk)

	http.HandleFunc("/yarik", handlers.Yarik)
	
	http.HandleFunc("/maccofee", handlers.MacCofee)

	log.Println("Starting server...")
	l, err := net.Listen("tcp", "localhost:443")
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.Serve(l, nil))

	// log.Println("Sending request...")
	// res, err := http.Get("http://localhost:8080/hello")

	// log.Println("Reading response...")
	// if _, err := io.Copy(os.Stdout, res.Body); err != nil {
	// 	log.Fatal(err)
	// }
	for {
	}
}
