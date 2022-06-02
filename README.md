<h1 align="center">Bank Account API</h1>  

### A aplicação representa um sistema de contas bancárias com operações básicas como criação de conta, depósito, saque, transferência e extrato.

Cada operação possui uma taxa de seviço:

- Depósito: 1% sobre o valor depositado 
- Saque: R$ 4,00 sobre o valor do saque 
- Transferência: R$ 1,00 

## Ferramentas utilizadas
- `gorilla.mux`
- `gorm-io`
- `postgreSQL`
- `pgAdmin4`
          
## Endpoints:
- `/api/customer/`
- `/api/account/{accountNumber}`
- `/api/deposit`
- `/api/withdrawal`
- `/api/transfer`
- `/api/statement`

## Configurando o ambiente de desenvolvimento

No terminal, execute o comando abaixo para criar o contêiner com as imagens do `pgAdmin` e `postgres`.

```shell
docker-compose up
```
Em seguida, execute os comandos abaixo para buscar o hostname do postgreSQL:
```shell
docker-compose exec postgres sh
# hostname -i
```
Abra o navegador, digite `http://localhost:12345/` e insira os valores:

- email: email@email.com
- senha: 123456

Para configurar o server no pgAdmin4 utilize o hostname e as seguintes variáveis:

```shell
- POSTGRES_USER=root
- POSTGRES_PASSWORD=root
- POSTGRES_DB=root
```

![Screenshot_3](https://user-images.githubusercontent.com/71775013/171734396-8c68a95c-bf5a-4d35-afab-b89762a8a270.png)

O projeto possui o arquivo que criará o modelo de tabelas utilizado pela aplicação.

```shell
docker-database-initial.sql
```

## Iniciando a aplicação

Para iniciar a execução da aplicação, insira o comando no terminal:

```shell
go run main.go
```

## Cadastrar cliente

Para cadastrar um cliente, execute uma request do tipo POST para o endpoint `api/customer` seguindo o modelo abaixo: 

*Request*:
```json
{
	"Name": string, 
	"Document": string,
	"Email": string,    
}
```
*Response*
```json
{
	"AccountId": number,
	"AccountNumber": string,
	"Balance": number,
	"Holder": {
		"Id": number,
		"Name": string,
		"Document": string,
		"Email": string
	}
}
```
## Buscar dados da conta por número de conta

Para buscar uma conta por número de conta, execute uma request do tipo GET para o endpoint `api/account/{accountNumber}` seguindo o modelo abaixo:

*Response*
```json
{
	"AccountId": number,
	"AccountNumber": string,
	"Balance": number,
	"Holder": {
		"Id": number,
		"Name": string,
		"Document": string,
		"Email": string
	}
}
```

## Realizar depósito

Para realizar um depósito, execute uma request do tipo POST para o endpoint `api/deposit` seguindo o modelo abaixo:  

*Request*:
```json
{
	"AccountNumber": string,
	"Amount": number
}
```
*Response*
```json
{
	"Deposit made successfully"
}
```
## Realizar saque

Para realizar um saque, execute uma request do tipo POST para o endpoint `api/withdrawal` seguindo o modelo abaixo:

*Request*:
```json
{
	"AccountNumber": string,
	"Amount": number
}
```
*Response*
```json
{
	"Withdrawal made successfully"
}
```

## Realizar transferência

Para realizar uma transferência, execute uma request do tipo POST para o endpoint `api/transfer` seguindo o modelo abaixo:

*Request*:
```json
{
	"OriginAccountNumber": string,
	"DestinationAccountNumber": string,
	"Amount": number
}
```
*Response*
```json
{
	"Transfer made successfully"
}
```

## Emitir extrato

Para emitir um extrato, execute uma request do tipo POST para o endpoint `api/statement` seguindo o modelo abaixo:

*Request*:
```json
{
	"AccountNumber": string
}
```
*Response* 
```json
{
	"AccountId": number,
	"CurrentBalance": number,
	"Statement": [
		{
			"OperationId": number,
			"OriginAccountId": number,
			"DestinationAccountId": number,
			"OperationType": "Deposit",
			"Amount": number,
			"ServiceCharge": number,
			"OperationDate": "2000-01-01T00:00:00Z"
		},
		{
			"OperationId": number,
			"OriginAccountId": number,
			"DestinationAccountId": number,
			"OperationType": "Withdrawal",
			"Amount": number,
			"ServiceCharge": number,
			"OperationDate": "2000-01-01T00:00:00Z"
		},
		{
			"OperationId": number,
			"OriginAccountId": number,
			"DestinationAccountId": number,
			"OperationType": "Transfer",
			"Amount": number,
			"ServiceCharge": number,
			"OperationDate": "2000-01-01T00:00:00Z"
		}
	]
}
```
