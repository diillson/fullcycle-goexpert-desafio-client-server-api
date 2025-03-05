package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"time"
)

type Cotacao struct {
	USDBRL struct {
		Bid string `json:"bid"`
	} `json:"USDBRL"`
}

func main() {
	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criar tabela se não existir
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS cotacoes (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        bid TEXT,
        data DATETIME
    )`)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/cotacao", func(w http.ResponseWriter, r *http.Request) {
		cotacao, err := getBid()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = saveCotacao(db, cotacao.USDBRL.Bid)
		if err != nil {
			log.Printf("Erro ao salvar no banco: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"bid": cotacao.USDBRL.Bid})
	})

	http.ListenAndServe(":8080", nil)
}

func getBid() (*Cotacao, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cotacao Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		return nil, err
	}

	return &cotacao, nil
}

func saveCotacao(db *sql.DB, bid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	query := "INSERT INTO cotacoes (bid, data) VALUES (?, datetime('now'))"
	_, err := db.ExecContext(ctx, query, bid)
	if err != nil {
		return fmt.Errorf("erro ao salvar no banco: %v", err)
	}

	return nil
}
