+++
title = "Hextest REST API testing"
postname = "hextestv01"
description = "Meu projeto principal atualmente, Testar REST Api's de maineira fácil e simples"
created_at = "18 Sep 2025"
last_updated = "18 Sep 2025"
author = "Poroca"
tags = ["daedalus"]
cover = "/images/hextest.jpg"
+++

# Sobre testar Apis REST

Logo que entrei no meu emprego atual, tive a sensação que fui chamado para resolver os problemas na automação do projeto. Não tinha tanta esperiência com o framework de testes que eles utilizavam (Mocha Chai), e demorou um certo tempo para ter eles voltando a funcionar aos poucos.

### Só eu mexia naquilo

Aparentemente antes de eu ser chamado, outra Engenheira QA trabalhou nos testes automatizados mas desde o dia que ela foi demitida, até o dia que eu saí desse projeto, só eu tocava e interagia com esse repositório, isso fez com que, quando eu fui realocado de projeto, novamente outra pessoa teve que passar pelo mesmo processo penoso de tentar entender o que estava acontecendo ali. E isso não era limitado pela senioredade da pessoa, e sim pelos próprios testes.

## Quero voltar a programar

Em agosto desse ano me veio uma realização de que eu estava acomodado, e queria criar um projeto para voltar a por em prática as coisas relacionadas a programação. Meu lado desenvolvedor estava adormecido, e sempre que eu precisava de algum código eu primeiro tentava conseguir ele através da IA.

<figure>
    <img src="/images/gopher.jpg" alt="photo">
    <figcaption>
        <p>Gopher</p>
    </figcaption>
</figure>

Escolhi Go como linguagem por causa do charme, sempre gostei de linguagens mais low level, mas não queria escolhe algo como C ou Zig por causa da aplicação que eu queria construir, e Rust me dá medo sinceramente. Go é rápido e simples, e me faz focar em resolver os problemas e não na linguagem em si. Sou um tanto odioso com toda a cultura de desenvolvimento web hoje em dia e seus frameworks, e não sou fã de javascript também. Então Go it is.

## A ferramenta

O plano era claro, uma ferramenta que lêsse um JSON que estivesse com os dados do projeto e dos testes nele presentes, montasse um struct dentro da ferramenta, fizesse as requisições http e voltasse com os resultados. O resultado está bem próximo disso mas como chegamos a (Projeto)->(Suite)->(Teste)->(Assert)->(Check) só as madrugadas programando que conseguem responder...

### Atual

Eu considero que a ferramenta está numa versão pré alpha -1.0 (isso mesmo, versionamento negativo, cadê seu deus agora?) mas temos um front end para editar os testes sem ter que mexer manualmente no json para quem não tem interesse, e principalmente, os testes rodam. Mas ainda há muito a se melhorar, o front é horrendo e pode ter uma UX infinitamente melhor, o back poderia gerar relatorios melhores em vez de fazer um printf dos resultados, e muitas outras coisas, mas eu sou só um infelizmente.
 Vou atualizando por aqui o processo de desenvolvimento dessa ferramenta!

<figure>
    <img src="/images/hextest.jpg" alt="photo">
</figure>
