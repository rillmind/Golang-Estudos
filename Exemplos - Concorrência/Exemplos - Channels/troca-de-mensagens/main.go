package main

import (
	"fmt"
	"sync"
	"time"
)

// 1. Define a mensagem de Requisição.
// Ela inclui um canal de "resposta" para o Coordenador enviar a Concessão.
type Requisicao struct {
	idCliente      int
	canalConcessao chan bool // Canal para onde o Coordenador enviará o "GRANT"
}

// 2. O Coordenador (Servidor)
// Ele gerencia o acesso em um loop infinito, serializando as requisições.
func coordenador(canalRequisicao <-chan Requisicao, canalLiberacao <-chan int) { // Recebe o pedido e a liberações
	fmt.Println("[Coord] Coordenador iniciado. Aguardando pedidos.")

	for {
		// Fase 1: Aguarda uma REQUISIÇÃO
		// (Bloqueia aqui até um cliente enviar um pedido)
		req := <-canalRequisicao
		fmt.Printf("[Coord] Requisição recebida do Cliente %d.\n", req.idCliente)

		// Fase 2: Envia a CONCESSÃO (GRANT)
		// (Envia o "OK" de volta ao cliente específico)
		fmt.Printf("[Coord] Concedendo acesso ao Cliente %d.\n", req.idCliente)
		req.canalConcessao <- true

		// Fase 3: Aguarda a LIBERAÇÃO (RELEASE)
		// (Bloqueia aqui até o cliente que teve acesso terminar)
		idClienteLiberou := <-canalLiberacao
		fmt.Printf("[Coord] Cliente %d liberou o recurso.\n", idClienteLiberou)
	}
}

// 3. O Cliente
func cliente(
	id int,
	canalRequisicao chan<- Requisicao,
	canalLiberacao chan<- int,
	wg *sync.WaitGroup) {

	defer wg.Done()

	// Cada cliente cria seu próprio canal privado para receber a concessão
	meuCanalConcessao := make(chan bool)

	// --- INÍCIO DA REGIÃO CRÍTICA (PROTOCOLO DE ENTRADA) ---

	// 1. Envia REQUISIÇÃO ao coordenador
	fmt.Printf("[Cliente %d] Solicitando acesso...\n", id)
	requisicao := Requisicao{
		idCliente:      id,
		canalConcessao: meuCanalConcessao,
	}
	canalRequisicao <- requisicao

	// 2. Aguarda CONCESSÃO do coordenador
	// (Bloqueia aqui até o coordenador responder)
	<-meuCanalConcessao
	fmt.Printf("[Cliente %d] >>> Concessão recebida. ENTROU na região crítica.\n", id)

	// --- REGIÃO CRÍTICA ---
	// (Simula o trabalho que precisa de exclusão mútua)
	fmt.Printf("[Cliente %d] ...Trabalhando...\n", id)
	time.Sleep(500 * time.Millisecond)
	// --- FIM DA REGIÃO CRÍTICA ---

	fmt.Printf("[Cliente %d] <<< Saindo da região crítica.\n", id)

	// 3. Envia LIBERAÇÃO para o coordenador
	canalLiberacao <- id
}

// 4. Main
func main() {
	// Canais centrais
	canalRequisicao := make(chan Requisicao)
	canalLiberacao := make(chan int)

	// WaitGroup para esperar todos os clientes terminarem
	var wg sync.WaitGroup

	// Inicia o Coordenador (apenas um)
	go coordenador(canalRequisicao, canalLiberacao)

	// Inicia 5 clientes
	numClientes := 5
	wg.Add(numClientes)
	for i := 1; i <= numClientes; i++ {
		go cliente(i, canalRequisicao, canalLiberacao, &wg)
		time.Sleep(time.Second * 1) // Apenas para espaçar os pedidos
	}

	// Espera todos os clientes finalizarem seus trabalhos
	wg.Wait()
	fmt.Println("Todos os clientes terminaram.")
}
