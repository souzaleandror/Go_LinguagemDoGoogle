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

Sabemos que o Google é grande aliado do programador, normalmente sendo o primeiro local aonde buscamos ajuda quando estamos com dificuldade em algum problema ou travados na resolução de um bug.
Porém, se você for buscar pelas suas dúvidas de programação em Go no Google, pode ser que dependendo do que você busque , você não encontre resultados esperados.

A palavra Go é muito comum na lingua inglesa, o que faz com que o Google muitas vezes não entenda direito o contexto da pergunta que estamos pesquisando e nos mostre resultados que não tem nada haver com programação!

E nem adianta recorrer para o Google em português, pois buscando por Go ele nos retorna o estado de Góias!

Goias no google

Por isto, a dica que dou para vocês é sempre que for realizar buscas online relacionados a linguagem de programação Go, utilizem o termo golang que é o termo comumente adotado por desenvolvedores do mundo quando estão falando da linguagem Go, assim suas buscas no Google e no Stackoverflow serão mais precisas!

Buscando por golang

#### 13/09/2023

@03-Controle de fluxo com Go

@@01
Controle de fluxo com if

Dando sequência à construção da nossa aplicação, agora que já conseguimos capturar o comando escolhido pelo usuário, podemos fazer algo de acordo com o comando digitado.
Se o usuário escolheu o comando 1, vamos fazer algo, se ele escolheu o comando 2, fazemos outra coisa, e assim por diante. Quando queremos saber se o usuário escolheu tal comando, precisamos utilizar a instrução if da programação.

A condição do if, no Go, não fica entre parênteses, como já é prática de outras linguagens:

if comando == 1 {

}COPIAR CÓDIGO
E a condição deve sempre retornar um booleano, ou seja, true ou false. Como queremos testar os 3 comandos, vamos colocá-lo no if:

package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)
    fmt.Println("O comando escolhido foi", comando)

    if comando == 1 {

    } else if comando == 2 {

    } else if comando == 0 {

    }
}COPIAR CÓDIGO
E por último, se não for nenhum comando conhecido, vamos imprimir uma mensagem, demonstrando que não conhecemos o comando digitado. Vamos aproveitar e colocar as mensagens dos outros comandos:

package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)
    fmt.Println("O comando escolhido foi", comando)

    if comando == 1 {
        fmt.Println("Monitorando...")
    } else if comando == 2 {
        fmt.Println("Exibindo Logs...")
    } else if comando == 0 {
        fmt.Println("Saindo do programa...")
    } else {
        fmt.Println("Não conheço este comando")
    }
}COPIAR CÓDIGO
Com isso, concluímos os ifs, que não possui muito mistério, as diferenças para as outras linguagens é o não uso de parênteses na condição, que deve sempre retornar um booleano.

@@02
Discutindo maioridade...

Para qual das alternativas abaixo o Go vai permitir compilarmos o if abaixo?
if maiorDeIdade {

}COPIAR CÓDIGO
Alternativa correta
maiorDeIdade := true
 
Já que a variável é avaliada pelo Go como booleana por causa do valor de inicialização true, o if funcionará.
Alternativa correta
var idade int
maiorDeIdade = idade >= 21
 
Alternativa correta
maiorDeIdade := 23
 
Alternativa correta
idade int
maiorDeIdade := idade > 21
 
Lembre-se que o fluxo condicional if no Go só permite expressões booleanas.

@@03
Controle de fluxo com switch

Além do if, sabemos que existem outras instruções para controlarmos o fluxo da nossa aplicação. E uma que se adéqua melhor ao nosso código é a instrução switch.
Essa instrução recebe uma variável e dá possíveis situações para cada valor dessa variável. Nós dizemos quais são os casos, e para cada um há uma situação. Caso a variável valha 0, acontece algo, caso valha 1, acontece outra, e assim por diante.

switch comando {
case 1:
    fmt.Println("Monitorando...")
case 2:
    fmt.Println("Exibindo Logs...")
case 0:
    fmt.Println("Saindo do programa...")
}COPIAR CÓDIGO
Mas se o valor da variável não estiver em nenhum dos casos listados? Para isso, existe o caso default, que é o que será executado se os nossos casos não forem atendidos:

package main

import "fmt"

func main() {
    nome := "Douglas"
    versao := 1.1
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)
    fmt.Println("O comando escolhido foi", comando)

    switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa...")
    default:
        fmt.Println("Não conheço este comando")
    }
}COPIAR CÓDIGO
Ao executar o programa, ele funciona como antes, mas desta vez com uma nova instrução de controle de fluxo.

Uso do break
Para quem vem de outras linguagens de programação, pode estranhar o não uso do break, ao final do código de cada caso do switch. O break serviria para evitar a execução do código de mais de um caso, se mais de um caso for atendido.

No Go, ele não possui o break, pois somente um caso pode ser atendido. O primeiro caso que for atendido, terá o seu código executado e o switch será encerrado.

@@04
Cardápio apetitoso

Mário foi contratado por um restaurante para desenvolver um programa que solicita a escolha do cardápio. Ele implementou na linguagem Go e seu código está listado abaixo.
package main

import "fmt"

func main() {
    var prato string
    fmt.Println("Digite seu prato preferido...")
    fmt.Println("P - Pizza")
    fmt.Println("H - Hambúrguer")
    fmt.Println("B - Bife com fritas")
    fmt.Println("S - Salada Caesar")
    fmt.Println("F - Salada de Frutas")
    fmt.Println("E - Estrogonofe")
    fmt.Println("O - Outros")
    fmt.Scan(&prato)

    switch prato {
    case "B":
        fmt.Println("Com batatas Palito ou Noisete?")
    case "H":
        fmt.Println("Com Queijo ou com Ovo?")
    case "P":
        fmt.Println("Calabresa ou Napolitana?")
    case "S":
        fmt.Println("Alface ou Rúcula?")
    case "F":
        fmt.Println("Kiwi ou Frango?")
    case "E":
        fmt.Println("Carne ou Frango?")
    case "O":
        fmt.Println("Não gostou de nosso cardápio?")
    default:
        fmt.Println("Não entendi seu paladar...")
    }
}COPIAR CÓDIGO
Quando o cliente digitar b, o que será impresso?

Alternativa correta
Com batatas Palito ou Noisete?
 
Alternativa correta
Não entendi seu paladar...
 
Como não houve nenhum tratamento para reconhecer a caixa alta do valor entrado pelo usuário, a alternativa default será executada.
Alternativa correta
Bife com fritas
 
A palavra reservada switch permite incluir várias alternativas de execução para o programa. Cada alternativa deve utilizar case, e a alternativa executada quando nenhuma condição for atendida é default.

@@05
Como dizia Tim Maia, olha o breque!

Em outras linguagens, após cada alternativa é necessário colocar a palavra reservada break. Quais são os comportamentos do compilador da linguagem Go?
Alternativa correta
Não precisa colocar break.
 
No Go no uso do break não é obrigatório.
Alternativa correta
Se o desenvolvedor colocar break, o compilador não vai reclamar.
 
No Go no uso do break não é obrigatório.
Alternativa correta
O compilador obriga a colocar o break no final de cada alternativa.
 
Se o desenvolvedor colocar o break, o compilador não vai reclamar. Por exemplo na listagem abaixo.
switch comando {
case 1:
    fmt.Println("Monitorando... ")
    break
case 2:
    fmt.Println("Exibindo logs...")
    break
case 0:
    fmt.Println("Saindo do programa")
    break
default:
    fmt.Println("Não conheço esse comando")
    break
}COPIAR CÓDIGO
Porém, ele não é obrigatório. Apenas uma alternativa é executada por avaliação switch.

@@06
Introdução às funções

Queremos agora construir as funções do nosso programa, isso inclui a função que monitora os sites, que exibe os logs, que sai do programa, entre outras. Mas se implementarmos o código dentro de cada caso nosso, teremos uma centralização do nosso código na função main, ou seja, teremos pouco reuso de código, tornando-o pouco modularizado.
Por isso, chegou a hora de organizarmos o nosso código, quebrando-o em funções. Teremos uma função específica para exibir a mensagem de boas vindas, uma para capturar o comando do usuário, e uma função para cada caso do nosso programa.

Função para exibir a introdução do programa
Já sabemos como declarar uma função, já que já declaramos a função main, então vamos criar a função exibeIntroducao, com o nosso código de boas vindas:

package main

import "fmt"

func main() {

    exibeIntroducao()

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    var comando int
    fmt.Scan(&comando)
    fmt.Println("O endereço da minha variavel comando é", &comando)
    fmt.Println("O comando escolhido foi", comando)

    switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa...")
    default:
        fmt.Println("Não conheço este comando")
    }
}

func exibeIntroducao() {
    nome := "Douglas"
    versao := 1.1
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)
}COPIAR CÓDIGO
É comum em Go, utilizarmos o padrão camelCase para nome das funções, ou seja, a primeira letra minúscula e a cada nova palavra, sua primeira letra será maiúscula, sem espaço entre elas.

