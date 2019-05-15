package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/alibaba/pouch/apis/types"
	"github.com/stretchr/testify/assert"
)

func TestImageSearchServerError(t *testing.T) {
	client := &APIClient{
		HTTPCli: newMockClient(errorMockResponse(http.StatusInternalServerError, "Server error")),
	}
	term, registry := "", "nginx"
	_, err := client.ImageSearch(context.Background(), term, registry)
	if err == nil || !strings.Contains(err.Error(), "Server error") {
		t.Fatalf("expected a Server Error, got %v", err)
	}
}

func TestImageSearchOK(t *testing.T) {
	expectedURL := "/images/search"

	httpClient := newMockClient(func(req *http.Request) (*http.Response, error) {
		if !strings.HasPrefix(req.URL.Path, expectedURL) {
			return nil, fmt.Errorf("expected URL '%s', got '%s'", expectedURL, req.URL)
		}
		if req.Method != "POST" {
			return nil, fmt.Errorf("expected POST method, got %s", req.Method)
		}

		searchResults, err := json.Marshal([]types.SearchResultItem{
			{
				Description: "nginx info",
				IsAutomated: false,
				IsOfficial:  true,
				Name:        "nginx",
				StarCount:   1233,
			},
		})
		if err != nil {
			return nil, err
		}

		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader([]byte(searchResults))),
		}, nil
	})

	client := &APIClient{
		HTTPCli: httpClient,
	}

	searchResults, err := client.ImageSearch(context.Background(), "nginx", "")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, searchResults[0].StarCount, int64(1233))
	assert.Equal(t, searchResults[0].Name, "nginx")
}
