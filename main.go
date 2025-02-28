package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/time", handler)
	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":8795", nil)
	if err != nil {
		fmt.Println("An error has occurred. The server is down...")
	}
}

func handler(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	resp := map[string]string{
		"time": time.Now().Format(time.RFC3339),
	}
	jresp, _ := json.Marshal(resp)
	_, err := writer.Write(jresp)
	if err != nil {
		fmt.Println("An error occurred while sending the json.")
	}
}