Retornando um valor em uma função
Agora, vamos criar a função leComando, para ler o comando digitado pelo usuário e retornar o seu valor para nós. No Go, colocamos o tipo do retorno da função após o nome da mesma:

func leComando() int {

}COPIAR CÓDIGO
E para retornar um valor, utilizarmos o return:

func leComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O comando escolhido foi", comandoLido)

    return comandoLido
}COPIAR CÓDIGO
Agora, na função main, nós chamamos essa função, já atribuindo-a à variável comando, utilizando o operador de declaração de variável curta:

package main

import "fmt"

func main() {

    exibeIntroducao()

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    comando := leComando()

    switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa...")
    default:
        fmt.Println("Não conheço este comando")
    }
}

func exibeIntroducao() {
    nome := "Douglas"
    versao := 1.1
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O comando escolhido foi", comandoLido)

    return comandoLido
}COPIAR CÓDIGO
O Go verá que a variável comando será o retorno da função leComando, que é um inteiro, logo a variável também será um inteiro.

Função para exibir o menu
Do mesmo jeito, vamos extrair para uma função o código que exibe o menu do nosso programa:

package main

import "fmt"

func main() {

    exibeIntroducao()
    exibeMenu()
    comando := leComando()

    switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa...")
    default:
        fmt.Println("Não conheço este comando")
    }
}

func exibeIntroducao() {
    nome := "Douglas"
    versao := 1.1
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}

func leComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O comando escolhido foi", comandoLido)

    return comandoLido
}COPIAR CÓDIGO
Agora que já temos uma noção de como construir funções básicas, vamos nos aprofundar nisso ao longo do treinamento, mas já podemos começar pela função que sai do programa.

Saindo do programa
Podemos criar uma função para sair do programa, mas o seu código possuiria somente uma linha, vamos fazer isso diretamente no seu caso, dentro do switch. Para sair do programa, é uma boa prática retornarmos um status 0 para o sistema operacional.

Para fazer isso com Go, precisamos importar o pacote que se comunica com o sistema operacional, o pacote os. Importando esse pacote, chamamos a sua função Exit, passando o valor 0 para ela, indicando para o sistema operacional que o programa foi encerrado com sucesso. A função main ficará assim:

package main

import "fmt"
import "os"

func main() {

    exibeIntroducao()
    exibeMenu()
    comando := leComando()

    switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa...")
        os.Exit(0)
    default:
        fmt.Println("Não conheço este comando")
    }
}COPIAR CÓDIGO
Do mesmo jeito, há uma forma de informar o sistema operacional que ocorreu algum problema na execução do programa, como por exemplo quando o usuário digita um comando desconhecido. Para isso, passamos o valor -1 para a função Exit. Então a função main ficará assim:

package main

import "fmt"
import "os"

func main() {

    exibeIntroducao()
    exibeMenu()
    comando := leComando()

    switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa...")
        os.Exit(0)
    default:
        fmt.Println("Não conheço este comando")
        os.Exit(-1)
    }
}COPIAR CÓDIGO
Quando o usuário digitar um comando desconhecido, ao invés de encerrar o programa, podemos pedir para o usuário digitar um novo comando. Veremos isso mais à frente.

@@07
Retornando valores numa função

Observe o código abaixo.
package main

import "fmt"

func main() {
    palavraDigitada := lePalavra()
    fmt.Println("A palavra digitada foi", palavraDigitada)
}

func lePalavra() string {
    var palavra string
    fmt.Print("Digite uma palavra: ")
    fmt.Scan(&palavra)
    return palavra
}COPIAR CÓDIGO
O que está faltando para que ele compile e a função retorne a palavra digitada para a função principal?

Alternativa correta
Nada, pois o código está compilando normalmente.
 
Isso aí! A função lePalavra retorna uma string com o valor digitado pelo usuário.
Alternativa correta
Colocar o símbolo & na variável de retorno da função, para indicar que queremos passar a referência da mesma.
return &palavra
 
Alternativa correta
Colocar o tipo antes da declaração do nome da função.
func string lePalavra() {

@@08
Saindo com graça

O que precisamos adicionar em nosso programa para sair dele com código de status de sucesso?
Alternativa correta
import "os"
os.Exit(0)
 
Isso aí! A função Exit fica no pacote os, que deve ser importado.
Alternativa correta
import "system"
os.Exit(0)
 
Alternativa correta
import "os"
os.Quit(0)
 
Alternativa correta
import "os"
os.Exit(-1)
 
Precisamos usar a função Exit que está disponível no pacote os.
Para isso é necessário importar o pacote os

import "os"COPIAR CÓDIGO
E no local devido chamar a função Exit com código 0, que representa uma saída bem sucedida.

os.Exit(0)

@@09
Mãos na Massa: Organização e Fluxo

Começando deste ponto? Você pode fazer o DOWNLOAD completo do projeto do capítulo anterior e continuar seus estudos a partir deste capítulo.
Neste capitulo vamos determinar para qual fluxo do nosso programa o usuário deve seguir a partir do comando que ele escolheu, e além disso vamos começar a organizar nosso código em pequenas funções.

1- O primeiro passo vai ser criar um switch com todas as opções que oferecemos ao usuário até agora, e um caso de default para quando o usuário colocar algum comando desconhecido:

//hello.go

package main

import "fmt"

func main() {
    // ... restante da função main
    var comando int
    fmt.Scan(&comando)

    fmt.Println("O valor da variável comando é:", comando)

    switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa...")
    default:
        fmt.Println("Não conheço este comando")
    }
}COPIAR CÓDIGO
_ Aqui colocamos uma opção no switch para cada item do menu que tinhamos no menu anteriormente_

2- Agora vamos começar a organizar nosso código em pequenas funções, cada uma com sua responsabilidade. A primeira a ser criada será a função que vai exibir a introdução de boas vindas para o usuário. Extrai este código da função main e coloque na função exibeIntroducao():

//hello.go
func main() {

    exibeIntroducao()

    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")

    // ...  restante da função main
}
func exibeIntroducao(){
    nome := "Douglas"
    versao := 1.1
    fmt.Println("Olá, sr(a).", nome)
    fmt.Println("Este programa está na versão", versao)
}COPIAR CÓDIGO
_ Não esqueça de chamar a função exibeIntroducao() na função main. _

3- Vamos extrair também o código que exibe o menu de opções do usuário para uma função externa, chamda de exibeMenu:

func main() {

    exibeIntroducao()
    exibeMenu()

    var comando int
    fmt.Scan(&comando)
    fmt.Println("O valor da variável comando é:", comando)
    //... restante da função

}

func exibeIntroducao() {
    //... restante da função
}

func exibeMenu(){
    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")    
}COPIAR CÓDIGO
_ Não esqueça de chamar a função exibeMenu() na função main. _

4- Por último, vamos criar uma função responsável por ler os comandos do usuário e exportar esta função da main também. A função leComando deve retornar um int com o comando lido pelo usuário:

func main() {

    exibeIntroducao()
    exibeMenu()
    comando := leComando()

    switch comando {        
    //... restante da função
    }

}

func exibeIntroducao() {
    //... restante da função
}

func exibeMenu(){
    //... restante da função
}

func leComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O valor da variável comando é:", comandoLido)

    return comandoLido
}COPIAR CÓDIGO
_ Não esqueça de capturar o comando lido na função main, para que ele possa ser utilizado pelo switch _

5- Vamos por último adicionar nos itens finais do switch um ponto de saída para o nosso script, retornando para o sistema operacional se tudo correu bem ou não. Importe o pacote os do Go e utilize a função Exit() para informar o status de saída do programa:

package main

import "fmt"
import "os"

func main() {

    exibeIntroducao()
    exibeMenu()
    comando := leComando()

    switch comando {
    case 1:
        fmt.Println("Monitorando...")
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa...")
        os.Exit(0)
    default:
        fmt.Println("Não conheço este comando")
        os.Exit(-1)
    }

}
    //... restante do arquivoCOPIAR CÓDIGO
Pronto, nosso código está mais organizado e agora só falta implementar as funções de cada um dos casos do nosso switch.

https://s3.amazonaws.com/caelum-online-public/624-golang/02/projetos/alura-golang-stage-fim-cap02.zip

@@10
O que aprendemos?

O que aprendemos?
Controle de fluxo com if
Sua condição não fica entre parênteses e deve sempre retornar um booleano
Controle de fluxo com switch
Se os casos não forem atendidos, será executado o código do caso default
Introdução às funções
Pacote os, para encerrar o programa

#### 

@04-Fazendo requisições para a web

@@01
Iniciando o monitoramento

Começando deste ponto? Você pode fazer o DOWNLOAD completo do projeto do capítulo anterior e continuar seus estudos a partir deste capítulo.
Agora que já começamos a separar o nosso código em funções menores, já podemos dar início ao principal da aplicação, que é monitorar os sites, para ver se os mesmos estão online ou não. Então, vamos começar criando a função iniciarMonitoramento:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
}COPIAR CÓDIGO
E já vamos chamá-la caso o usuário digite o comando 1:

