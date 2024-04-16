# Clean Architecture Challenge

Olá devs!
Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:

- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
  Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando
docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada
serviço.

## Resolução

Para resolução deste desafio foi realizado um fork do código fonte construído durante o módulo de Clean Architecture e
foi realizado as seguintes alterações:

- Adicionado método GetAll no `orders.reporitory.go` + adicionado testes
- Adicionado um novo usecase `get_orders.go` com um novo método Execute
- Adicionado nova rota HTTP no webservice, novos protos e novos graphql schemas para o novo serviço
- Atualizado injeção de dependências no wire
- Adicionado novos services nas camadas de webservice, gRPC e graphql que utilizam o novo usecase acima citado
- Adicionado `Makefile` para auxiliar na iniciação do projeto
- Adicionado package `migrations` que possibilita executar as migrations necessarias no banco antes de iniciar o servidor
- Alterado versão do MYSQL de 5.7 para 8 no `docker-compose.yaml` devido a problemas de compatibilidade da versão 5.7
  com meu MAC arm64

* OBS: o nome para o método de listar orders acabou sendo `GetOrders` e não `ListOrders` conforme mencionado no enunciado, porém a funcionalidade foi desenvolvida conforme indicado

### Rodando novo Serviço

Rodar o servidor para testar o novo serviço requer Docker instalado para possibilidar o uso das dependências MYSQL e
RabbitMQ - [Obter Docker](https://docs.docker.com/engine/install/)

Com o Docker instalado, seguir os seguintes passos

- Na raíz do projeto, executar no terminal `go mod tidy` para instalar todas as dependências
- Executar o comando `make start` no terminal. Ele execuratá o docker-compose e iniciará a aplicação GO

### Executando novo Serviço

Com o servidor em execução:

- Pelo WebServer HTTP
    - Abrir arquivo `create_order.http` dentro da pasta `api`
    - Crie algumas orders através da chamada `POST http://localhost:8000/order` com diferentes IDs e preços
    - Faça a chamada `GET http://localhost:8000/order/get HTTP/1.1` e verifique se as orders criadas foram retornadas
      com sucesso

- Pelo Servidor GraphQL
    - Acessar no navegador `http://localhost:8080/`
    - Digite a query abaixo e execute-a (remova/adicione os campos da query para verificar os diferentes retornos)
  ```
  query queryOrders {
    orders {
      id,
      Price,
      Tax,
      FinalPrice,
    }
  }
  ```
  
- Pelo Servidor gRPC
  - Caso ainda não possua, instale o [Evans GRPC Client](https://github.com/ktr0731/evans)
  - Com o Evans instalado, executar comando no terminal `evans -r repl`
  - Acessar serviço através do comando `service OrderService`
  - Chamar serviço de listar orders através do comando `call GetOrders`

Obs: como já criamos as orders através do web server, nos outros servidores listei apenas os passos para listá-las, porém também é possível criar orders através dos respectivos serviços também disponíveis nestes servidores
