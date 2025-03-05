package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar: %s", response.Bid))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Cotação do dólar salva com sucesso: %s\n", response.Bid)
}