package main

import "fmt"
import "os"

func main() {

    exibeIntroducao()
    exibeMenu()
    comando := leComando()

    switch comando {
    case 1:
        iniciarMonitoramento()
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa")
        os.Exit(0)
    default:
        fmt.Println("Não conheço este comando")
        os.Exit(-1)
    }
}

// restante do código omitidoCOPIAR CÓDIGO
Se queremos que o programa fique testando se o site está online, ou caiu, ele precisa acessar o site. Se queremos acessar um site, precisamos realizar uma requisição web, utilizando a linguagem Go.

Para fazer requisições web, existe um pacote especialista nisso, dentro do net, pacote de internet do Go, há o http, pacote responsável pelas requisições web:

package main

import "fmt"
import "os"
import "net/http"

// restante do código omitidoCOPIAR CÓDIGO
Agora, vamos monitorar o site da Alura, então, na função iniciarMonitoramento, vamos definir uma variável com o nome do site:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    site := "https://www.alura.com.br"
}COPIAR CÓDIGO
Com o site em mãos, vamos fazer uma requisição GET para o mesmo, utilizando a função Get, de http:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    site := "https://www.alura.com.br"
    http.Get(site)
}COPIAR CÓDIGO
Quando acessamos o site da Alura através do browser, obtemos uma resposta, que é carregada no navegador. A mesma coisa acontece quando carregamos o site através do Go, essa resposta vem através de um retorno da função Get, que iremos guardar na variável resp:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    site := "https://www.alura.com.br"
    resp := http.Get(site)
}COPIAR CÓDIGO
Ao salvar o arquivo, vemos que o Visual Studio Code aponta um erro, isso acontece porque a função Get retorna mais de um valor. Sim, existem funções no Go que retornam mais de um valor e a Get é uma delas, além da resposta, ela também retorna um possível erro que possa ter acontecido na requisição:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    site := "https://www.alura.com.br"
    resp, err := http.Get(site)
}COPIAR CÓDIGO
Novidade isso, né? Vamos avaliar essa questão de funções que retornam mais de um valor no próximo vídeo.

https://s3.amazonaws.com/caelum-online-public/624-golang/03/projetos/alura-golang-stage-fim-cap03.zip

@@02
Pacote para acesso a internet

Na construção do projeto de monitorização de sites teremos que efetuar uma comunicação com a Web. Isso é facilmente resolvido por que no Go já temos o seguinte pacote:
Alternativa correta
"net/http"
 
"net/http" é pacote mais específico à nossa necessidade. Já que nele temos funções para realizar requisições Get e Post.
Alternativa correta
"net/web"
 
Alternativa correta
"http"
 
Alternativa correta
"net"
 
Alternativa correta
"io/ioutil"
 
A utilização de pacotes providos pela linguagem é extremamente comum no desenvolvimento de Software. Com o Go não é diferente! Existem diversos pacotes para os mais variados propósitos.
Para uma comunicação web com um determinado site é necessário a utilização do pacote "net/http". Que pela própria definição dele tem como objetivo fornecer a implementações de cliente e servidor HTTP.

É importante saber que temos vários subdiretórios dentro do "net". Se quiséssemos fazer um envio de email poderíamos usar o "net/smtp".

Como referência para sabermos os pacotes da linguagem temos: https://golang.org/pkg/

@@03
Funções com múltiplos retornos

Para entendermos melhor a questão das funções que retornam mais de um valor, nada melhor do que aprendermos isso fazendo uma função desse tipo. Vamos criar a função devolveNome, que retorna um nome, logo uma string:
// restante do código omitido

func devolveNome() string {
    nome := "Douglas"
    return nome
}COPIAR CÓDIGO
E na função main, vamos guardar o retorno dessa função em uma variável e imprimi-la. Além disso, vamos comentar as chamadas das funções exibeIntroducao e exibeMenu, somente para não ficarmos com muita informação na tela:

// restante do código omitido

func main() {

    //exibeIntroducao()
    //exibeMenu()

    nome := devolveNome()
    fmt.Println(nome)

    comando := leComando()

    switch comando {
    case 1:
        iniciarMonitoramento()
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa")
        os.Exit(0)
    default:
        fmt.Println("Não conheço este comando")
        os.Exit(-1)
    }
}COPIAR CÓDIGO
Agora, além de devolver o nome, vamos fazer com quê a função também retorne a idade, para isso, basta separar os retornos por vírgula:

// restante do código omitido

func devolveNomeEIdade() string {
    nome := "Douglas"
    idade := 24
    return nome, idade
}COPIAR CÓDIGO
Além disso, precisamos adicionar o tipo int como retorno da função. Nesse caso, quando temos mais de um tipo, precisamos colocá-lo entre parênteses:

// restante do código omitido

func devolveNomeEIdade() (string, int) {
    nome := "Douglas"
    idade := 24
    return nome, idade
}COPIAR CÓDIGO
Logo, o primeiro valor que a função retorna é uma string, e o segundo é um número inteiro. E como agora a função retorna dois valores, nós precisamos criar duas variáveis:

// restante do código omitido

func main() {

    //exibeIntroducao()
    //exibeMenu()

    nome, idade := devolveNomeEIdade()
    fmt.Println(nome, "tem", idade, "anos")

    comando := leComando()

    switch comando {
    case 1:
        iniciarMonitoramento()
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa")
        os.Exit(0)
    default:
        fmt.Println("Não conheço este comando")
        os.Exit(-1)
    }
}COPIAR CÓDIGO
Com isso, conseguimos criar uma função que retorna dois valores, uma string e um inteiro.

Como ignorar algum retorno da função?
Isso é bem interessante no Go, e o mesmo ocorre com a função Get, ela retorna a resposta da requisição e um possível erro que possa ter ocorrido. Mas e se só estamos interessados em somente um dos retornos, como devemos fazer?

Quando não queremos saber de um dos retornos, queremos ignorá-lo, nós utilizamos o operador de identificador em branco (_):

// restante do código omitido

func main() {

    //exibeIntroducao()
    //exibeMenu()

    _, idade := devolveNomeEIdade()
    fmt.Println(idade)

    comando := leComando()

    switch comando {
    case 1:
        iniciarMonitoramento()
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa")
        os.Exit(0)
    default:
        fmt.Println("Não conheço este comando")
        os.Exit(-1)
    }
}COPIAR CÓDIGO
Esse operador diz para o Go ignorar o retorno, no nosso caso, o primeiro deles, pois só estamos interessados no segundo retorno. Então, se não queremos algum retorno de uma determinada função, no momento em que formos declarar as variáveis, basta utilizar esse operador no seu lugar.

Do mesmo jeito, como não queremos tratar possíveis erros de requisição nesse momento, podemos ignorar esse retorno:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)
    fmt.Println(resp)
}COPIAR CÓDIGO
Feito isso, podemos voltar com o código que havíamos comentado, e remover a função que criamos, visto que ela foi feita somente para entendermos as funções de múltiplos retornos:

package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {

    exibeIntroducao()
    exibeMenu()
    comando := leComando()

    switch comando {
    case 1:
        iniciarMonitoramento()
    case 2:
        fmt.Println("Exibindo Logs...")
    case 0:
        fmt.Println("Saindo do programa")
        os.Exit(0)
    default:
        fmt.Println("Não conheço este comando")
        os.Exit(-1)
    }
}

func exibeIntroducao() {
    nome := "Douglas"
    versao := 1.2
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
    fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}

func leComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O comando escolhido foi", comandoLido)

    return comandoLido
}

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)
    fmt.Println(resp)
}COPIAR CÓDIGO
Com isso, podemos seguir com a construção da nossa aplicação, a partir do próximo vídeo.

@@04
Capital do Estado

Manoela está fazendo um programa para imprimir informações sobre determinada cidade. Por enquanto o código está assim:
package main

import "fmt"

func main() {
    cidade, populacao, capital := devolveCidadeEPopulacao()
    if capital {
        fmt.Println("A capital ", cidade, "tem", populacao, "habitantes")
    } else {
        fmt.Println("A cidade ", cidade, "tem", populacao, "habitantes")
    }
}

func devolveCidadeEPopulacao() (string, int, bool) {
    return "Vila Sem Nome", 4328, true
}COPIAR CÓDIGO
Quando ela executar o programa, o que será impresso no terminal?

Alternativa correta
A capital Vila Sem Nome tem 4328 habitantes
 
Muito bem! Repare que o valor retornado pela função devolveCidadeEPopulacao para a variável capital é true.
Alternativa correta
Nenhum, porque o compilador não permite funções com mais de dois valores retornados.
 
Alternativa correta
A cidade Vila Sem Nome tem 4328 habitantes
 
Em funções com múltiplos valores, temos que informar os tipos de cada retorno entre parênteses. E declarar variáveis separadas por vírgula para receber os valores de saída. No exemplo, a função retornou valores para três tipos diferentes.

@@05
Mas e a sigla do estado?

