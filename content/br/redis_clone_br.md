+++
title = "Fazendo um clone do Redis"
postname = "redis_clone"
description = "Aprendendo sobre o Redis programando um clone dele em Go"
created_at = "30 Aug 2026"
last_updated = "30 Aug 2026"
author = "Poroca"
tags = ["daedalus"]
cover = "/images/redis_go.png"
+++

# Bem vindo de volta Eu, 

Estou voltando com posts no blog e pretendo continuar a criar e desenvolver minhas habilidades, principalmente como desenvolvedor, continuo trabalhando como QA, mas talvez venha uma mudança de carreira futura.

## O projeto:
Esse projeto é para aprender um pouco sobre o Redis, como funciona e como usar, e também voltar a desenvolver algo pessoal, por simples motivo de aprendizado. No momento estou implementando as features básicas que um banco tem que ter, você pode acompanhar esse projeto no [GitHub](https://github.com/GuPoroca/Redis_clone).

## Primeiros passos:
Eu não tinha muito conhecimento sobre o que era o Redis além de que ele servia como cache e para armazenar dados. após uma pesquisa rápida aprendi sobre o protocolo RESP2, redis-cli, e os comandos básicos do Redis, e como ele otimiza utilizando a própria memória dinâmica como um banco de dados.

## Para o clone:
Vou começar apenas com o básico, comandos de SET, GET, DELETE, EXPIRE, PERSIST... e no momento ele só suporta armazenamento estilo dicionário com strigs:strings, mas isso vai mudar num futuro próximo.

## Uso de IA
Nesse projeto estou tentando integrar parcialmente o meu development flow com o uso de agentes para me ajudar principalmente a entender o código e me guiar em próximos passos enquanto estou desbravando águas desconhecidas. Como desenvolvo no linux, atualmente testando entre o Neovim e o [Pardes](https://git.sr.ht/~gbrls/pardes) (editor do meu amigo gbrls), fazer o workflow funcionar com agentes tem sido um pouco diferente do que estou acostumado no trabalho.

## Por hoje é isso pessoal!
Pretendo ir atualizando com outros blogposts conforme o projeto for tomando forma, mas por enquanto é isso, também gostaria de fazer um post sobre o próprio site, a pipeline de deploy dele e como eu recriei o Hugo sem querer, mas isso fica para outro post.
