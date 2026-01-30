# ⚔️ Go RPG CLI

Um gerenciador de jogo de RPG simples via terminal (CLI) desenvolvido em **Go**. Este projeto foi criado para praticar conceitos fundamentais da linguagem, como Structs, Maps, manipulação de arquivos (I/O) e serialização JSON.

## 📋 Funcionalidades

O sistema permite gerenciar o estado de um jogo, incluindo:

* **👤 Gerenciamento de Jogador**: Criação de personagem com Nickname, Classe e Nível.
* **📜 Sistema de Missões**:
    * Cadastro de novas missões (ID, Nome, Dificuldade, Recompensa).
    * Listagem tabular de todas as missões cadastradas.
* **💾 Persistência de Dados**:
    * **Salvar**: Grava o estado atual (Jogador e Missões) em um arquivo `data.json`.
    * **Carregar**: Restaura o progresso salvo anteriormente.

## 🚀 Como Rodar

Certifique-se de ter o **Go** instalado na sua máquina (versão 1.25 ou superior recomendada).

1.  Clone este repositório (ou baixe os arquivos):
    ```bash
    git clone <https://github.com/JpSeveroo/go-projects/tree/main/go-rpg>
    cd go-rpg
    ```

2.  Execute o projeto:
    ```bash
    go run .
    ```

## 🎮 Como Usar

Ao iniciar o programa, você verá o seguinte menu:

```text
1 - Criar jogador   -> Define os atributos do seu herói
2 - Criar missão    -> Adiciona uma nova missão ao quadro
3 - Listar missões  -> Mostra todas as missões disponíveis
4 - Salvar          -> Salva todo o progresso no disco (data.json)
5 - Carregar        -> Lê o arquivo data.json e restaura os dados
0 - Sair            -> Encerra o programa