Manoela se empolgou com seu pequeno programa e previu uma melhoria. Agora ela quer que a função retorne também o estado da cidade.
func devolveCidadeEPopulacao() (string, int, bool, string) {
    return "Vila Sem Nome", 4328, true, "RJ"
}COPIAR CÓDIGO
Mas por enquanto não quer utilizá-la no programa. O que ela deve fazer na função main abaixo para que o programa compile sem utilizar o estado?

func main() {
    cidade, populacao, capital := devolveCidadeEPopulacao()
    if capital {
        fmt.Println("A capital ", cidade, "tem", populacao, "habitantes")
    } else {
        fmt.Println("A cidade ", cidade, "tem", populacao, "habitantes")
    }
}COPIAR CÓDIGO
Alternativa correta
cidade, populacao, capital, _ := devolveCidadeEPopulacao()
 
Isso aí!
Alternativa correta
cidade, populacao, capital, null := devolveCidadeEPopulacao()
 
Alternativa correta
Nada. Manoela vai precisar necessariamente declarar uma variável para armazenar o estado, e em seguida utilizá-lo no programa. Assim:
cidade, populacao, capital, estado := devolveCidadeEPopulacao()
fmt.Println("Estado é", estado)
 
Quando queremos ignorar um valor retornado em uma função que retorna valores múltiplos, usamos o símbolo _ (underscore).

@@06
Monitorando os nossos sites

Agora que entendemos a questão das funções com múltiplos retornos, podemos ver o conteúdo da resposta da nossa requisição. Ao executar o programa e digitar o comando 1, temos uma resposta semelhante a essa:
&{200 OK 200 HTTP/2.0 2 0 map[X-Cloud-Trace-Context:[6f3fa7e590ac68bd43d76c82a67df476] Date:[Tue, 13 Jun 2017 21:20:36
GMT] Server:[Google Frontend] X-Ua-Compatible:[IE=edge,chrome=1] Expires:[Tue, 13 Jun 2017 21:50:36 GMT] Content-Type:[
text/html] Cache-Control:[public, max-age=1800] Age:[1298] X-Dns-Prefetch-Control:[on]] 0xc4200d4900 -1 [] false true m
ap[] 0xc42000a800 0xc4203a3080}COPIAR CÓDIGO
Nessa resposta, temos o status da requisição, a data, os cabeçalhos, entre outras informações. O que nos informa se um site carregou com sucesso, ou teve algum problema, é o status da requisição. Quando temos um status 200, significa que o site foi carregado perfeitamente. Então, se obtivermos algum status diferente de 200, significa que o nosso site está com problema.

Para saber o status da resposta, podemos acessar a sua propriedade StatusCode. Logo, podemos fazer um if, se o status for 200, nós imprimimos uma mensagem de sucesso, mas se não for, nós imprimimos uma mensagem dizendo que o site está com problema, imprimindo o status em seguida:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}COPIAR CÓDIGO
Ao executar o programa, recebemos uma mensagem de sucesso. Para não termos que ficar dependendo do site da Alura cair, podemos inventar uma URL inexistente, por exemplo:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    // site com URL inexistente
    site := "https://www.alura.com.br/askjdbahsbciahsbca"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}COPIAR CÓDIGO
Assim recebemos um status code 404. Ou podemos utilizar o site https://random-status-code.herokuapp.com/, que nos retorna um status diferente à cada requisição.

Colocando o nosso programa em loop
Conseguimos monitorar o site, mas assim que monitoramos, o programa é encerrado. O ideal é que o programa fique rodando em loop, eternamente, e só pare de rodar quando nós quisermos.

Em outras linguagens de programação, poderíamos utilizar o while, mas ele não existe no Go! Para isso, vamos utilizar a instrução for, sem passar nada para ela, pois assim ela ficará em loop eternamente:

package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    exibeIntroducao()

    for {
        exibeMenu()
        comando := leComando()

        switch comando {
        case 1:
            iniciarMonitoramento()
        case 2:
            fmt.Println("Exibindo Logs...")
        case 0:
            fmt.Println("Saindo do programa")
            os.Exit(0)
        default:
            fmt.Println("Não conheço este comando")
            os.Exit(-1)
        }
    }
}

// outras funções omitidasCOPIAR CÓDIGO
Agora, ao rodar o nosso programa, vemos que após digitar o comando e ter o seu código executado, o menu volta a ser exibido, logo o nosso código está em loop. Para sair do loop, basta que digitemos a opção 0.

Com isso, avançamos mais na nossa aplicação, nos próximos capítulos colocaremos mais sites para o programa monitorar, escrever no log, entre outras funcionalidades!

@@07
Resultado de uma requisição HTTP

Suponha que capturemos o resultado da função http.Get assim:
resultado, _ := http.Get("https://alura.com.br")COPIAR CÓDIGO
Qual propriedade usaremos para verificar o código de retorno de uma requisição HTTP?

Alternativa correta
resultado.HttpStatus
 
Alternativa correta
resultado.StatusCode
 
Isso aí!
Alternativa correta
http.StatusCode

@@08
Ao infinito e além!

O que Fabi deve colocar no seu programa Go para que ele tenha um loop infinito?
Alternativa correta
for {

}
 
Boa!
Alternativa correta
for () {

}
 
Alternativa correta
while true {

}

@@09
Mãos na Massa: Monitorando um site

Começando deste ponto? Você pode fazer o DOWNLOAD completo do projeto do capítulo anterior e continuar seus estudos a partir deste capítulo.
Neste exercício vamos iniciar o monitoramento de nossos sites e colocar o nosso programa para funcionar até que solicitemos o contrário.

1- Vamos implementar a primeira funcionalidade do nosso script, que é a de monitoramento de um site. Crie a função iniciarMonitoramento, que fará uso do pacote net/http para fazer uma requisição para um site da web. Depois faça a chamada dela no primeiro case do switch.

package main

import (
    "fmt"
    "net/http"
    "os"
)

func main(){
    //restante da função 

    switch comando {
    case 1:
        iniciarMonitoramento()
    //restante do switch
    }

}

// outras funções

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    site := "https://www.alura.com.br/"
    resp, _ := http.Get(site)
}COPIAR CÓDIGO
2- Vamos fazer uma verificação através do status code da requisição para averiguar se o site está online ou sofreu alguma queda. Teste para ver se o status code é 200 e exiba mensagens de acordo. Utilize o operador _ para não ter lidar com o segundo parâmetro de erro agora.

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}COPIAR CÓDIGO
3- Agora com a função de monitoramento pronta, vamos colocar um loop infinito na função main, para que o usuário possa acessar o menu repetidas vezes até que ele escolha sair. Para isto, utilize a instrução for pura:

func main() {

    exibeIntroducao()

    for {
        exibeMenu()
        comando := leComando()

        switch comando {
        case 1:
            iniciarMonitoramento()
        // restante do switch
        }

    }
}
COPIAR CÓDIGO
Pronto, agora o nosso usuário já consegue monitorar pelo menos um site até quando ele desejar sair do programa!

https://s3.amazonaws.com/caelum-online-public/624-golang/03/projetos/alura-golang-stage-fim-cap03.zip

@@10
O que aprendemos?

O que aprendemos?
Pacote net/http, com funcionalidades de acesso à internet, de requisições web
Entre elas, a função http.Get, para fazer uma requisição GET para um site
Funções com múltiplos retornos
Identificador em branco, para ignorar um ou mais retornos de uma função
Status de uma requisição
Uma requisição de sucesso possui status code 200
A instrução for, para deixar o nosso programa em loop eterno

#### 14/09/2023

@05-As principais coleções do Go

@@01
Arrays em Go

Neste capítulo, vamos evoluir o monitoramento do nosso programa, monitorando mais de um site ao mesmo tempo. Como programadores, não vamos criar uma variável para cada site que desejamos monitorar, e sim uma estrutura de dados responsável por conter várias strings, vários sites, o já conhecido Array.
Então, vamos ver como trabalhamos com coleções, principalmente Arrays e Slices, na linguagem Go.

Declarando um array
Primeiramente, vamos criar um array com a estrutura clássica, com o var. Para isso, além do var, nós devemos informar o nome do array, colchetes e o tipo de dados que ele guardará. Além disso, dentro dos colchetes, devemos informar o tamanho do array:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
    var sites [4]string
    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}COPIAR CÓDIGO
É importante nos atentar ao tipo de dados do array, não há espaço entre ele e os colchetes, na hora da declaração do array.

Colocando um valor dentro do array
Para colocar um valor no array, não há mistério, basta atribuir um valor a algum dos seus índices:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    var sites [4]string
    sites[0] = "https://random-status-code.herokuapp.com/"
    sites[1] = "https://www.alura.com.br"
    sites[2] = "https://www.caelum.com.br"

    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}COPIAR CÓDIGO
Mas o fato do array ter um tamanho fixo, nos limita um pouco, pois se quisermos adicionar 5 itens no array, teremos que alterar o seu tamanho na sua declaração. Por isso, em Go, geralmente não trabalhamos com arrays, e sim com uma outra estrutura de dados, chamada Slice, que funciona em cima do array, mas não tem tamanho fixo.

