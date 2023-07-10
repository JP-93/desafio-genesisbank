# Desafio Genesisbank

<h2 align="center"> API conversão de moedas </h1>

# Inciando projeto
- Certifique-se que tenha o docker intalado e rodando.
- Pode usar o programa que preferir para fazer a chama rest.
- Para que o projeto compile precisamos de um banco de dados, optei por postgreSQL

# Run
- Navegue ate a raiz do projeto e em seu terimal use os comandos:
    - docker-compese build
    - docker-compose up
    - Após esses comando o nosso banco de dados está pronto para ser populado.

- Navegue atá a pasta onde contém nosso main.go
    - cmd/main.go >> comando >> go run main.go
    - A aplicação já está funcionado para fazer as convesões.

- Endpoint para convesão:
    - GET http://localhost:8000/exchange/{amount}/{from}/{to}/{rate}
    - substitua o amout pelo valor desejado
    - from - moeda origem 
    - to -  moeda destino
    - rate -  cotação atual da moeda

- Endpoint para consulta geral:
    - GET http://localhost:8000/exchanges




