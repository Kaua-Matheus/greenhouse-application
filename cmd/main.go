package main

import (
	"fmt"
	"io"
	"net/http"
)

// "github.com/kaua-matheus/greenhouse-application/server"

func main() {
	resp, err := http.Get("http://localhost:8080/test"); if err != nil {
		fmt.Printf("Error trying to execute http.Get: %s\n", err);
		return;
	};
	defer resp.Body.Close();

	body, err := io.ReadAll(resp.Body); if err != nil {
		fmt.Printf("An error occurs trying to read the resp.Body: %s\n", err);
		return;
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode);
	fmt.Printf("Full Body: %s\n", string(body));
}