Veremos mais sobre isso no próximo vídeo.

@@02
Seu Valdomiro foi à feira
PRÓXIMA ATIVIDADE

Seu Valdomiro foi à feira e comprou 3 frutas. Sua bolsa tinha capacidade para guardar 4 frutas.
var frutas [4] string
frutas[0] = "Abacaxi"
frutas[1] = "Laranja"
frutas[2] = "Morango"COPIAR CÓDIGO
O que vai acontecer se executarmos o código abaixo?

fmt.Println(frutas[3])COPIAR CÓDIGO
Alternativa correta
Será impresso uma string vazia.
 
Isso aí!
Alternativa correta
O programa não vai compilar porque a "variável" alocada para a última posição do array não foi usada.
Alternativa correta
Acontecerá um erro na execução.
Quando os arrays são criados, eles assumem os valores padrão para os tipos de seus elementos. No caso, o tipo do array frutas é string e o valor padrão para cada posição do array será vazio. Portanto, o valor impresso será uma string vazia.

@@03
Para saber mais: funções que retornam arrays
PRÓXIMA ATIVIDADE

Considere o código abaixo.
package main

import (
    "fmt"
)

func main() {
    estados := devolveEstadosDoSudeste()
    fmt.Println(estados)
}

func devolveEstadosDoSudeste() [4]string {
    var estados [4]string
    estados[0] = "RJ"
    estados[1] = "SP"
    estados[2] = "MG"
    estados[3] = "ES"
    return estados
}COPIAR CÓDIGO
O que será impresso no terminal?

Alternativa correta
Nada, porque o programa não compila, uma vez que não é possível retornar um array através de uma função.
 
Alternativa correta
Nada, porque o programa não compila, uma vez que a declaração de retorno da função não exige a quantidade de posições do array.
 
Alternativa correta
[RJ SP MG ES]
 
Isso aí. O código acima cria uma função que retorna um array com quatro posições, que é atribuído a uma variável estado, e em seguida ela é impressa no terminal.

@@04
Slices em Go

Para entendermos como funciona os Slices, essas abstrações do array em Go, vamos criar um, para aprendermos fazendo.
A primeira grande vantagem dele, é que, como foi falado no vídeo anterior, seu tamanho não é fixo, é dinâmico, indeterminado. Então, como exemplo, vamos criar a função exibeNomes, para testar a utilização dessa estrutura de dados.

Para criar um slice, podemos utilizar a declaração curta de variáveis. Sua declaração é bem parecida com a de um array:

func exibeNomes() {
    nomes := []string
}COPIAR CÓDIGO
Além disso, já na sua declaração, podemos preencher os seus dados, passando-os dentro de chaves, separados por vírgula:

func exibeNomes() {
    nomes := []string{"Douglas", "Daniel", "Bernardo"}
}COPIAR CÓDIGO
O slice infere o seu tamanho de acordo com a sua quantidade de elementos. Podemos verificar isso imprimindo a quantidade de itens contidos nele, através da função len:

func exibeNomes() {
    nomes := []string{"Douglas", "Daniel", "Bernardo"}
    fmt.Println(len(nomes))
}COPIAR CÓDIGO
Ao chamar essa função, vemos que o slice possui 3 itens.

Diferenças entre Array e Slice
Além do tamanho, que no array é fixo, e no slice não, quais são as outras diferenças entre eles? Na verdade, o slice nada mais é do que um array, é uma abstração que funciona acima do array. Quando criamos um array, ele é criado na memória com as suas posições, e com o tipo especificado. Quando atribuímos valores aos seus índices, eles vão sendo preenchidos, e os índices que não utilizarmos vão continuar em branco, esperando serem preenchidos.

Já no caso do slice, ele cria um array de acordo com os elementos que passamos para ele. Por exemplo, na função que criamos, um array de 3 índices foi criado, e cada índice foi preenchido com um nome.

Mas podemos adicionar itens no slice, através da função append, que recebe o slice e o item a ser adicionado. Vamos adicionar o retorno dessa função no próprio slice e imprimir novamente a quantidade de itens contidos no slice:

func exibeNomes() {
    nomes := []string{"Douglas", "Daniel", "Bernardo"}
    fmt.Println("O meu slice tem", len(nomes), "itens")

    nomes = append(nomes, "Aparecida")
    fmt.Println("O meu slice tem", len(nomes), "itens")
}COPIAR CÓDIGO
Mas a função len informa a quantidade de itens contidos nele, e não a sua capacidade. Para descobrir a sua capacidade, devemos utilizar a função cap:

func exibeNomes() {
    nomes := []string{"Douglas", "Daniel", "Bernardo"}
    fmt.Println("O meu slice tem", len(nomes), "itens")
    fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

    nomes = append(nomes, "Aparecida")
    fmt.Println("O meu slice tem", len(nomes), "itens")
    fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
}COPIAR CÓDIGO
Ao executar a função, temos a seguinte saída:

O meu slice tem 3 itens
O meu slice tem capacidade para 3 itens
O meu slice tem 4 itens
O meu slice tem capacidade para 6 itensCOPIAR CÓDIGO
Ou seja, o slice dobrou de tamanho quando adicionamos um novo item! Então, sempre que estourarmos a capacidade máxima do slice, do array abaixo dele, ele dobra de tamanho.

Logo, o slice nada mais é do que o Go cuidando do array para nós, pois eles não funcionam de forma diferente. O slice é um array com algumas coisas abstraídas, evitando com que nos preocupemos com o tamanho e capacidade do array, focando somente em trabalhar com os dados.

@@05
Planning Poker
PRÓXIMA ATIVIDADE

Planning Poker é uma técnica de estimativa geralmente usada nas reuniões da metodologia Scrum. É chamada de poker porque utiliza um baralho com cartas numeradas.
Reinaldo é um Scrum Master que deseja implementar a idéia no Go. Começou por montar um slice com os pontos do Scrum.

pontosPlanningPoker := [] int {1, 2, 3, 5, 8, 13, 21}COPIAR CÓDIGO
Como Reinaldo pode saber quantas cartas ele vai precisar para fazer a técnica? Dica: verifique o tamanho do slice.

Alternativa correta
size(pontosPlanningPoker)
 
Alternativa correta
count(pontosPlanningPoker)
 
Alternativa correta
len(pontosPlanningPoker)
 
Isso aí! A função que usamos para descobrir o tamanho de uma slice é len().

@06
Mais cartas no baralho
PRÓXIMA ATIVIDADE

No exercício anterior, Reinaldo criou um slice que representava os pontos utilizados numa técnica de estimativa chamada Planning Poker.
pontosPlanningPoker := [] int {1, 2, 3, 5, 8, 13, 21}COPIAR CÓDIGO
Contudo, ele se esqueceu de colocar uma última pontuação, 40. Em vez de mudar o código de inicialização do slice, usou a função append.

pontosPlanningPoker = append(pontosPlanningPoker, 40)COPIAR CÓDIGO
O que será impresso no terminal quando ele executar o seguinte código:

fmt.Println(cap(pontosPlanningPoker))COPIAR CÓDIGO
Alternativa correta
Não é possível determinar porque o slice usa um número randômico para aumentar sua capacidade.
 
Alternativa correta
Vai imprimir o número 8.
 
Alternativa correta
Vai imprimir 14.
 
Isso aí! Quando é necessário colocar mais elementos do que sua capacidade atual, o slice dobra a capacidade.

@@07
A instrução for

Vamos apagar a função exibeNomes, que foi útil para aprendermos o funcionamento do slice, e começar a utilizá-lo na nossa aplicação, para melhorar o armazenamento dos sites a serem monitorados:
// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/", 
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}COPIAR CÓDIGO
Agora que temos uma coleção de sites, podemos monitorar mais de um site ao mesmo tempo. Para isso, devemos percorrer os elementos nosso slice, acessando e monitorando cada site.

Percorrendo os itens do slice
Como vimos na aula anterior, não existe no Go outra estrutura de repetição além do for, então vamos utilizá-lo para percorrer os itens do slice de sites. Para cada item do slice, nós vamos mandar uma requisição e testar o seu status code.

Vimos o for infinito, que faz com o que o código seja repetido para sempre, mas também podemos utilizar o for "tradicional", onde declaramos uma variável e vamos incrementado-a até o tamanho de itens do slice, por exemplo:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/", 
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i := 0; i < len(sites); i++ {
        fmt.Println(sites[i])
    }

    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}COPIAR CÓDIGO
Mas há uma forma mais enxuta de fazer isso em Go, utilizando o range. Ela é como se fosse um operador de iteração do Go, nos dando acesso a cada item do array, ou do slice, e ele nos retorna dos valores, a posição do item iterado e o próprio item daquela posição:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/", 
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i, site := range sites {
        fmt.Println("Estou passando na posição", i,
            "do meu slice e essa posição tem o site", site)
    }

    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}COPIAR CÓDIGO
Agora que conseguimos passar por cada item do slice, basta fazer com cada um deles o que já fazemos fora do for, ou seja, fazer uma requisição para o site e verificar o seu status code.

