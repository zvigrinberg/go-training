package main
import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("key1","value1")
		w.Header().Set("key2","value2")
		fmt.Fprint(w, "Hello from go")

	})

	log.Println("Starting server...")
	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Fatal(http.Serve(l, nil))
	}()

	log.Println("Sending request...")
	res, err := http.Get("http://localhost:8080/test")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Reading response...")
	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatal(err)
	}
	log.Println("Reading Headers...")
	for key, values := range res.Header {
         fmt.Printf("key name: %v, values: %v",key,values)
		 fmt.Println("")
	}

}
