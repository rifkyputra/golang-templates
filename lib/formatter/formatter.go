package libs

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func BeautifyJSON(jsonStr string) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(jsonStr), "", "\t")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\n", string(prettyJSON.String())), nil
}