Faremos isso no próximo vídeo.

@@08
Percorrendo o slice
PRÓXIMA ATIVIDADE

Considere o código abaixo
package main

import "fmt"

func main() {
    pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21, 40}
    for i := 0; i < len(pontosPlanningPoker); i++ {
        fmt.Println(pontosPlanningPoker[i])
    }
}COPIAR CÓDIGO
Defina o conteúdo da próxima linha, de maneira que ela faça o programa compilar.

package main

import "fmt"

func main() {
    pontosPlanningPoker := []int{1, 2, 3, 5, 8, 13, 21, 40}
    for ________________________________ {
        fmt.Println("O ponto na posição", i, " tem o valor", ponto)
    }
}COPIAR CÓDIGO
Alternativa correta
i, ponto := range(pontosPlanningPoker)
 
Isso aí!
Alternativa correta
(ponto in pontosPlanningPoker)
 
Alternativa correta
i, ponto : pontosPlanningPoker

@@09
Testando múltiplas vezes

Queremos rodar o código que testa o site para cada um dos sites, então vamos exportá-lo ara uma função, a testaSite, que recebe o site por parâmetro:
// restante do código omitido

func testaSite(site string) {

    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}COPIAR CÓDIGO
Agora chamamos essa função dentro do for, para cada item do nosso slice. Vamos aproveitar e diminuir a mensagem de teste e adicionar uma linha em branco, para dar um espaçamento entre as mensagens e o menu:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/", 
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i, site := range sites {
        fmt.Println("Testando site", i, ":", site)
        testaSite(site)
    }

    fmt.Println("")
}COPIAR CÓDIGO
Ao executar o programa, todos os sites do slice são monitorados.

Mas os sites são testados uma única vez e para testarmos novamente, devemos digitar o comando 1. Então, vamos fazer com que os sites sejam testados 5 vezes a cada monitoramento. Para isso, criamos mais um loop:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/",
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i := 0; i < 5; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }        
    }

    fmt.Println("")
}COPIAR CÓDIGO
Aumentando o intervalo entre os monitoramentos
Mas o for executa muito rápido, então testaríamos cinco vezes cada site com um intervalo mínimo entre cada teste. Seria interessante termos um intervalo maior entre os testes, por exemplo de 5 em 5 minutos.

Para tal, a cada teste, podemos pedir para o Go esperar um pouco. Fazemos isso utilizando a função Sleep, do pacote time, passando para ela o quanto de tempo queremos esperar. Representamos o tempo através de constantes da própria biblioteca, como Second, Minute, entre outras:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/",
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i := 0; i < 5; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }
        time.Sleep(5 * time.Minute)
    }
    fmt.Println("")
}COPIAR CÓDIGO
No caso estamos testando de 5 em 5 minutos, mas esse tempo pode ser aumentado.

Criando constantes
Por último, vamos nos livrar de alguns números do nosso código, e exportá-los para constantes. Por que constantes? Pois elas não podem ser modificadas.

Os números que queremos atacar é o número 5 que está dentro do for, que representa o número de monitoramentos, e o número 5 dentro do Sleep, que representa o delay do nosso monitoramento.

Então, após os imports, criamos as constantes:

package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

const monitoramentos = 3
const delay = 5

// restante do código omitidoCOPIAR CÓDIGO
Agora, na função iniciarMonitoramento, utilizamos essas constantes:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/",
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i := 0; i < monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }

        time.Sleep(delay * time.Second)
    }

    fmt.Println("")
}COPIAR CÓDIGO
Agora, se quisermos alterar o delay ou a quantidade de monitoramentos, basta alterarmos diretamente na declaração das constantes. E para finalizar, vamos adicionar mais alguns espaçamentos, a cada monitoramento dos nossos sites e depois que o comando foi escolhido:

// restante do código omitido

func leComando() int {
    var comandoLido int
    fmt.Scan(&comandoLido)
    fmt.Println("O comando escolhido foi", comandoLido)
    fmt.Println("")

    return comandoLido
}

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/",
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i := 0; i < monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }
        time.Sleep(delay * time.Second)
        fmt.Println("")
    }
    fmt.Println("")
}COPIAR CÓDIGO
Com isso, conseguimos monitorar mais de uma vez múltiplos sites.

@@10
Cronômetro
PRÓXIMA ATIVIDADE

Fábio é um personal trainer e encomendou um programa para marcar o tempo de seus clientes.
O desenvolvedor está tendo dificuldades para fazer o seguinte programa compilar. Você pode ajudá-lo?

package main

import "fmt"

const aquecimento = 5
const corridaLeve = 10
const corridaForte = 10

func main() {
    fmt.Println("Período de alongamento...")
    time.Sleep(alongamento * time.Minute)

    fmt.Println("Período de aquecimento...")
    time.Sleep(aquecimento * time.Minute)

    fmt.Println("Período de corrida leve...")
    time.Sleep(corridaLeve * time.Minute)

    fmt.Println("Período de corrida forte...")
    time.Sleep(corridaForte * time.Minute)

    fmt.Println("Período de alongamento...")
    time.Sleep(alongamento * time.Minute)
}COPIAR CÓDIGO
Alternativa correta
Falta incluir a constante alongamento:
const alongamento = 1
 
Alternativa correta
Falta importar o pacote time...
import "time"
...e incluir a constante alongamento:
const alongamento = 1
 
Isso aí!
Alternativa correta
Falta importar o pacote time...
import "time"

@@11
Mãos na Massa: Testando múltiplos sites
PRÓXIMA ATIVIDADE

Começando deste ponto? Você pode fazer o DOWNLOAD completo do projeto do capítulo anterior e continuar seus estudos a partir deste capítulo.
Neste exercício vamos melhorar o monitoramento de nosso programa, para que ele monitore mais de um site ao mesmo tempo.

1- Nosso primeiro passo é substituir a nossa string site pelo slice sites, que conterá os endereços dos vários sites a serem testados. Vamos testar pelo menos 3 sites, o da Alura (https://www.alura.com.br) , o da Caelum (https://www.caelum.com.br) e o último do Random Status Code (https://random-status-code.herokuapp.com), para ter um falhe de vez em quando. Crie o slice com estes 3 sites:

//hello.go

//restante do arquivo

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/", 
        "https://www.alura.com.br", "https://www.caelum.com.br"}
    //restante da função
}COPIAR CÓDIGO
2- Como queremos que cada um destes sites seja testado uma vez, vamos um for para percorrer o slice inteiro, utilize o range para facilitar:

//hello.go

//restante do arquivo

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/", 
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i, site := range sites {
        // restante 
    }
}COPIAR CÓDIGO
3- Vamos agora extrair o nosso código que testa um site para uma função externa, para facilitar na hora que estamos iterando sobre o slice, aonde cada site deve executar uma vez a função. Crie a função testaSite que deve receber uma string com o site a ser testado e mova o código go http.Get para lá:

//hello.go 
//restante do arquivo

func testaSite(site string) {

    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}COPIAR CÓDIGO
4- Vamos agora chamar a testa site dentro do nosso for que está varrendo o slice:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/", 
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i, site := range sites {
        fmt.Println("Testando site", i, ":", site)
        testaSite(site)
    }

    fmt.Println("")
}COPIAR CÓDIGO
_ Adicione o fmt.Println("") vazio ao final, e o que está dentro do for também, para que o nosso script vá exibindo mensagens para o nosso usuário enquanto ele é executado._

Execute o nosso script e veja que agora conseguimos testar diversos sites, e ao final do teste o nosso Menu surge para escolhermos a opção de monitorar novamente.

https://s3.amazonaws.com/caelum-online-public/624-golang/04/projetos/alura-golang-stage-fim-cap04.zip

https://www.alura.com.br/?_gl%253D1%252A106udxe%252A_ga%252AMTUyNTg3MjAxMy4xNjk0Njc1MDk2%252A_ga_1EPWSW3PCS%252AMTY5NDY5MzY1MC4zLjEuMTY5NDY5NDI2Ny4wLjAuMA..%252A_fplc%252AUEhJQ3lwQ0Zpd0RqSjJOZDBXNDNmUkswTWY1UzdTbnhkNGFNeWpZUzduSE9jSTl4WG5jJTJGOTZtRyUyRnNBVkVuNzVjNE5HaEJSODRib1JRJTJGcGthSWtTVCUyRiUyRkZ1WWxiJTJGaldIYXFWdmdDY0dBRjBZSG5tVk5VSDJTNzNMM1Y0SHBnJTNEJTNE

https://www.caelum.com.br/

https://random-status-code.herokuapp.com/

@@12
Mãos na Massa: Delay e Aumentando o número de testes
PRÓXIMA ATIVIDADE

Agora vamos fazer com que a cada vez que o usuário selecione a opção de monitoramento, cada site seja testado mais de uma vez, de acordo com o que o usuário setar nas constantes.
1- Primeiro crie a constante monitoramentos que indicará quantas vezes o site será testado. Coloque o número que desejar, mas lembre-se que quanto mais você testar, mais o script vai demorar:

//hello.go
package main

import (
    "fmt"
    "net/http"
    "os"    
)

const monitoramentos = 3

// restante do arquivoCOPIAR CÓDIGO
2- Agora vamos colocar um for dentro da função iniciarMonitoramentos para que o slice seja percorrido o número de vezes que configuramos na constante:

//hello.go

//restante do arquivo

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/",
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    //Adição aqui
    for i := 0; i < monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }
    }

    fmt.Println("")
}COPIAR CÓDIGO
3- Para que haja uma pausa entre cada um dos testes que faremos, vamos adicionar um pequeno delay entre cada iteração de monitoramento. Utilize a função Sleep do pacote time para dar uma pausa entre monitoramentos. Para que o tamanho desta pausa fique configurável, vamos criar uma constante que vai dizer o tamanho do delay:

//hello.go
package main

import (
    "fmt"
    "net/http"
    "os"    
)

const monitoramentos = 3
const delay = 5

// restante do arquivoCOPIAR CÓDIGO
4- E a partir do valor da constante, vamos adicionar um time.Sleep na função iniciarMonitoramento. Vamos multiplicar a nossa contante (delay) pela constante que representa o número de segundos no Go(time.Seconds):

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := []string{"https://random-status-code.herokuapp.com/",
        "https://www.alura.com.br", "https://www.caelum.com.br"}

    for i := 0; i < monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }

        // adição AQUI!
        time.Sleep(delay * time.Second)
        fmt.Println("")
    }
    fmt.Println("")
}COPIAR CÓDIGO
_ Adicione o outro fmt.Println("") para que a exibição fique mais organizada para o seu usuário_

Agora o seu script deve estar testando os seus sites de acordo com o número de vezes que você setou na constante monitoramentos, e entre cada testada ele deve dar uma pausa de acordo com o número de segundos que você configurou na constante delay.

@@13
O que aprendemos?

O que aprendemos?
Coleções: Arrays e Slices
Trabalhar com Arrays e Slices
Diferenças entre Arrays e Slices
A estrutura de repetição for
Operador de iteração range
O range nos dá acesso a cada item da coleção, nos retornando a posição do item iterado e o próprio item daquela posição
Constantes
Trabalhando com o pacote time

####

@06-Leitura de Arquivos

@@01
Arquivos em Go

O que queremos fazer agora é não ficar dependente do slice, de só monitorar os sites contidos nele. Para isso, teremos um arquivo de texto com todos os sites que queremos monitorar. Adicionaremos vários sites em um arquivo de texto e ao iniciar nosso programa, ele monitorará todos os sites especificados dentro desse arquivo.
Para trabalhar com isso, devemos saber como trabalhamos com arquivo em Go, mais especificamente como lê-los. Para tal, criaremos a função leSitesDoArquivo, com essa responsabilidade. Além de ler os dados do arquivo, essa função os retornará em um slice:

// restante do código omitido

func leSitesDoArquivo() []string {

}COPIAR CÓDIGO
Agora, na função iniciarMonitoramento, ao invés de criar o slice na mão, nós criamos a partir do retorna dessa função:

// restante do código omitido

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    // criando o slice a partir da função leSitesDoArquivo()
    sites := leSitesDoArquivo()

    for i := 0; i < monitoramentos; i++ {
        for i, site := range sites {
            fmt.Println("Testando site", i, ":", site)
            testaSite(site)
        }
        time.Sleep(delay * time.Second)
        fmt.Println("")
    }
    fmt.Println("")
}COPIAR CÓDIGO
Pronto, agora falta implementarmos a função. Antes de ler os dados do arquivo, devemos saber como abri-lo.

Abrindo um arquivo em Go
Para abrir um arquivo em Go, precisamos pedir ao sistema operacional que abra o arquivo para nós. E quem é o representante do sistema operacional em Go? O pacote os, que já trabalhamos anteriormente. Nele, há a função Open, que abre o arquivo e nos retorna a sua representação em memória e um possível erro que possa ocorrer, que iremos ignorar:

// restante do código omitido

func leSitesDoArquivo() []string {
    arquivo, _ := os.Open("sites.txt")
}COPIAR CÓDIGO
Vamos fazer um teste agora, vamos imprimir o arquivo, criar um slice vazio e retorná-lo:

// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string
    arquivo, _ := os.Open("sites.txt")
    fmt.Println(arquivo)

    return sites
}COPIAR CÓDIGO
Agora, ao executar o programa e iniciar o monitoramento, vemos uma impressão <nil>. O que isso significa? Isso equivale ao null de outras linguagens de programação, ou seja, nosso arquivo está nulo. Isso acontece pois estamos abrindo um arquivo que não existe!

Mas como sabemos disso? Sabemos pois foi um erro proposital, já que estamos sendo descuidados, pois estamos ignorando os erros. Em nenhum momentos estamos tratando os erros da linguagem de programação, já que, ao tentar abrir um arquivo que não existe, era para o sistema operacional nos avisar isso, mas estamos ignorando esses avisos, tanto no momento de abrir um arquivo, quanto no momento de realizar as requisições para os sites.

Então, vamos ver como tratar esses erros no próximo vídeo.

@@02
Abrindo arquivo no Go
PRÓXIMA ATIVIDADE

Qual o código para abrir corretamente um arquivo no Go?
Alternativa correta
arquivo, error = os.Open("sites.txt")
 
Alternativa correta
arquivo, error = os.ReadFile("sites.txt")
 
Alternativa correta
arquivo, _ := os.Open("sites.txt")
 
Isso aí! Podemos utilizar a função os.Open quando queremos abrir um arquivo em Go.

@@03
Detectando erros

Para tratar os erros, primeiramente devemos substituir o operador de identificador em branco (_) pela variável err, padrão para os erros em Go:
// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string
    arquivo, err := os.Open("sites.txt")
    fmt.Println(arquivo)

    return sites
}COPIAR CÓDIGO
E agora, se err não for nulo, significa que houve um erro, então vamos imprimi-lo:

// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string
    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(arquivo)

    return sites
}COPIAR CÓDIGO
Agora, ao executar o programa e iniciar o monitoramento, vemos a seguinte impressão:

Ocorreu um erro: open sites.txt: no such file or directoryCOPIAR CÓDIGO
Agora está bem claro o que aconteceu, que o arquivo, ou o diretório, não foi encontrado. Agora, vamos tratar do outro erro que estamos ignorando, no momento em que fazemos a requisição para o site, na função testaSite:

// restante do código omitido

func testaSite(site string) {

    resp, err := http.Get(site)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}COPIAR CÓDIGO
Agora que já estamos detectando os erros do programa, vamos prosseguir com a leitura do arquivo.

@@04
Capturando erros
PRÓXIMA ATIVIDADE

Suponha que o arquivo "sites.txt" não existe. Quando o código abaixo é executado, a variável err conterá o erro de arquivo inexistente.
arquivo, err := os.Open("sites.txt")COPIAR CÓDIGO
O que o programador deve fazer para tratar esse erro?

Alternativa correta
Verificar se a variável err é diferente de nil, e mostrar a mensagem de erro no terminal.
if err != nil {
  fmt.Println("Ocorreu um erro:", err)
}
 
Isso aí!
Alternativa correta
Usar a estrutura 'try capture'.
try {
  arquivo, err := os.Open("sites.txt")
} capture {
  fmt.Println("Ocorreu um erro:", err)
}
 
Alternativa correta
Usar um try catch.
try {
  arquivo, err := os.Open("sites.txt")
} catch {
  fmt.Println("Ocorreu um erro:", err)
}

@@05
Lendo de um arquivo

Se o erro que está ocorrendo é que o sites.txt não existe, então vamos criá-lo dentro da pasta hello, com o conteúdo que estava no nosso slice:
https://random-status-code.herokuapp.com/
https://www.alura.com.br
https://www.caelum.com.brCOPIAR CÓDIGO
Agora, ao executar o programa e iniciar o monitoramento, vemos um código estranho sendo impresso, isso acontece pois o que a função Open nos retorna nada mais é do que um ponteiro para o arquivo em si.

Então, precisamos ler o arquivo de outro jeito, e há diversas maneiras para isso. Uma delas, um jeito fácil e rápido para ler um arquivo inteiro, sem ter que ler linha a linha, é utilizar o pacote io/ioutil.

Lendo os dados de um arquivo
Através do pacote io/ioutil, chamamos a função ReadFile, para ler o arquivo passado:

// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string
    arquivo, err := ioutil.ReadFile("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(arquivo)

    return sites
}COPIAR CÓDIGO
Essa função nos retorna o arquivo em um array de bytes, então basta convertê-los para uma string através da função string:

// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string
    arquivo, err := ioutil.ReadFile("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(string(arquivo))

    return sites
}COPIAR CÓDIGO
Ao executar o programa, iniciar o monitoramento, vemos o conteúdo do arquivo sendo impresso no console! Mas esse é um jeito bom para quem quer exibir o conteúdo todo do arquivo, o que não é muito útil para nós, já que queremos ler cada site de uma vez, para salvar cada um dentro do slice.

