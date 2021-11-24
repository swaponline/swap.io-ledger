package AgentHandler

import (
	"fmt"
	"net/http"
)

func (ah *AgentHandler) Subscribe(address string) error {
	_, err := http.Get(fmt.Sprintf(
		`http://%v/getCursorTransactions?token=%v&address=%v`,
		ah.baseUrl,
		ah.apiKey,
		address,
	))

	return err
}