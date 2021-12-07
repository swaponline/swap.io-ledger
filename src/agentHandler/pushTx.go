package AgentHandler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func (ah *AgentHandler) PushTx(hex []byte) ([]byte, error) {
	resp, err := http.Post(
		fmt.Sprintf(
			`http://%v/pushTx?token=%v`,
			ah.baseUrl,
			ah.apiKey,
		),
		"text/plain",
		bytes.NewReader(hex),
	)
	if err != nil {
		return nil, err
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}