Lendo a primeira linha do arquivo
Para ler o arquivo linha a linha, vamos ver uma terceira forma de ler arquivos em Go, utilizando o pacote bufio. Para isso, vamos voltar à primeira leitura do arquivo, utilizando os.Open:

// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string
    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    return sites
}COPIAR CÓDIGO
E com o bufio, nós criamos um leitor através da função NewReader, que recebe o arquivo a ser lido:

// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string
    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    return sites
}COPIAR CÓDIGO
Com esse leitor, temos funções que nos ajudam a ler o arquivo, lendo linha a linha, como por exemplo a ReadString, que lê uma linha do arquivo e nos retorna em forma de string. Essa função deve receber um byte delimitador, para saber até onde queremos ler a linha.

No nosso caso, queremos ler a linha inteiro, ou seja, até a quebra de linha, que é representada pelo \n. Como queremos representar um byte, utilizaremos aspas simples:

// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string
    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    leitor.ReadString('\n')

    return sites
}COPIAR CÓDIGO
Essa função nos retorna a linha e um possível erro, que vamos detectar. Além disso, vamos imprimir a linha:

// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    linha, err := leitor.ReadString('\n')
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    fmt.Println(linha)

    return sites
}COPIAR CÓDIGO
Quando iniciamos o monitoramento, a primeira linha do arquivo é impressa. Mas queremos ler todas as linhas do arquivo, então como fazemos isso? Veremos no próximo vídeo.

@@06
Obtendo o conteúdo de um arquivo
PRÓXIMA ATIVIDADE

Observe o código abaixo, que abre o arquivo existente "sites.txt".
package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  arquivo, err := os.Open("sites.txt")
  if err != nil {
    fmt.Println("Ocorreu um erro:", err)
  }
  leitor := bufio.NewReader(arquivo)
}COPIAR CÓDIGO
O que você deve fazer para ler a primeira linha do arquivo?

Alternativa correta
linha, _ := leitor.ReadString('\n')
 
Isso aí!
Alternativa correta
linha, _ := bufio.ReadString(leitor, '\n')
 
Alternativa correta
linha, _ := os.read(leitor)
 
Alternativa correta
linha, _ := leitor.ReadString()

@@07
Lendo múltiplas linhas

Se queremos que o nosso programa leia além da primeira linha do arquivo, o código responsável por fazer a leitura deve ser executado mais de uma vez. Então, o que vamos fazer é colocar esse código dentro de um for, até dar um erro específico, o erro de EOF (End Of File), que acontece quando não há mais linha a serem lidas.
// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    for {
        linha, err := leitor.ReadString('\n')
        fmt.Println(linha)
        if err == io.EOF {
            fmt.Println("Ocorreu um erro:", err)
        }
    }

    return sites
}COPIAR CÓDIGO
Mas como saímos do for? Se houver o erro, que já estamos verificando, nós damos um break, que faz com que o código saia do loop:

// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    for {
        linha, err := leitor.ReadString('\n')
        fmt.Println(linha)
        if err == io.EOF {
            break
        }
    }

    return sites
}COPIAR CÓDIGO
Ao iniciar o monitoramento, vemos a seguinte impressão:

https://random-status-code.herokuapp.com/

https://www.alura.com.br

https://www.caelum.com.brCOPIAR CÓDIGO
A impressão está sendo desse jeito pois estamos pulando linha no arquivo, então ela é lida também. O leitor lê a linha, incluindo o \n, por isso que fica esse pulo de linha na hora de impressão.

Logo, devemos remover essa quebra de linha da linha lida, antes de adicioná-la ao slice. Para isso, existe a função TrimSpace, do pacotes strings, que remove as quebras de linha e espaços ao final de uma string:

// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    for {
        linha, err := leitor.ReadString('\n')
        linha = strings.TrimSpace(linha)
        fmt.Println(linha)
        if err == io.EOF {
            break
        }
    }

    return sites
}COPIAR CÓDIGO
Agora, as linhas são impressas sem quebra de linha. Logo, já podemos adicioná-las ao slice:

// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    for {
        linha, err := leitor.ReadString('\n')
        linha = strings.TrimSpace(linha)
        sites = append(sites, linha)
        if err == io.EOF {
            break
        }
    }

    return sites
}COPIAR CÓDIGO
Ao executar o programa novamente e iniciar o monitoramento, podemos perceber que os sites são monitorados corretamente! Agora, se quisermos adicionar mais sites, basta colocá-los no arquivo sites.txt.

Por último, devemos ser educados com o sistema operacional, se abrimos um arquivo com os.Open, após lê-lo, é uma boa prática fechá-lo com a função Close:

// restante do código omitido

func leSitesDoArquivo() []string {

    var sites []string

    arquivo, err := os.Open("sites.txt")
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)

    for {
        linha, err := leitor.ReadString('\n')
        linha = strings.TrimSpace(linha)
        sites = append(sites, linha)
        if err == io.EOF {
            break
        }
    }

    arquivo.Close()

    return sites
}

@@08
Mãos na Massa: Lendo sites do arquivo
PRÓXIMA ATIVIDADE

Começando deste ponto? Você pode fazer o DOWNLOAD completo do projeto do capítulo anterior e continuar seus estudos a partir deste capítulo.
Neste exercício vamos fazer com que os sites monitorados venham de um arquivo .txt para que fique fácil adicionar ou remover um site do monitoramento, além disto vamos colocar algumas detecções de erro em nossas funções.

1- Primeiro, vamos colocar nosso aprendizado sobre erros deste capítulo para capturar os erros, o primeiro deles de uma função que já utilizamos, a http.Get() da funçao testaSite:

// restante do código omitido

func testaSite(site string) {

    //Capturando o segundo parâmetro
    resp, err := http.Get(site)
    //Verificando se houve algum erro
    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
}COPIAR CÓDIGO
2- Agora vamos começar a ler os nossos sites de um arquivo .txt. Crie um arquivo de texto chamado sites.txt e coloque os sites que você quer monitorar lá, um em cada linha:

sites.txt

https://www.alura.com.br    
https://random-status-code.herokuapp.com 
https://www.caelum.com.br
https://www.casadocodigo.com.brCOPIAR CÓDIGO
3- Para ler este arquivo, vamos criar a função leSitesDoArquivo, que vai retornar para o nosso slice de sites preenchido de acordo com os sites do arquivo sites.txt. Crie a função leSitesDoArquivo que vai retornar um slice de strings:

 func leSitesDoArquivo() []string {

    var sites []string


    return sites
}COPIAR CÓDIGO
4- Agora vamos abrir o arquivo de sites.txt e detectar caso algum erro aconteça na abertura. Vamos utilizar o pacote os para abrir com a função `Open:

func leSitesDoArquivo() []string {

    var sites []string

    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }


    arquivo.Close()
    return sites
}COPIAR CÓDIGO
Não esqueça de fechar o arquivo no final com a função Close() do arquivo

5- Agora vamos criar um leitor com o pacote bufio, para que facilite o processo de percorrer o arquivo sites.txt:

func leSitesDoArquivo() []string {

    var sites []string

    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)


    arquivo.Close()
    return sites
}COPIAR CÓDIGO
6- Agora vamos utilizar um for e ler linha a linha até que o leitor encontre o EOF, o que nos dará um erro e utilizaremos a instrução break para sair do loop, indicando que chegamos ao final. Para ler cada linha utilizaremos a função ReadString do leitor, lendo até o caractere \n que indica o final da linha:

func leSitesDoArquivo() []string {

    var sites []string

    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)
    for {
        linha, err := leitor.ReadString('\n')


        if err == io.EOF {
            break
        }
    }

    arquivo.Close()
    return sites
}COPIAR CÓDIGO
7- Por último vamos dar um trim em cada linha lida para remover caracteres especiais, como \n , espaços e tabs. E claro, após isto vamos adicionar a linha ao slice de sites:

func leSitesDoArquivo() []string {

    var sites []string

    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)
    for {
        linha, err := leitor.ReadString('\n')
        linha = strings.TrimSpace(linha)

        sites = append(sites, linha)

        if err == io.EOF {
            break
        }

    }

    arquivo.Close()
    return sites
}COPIAR CÓDIGO
8- Agora vamos chamar a nossa função récem criada dentro função iniciarMonitoramento, para que ao invés de obtermos os sites de um slice fixo, obteremos do arquivo sites.txt:

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")

    sites := leSitesDoArquivo()

    // restante da função
}COPIAR CÓDIGO
Faça o teste, adicione novos sites ao arquivos sites.txt e verifique se os sites estão sendos lidos e testados corretamente.

https://s3.amazonaws.com/caelum-online-public/624-golang/05/projetos/alura-golang-stage-fim-cap05.zip

@@09
O que aprendemos?

O que aprendemos?
Detectar erros
Trabalhar com arquivos
Inclusive abrindo e fechando arquivos
Ler um arquivo por completo com o pacote io/ioutil
Ler linha a linha com o pacote bufio
Criação de leitores
Limpando strings com TrimSpace