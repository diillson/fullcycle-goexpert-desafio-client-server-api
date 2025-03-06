# fullcycle-goexpert-desafio-client-server-api

Para executar o sistema:

    Primeiro, você precisará instalar as dependências, dentro da pasta do projeto:

go mod tidy

    Em um terminal, execute o servidor:

go run server.go

    Em outro terminal, execute o cliente:

go run client.go

Este código implementa todos os requisitos solicitados:

    O servidor (server.go):
        Consome a API de cotação do dólar
        Usa context com timeout de 200ms para a chamada à API
        Usa context com timeout de 10ms para persistência no banco SQLite
        Retorna apenas o campo "bid" para o cliente
        Expõe o endpoint /cotacao na porta 8080

    O cliente (client.go):
        Faz requisição HTTP ao servidor
        Usa context com timeout de 300ms
        Salva a cotação em um arquivo "cotacao.txt"
        Trata erros e timeouts apropriadamente

Os logs de erro serão exibidos caso algum dos timeouts seja excedido.