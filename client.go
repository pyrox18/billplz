package billplz

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client

	APIKey string
}

func NewClient(httpClient *http.Client, apiKey string, sandbox bool) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		httpClient: httpClient,
		APIKey:     apiKey,
	}

	var err error
	if sandbox {
		c.baseURL, err = url.Parse(endpointStaging)
	} else {
		c.baseURL, err = url.Parse(endpointProdV3)
	}

	return c, err
}

func (c *Client) CreateCollection(collection Collection) (*Collection, error) {
	if collection.SplitPayment == nil {
		collection.SplitPayment = &SplitPayment{}
	}
	err := collection.Validate()
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodPost, "/collections", collection)
	if err != nil {
		return nil, err
	}

	var result Collection
	_, err = c.do(req, &result)
	return &result, err
}

func (c *Client) GetCollection(id string) (*Collection, error) {
	req, err := c.newRequest(http.MethodGet, "/collections/"+id, nil)
	if err != nil {
		return nil, err
	}

	var result Collection
	_, err = c.do(req, &result)
	return &result, err
}

func (c *Client) GetCollectionIndex(page int, status string) (*CollectionIndexResult, error) {
	if page <= 0 {
		page = 1
	}
	if status != "active" && status != "inactive" {
		status = ""
	}

	req, err := c.newRequest(http.MethodGet, "/collections", nil)
	if err != nil {
		return nil, err
	}

	var q = req.URL.Query()
	q.Set("page", strconv.Itoa(page))
	if status != "" {
		q.Set("status", status)
	}
	req.URL.RawQuery = q.Encode()

	var result CollectionIndexResult
	_, err = c.do(req, &result)
	return &result, err
}

func (c *Client) CreateOpenCollection(o OpenCollection) (*OpenCollection, error) {
	if o.SplitPayment == nil {
		o.SplitPayment = &SplitPayment{}
	}
	err := o.Validate()
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodPost, "/open_collections", o)
	if err != nil {
		return nil, err
	}

	var result OpenCollection
	_, err = c.do(req, &result)
	return &result, err
}

func (c *Client) GetOpenCollection(id string) (*OpenCollection, error) {
	req, err := c.newRequest(http.MethodGet, "/open_collections/"+id, nil)
	if err != nil {
		return nil, err
	}

	var result OpenCollection
	_, err = c.do(req, &result)
	return &result, err
}

func (c *Client) GetOpenCollectionIndex(page int, status string) (*OpenCollectionIndexResult, error) {
	if page == 0 {
		page = 1
	}
	if status != "active" && status != "inactive" {
		status = ""
	}

	req, err := c.newRequest(http.MethodGet, "/open_collections", nil)
	if err != nil {
		return nil, err
	}

	var q = req.URL.Query()
	q.Set("page", strconv.Itoa(page))
	if status != "" {
		q.Set("status", status)
	}
	req.URL.RawQuery = q.Encode()

	var result OpenCollectionIndexResult
	_, err = c.do(req, &result)
	return &result, err
}

func (c *Client) DeactivateCollection(id string) error {
	req, err := c.newRequest(http.MethodPost, "/collections/"+id+"/deactivate", nil)
	if err != nil {
		return err
	}

	res, err := c.do(req, struct{}{})
	if res.StatusCode == http.StatusUnprocessableEntity {
		return ErrCannotDeactivateCollection
	}
	return nil
}

func (c *Client) ActivateCollection(id string) error {
	req, err := c.newRequest(http.MethodPost, "/collections/"+id+"/activate", nil)
	if err != nil {
		return err
	}

	res, err := c.do(req, struct{}{})
	if res.StatusCode == http.StatusUnprocessableEntity {
		return ErrCannotActivateCollection
	}
	return nil
}

func (c *Client) CreateBill(b Bill) (*Bill, error) {
	err := b.Validate()
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodPost, "/bills", b)
	if err != nil {
		return nil, err
	}

	var result Bill
	_, err = c.do(req, &result)
	return &result, err
}

func (c *Client) GetBill(id string) (Bill, error) {
	return Bill{}, nil
}

func (c *Client) DeleteBill(id string) error {
	return nil
}

func (c *Client) CheckRegistration(accountNumber string) (bool, error) {
	return true, nil
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

func (c *Client) GetBankAccountIndex(accountNumbers []string) ([]BankAccount, error) {
	return []BankAccount{}, nil
}

func (c *Client) GetBankAccount(accountNumber string) (BankAccount, error) {
	return BankAccount{}, nil
}

func (c *Client) CreateBankAccount(b *BankAccount) (BankAccount, error) {
	return BankAccount{}, nil
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	u := c.baseURL
	u.Path = u.Path + path

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
	req.SetBasicAuth(c.APIKey, "")
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
