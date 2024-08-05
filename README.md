# ingestion-txt-file

## Como rodar o projeto
### Requisitos
- Docker Engine instalado na versão mais recente
### Execução
- Para rodar o projeto, é necessário apenas entrar na pasta em que o arquivo docker-compose.yml se encontra e utilizar o comando `docker-compose up --build`

## Lógica utilizada
- Inicialmente a main.go realiza os seguintes passos para inicializar o backend corretamente -> Carrega as variáveis de ambiente, abre a conexão com o banco de dados, inicia a migração da tabela que será utilizada, abre o arquivo que será lido e chama a função que irá realizar a ingestão dos dados e posteriormente salvar no banco de dados.
- Após a main.go chamar a função `CreateDataIngestion` na camada de business, a camada de business escaneia o arquivo, divide cada linha em um array onde cada campo ocupa uma posição e em seguida chama a função de realizar o parse das colunas para as variáveis na estrutura.
- A função que realiza o parse dos dados do arquivo é a `parseData`, que recebe um array de strings e realiza o parse dos dados para a estrutura `Ingestion`.
- Logo em seguida, a parseData realiza a chamada da função `validateFields`, que é responsável por validar os campos da ingestão, verificando se os campos `CPF`, `LastPurchaseStore` e `MostFrequentStore` são válidos.
- Após as devidas validações, o método CreateDataIngestion finaliza a ingestão chamando o método CreateDataIngestion do repositório de dados para criar os dados na tabela `Ingestion` no banco de dados

## Tecnologias utilizadas
- Linguagem: Golang com DDD
- Banco de dados: PostgreSQL
- Docker
