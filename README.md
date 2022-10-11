# Arquitetura Baseada em Microsserviços

[Mind Map](https://whimsical.com/arquitetura-baseada-em-microsservicos-UU7VzzjVnofKr2at3iZ4Rx)

### 1. Introdução

### 2. Conceitos básicos
- Microsserviço é uma aplicação comum
- Possuem objetivos bem definidos
- Faz parte de um ecossistema
- São independentes ou autônomos; possui banco próprio
- Microsserviços se comunicam o tempo todo

### 3. Microsserviços vs  Monolíticos
- Objetivos / Dominio bem definido:
  - Microsserviços: Objetivos definidos. Contexto definido
  - Monolíticos: Toda aplicação. Todos os contextos dentro do mesmo sistema.
- Linguagens de Programação:
  - Microsserviços: Diversas tecnologias
  - Monolíticos: Única tecnologia
- Deploy
  - Microsserviço: Menor risco de indisponibilidade
  - Monolitico: Risco maior de tudo cair
- Organização das equipes:
  - Microsserviço: Equipes definidas por contexto
  - Monolitico: Todos no mesmo sistema.
- Começar um projeto / POC:
  - Monolitico: É mais simples.

### 4. Quando utilizar Microsserviços /  Monolíticos

|                                                     | Microsserviços   | Monolíticos       |
| --------------------------------------------------- | ---------------- | ----------------- |
| Inicio de projeto                                   | Complexo         | Simples           |
| Escalar Times                                       | Simples          | Pode ser complexo |
| Contextos bem definidos / Área de negócio           | Pode valer apena | -                 |
| Tem manuridade nos processos de entrega?            | Pode ser viavel  | -                 |
| Tem manuridade técnica dos times?                   | Pode ser viavel  | -                 |
| Necessidade de escalar apenas uma parte do sistema? | Melhor opção     | -                 |
| Tecnologias especificas em partes do sistema?       | Melhor opção     | -                 |
| POC                                                 | -                | Melhor opção      |
| Sem conhecimento completo do Dominio?               | -                | Melhor opção      |
| Governança simplificada sobre tecnologias           | -                | Melhor opção      |
| Facilidade na contratação                           | -                | Melhor opção      |
| Facilidade no treinamento dos devs                  | -                | Melhor opção      |
| Compartilhamento facilidade de Libs (Shared kernel) | -                | Melhor opção      |

### 5. Migração de Monolito para Microsserviços

- Separação de Contextos (DDD)
- Evite excesso de granulidade
- Verifique dependencias.
- Planeje o processo de migração dos bancos de dados
- Eventos (Comunicação assíncrona)
- Não tenha medo de duplicação de dados
- Lidar com consistencia eventual
- CI/CD/Testes/Ambientes
- Comece pelas beiradas
- Padrão estrangulamento

### 6. 9 Caracteristicas dos Microsserviços por Martin Fowler

- **Componentização via serviços:** é uma unidade de software que é substituível e atualizável independentemente. **Serviços** são componentes fora do processo que se comunicam através de requisições ou chamadas procedure call. **Serviços** possui deploy independente.
- **Organizado em torno das areas de negócio:** Organizar os times por areas de negócio e não por funções
- **Produtos e não Projetos:**
- **Smart endpoints and dumb pipes**
- **Governança descentralizada**
- **Gerenciamento de dados descentralizados**
- **Automação de Infraestrutura**
- **Desenhado para falhar**
- **Design evolutivo**

### 7. Resiliência

Em algum momento todo sistema vai falhar...

Quais as praticas/estratégias para mitigar riscos e ajudar o sistema a se tornar mais resiliênte em momentos de falhas?

**O que é resiliência?**
- É um conjunto de estratégias adotadas intencionalmente para a **adaptação** de um sistema quando uma falha ocorre.
- Ter estratégias de resiliência nos possibilita minimizar os riscos de perda de dados e transações importantes para o negócio.

**Quais são estas estratégias?**
- **Proteger e ser protegido:**
  - Um sistema em uma arquitetura distribuida precisa adotar mecanismos de **autopreservação** para garantir ao máximo sua operação com qualidade.
  - Um sistema não pode ser "egoísta" ao ponto de realizar mais requisições em um sistema que está falhando.
  - Um sistema lento no ar, muitas vezes, é pior que um sistema fora do ar. (Efeito dominó)
- **Health Check:**
  - Sem sinais vitais, não é possivel saber a "saúde" de um sistema.
  - Um sistema que não está saudável possui uma chance de se recuperar caso o tráfego pare de ser direcionado a ele temporariamente. **(Self Healing) - Auto Cura**
  - Health check de qualidade
- **Rate Limit:**
  - Protege o sistema baseado no que ele foi projetado para suportar
  - Preferência programada por tipo de client
- **Circuit Break:**
  - Protege o sistema negando requisições. Ex: 500
  - Circuito fechado = Requisições chegam
  - Circuito aberto = Requisições não chegam
  - Meio aberto = Permite uma quantidade limitada de requisições e verifica se o sistema possui condições de voltar ao ar integralmente

#### API Gateway
-  Garante que requisições "inapropriadas" cheguem até o sistema: Ex.: usuário não autenticado
-  Implementa politicas de Rate Limiting, Health check, etc

#### Service Mesh
- Controla o tráfego de rede
- Evita implementações de proteção pelo próprio sistema
- mTLS
- Circuit breaker, retry, timeout, fault injection, etc

#### Trabalhar de forma assíncrona
- Evita perda de dados
- Não há perda de dados no envio de uma transação se o server estiver fora
- Servidor pode processar a transação em seu tempo quanto estiver online
- Entender com  profundidade o message broker/sistema de stream

#### Garantias de entrega: Retry
- Linear backoff
- Exponential backoff
- Exponential backoff - Jitter

#### Garantias de entrega: Kafka

1) Producer apenas envia mensagem
[Producer] ---> (Ack 0 : none) ---> [Broker A] Leader

