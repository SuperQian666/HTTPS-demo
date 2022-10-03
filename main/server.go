package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Ans struct {
	CurrentTime time.Time   `json:"current_time"`
	TimeList    []time.Time `json:"time_list"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().Unix()
	res := time.Unix(timestamp, 0)
	gap := 10012
	list := make([]time.Time, 7)
	for i := 0; i < 7; i++ {
		list[i] = time.Unix(timestamp+int64(i*gap), 0)
	}
	ans := &Ans{
		CurrentTime: res,
		TimeList:    list,
	}

	ret, _ := json.MarshalIndent(ans, "", "")
	fmt.Fprintf(w, "%s", ret)
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServeTLS("localhost:443", "../server.pem", "../server.key", nil); err != nil {
		fmt.Println(err)
	}
}
