package AgentHandler

import (
	"bytes"
	"fmt"
	"net/http"
)

func (ah *AgentHandler) Subscribe(address string) error {
	_, err := http.Post(fmt.Sprintf(
		`http://%v/subscribe?token=%v&address=%v`,
		ah.baseUrl,
		ah.apiKey,
		address,
	), "", bytes.NewReader([]byte{}))

	return err
}