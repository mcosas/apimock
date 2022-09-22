package banana

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func setupHandlers(apis []string) {
	http.HandleFunc(apis[0], response)
}

func response(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI
	fmt.Println(path)
	log.Printf("processing... %s", path)

	w.Header().Set("Content-Type", "application/json")

	content, err := ioutil.ReadFile("./payloads/users/12345.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	w.Write(content)
}
