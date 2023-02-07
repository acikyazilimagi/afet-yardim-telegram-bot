package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const envAPIKey = "ADDRESS_EXTRACT_API"

const (
	connectTimeout = 1 * time.Minute
	requestTimeout = 1 * time.Minute
)

type AddressExtractRequest struct {
	Data []string
}

type AddressExtractResponse struct {
	Data []string
}

type AddressDetail struct {
	Neighbourhood string `json:"neighbourhood"`
	Street        string `json:"street"`
	No            string `json:"no"`
	NameSurname   string `json:"name_surname"`
	Address       string `json:"address"`
	City          string `json:"city"`
	District      string `json:"district"`
	Tel           string `json:"tel"`
}

type BackendAddressFormat struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	City        string `json:"city"`
	District    string `json:"district"`
	FullAddress string `json:"full_address"`
	Tel         string `json:"tel"`
}

func sendExtractAddressResponse(ctx context.Context, text string) (*AddressDetail, error) {
	req := AddressExtractRequest{
		Data: []string{text},
	}

	addressExtractApiAddress := os.Getenv(envAPIKey)
	if addressExtractApiAddress == "" {
		err := fmt.Errorf("%s is not found in environment variables", envAPIKey)
		return nil, err
	}

	serialized, err := json.Marshal(req)
	if err != nil {
		err = fmt.Errorf("error occurred while serializing request: %w", err)
		return nil, err
	}

	tctx, cancel := context.WithTimeout(ctx, requestTimeout)
	defer cancel()

	request, err := http.NewRequestWithContext(tctx, http.MethodPost, addressExtractApiAddress, bytes.NewReader(serialized))
	if err != nil {
		err = fmt.Errorf("error occurred while creating request: %w", err)
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{
		Timeout: connectTimeout,
	}

	resp, err := client.Do(request)
	if err != nil {
		err = fmt.Errorf("error occurred while sending address request: %w", err)
		return nil, err
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("status returned not okay address request: %d", resp.StatusCode)
		_, _ = io.Copy(io.Discard, resp.Body) // Drain the body to let the Transport reuse the connection.
		return nil, err
	}

	var addressExtractResponse AddressExtractResponse

	readedRespBody, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("error occurred while reading response body: %w", err)
		return nil, err
	}

	if err := json.Unmarshal(readedRespBody, &addressExtractResponse); err != nil {
		err = fmt.Errorf("error occurred while deserializing response body: %w", err)
		return nil, err
	}

	serializedAddressExtractResp, err := json.Marshal(addressExtractResponse.Data)
	if err != nil {
		err = fmt.Errorf("error occurred while serializing extract response: %w", err)
		return nil, err
	}

	var addressDetail AddressDetail
	if err := json.Unmarshal(serializedAddressExtractResp, &addressDetail); err != nil {
		err = fmt.Errorf("error occurred while deserializing addressDetail: %w", err)
		return nil, err
	}

	return &addressDetail, nil
}

// func sendDataToBackend() {
// 	req, err := http.NewRequest(http.MethodPost, "", nil)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "could not prepare request data to backend: %s", err.Error())
// 		return
// 	}

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "could not send data to backend: %s", err.Error())
// 		return
// 	}

// 	// TODO: backend status control
// 	if resp.StatusCode != http.StatusOK {
// 		fmt.Fprintf(os.Stderr, "an error on backend response: %s", err.Error())
// 		return
// 	}
// }
