# t

`t` é uma ferramenta minimalistica de anotação

## Filosofia

Ao invés de usar um editor web ou algum navegador inteiro disfarçado de
uma aplicação nativa, ou até mesmo um aplicativo lento e pesado para
simplesmente escrever texto e armazenar isso na "nuvem" de alguém usando
os programas e APIs proprietárias, optamos  por um caminho bem mais
simples, estar sobre os ombros de duas ferramentas altamente confiáveis,
bem documentadas e livres/abertas: *git* e o formato de texto *markdown*.

## Arquitetura

`t` é feito para ser completamente simples e fora do teu caminho.
Especialmente quando escrevendo anotações, tudo o que você precisa é um
lugar para jogar um monte de pensamentos ali e deu.

Quando você está escrevendo uma anotação, `t` vai abrir o editor de
texto configurado na variável de ambiente `EDITOR` (ou se não estiver
declarado, abrirá o vim). Assim que você terminar de escrever e fechar
o aplicativo, `t` vai jogar as suas anotações para o servidor git
configurado. É isto.

## Uso

`t` tem uma pequena lista de comandos:

- `list` ou `l` vai listar todas as anotações indexado pela data de
modificação, enquanto passa essa lista adiante para o que estiver
configurado na variável de ambiente `PAGER` no seu sistema (geralmente
`less (1)`).

- `create` ou `c` vai criar uma nova anotação usando a data atual. Caso
você queira você pode passar um título para a anotação.

- `edit` ou `e` vai reabrir a anotação no seu `EDITOR`, e por padrão vai
abrir a ultima anotação feita. Você pode passar um indice para esse
comando.

- `show` ou `s` vai jogar a anotação para dentro do [glow][glow]
e mostrar na sua tela. Esse comando também recebe um índice.

- `share` ou `sh` vai passar a anotação para o serviço de
compartilhamento configurado, por padrão ele manda para uma instancia
pública do [rascunho][rascunho] hospedado em
[eletrotupi.com][eletrotupi], mas você pode customizar para utilizar
o seu serviço favorito.

[rascunho]: https://sr.ht/~porcellis/rascunho
[glow]: https://github.com/charmbracelet/glow
[eletrotupi]: https://eletrotupi.com

Consulte o manual local e offline `man t` se você precisar de alguma
ajuda.

## Instalando

Por enquanto, a única forma de instalar é clonar esse repositório
e rodar os comandos de instalação.

Então, abra seu terminal e digite esses comandos:
- `git clone https://git.sr.ht/~porcellis/t`
- `make && sudo make install`

Se você tem algum interesse em empacotar esse projeto para sua
distribuição, por favor me avise para que eu possa atualizar as
instruções aqui.

**Aviso**: Por enquanto, o comando `show` depende de já ter sido
efetuada a instalação do [glow][glow].

## Contribuindo

A forma padrão, boa e velha do git.

Existe algumas dependências, em especial o `git`, `go`, `less` e o já
mencionado [glow][glow]. Todas essas ferramentas, provavelmente, estão
disponíveis de serem instaladas usando o seu gerenciador de pacotes da
sua distribuição (apk no Alpine, pacman no Arch Linux, apt em
distribuições baseadas no Debian, etc). Depois disso, você pode usar
o comando `make get` dentro da pasta do repositório para baixar todas as
dependências e bibliotecas usadas pelo `go get` e começar a modificar.

[Envie seus patches e modificações][git-send-email] para minha caixa de
entrada pública: [~porcellis/public-inbox@lists.sr.ht][public-inbox]

(Lembre-se de usar [apenas texto][plain-text])!

[public-inbox]: mailto:~porcellis/public-inbox@lists.sr.ht
[plain-text]: https://apenastexto.eletrotupi.com
[git-send-email]: https://git-send-email.io

## Futuro

Tem várias coisas que eu gostaria de implementar, algumas delas são:

- Adicionar uma forma melhorada de visualização quando listar
anotações, provavelmente usando ncurses
- ~~Oferecer suporte ao Pager do sistema operacional quando listar
anotações, hoje pode ser resolvido paliativamente com o comando `t list
| less -r`~~
- Deletar uma anotação
- Adicionar um comando para configurar o projeto passo-a-passo, criar
repositório, adicionar um servidor, etc.
- Adicionar suporte a encriptação das anotações (isso é prioridade)
- Permitir editar uma anotação pelo seu título
- Aplicativo nativo desktop [tw](https://git.sr.ht/~porcellis/tw)

## Licença

Sob os termos da licença GNU GPL-3.0, verifique o arquivo
[LICENSE](https://git.sr.ht/~porcellis/t/tree/master/LICENSE) para mais
informações.
