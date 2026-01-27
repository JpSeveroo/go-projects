# 📊 Go Logs Analyzer

Um analisador de logs brutos desenvolvido em **Go** (Golang) focado em eficiência e simplicidade. A ferramenta extrai estatísticas, identifica falhas críticas e gera visualizações em ASCII diretamente no terminal.

Este projeto faz parte dos meus estudos em **Engenharia de Software** e foca em conceitos como manipulação de buffers, processamento de strings e estruturação modular de código.

## 🚀 Funcionalidades

- **Interface Interativa:** Banner de boas-vindas com efeito de digitação real.
- **Entrada Multi-linha:** Captura grandes blocos de dados via `os.Stdin` (Copy-Paste) até o sinal de EOF (`Ctrl+D` ou `Ctrl+Z`).
- **Análise de Dados:** Identifica e contabiliza ocorrências de `ERROR` e `INFO`.
- **Relatório Detalhado:** Exibição das linhas específicas onde erros foram detectados.
- **Gráfico ASCII:** Representação visual proporcional da saúde dos logs através de barras de progresso no terminal.

## 🛠️ Tecnologias e Conceitos

- **Linguagem:** Go 1.2x+
- **Módulos:** `bufio`, `strings`, `time`, `fmt`.
- **Estruturas:** Uso de `Structs` e `Slices` para gerenciamento de dados.

## 📋 Como Executar

### Pré-requisitos
- Ter o [Go](https://go.dev/dl/) instalado.

### Instalação e Execução
1. Clone o repositório:
   git clone https://github.com/Jpseveroo/go-analyzer-logs.git

2. Acesse o diretório e execute o programa:
   cd go-analyzer-logs
   go run .

### Como usar
1. Cole o conteúdo do seu log após o banner inicial.
2. Pressione **ENTER** para garantir uma linha vazia.
3. Pressione **Ctrl+D** (Linux/Mac) ou **Ctrl+Z** seguido de **ENTER** (Windows).
4. Digite `1` para gerar o gráfico visual ou `0` para encerrar.

## 🏗️ Estrutura do Código

- `main.go`: Maestro que orquestra o fluxo entre arquivos.
- `entrada.go`: Gerencia a interface e o buffer de entrada.
- `filtro.go`: Lógica de processamento e filtragem.
- `grafico.go`: Matemática de proporção e desenho ASCII.

---
*Desenvolvido por João Pedro - Estudante de Engenharia de Software.*