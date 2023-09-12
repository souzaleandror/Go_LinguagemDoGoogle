##### 12/09/2023

CURSO Go: a linguagem do Google

@01-Introdução

@@01
Introdução

Seja bem-vindo ao curso de Go, a linguagem de programação do Google, aqui da Alura.
O Google estava com um problema, pois muitos dos seus sistemas eram feitos em C++ e em C, e o processo de compilar esses programas, para gerar um executável, era complicado e demorado. Com isso, os engenheiros do Google tiveram a ideia de criar uma nova linguagem de programação, surgindo daí o Go.

Essa linguagem, diferentemente de outras linguagens, como C, Python e Java, é uma linguagem moderna, visto que ela foi criada em 2009, aproximadamente 20 anos depois que a maioria das outras linguagens foi criada. E nesses 20 anos, a computação foi evoluindo, evoluções essas que foram incorporadas no Go.

A primeira dessas evoluções faz com que o processo de compilação de um programa em Go é bem rápido, justamente o problema do Google. Um outra questão é a modularização da linguagem, com as suas funcionalidades espalhadas em pacotes, que são importados na nossa aplicação conforme a nossa necessidade.

Além disso, a sua sintaxe tem em torno de 25 palavras-chave, ou seja, uma sintaxe bem curta, fazendo com que o nosso código fique simples de se entender. Apesar da linguagem ser simples, ela é bem convencionada, já que o Go foi feito para ser uma linguagem rápida de desenvolvimento, então há certas convenções (algumas delas forçadas pela linguagem) que fazem com que o foco do programador seja em desenvolver código, e não em discutir como ele deve ser feito.

Outra coisa interessante é o jeito com que o Go trabalha com concorrência, paralelismo da linguagem. Essa é uma parte avançada, que não é o nosso foco, que é a introdução à linguagem, vendo a sua sintaxe, convenções, pacotes, tudo isso criando um projeto, um monitorador de sites, para colocar todos os sites que administramos e ficar verificando se o site está online ou não, e guardar um log para sabermos o horário em que o site estava online ou offline.

Então, vamos criar uma aplicação capaz de se comunicar com os sites, e ficar monitorando para ver se eles estão online ou caíram.

Estão prontos? Então vamos começar!

@@02
Estrutura do Curso

A estrutura do curso
Além dos vídeos com a transcrição, neste curso temos alguns tipos de exercício para você consolidar seus conhecimentos:

Os exercícios de múltipla escolha são para você rever os conceitos dos vídeos, para você relembrar os detalhes de cada função ou conceito e fixar bem os conhecimentos.

Os exercícios Mãos na Massa são para você implementar o que fizemos em vídeo, com um texto te guiando para que você não esqueça nenhum detalhe.

O Download do projeto
E no primeiro exercício Mão na Massa de cada capítulo também temos o link para Download do projeto completo até aquele ponto, para caso você precise pular alguma coisa ou comparar com o código que você mesmo fez!

@@03
Instalação do Go

