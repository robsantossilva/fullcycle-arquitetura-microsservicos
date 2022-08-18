# Arquitetura Baseada em Microsserviços

[Mind Map](https://whimsical.com/arquitetura-baseada-em-microsservicos-UU7VzzjVnofKr2at3iZ4Rx)

## Introdução

#### Conceitos básicos
- Microsserviço é uma aplicação comum
- Possuem objetivos bem definidos
- Faz parte de um ecossistema
- São independentes ou autônomos; possui banco próprio
- Microsserviços se comunicam o tempo todo

#### Microsserviços vs  Monolíticos
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

#### Quando utilizar Microsserviços /  Monolíticos

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

#### Migração de Monolito para Microsserviços

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

#### 9 Caracteristicas dos Microsserviços por Martin Fowler

- **Componentização via serviços:** é uma unidade de software que é substituível e atualizável independentemente. **Serviços** são componentes fora do processo que se comunicam através de requisições ou chamadas procedure call. **Serviços** possui deploy independente.
- **Organizado em torno das areas de negócio:** Organizar os times por areas de negócio e não por funções
- **Produtos e não Projetos:** 