package servers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Allows to change this url
var DDNetHTTPMasterUrl = "https://master1.ddnet.org/ddnet/15/servers.json"

func GetAllServers(ctx context.Context) ([]Server, error) {
	return GetServers(ctx, DDNetHTTPMasterUrl)
}

func GetServers(ctx context.Context, url string) ([]Server, error) {
	client := &http.Client{}

	var intermediary struct {
		Servers []json.RawMessage `json:"servers"`
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request = request.WithContext(ctx)
	request.Header.Set("Accept", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		data, _ := io.ReadAll(resp.Body) // explicitly ignore error
		return nil, fmt.Errorf("error while fetching servers: %s: %s", resp.Status, string(data))
	}

	err = json.NewDecoder(resp.Body).Decode(&intermediary)
	if err != nil {
		return nil, err
	}

	result := make([]Server, 0, len(intermediary.Servers))

	for _, data := range intermediary.Servers {
		var server Server
		// skip weird servers in that list
		err = json.Unmarshal(data, &server)
		if err != nil {
			continue
		}

		result = append(result, server)
	}

	return result, nil
}
