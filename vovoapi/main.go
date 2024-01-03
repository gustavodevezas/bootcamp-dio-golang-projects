// Este programa implementa uma API simples para gerenciar informações sobre clientes da vovó.
// Foi desenvolvido como parte do Bootcamp DIO em GoLang, seguindo as convenções de nomenclatura em Go (camel case para funções e variáveis).

// Pacote principal do programa em Go.
package main

// Importação de Pacotes, Bibliotecas e Frameworks.
import (
	// Importação do pacote "encoding/json" para lidar com codificação e decodificação JSON.
	"encoding/json"

	// Importação do pacote "log" para registrar mensagens de log.
	"log"

	// Importação do pacote "net/http" para criar um servidor HTTP.
	"net/http"

	// Importação do pacote "strconv" que fornece funções para conversão de valores entre tipos numéricos e representações de string.
	"strconv"

	// Importação do pacote "github.com/gorilla/mux".
	// Pacote externo utilizado para facilitar o roteamento HTTP.
	"github.com/gorilla/mux"
)

// Slice para armazenar os clientes da vovó.
var clientesDaVovo []ClienteDaVovo

// Definir a estrutura de dados dos clientes da vovó.
// Pode incluir campos adicionais, como email, CEP, cidade, etc.
type ClienteDaVovo struct {
	ID        string `json:"id,omitempty"`
	Nome      string `json:"nome,omitempty"`
	NomeUnico string `json:"nomeUnico,omitempty"`
	Idade     int    `json:"idade,omitempty"`
	Telefone  string `json:"telefone,omitempty"`
}

// Função auxiliar para verificar se o nome é único.
func isNomeUnico(nome string) bool {
	for _, cliente := range clientesDaVovo {
		if cliente.Nome == nome {
			return false
		}
	}
	return true
}

// Função auxiliar para verificar se o nome é único ao atualizar.
func isNomeUnicoAoAtualizar(novoNome, clienteAtualizadoID string) bool {
	for _, cliente := range clientesDaVovo {
		// Verificar se o nome já está em uso por outro cliente (diferente do cliente sendo atualizado).
		if cliente.Nome == novoNome && cliente.ID != clienteAtualizadoID {
			return false
		}
	}
	return true
}

// ObterClientesDaVovo obtém todos os clientes da vovó.
func ObterClientesDaVovo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Iterar sobre todos os clientes e escrever no ResponseWriter.
	for _, cliente := range clientesDaVovo {
		if err := json.NewEncoder(w).Encode(cliente); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("\n")) // Adicionar uma quebra de linha entre os clientes.
	}
}

// ObterClienteDaVovo Obtém um cliente da vovó por ID.
func ObterClienteDaVovo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Iterar sobre os clientes e verificar se há um com o ID fornecido.
	for _, cliente := range clientesDaVovo {
		if cliente.ID == params["id"] {
			json.NewEncoder(w).Encode(cliente)
			return
		}
	}

	// Se o cliente não for localizado, retornar um erro com uma mensagem instrutiva.
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Cliente não encontrado. Por favor, forneça um ID válido."})
}

// CriarClienteDaVovo cria um novo cliente da vovó atribuíndo os parâmetros.
func CriarClienteDaVovo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var novoClienteDaVovo ClienteDaVovo

	// Decodificar o corpo da requisição JSON para obter o novo cliente.
	err := json.NewDecoder(r.Body).Decode(&novoClienteDaVovo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Verificar se o nome é único, para evitar duplicidade.
	if !isNomeUnico(novoClienteDaVovo.Nome) {
		http.Error(w, "O nome não é único.", http.StatusBadRequest)
		return
	}

	// Encontrar o próximo ID disponível.
	novoClienteDaVovo.ID = strconv.Itoa(len(clientesDaVovo) + 1)

	// Adicionar o novo cliente à lista.
	clientesDaVovo = append(clientesDaVovo, novoClienteDaVovo)

	// Retornar o novo cliente com o ID atribuído.
	json.NewEncoder(w).Encode(novoClienteDaVovo)
}