Se queremos aprender a trabalhar com a linguagem de programação Go, primeiramente devemos instalá-la. Para isso, no site da linguagem (https://golang.org/), há um link de download (https://golang.org/dl/).
Então, basta fazer o download e instalar a última versão da linguagem para o nosso sistema operacional.

Para saber mais, disponibilizamos este exercício com mais detalhes da instalação para cada sistema operacional.
Ao término da instalação, o Go será habilitado para ser utilizado na linha de comando. Para verificar a sua instalação, executamos no terminal/linha de comando:

go versionCOPIAR CÓDIGO
Executando o comando acima, teremos uma saída parecida com esta:

go version go1.8.1 linux/amd64COPIAR CÓDIGO
Mas além da instalação do Go, devemos seguir outros padrões para trabalhar com essa linguagem.

Go Workspace
No momento da instalação é definido um local onde ficará todo o nosso código Go, esse local é o Go Workspace. Por padrão essa pasta deve ser chamada go e deve ficar na raiz do nosso usuário. Isso vale para todos os sistemas operacionais, ou seja, na pasta do nosso usuário, nós criamos a pasta go.

Além disso, o Go Workspace deve possuir três pastas. A primeira delas é a pkg, onde ficará os pacotes compartilhados das nossas aplicações, pois o Go é uma linguagem bastante modular, dependendo bastante de pacotes que vamos importando no nosso código, veremos isso mais à frente. Além da pasta pkg, também deve ter a pasta src, onde escreveremos o código fonte de cada aplicação, e a pasta bin, onde ficará os compilados do nosso código Go.

A estrutura final do Go Workspace ficará assim:

pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── srcCOPIAR CÓDIGO
Mas se quisermos que o Go Workspace seja em outro lugar, pode? Pode ser, mas por padrão o instalador do Go já espera que essa pasta esteja na raiz do usuário, então, para nos facilitar, vamos seguir os padrões da linguagem.

Editor de texto
Mas além disso, precisamos de um editor de texto que nos ajude a codificar em Go. Um editor de texto que é bastante integrado com a linguagem, que será utilizado no treinamento, é o Visual Studio Code.

Sabendo disso, podemos fazer o nosso primeiro programa em Go, o famoso Hello World. Então, dentro da pasta go/src, vamos criar a pasta hello. E através do Visual Studio Code, nós abrimos essa pasta e dentro dela criamos o programa hello.go.

Assim que criamos o arquivo, recebemos uma notificação do Visual Studio Code: The 'Go' extension is recommended for this file type. Ou seja, existem extensões do editor que nos facilita o trabalho com a linguagem. Ao clicar em Show Recommendations, nós vemos essas extensões, e uma delas é a Go, do autor lukehoban.

Nós podemos instalar essa extensão, clicando em Install e logo em seguida em Reload, para recarregarmos a janela. No momento em que a janela for recarregada, a extensão será habilitada, mas ela tem algumas dependências, conforme é exibida na notificação do Visual Studio Code. Então, basta clicar em Install All, para instalar todas as dependências, que facilitarão a nossa codificação em Go.

Por isso é uma recomendação utilizar o Visual Studio para programar Go, pois ele facilita muito o nosso processo. Algumas funcionalidade desse editor serão exploradas ao londo do treinamento

Após o término da instalação das dependências, no próximo vídeo vamos dar início ao nosso primeiro programa em Go, o Hello World.

https://golang.org/

https://golang.org/dl/

@@04
Preparando o ambiente: Instalando o Go

Para instalarmos o Go em nosso computador primeiramente precisamos fazer o Download do binário da linguagem adequado para cada sistema operacional. Visite a página de download oficial da linguagem e faça o download adequado para o seu sistema. Neste curso usaremos a versão 1.8.X do Go:
Imagem dos downloads

Agora vamos seguir para a instalação para cada uma das plataformas:

Instalando para o Windows
Abra o arquivo MSI que contêm a instalação e siga as instruções na tela para instalar o Go. O instalador é bem simples e basta apenas confirmar a instalação com todas as opções que vierem por padrão. O instalador coloca a instalação do Go em c:\Go e não precisamos interferir nisto.

O instalador deve colocar o diretório c:\Go\bin na variável de ambiente PATH do seu sistema. É necessário reiniciar qualquer prompt de comando aberto para que a alteração da variável entre em vigor e você possa testar a instalação.

Ao final da instalação , abra um novo prompt de comando e execute o comando go version e veja se a versão e os comandos do Go foram exibidos corretamente.

Instalando para o MacOS
Baixe o arquivo .pkg, abra-o, e siga as instruções para instalar o Go. Assim como no Windows, o instalador para MacOS é bem simples e apenas devemos confirmar as opções que vierem por padrão e seguir com a instalação. O pacote instala a linguagem Go em /usr/local/go.

O instalador normalmente coloca o diretório /usr/local/go bin na sua variável de ambiente PATH. Talvez seja necessário reiniciar quaisquer sessões abertas do Terminal para que a alteração entre em vigor.

Ao final da instalação , abra um novo prompt de comando e execute o comando go version e veja se a versão e os comandos do Go foram exibidos corretamente.

Instalando para o Ubuntu
Faça o download do arquivo .tar do Go e navegue com o terminal até a pasta aonde o arquivo foi baixado. Use o comando abaixo para extraí-lo em sua pasta /usr/local:

sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gzCOPIAR CÓDIGO
Repare que neste exemplo estamos utilizando a versão 1.8.3 da linguagem, caso você faça o Download de uma versão mais nova você deve alterar o comando para que o nome do arquivo corresponda ao que foi baixado.

Em seguida precisamos adicionar o caminho /usr/local/go/bin no PATH do sistema. Você pode fazer isto adicionando uma linha extra no arquivo /etc/profile, por exemplo com o editor de textos gedit. Abra o arquivo /etc/profile com o gedit utilizando o comando abaixo:

sudo gedit /etc/profileCOPIAR CÓDIGO
E em seguida adicione a seguinte linha no final do arquivo:

export PATH=$PATH:/usr/local/go/binCOPIAR CÓDIGO
Salve a alteração e feche o editor.

Agora utilize o comando:

source /etc/profileCOPIAR CÓDIGO
E o Go já deve estar funcionando com sucesso! Faça o teste em seu terminal com o comando go version e veja se a versão e os comandos do Go foram exibidos corretamente.

https://golang.org/dl/

@@05
Preparando o ambiente: Criando o Workspace

Criando o seu workspace
Para utilizarmos a linguagem de programação Go devemos seguir algumas convenções, e a principal delas é a sua organização de pastas no seu Workspace.

O Workspace padrão do Go é um diretório chamado /go que fica na pasta do seu usuário em seu sistema operacional. No Windows esta pasta normalmente se encontra em C:/Users/seu-usuario/ , e nos sistemas Unix ( MacOS e distribuições do Linux) normalmente se encontra em /home/seu-usuario/.

Dentro deste diretório devem conter as seguintes pastas:

pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── srcCOPIAR CÓDIGO
A bin , onde ficará os compilados do nosso código Go.

A pkg, onde ficará os pacotes compartilhados das nossas aplicações, pois o Go é uma linguagem bastante modular, dependendo bastante de pacotes que vamos importando ao longo de nossos códigos.

A src, onde escreveremos o código fonte de cada aplicação.

Crie estas pastas como indicado acima, para que o seu Go Workspace fique funcionando corretamente.

@@06
Preparando o ambiente: Visual Studio Code

Instalando o Visual Studio Code
Neste treinamento recomendamos que você utilize o Visual Studio Code para os seus códigos em Go, visto que o mesmo tem excelentes plugins para a linguagem que facilitam o seu seu desenvolvimento, principalmente neste fase inicial.

Você pode fazer o download do Visual Studio Code AQUI , ele é gratuito e está disponível para diversos sistemas operacionais.

Plugin do Go
Para sua vida seja facilitada quando você está trabalhando com Go no Visual Studio Code, você deve instalar os plugins recomendados por ele. Se você criar qualquer arquivo com a extensão .go, uma pequena pop up surgirá perguntando se você deseja instalar a extensão.

Antes de prosseguir com a instalação, é importante ter o Git instalado em sua máquina , pois o Visual Studio Code utiliza-o para baixar e instalar o plugin. Se você ainda não tem o Git em seu computador, você pode conferir como se instala AQUI no curso de Git da Alura.

Ao clicar em Show Recommendations você verá no Menu lateral a extensão chamada "Go", do autor lukehoban. Selecione-a clicando em Install e em seguida ao final da instalação em Reload para recarregamos a janela.

No momento em que a janela for recarregada, a extensão será habilitada, mas ela tem algumas dependências, conforme é exibida na notificação do Visual Studio Code. Então, basta clicar em Install All, para instalar todas as dependências, que facilitarão a nossa codificação em Go.

https://code.visualstudio.com/

https://cursos.alura.com.br/course/git-github-controle-de-versao/task/57007

@@07
Criando o Hello World

Para construirmos o nosso primeiro programa em Go, primeiramente devemos entender algumas instruções iniciais dele. Se queremos que o nosso programa seja um executável, ou seja, se queremos que depois o programa seja executado no nosso computador, devemos declarar o programa assim:
package mainCOPIAR CÓDIGO
Essa linha informa que o nosso programa será o principal pacote da nossa aplicação, e que o nosso código deve começar por ele. Além disso, todo programa principal tem a função principal. As funções em Go são declaradas com func, seguida do seu nome e os argumentos entre parênteses, com o bloco de código ficando entre chaves:

package main

func main() {

}COPIAR CÓDIGO
Uma das características da função principal é não retornar nada e nem receber argumento nenhum. Quando executamos o programa, essa função é iniciada, e quando ela é finalizada, o programa termina.

Mas como imprimimos a mensagem de Hello World?

Exibindo uma mensagem
Se queremos imprimir uma mensagem, devemos utilizar a função Println, mas para utilizá-la, devemos importá-la, do pacote fmt. Esse é um pacote de formatação, possuindo funções de impressão e escaneamento do texto:

package main

import "fmt"

func main() {

}COPIAR CÓDIGO
Com o pacote importado, podemos utilizar a função:

package main

import "fmt"

func main() {
    fmt.Println()
}COPIAR CÓDIGO
O fato da função Println estar com a primeira letra maiúscula pode causar uma estranheza para quem vem de uma outra linguagem de programação. Quando estamos trabalhando com Go, a função que vem de um pacote externo, ou seja, uma função que não está declarada no nosso arquivo, é utilizada com a primeira letra maiúscula. Nós chamamos o pacote externo (no nosso caso, o fmt), e a função com a primeira letra maiúscula (no nosso caso, Println). Isso é uma convenção da linguagem, que faz com que saibamos que a função veio de um pacote externo

Agora, falta somente passarmos a mensagem que queremos exibir como argumento para a função Println:

package main

import "fmt"

func main() {
    fmt.Println("Olá Mundo")
}COPIAR CÓDIGO
Resta agora executar esse programa.

Executando um programa em Go
Como Go é uma linguagem compilada, para executar um programa seu, devemos compilá-lo para um executável, e para isso, nós devemos utilizar o terminal/linha de comando.

Nele, nós entramos dentro da pasta go/src/hello e executamos o comando go build seguido do nome do programa que queremos executar:

go build hello.goCOPIAR CÓDIGO
Caso não dê nenhum erro, ao verificar o conteúdo da pasta, temos o seguinte:

alura@alura-01:~/go/src/hello$ ls
hello  hello.goCOPIAR CÓDIGO
Ou seja, o executável foi criado. Agora basta executá-lo:

alura@alura-01:~/go/src/hello$ ./hello 
Olá MundoCOPIAR CÓDIGO
Toda vez que alterarmos o nosso código, devemos compilá-lo para depois executá-lo, mas o Go é uma linguagem que facilita muita coisa para nós, e isso inclui o processo de compilação. Em vez de executar o go build, para depois rodar o executável gerado, podemos executar o comando go run seguido do nome do programa que queremos executar:

alura@alura-01:~/go/src/hello$ go run hello.go 
Olá Mundo, alunos!!!COPIAR CÓDIGO
O comando go run compila e executa o programa de uma vez só! Isso evita com que tenhamos sempre que executar dois comandos para executar um simples programa.

Terminal embutido do Visual Studio Code
Por último, estamos trabalhando com dois programas, o Visual Studio Code e o terminal, mas o Visual Studio Code já possui um terminal embutido! No menu superior, acessando View -> Integrated Terminal.

Nele, nós podemos trabalhar do mesmo jeito como estávamos trabalhando antes, com a facilidade de ter tudo em uma mesma tela, de programar Go sem sair do Visual Studio Code.

@@08
Run, Forrest, Run!

Forrest desenvolveu um programa em Go em arquivo chamado forrest.go. Que comando ele deverá executar no terminal para ao mesmo tempo compilar e executar o programa recém criado?
Alternativa correta
go forrest.go
 
Alternativa correta
forrest
 
Alternativa correta
go build forrest.go
 
Alternativa correta
go run forrest.go
 
Isso aí! Com esse comando compilamos e executamos o programa forrest.go numa tacada só!
Uma das características de projeto da linguagem é a facilidade de utilização. O comando go run foi criado para economizar o trabalho de usar um comando para compilar e outro para executar um programa.

@@09
Convenções da linguagem

Antes de encerrar o Hello World, podemos reparar que, diferente do C e do Java, nós não colocamos ponto e vírgula ao final da instrução. Em Go, o ponto e vírgula ao final da instrução é opcional e se tentarmos colocá-lo no nosso programa, ao salvá-lo, veremos que ele é removido pelo Visual Studio Code!
O que acontece é que uma das extensões que instalamos no primeiro vídeo, faz nós termos boas práticas de programação. Apesar do uso de ponto e vírgula ser opcional, é uma recomendação não utilizá-lo, por isso ele é removido pela extensão no momento em que salvamos o nosso programa.

Uso das chaves
Outra convenção é sobre o uso das chaves, elas devem ficar do lado da função. Se tentarmos colocá-la em baixo, por exemplo:

package main

import "fmt"

func main() 
{
    fmt.Println("Olá Mundo")
}COPIAR CÓDIGO
O próprio editor já acusará erro de sintaxe. Ao tentar executar o programa, o mesmo erro será apontado no terminal.

O Go tem convenções que fazem com que o desenvolvedor não fique perdendo tempo discutindo, por exemplo se a chave precisa ser do lado ou embaixo. Go é uma linguagem para simplificar, então ele já diz que a convenção dele é a chave ser do lado da função, e se não fizer isso, ocasionará um erro.

No caso do ponto e vírgula, se não estivéssemos utilizando o Visual Studio Code e tentássemos compilar o programa com ponto e vírgula, até iria funcionar, mas é uma recomendação não utilizá-lo.

Ao longo do treinamento, veremos outas convenções da linguagem.

@@10
Principal

Observe a listagem abaixo. Marque as alternativas que representam as correções que precisam ser feitas para que o código compile e seja executado como um programa.
function main() {
  fmt.Println("Hello, Gopher!");
}COPIAR CÓDIGO
Alternativa incorreta
Importar o pacote que possui a função para imprimir uma linha no terminal `import "fmt".
 
A linguagem Go é bastante modularizada. Por isso, para utilizar seus pacotes em seus programas, eles devem ser devidamente importados. O pacote fmt possui funções de impressão no terminal e formatação.
Alternativa incorreta
Substituir function por func
 
func é a palavra reservada que representa uma função no Go.
Alternativa incorreta
Incluir a definição de pacote principal package main no começo do arquivo.
 
A palavra reservada package define o pacote e o nome main define que esse arquivo é um programa.
Alternativa incorreta
Importar o pacote que possui a função para imprimir uma linha no terminal `use "fmt".
 
A palavra reservada usada no Go para importar pacotes é import.
Alternativa incorreta
Remover o ponto e vírgula após a instrução Println.
 
A linguagem Go possui várias convenções que ajudam o desenvolvedor a focar no problema a ser resolvido.

@@11
Função principal

Roberto está acompanhando o primeiro capítulo do curso de Go e escreveu o código abaixo.
package main

import "fmt"

func principal() {
  fmt.Println("Hello, world!")
}COPIAR CÓDIGO
Infelizmente o programa não está compilando. O que precisa ser feito para que o código passe a compilar?

Alternativa correta
Usar o comando go run hello.go
 
Alternativa correta
Importar o pacote de impressão no terminal através do comando import "terminal"
 
Alternativa correta
Mudar o nome da função principal para main.
 
A função principal de um programa deve ser chamada main.

@@12
Mãos na Massa: Hello World com Go

Neste primeiro capítulo vamos nos focar em criar o nosso Hello World inicial com a linguagem.
O primeiro passo é criar a pasta hello, que irá conter o nosso primeiro script em Go. Crie a pasta/hello dentro da pasta src do seu Go Workspace, que criamos anteriormente no exercício de Preparando o Ambiente.

pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── src
        └── helloCOPIAR CÓDIGO
Todos os projetos que você for desenvolver na linguagem Go ficarão em suas próprias pastas no diretório src.

Agora vamos criar o primeiro arquivo , o hello.go dentro pasta /hello:

pasta-do-usuario/
└── go
    ├── bin
    ├── pkg
    └── src
        └── hello
           └── hello.goCOPIAR CÓDIGO
E neste arquivo vamos criar o nosso primeiro programa.

Como todo projeto em Go, precisamos definir qual será o pacote inicial com a instrução:

//hello.go
package mainCOPIAR CÓDIGO
E também precisamos definir a função principal do programa,a nossa função main:

//hello.go
package main

func main(){

}COPIAR CÓDIGO
Como desejamos exibir uma mensagem, precisamos importar o pacote fmt, que contêm as funções de formatação da linguagem inclusive a função que imprime uma mensagem, a função fmt.Println():

//hello.go
package main

import "fmt"

func main(){

}COPIAR CÓDIGO
Para não dar azar e começar nossa jornada no mundo do Go com o pé direito, vamos começar com o Hello World clássico de todas as linguagens:

//hello.go
package main

import "fmt"

func main(){
    fmt.Println("Hello World com Go!")
}COPIAR CÓDIGO
Para executar nosso código, basta utilizarmos o comando go run hello.go no terminal dentro da pasta que contêm nosso arquivo com o código fonte do programa que o executável será automaticamente criado e executado :

// Terminal
go run hello.go
Hello World com Go!COPIAR CÓDIGO
Nosso primeiro programa está completo!

@@13
O que aprendemos?

O que aprendemos?
Instalação do Go
Go Workspace
A pasta bin para os arquivos binários, compilados
A pasta src para o código fonte
A pasta pkg para os pacotes compartilhados entre as aplicações
Instalação da extensão do Go no Visual Studio Code
Com isso temos autocomplete, detecção de erros, etc
Convenções da linguagem
Implementação do Hello World
Compilando e executando um programa em Go

@@14
Para saber mais: O mascote Gopher

Se você já buscou qualquer coisa sobre a linguagem de programação Go na internet, você deve ter encontrado este pequeno carinha abaixo:
Imagem do Gopher

Este é o Gopher, o icônico mascote da linguagem Go.

Sua história vem de quando a linguagem Go estava sendo criada, os engenheiros do projeto no Google precisavam de um logo e solicitaram à artista Renée French para que criasse um. O logo inicial da linguagem criado por ela foi o abaixo:

Logotipo inicial do logo

Em 2009, quando a linguagem ia ser disponibilizada para o mundo como um projeto Open Source, a artista Renée sugeriu adaptar o mascote que ela havia criado para uma camiseta de uma promoção em uma rádio, há mais de 15 anos, como o mascote da linguagem. Assim, o Gopher original foi criado:

Gopher original

Em 2011, no evento de tecnologia Google/IO, ia acontecer um grande lançamento na cloud da Google para um runtime que executasse a Golang, então eles encomendaram vários bonecos de pelúcia do Gopher, que foi quando ele surgiu pela primeira vez com a cor azul:

Gopher Plush

Desde então ele já surgiu em diversas formas, já que a artista Renée disponibilizou o Gopher com uma licença Creative Commons 3.0 ,em que você pode modificar a vontade o pequeno bichinho, contando que continue mantendo os créditos do original a ela.

Variações do Gopher

Por ser pequeno, simples e muito fofo o Gopher hoje é o maior ícone da linguagem Go e mais uma das características interessantes e únicas desta linguagem.

https://en.wikipedia.org/wiki/Ren%25C3%25A9e_French

#####

@02-Trabalhando com variáveis

@@01
Tipos das variáveis

Tá na hora de começarmos a construir o nosso monitorador de sites. Começaremos exibindo a mensagem inicial, no estilo "Olá, sr. Douglas", mais a versão do sistema. Começando pelo nome, ele irá variar, e não queremos que ele fique preso na frase, logo vamos extrai-lo para uma variável
Declarando uma variável em Go
Para declarar uma variável em Go, utilizamos a palavra var seguida do nome da variável mais o seu tipo. Como é um texto, o seu tipo será string, e vamos logo inicializá-la:

package main

import "fmt"

func main() {
    var nome string = "Douglas"
    fmt.Println("Olá sr.")
}COPIAR CÓDIGO
Agora, para concatenar a frase com a variável, passamos a mesma para a função Println, separando os elementos por vírgula:

package main

import "fmt"

func main() {
    var nome string = "Douglas"
    fmt.Println("Olá sr.", nome)
}COPIAR CÓDIGO
Falta agora exibirmos a versão do sistema, que será no modelo 1.0, 1.1, etc. Logo, a versão é um número flutuante, que no Go tem duas versões, o float32 para 32 bits, e float64 para 64 bits. Como o nosso número será pequeno, utilizaremos a versão de 32 bits:

package main

import "fmt"

func main() {
    var nome string = "Douglas"
    var versao float32 = 1.1
    fmt.Println("Olá sr.", nome)
    fmt.Println("Este programa está na versão", versao)
}COPIAR CÓDIGO
Variável sem valor atribuído
Do mesmo jeito, podemos exibir a idade, que é um número inteiro, logo um int:

package main

import "fmt"

func main() {
    var nome string = "Douglas"
    var idade int
    var versao float32 = 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)
}COPIAR CÓDIGO
A variável idade está vazia propositalmente, e ao executar o programa, temos a seguinte saída:

