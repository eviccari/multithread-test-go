## Multithread Test - Go Version

    Este projeto tem por objetivo analisar vantagens e desvantagens ao utilizar [Golang](https://go.dev/) em processos/rotinas Batch oa invés de Java com [Spring Boot](https://spring.io/projects/spring-boot).

## Pontos Avaliados
    - Performance (tempo de processamento)
    - Consumo de recursos (memória e cpu)
    - Tamanho de Imagem Docker
    - Legibilidade + Manutenibilidade (um pouco subjetivo, deve ser analisado por pares)

## Tecnologias
    - Golang 1.19
    - MySQL ${latest}
    - Docker Engine
    - Docker Compose

## Arquitetura Alvo
    - Domain Driven Design 
    - Arquitetura Hexagonal      

![Clean Architecture](./docs/images/Clean_Architecture.png)
<p style="text-align:center;">https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html</p>

## Escopo Funcional

    Este projeto FAKE deverá importar dados de usuários do github a partir da API oficial/aberta (ou mockada) para um banco MySQL. O processo de leitura será sequencial ordenado pelo ID do usuário, respeitando o limite de paginação da API. O processo de insert deverá ser executado utilizando threads independentes.
    A quantidade de usuários a ser importada e o limite de threads abertas devem ser parametrizáveis.

## Estrutura de Projeto   




````bash
/cmd -> package main (executável)
/configs -> package responsável por carregar configurações e variáveis de ambiente
/docs -> contém imagens e packages relacionadas a documentação do job
/internal -> packages internas da aplicação 
    /adapters -> oferece interfaces p/ desacoplar componentes de baixo nível 
    /app -> 
        /domain -> domínio da aplicação com entidades, regras de negócio e orquestração de casos de uso
            /dtos -> objetos para transportar dados entre as camadas 
            /models -> entidades e políticas de negócio com nomenclatura alinhada entre TI e Negócio (Linguagem Ubíqua)
            /orchestrators -> 
````
 

## Execução do job (executar na raiz do projeto)
````bash
docker-compose build && docker-compose up -d
````

    A configuração das imagens fará com que o job sempre "restarte" ao finalizar, possibilitando avaliar os KPIs ao longo da execução.  


    

