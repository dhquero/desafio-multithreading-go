# Mutithreading

Para executar o programa, execute:

```bash
> go run main.go
```

## Execução

Será solicitado a digitação do CEP para busca das informações em 2 APIs em paralelo. Será exibido as informações da API que retornar primeiro, sendo o timeout para a leitura dos dados de 1 segundo.

APIs: Via Cep e Brasil API

Os dados que serão exibidos na tela são: CEP, estado, cidade, bairro, rua.


### Via CEP

```bash
> go run .\main.go
CEP: 17500100
Via CEP. CEP: 17500-100. Estado: SP. Cidade: Marília. Bairro: Centro. Rua: Rua Coronel Galdino de Almeida
```

### Brasil API

```bash
> go run .\main.go
CEP: 17500100
Brasil API. CEP: 17500100. Estado: SP. Cidade: Marília. Bairro: Centro. Rua: Rua Coronel Galdino de Almeida
```