alura@alura-01:~/go/src/hello$ go run hello.go
Olá sr. Douglas sua idade é 0
Este programa está na versão 1.1COPIAR CÓDIGO
No Go, quando não atribuímos um valor a uma variável, ela assume um valor zerado, ou seja, se for um inteiro, seu valor será 0, se for um número flutuante, seu valor será 0.0, e se for uma string, seu valor será uma string vazia.

Não utilização de uma variável
Outra característica do Go é não podermos declarar uma variável e não utilizá-la. Por exemplo, se removermos a impressão da idade:

package main

import "fmt"

func main() {
    var nome string = "Douglas"
    var idade int
    var versao float32 = 1.1
    fmt.Println("Olá sr.", nome)
    fmt.Println("Este programa está na versão", versao)
}COPIAR CÓDIGO
Temos a seguinte saída:

alura@alura-01:~/go/src/hello$ go run hello.go
# command-line-arguments
./hello.go:7: idade declared and not usedCOPIAR CÓDIGO
Essa é mais uma convenção do Go, temos sempre que utilizar as variáveis que declaramos. Até por que, se não estamos utilizando a variável, não tem motivo dela estar ali.

@@02
Criando um número flutuante

George é dono de uma padaria e criou um programa em Go para exibir os produtos vendidos nela:
package main