// AtualizarClienteDaVovo atualiza os detalhes de um cliente da vovó existente.
func AtualizarClienteDaVovo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Iterar sobre os clientes e encontrar o cliente com o ID fornecido.
	for index, cliente := range clientesDaVovo {
		if cliente.ID == params["id"] {
			// Salvar o nome atual do cliente antes da atualização.
			nomeAntigo := cliente.Nome

			// Criar uma cópia do cliente antes da atualização.
			clienteAntigo := clientesDaVovo[index]

			// Atualizar os campos do cliente existente com os novos dados.
			err := json.NewDecoder(r.Body).Decode(&clientesDaVovo[index])
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// Manter o ID inalterado.
			clientesDaVovo[index].ID = params["id"]

			// Verificar se o nome foi alterado antes de verificar se é único.
			if clientesDaVovo[index].Nome != nomeAntigo {
				// Verificar se o novo nome é único.
				novoNome := clientesDaVovo[index].Nome
				if !isNomeUnicoAoAtualizar(novoNome, params["id"]) {
					// Reverter para os dados antigos, pois o novo nome não é único.
					clientesDaVovo[index] = clienteAntigo
					http.Error(w, "O novo nome não é único. Por favor, escolha outro e tente novamente.", http.StatusBadRequest)
					return
				}
			}

			// Retornar o cliente atualizado.
			json.NewEncoder(w).Encode(clientesDaVovo[index])
			return
		}
	}

	// Se o cliente não for localizado, retornar um erro com uma mensagem instrutiva.
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "Cliente não encontrado. Por favor, forneça um ID válido."})
}

// DeletarClienteDaVovo exclui um cliente da vovó com base no ID fornecido.
func DeletarClienteDaVovo(w http.ResponseWriter, r *http.Request) {
	// Configurar o cabeçalho da resposta para indicar o tipo de conteúdo JSON.
	w.Header().Set("Content-Type", "application/json")
	// Obter os parâmetros da solicitação, incluindo o ID do cliente a ser deletado.
	params := mux.Vars(r)

	// Encontrar a posição do cliente a ser removido na lista.
	var indexToRemove int = -1
	for index, cliente := range clientesDaVovo {
		if cliente.ID == params["id"] {
			indexToRemove = index
			break
		}
	}

	// Se o cliente não for localizado, retornar um erro com uma mensagem instrutiva.
	if indexToRemove == -1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Cliente não encontrado. Por favor, forneça um ID válido."})
		return
	}

	// Remover o cliente da lista.
	clientesDaVovo = append(clientesDaVovo[:indexToRemove], clientesDaVovo[indexToRemove+1:]...)

	// Reatribuir IDs sequenciais, mantendo a ordem.
	for i, cliente := range clientesDaVovo {
		cliente.ID = strconv.Itoa(i + 1)
		clientesDaVovo[i] = cliente
	}

	// Retornar a lista atualizada após a remoção do cliente.
	json.NewEncoder(w).Encode(clientesDaVovo)
}

// A função main é o ponto de entrada para o programa em Go.
func main() {
	// Inicializar o roteador gorilla.
	router := mux.NewRouter()

	// Adicionar alguns clientes para teste.
	// Estes são exemplos usados apenas para testar a API.
	clientesDaVovo = append(clientesDaVovo, ClienteDaVovo{ID: "1", Nome: "Cauã Boaventura", Idade: 18, Telefone: "11986533554"})
	clientesDaVovo = append(clientesDaVovo, ClienteDaVovo{ID: "2", Nome: "Eric Shinoda", Idade: 27, Telefone: "12986868766"})
	clientesDaVovo = append(clientesDaVovo, ClienteDaVovo{ID: "3", Nome: "Valdomiro Amaral", Idade: 52, Telefone: "16972152374"})

	// Definir as rotas da API.
	// Testei utilizando PowerShell no  Windows executando os comandos Invoke-RestMethod
	// Para cada operação do CRUD
	router.HandleFunc("/clientesdavovo", ObterClientesDaVovo).Methods("GET")
	router.HandleFunc("/clientesdavovo/{id}", ObterClienteDaVovo).Methods("GET")
	router.HandleFunc("/clientesdavovo", CriarClienteDaVovo).Methods("POST")
	router.HandleFunc("/clientesdavovo/{id}", AtualizarClienteDaVovo).Methods("PUT")
	router.HandleFunc("/clientesdavovo/{id}", DeletarClienteDaVovo).Methods("DELETE")

	// Inicializar o servidor na porta 8000.
	// Acesso pelo navegador http://localhost:8000/clientesdavovo
	log.Fatal(http.ListenAndServe(":8000", router))
}
