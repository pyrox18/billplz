package billplz

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client

	APIKey string
}

func (c *Client) CreateCollection(title string) (Collection, error) {
	return Collection{}, nil
}

func (c *Client) GetCollection(id string) (Collection, error) {
	return Collection{}, nil
}

func (c *Client) GetCollectionIndex(page int, status string) ([]Collection, error) {
	if page == 0 {
		page = 1
	}
	if status != "active" && status != "inactive" {
		status = ""
	}

	return []Collection{}, nil
}

func (c *Client) CreateOpenCollection(o *OpenCollection) (OpenCollection, error) {
	return OpenCollection{}, nil
}

func (c *Client) GetOpenCollection(id string) (OpenCollection, error) {
	return OpenCollection{}, nil
}

func (c *Client) GetOpenCollectionIndex(page int, status string) ([]OpenCollection, error) {
	if page == 0 {
		page = 1
	}
	if status != "active" && status != "inactive" {
		status = ""
	}

	return []OpenCollection{}, nil
}

func (c *Client) DeactivateCollection(id string) error {
	return nil
}

func (c *Client) ActivateCollection(id string) error {
	return nil
}

func (c *Client) CreateBill(b *Bill) (Bill, error) {
	return Bill{}, nil
}

func (c *Client) GetBill(id string) (Bill, error) {
	return Bill{}, nil
}

func (c *Client) DeleteBill(id string) error {
	return nil
}

func (c *Client) GetBillTransactions(id string) (BillTransactions, error) {
	return BillTransactions{}, nil
}

func (c *Client) GetPaymentMethodIndex(id string) ([]PaymentMethod, error) {
	return []PaymentMethod{}, nil
}

func (c *Client) UpdatePaymentMethods(m *[]PaymentMethod) ([]PaymentMethod, error) {
	return []PaymentMethod{}, nil
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