import "fmt"

func main() {
    var precoLeite float = 2.99
    var precoOvo float = 3.99
    var precoPao float = 0.99

    fmt.Println("O preço do leite é R$", precoLeite)
    fmt.Println("O preço do ovo é R$", precoOvo)
    fmt.Println("O preço do pão é R$", precoPao)
}COPIAR CÓDIGO
Mas ao executar o programa, ele recebe o seguinte erro:

# command-line-arguments
./hello.go:6: undefined: floatCOPIAR CÓDIGO
O que aconteceu no programa do George para gerar esse erro?

Alternativa correta
O erro acontece pois o tipo de variável float no Go não existe, e sim suas versões, 32float para números de até 32 bits, e 64float, para números de até 64 bits.
 
Alternativa correta
O erro acontece pois o tipo de variável float no Go não existe, e sim suas versões, float32 para números de até 32 bits, e float64, para números de até 64 bits.
 
Alternativa correta! O tipo de variável float no Go não existe, porque ele possui tamanhos, de 32 ou 64 bits, que devem ser especificados no momento da sua declaração. Se a variável for de 32 bits, o seu tipo será float32, mas se for de 64 bits, o seu tipo será float64.
Alternativa correta
O erro acontece pois o tipo de variável float no Go não existe, o correto é double.
 