1) Producer envia mensagem e recebe confirmação do Leader
[Producer] ---> (Ack 1 : Leader) ---> [Broker A] Leader
[Producer] <------------------------------- [Broker A] Leader

1) Producer envia mensagem, Leader recebe e replica para os Followers, e em seguida avisa o producer 
[Producer] ---> (Ack -1 : Leader) ---> [Broker A] Leader
[Broker A] ------------------------------>> [Broker B] Follower
[Broker A] ------------------------------>> [Broker C] Follower
[Producer] <<------------------------------ [Broker A] Leader

#### Situações complexas
- O que acontece se o message broker cair?
- Haverá perda de mensagens?
- Seu sistema ficará fora do ar?
- Como garantir resiliência?

#### Transactional outbox
Antes de mandar a mensagem para o Kafka salvar em um banco de dados.
O Kafka confirmando o recebimento exclui o registro
| MessageID | Key     | Topic               | Payload                 |
| --------- | ------- | ------------------- | ----------------------- |
| 1234      | Account | account_transaction | {"id":"1dwef1hg1e5g"... |

#### Garantia de recebimento
No RabbitMQ, após a aplicação ter tirado a mensagem para ser processada, ele exclui a mensagem.
E se a aplicação cair e não conseguir processar? A mensagem se perdeu.
Mas, é possivel dizer ao rabbitmq para não excluir usando o parametro:
Auto Ack = false
Dizer de forma manual para o RabbitMQ que a mensagem foi lida.

- **Auto Ack = false** e commit manual
- **Prefetch** alinhado a volumetria. Applicação recebe mensagens em batch

#### Microsserviços

- Idempotência: É a capacidade de conseguir ligar com duplicidade de dados.
- Independência: Ex: Banco de Dados
- Politicas claras de fallback

#### Observabilidade

- APM
- Tracing distribuido
- Métricas personalizadas
- Spans personalizados
- Open Telemetry

Exponential backoff and Jitter: https://aws.amazon.com/pt/blogs/architecture/exponential-backoff-and-jitter/
Remédio ou Veneno - https://www.youtube.com/watch?v=1MkPpKPyBps
OTEL - https://opentelemetry.io/

### 8. Coreografia vs Orquestração

#### Como funciona a Coreografia?
Coreografia é uma técnica para composição de serviços de forma distribuída e descentralizada, vista sob uma perspectiva global, onde não há um nó coordenador; cada nó sabe o que deve fazer e como colaborar com seus vizinhos na coreografia.
- **Microsserviços Independentes**

#### Dinamica de Orquestração
A Orquestração é capaz de integrar sistemas de forma melodica e harmônica. Ela dita o ritmo da integração, invocando o serviço certo no momento certo, informando cada uma das entradas. Mas esse ‘maestro’, o nosso músico, é totalmente dependente do serviço, ao ponto de limitar-se à tarefa de realizar uma requisição e obter ou não uma resposta.
- **Microsserviços Dependentes de um orquestrador**
- Garante a sequencia
- Planos de fallback

#### Estratégias de APIs
- Mini API Gateways por contexto de Microsserviços