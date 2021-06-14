package main

import (
	"log"
	"net/http"
	"runtime"
)

func main() {
	log.Printf("the number of available cpus is %d\n", runtime.NumCPU())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var id string = r.RemoteAddr
		log.Printf("new session from %s\n", id)
		buf := make([]byte, 4096*1024)
		defer r.Body.Close()
		var totalSize int = 0
		for {
			v, err := r.Body.Read(buf)
			if err != nil {
				break
			}
			totalSize += v
		}
		log.Printf("the total upload size is %dmb from %s\n", totalSize/1024/1024, id)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