Alternativa correta
O erro acontece pois George esqueceu de colocar ponto e vírgula ao final das linhas de código.

@@03
Inferência de tipos

Vamos atribuir um valor à idade e voltar com a sua impressão:
package main

import "fmt"

func main() {
    var nome string = "Douglas"
    var idade int = 24
    var versao float32 = 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)
}COPIAR CÓDIGO
No Visual Studio Code, podemos perceber que as variáveis nome e idade estão com um sublinhado verde:

Variáveis com sublinhado verde

O que isso significa?

Inferindo os tipos das variáveis
O Visual Studio nos dá esse aviso dizendo que nós podemos omitir o tipo da variável, pois o Go consegue inferir o tipo dessas variáveis. Ele consegue entender que, se a variável começa e termina com aspas, ela é uma string. Da mesma forma, se temos um número inteiro, sem casa decimal, o Go entenderá que a variável é do tipo inteiro:

package main

import "fmt"

func main() {
    var nome = "Douglas"
    var idade = 24
    var versao float32 = 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)
}COPIAR CÓDIGO
Ao executar o programa, ele continua funcionando da mesma forma. Mas por que o Visual Studio não faz a mesma sugestão para a variável versao?

Ele não faz a mesma sugestão pois há dois tipos flutuantes, o float32 e o float64, logo o Go pode inferir qualquer um dos dois tipos à nossa variável, mas nada impede que façamos isso:

package main

import "fmt"

func main() {
    var nome = "Douglas"
    var idade = 24
    var versao = 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)
}COPIAR CÓDIGO
Para saber se o Go conseguir inferir corretamente o tipo das variáveis, podemos descobri-los, importando o pacote reflect e chamando a sua função TypeOf, passando para ela a variável que queremos saber o tipo:

package main

import "fmt"
import "reflect"

func main() {
    var nome = "Douglas"
    var idade = 24
    var versao = 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("O tipo da variável idade é", reflect.TypeOf(versao))
}COPIAR CÓDIGO
E temos a seguinte saída:

alura@alura-01:~/go/src/hello$ go run hello.go
Olá sr. Douglas sua idade é 24
Este programa está na versão 1.1
O tipo da variável idade é float64COPIAR CÓDIGO
Assim vemos que, por padrão, o Go infere o maior tipo flutuante para a variável, o float64. Então, a menos que queiramos especificar que queremos o tipo float32, podemos omitir a declaração do tipo, que quando o Go ver um número decimal, ele será capaz de inferir o seu tipo.

Declaração curta de variáveis
Para deixar o nosso código mais limpo ainda, podemos remover a palavra var das variáveis. Podemos fazer isso pois o Go possui um segundo operador de atribuição de variáveis, um mais "curto", que é o :=. Quando utilizamos esse operador, estamos dizendo ao Go que estamos declarando uma variável e atribuindo um valor a ela:

package main

import "fmt"

func main() {
    nome := "Douglas"
    idade := 24
    versao := 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)
}COPIAR CÓDIGO
Quando vermos o operador :=, significa que está sendo declarada uma variável, atribuindo o seu valor e inferindo o seu tipo. Além disso, esse operador possui outras vantagens, que veremos ao longo do treinamento.

@@04
Declaração curta de variáveis

Para deixar o código mais limpo, pode-se utilizar o operador de declaração curta de variáveis, informando ao Go que uma variável está sendo declarada e, ao mesmo tempo, um valor está sendo atribuído a ela.
Qual das alternativas abaixo representa uma declaração curta de variáveis?

Alternativa correta
var idade := 25
 
Alternativa correta
idade := 25
 
Alternativa correta! O operador de declaração curta de variáveis é o :=, assim é declarada a variável idade, e ao mesmo tempo atribuindo o valor 25 a ela.
Alternativa correta
idade = 25
 
Alternativa correta
idade =: 25

@@05
Uso de variáveis

Veja o código abaixo:
package main

import "fmt"

func main() {
    var nome string = "José"
    var peso float32 = 75.4
    fmt.Println("Olá, eu sou o", nome)
}COPIAR CÓDIGO
Esse programa gera um erro. Você sabe identificar a causa?

Alternativa correta
A variável peso é declarada mas não é utilizada.
 
Alternativa correta! É uma convenção do Go ter que utilizar as variáveis declaradas, e como a variável peso não foi utilizada, um erro acontece.
Alternativa correta
Ausência do ponto e vírgula ao final das linhas de código.
 
Alternativa correta
O tipo float32 não existe.

@@06
Exibindo o menu

Agora, após as mensagens de boas vindas e de versão do programa, queremos dar as opções para o usuário, se ele quer monitorar os sites, exibir os logs, etc, e vamos mostrar tudo isso em um menu!
Disponibilizaremos o menu imprimindo as suas opções para o usuário, ele irá digitar a opção desejada, e nós capturaremos o que foi digitado. Então, primeiramente devemos imprimir as opções:

package main

import "fmt"

func main() {
    nome := "Douglas"
    idade := 24
    versao := 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}COPIAR CÓDIGO
Capturando a entrada do usuário
Para capturar o que o usuário escrever, existe a função Scanf, também do pacote fmt. Ela recebe dois argumentos, um modificador (o que queremos receber como entrada, um inteiro, string, etc) e um ponteiro para a variável que guardará a entrada do usuário.

Então, vamos declarar essa variável:

package main

import "fmt"

func main() {
    nome := "Douglas"
    idade := 24
    versao := 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
}COPIAR CÓDIGO
E agora nós passamos para a função Scanf o valor "%d", que significa que esperamos receber um número inteiro, e a variável comando:

package main

import "fmt"

func main() {
    nome := "Douglas"
    idade := 24
    versao := 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scanf("%d", &comando)
}COPIAR CÓDIGO
Daqui a pouco falaremos sobre o &, que se encontra à frente da variável comando. Ou seja, a função salvará na variável o que o usuário digitar. E para verificar se a entrada está mesmo sendo salva na variável comando, vamos imprimi-la:

package main

import "fmt"

func main() {
    nome := "Douglas"
    idade := 24
    versao := 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scanf("%d", &comando)

    fmt.Println("O comando escolhido foi", comando)    
}COPIAR CÓDIGO
Ao executar o programa, ele fica travado esperando que digitemos algo. Ao digitar, temos a seguinte saída:

alura@alura-01:~/go/src/hello$ go run hello.go
Olá sr. Douglas sua idade é 24
Este programa está na versão 1.1
1- Iniciar Monitoramento
2- Exibir Logs
0- Sair do Programa
2
O comando escolhido foi 2COPIAR CÓDIGO
Entendendo o ponteiro
Sobre o & visto antes, ele significa o endereço da variável que queremos salvar a entrada, pois a função Scanf não espera uma variável, e sim o seu endereço, um ponteiro para a variável.

A variável nada mais é do que uma "caixa", onde guardamos dados. Essa "caixa" está em algum lugar da memória do nosso computador, e esse lugar, o endereço da nossa "caixa", é o que chamamos de ponteiro.

Para descobrir o endereço da variável, basta colocar o & à frente dela.

Devemos especificar o que esperamos ser digitado?
A variável comando é do tipo inteiro, logo, só pode receber números inteiros. Se sabemos disso, por que devemos especificar que esperamos receber um número inteiro na função Scanf, através do modificador "%d"? Na verdade, nós não precisamos.

Alguém do Go também pensou nisso, por isso criou a função Scan (sem a letra f). Nela, nós não precisamos especificar o modificador que esperamos:

