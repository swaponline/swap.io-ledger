package AgentHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (ah *AgentHandler) GetFirstTxsCursor(address string) (
	*CursorTxs,
	error,
) {
	resp, err := http.Get(fmt.Sprintf(
		`http://%v/getFirstCursorTransactions?token=%v&address=%v`,
		ah.baseUrl,
		ah.apiKey,
		address,
	))
	if err != nil {
		return nil, err
	}

	var cursorTxs = new(CursorTxs)
	if err := json.NewDecoder(resp.Body).Decode(cursorTxs); err != nil {
		return nil, err
	}

	return cursorTxs, nil
}
func (ah *AgentHandler) GetTxsCursor(cursorId string) (
	*CursorTxs,
	error,
) {
	resp, err := http.Get(fmt.Sprintf(
		`http://%v/getCursorTransactions?token=%v&cursor=%v`,
		ah.baseUrl,
		ah.apiKey,
		cursorId,
	))
	if err != nil {
		return nil, err
	}

	var cursorTxs = new(CursorTxs)
	if err := json.NewDecoder(resp.Body).Decode(cursorTxs); err != nil {
		return nil, err
	}

	return cursorTxs, nil
}