package main

import "fmt"

func main() {
    nome := "Douglas"
    idade := 24
    versao := 1.1
    fmt.Println("Olá sr.", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)

    fmt.Println("O comando escolhido foi", comando)    
}COPIAR CÓDIGO
Ou seja, estamos esperando um dado, e quando o recebermos, colocamos dentro de uma variável inteira. Como a variável é inteira, a função sabe que deve esperar, receber valores inteiros.

Ao executar o programa, caso digitemos um valor não-inteiro, como uma letra, por exemplo, não será guardado nenhum valor no endereço da variável comando, logo ela continuará com o valor padrão atribuído a ela no momento da sua inicialização, que é 0. Ou seja, o ideal seria que detectássemos quando um valor não-inteiro fosse digitado.

É isso que veremos no próximo capítulo.

@@07
Leitura da entrada do usuário

No último vídeo, para capturar o que o usuário escrever, foi visto que existe a função Scanf, do pacote fmt. Ela recebe como argumento um modificador e um ponteiro para a variável que guardará a entrada do usuário, por exemplo:
var idade int
fmt.Scanf("%d", &idade)COPIAR CÓDIGO
A variável idade é do tipo inteiro, logo, só pode receber números inteiros. Então, como o Go sabe que a variável só pode receber números inteiros, o modificador %d, que representa um número inteiro, deixa de ser necessário. Por isso há uma outra função, que não recebe como argumento o modificador.

Essa função é a:

Alternativa correta
var idade int
fmt.Read(&idade)
 
Alternativa correta
var idade int
fmt.Check(&idade)
 
Alternativa correta
var idade int
fmt.Scan(&idade)
 
Alternativa correta! A função é a Scan, também do pacote fmt. Ela consegue inferir o tipo digitado pelo usuário, baseado no tipo da variável recebida.
Alternativa correta
var idade int
fmt.Sweep(&idade)

@@08
Mãos na Massa: Criando o Menu

Começando deste ponto? Você pode fazer o DOWNLOAD completo do projeto do capítulo anterior e continuar seus estudos a partir deste capítulo.
Neste capítulo vamos criar a mensagem de introdução do script e o menu que será exibido para o usuário, além de capturar a escolha que ele fez.

1 - Para começar a trabalhar com variáveis, vamos declarar o nome do usuário e sua versão em duas variáveis, utilizando o operador declaração de variáveis curto:

//hello.go
package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1
}COPIAR CÓDIGO
2 - Vamos exibir as mensagens de boas vindas com estas duas variáveis:

//hello.go
package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1

    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)
}COPIAR CÓDIGO
3 - Agora vamos exibir o Menu para o usuário escolher qual opção ele deseja. Serão 3 Println para exibir o Menu, um para cada opção:

//hello.go
package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1

    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}COPIAR CÓDIGO
4- Com o menu sendo exibido corretamente, precisamos capturar a escolha do usuário do nosso programa. Vamos colocar a opção escolhida na variável int comando e capturar o seu valor com a função fmt.Scan():

//hello.go
package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1

    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)
}COPIAR CÓDIGO
Atenção para não esquecer de passar o endereço (&) da variável comando para a função Scan

5- Por último imprima o valor da variável comando para garantir que estamos capturando com sucesso:

//hello.go
package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1

    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)

    fmt.Println("O valor da variável comando é:", comando)
}COPIAR CÓDIGO
Conseguimos criar nossa mensagem de introdução e o menu de opções do usuário, além de exibir qual opção ele escolheu. Já estamos avançando em nosso programa !

https://s3.amazonaws.com/caelum-online-public/624-golang/01/projetos/alura-golang-stage-fim-cap01.zip

@@09
O que aprendemos?

O que aprendemos?
Declaração de variáveis
Valor padrão das variáveis sem valor inicial
Inferência de tipos de variáveis
Descobrir o tipo da variável
Através da função TypeOf, do pacote reflect
Declaração curta de variáveis
Ler dados digitados do usuário
Através das funções Scanf e Scan, do pacote fmt
Mais convenções do Go
Variáveis e imports não utilizados são deletados

@@10
Para saber mais: Buscando melhor no Google
PRÓXIMA ATIVIDADE

Sabemos que o Google é grande aliado do programador, normalmente sendo o primeiro local aonde buscamos ajuda quando estamos com dificuldade em algum problema ou travados na resolução de um bug.
Porém, se você for buscar pelas suas dúvidas de programação em Go no Google, pode ser que dependendo do que você busque , você não encontre resultados esperados.

A palavra Go é muito comum na lingua inglesa, o que faz com que o Google muitas vezes não entenda direito o contexto da pergunta que estamos pesquisando e nos mostre resultados que não tem nada haver com programação!

E nem adianta recorrer para o Google em português, pois buscando por Go ele nos retorna o estado de Góias!

Goias no google

Por isto, a dica que dou para vocês é sempre que for realizar buscas online relacionados a linguagem de programação Go, utilizem o termo golang que é o termo comumente adotado por desenvolvedores do mundo quando estão falando da linguagem Go, assim suas buscas no Google e no Stackoverflow serão mais precisas!

Buscando por golang