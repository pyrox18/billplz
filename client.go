// Package billplz provides wrapper types and functions for interacting with the Billplz REST API.
package billplz

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

// Client represents the HTTP client that interacts with the Billplz API. The
// Client stores the API key used for authentication with the API.
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client

	APIKey string
}

// NewClient instantiates and returns a new Client.
// If a http.Client is not supplied, the function will use the default HTTP client.
// An API key must be provided for the client to authenticate with the API.
// The sandbox boolean value determines whether the client will communicate with
// the sandbox endpoint or the production endpoint.
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

// CreateCollection creates a new collection.
// An error will be returned if the supplied collection fails validation,
// or if the HTTP request fails.
func (c *Client) CreateCollection(collection Collection) (*Collection, error) {
	if collection.SplitPayment == nil {
		collection.SplitPayment = &SplitPayment{}
	}
	err := collection.validate()
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

// GetCollection retrieves a single collection with the given ID.
// An error will be returned if the collection is not found, or if
// the HTTP request fails.
func (c *Client) GetCollection(id string) (*Collection, error) {
	req, err := c.newRequest(http.MethodGet, "/collections/"+id, nil)
	if err != nil {
		return nil, err
	}

	var result Collection
	res, err := c.do(req, &result)
	if res.StatusCode == 404 {
		return nil, ErrCollectionNotFound
	}
	return &result, err
}

// GetCollectionIndex retrieves a set of collections. Up to 15 collections
// will be returned at a time.
// The page parameter determines the page of the collection set to retrieve,
// and defaults to 1.
// The status parameter determines whether to retrieve all collections, or
// only collections that are active or inactive. This parameter can take
// values of "", "active" or "inactive", and will default to "".
// An error will be returned if the HTTP request fails.
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

// CreateOpenCollection creates a new open collection.
// An error will be returned if the supplied open collection fails
// validation, or if the HTTP request fails.
func (c *Client) CreateOpenCollection(o OpenCollection) (*OpenCollection, error) {
	if o.SplitPayment == nil {
		o.SplitPayment = &SplitPayment{}
	}
	err := o.validate()
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

// GetOpenCollection retrieves a single open collection with the given ID.
// An error will be returned if the open collection is not found, or if
// the HTTP request fails.
func (c *Client) GetOpenCollection(id string) (*OpenCollection, error) {
	req, err := c.newRequest(http.MethodGet, "/open_collections/"+id, nil)
	if err != nil {
		return nil, err
	}

	var result OpenCollection
	res, err := c.do(req, &result)
	if res.StatusCode == 404 {
		return nil, ErrCollectionNotFound
	}
	return &result, err
}

// GetOpenCollectionIndex retrieves a set of open collections. Up to 15
// collections will be returned at a time.
// The page parameter determines the page of the open collection set to retrieve,
// and defaults to 1.
// The status parameter determines whether to retrieve all open collections, or
// only open collections that are active or inactive. This parameter can take
// values of "", "active" or "inactive", and will default to "".
// An error will be returned if the HTTP request fails.
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

// DeactivateCollection deactivates a collection with the given ID.
// An error will be returned if the collection could not be deactivated,
// or if the HTTP request fails.
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

// ActivateCollection activates a collection with the given ID.
// An error will be returned if the collection could not be activated, or
// if the HTTP request fails.
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

// CreateBill creates a new bill.
// An error will be returned if the supplied bill fails validation,
// or if the HTTP request fails.
func (c *Client) CreateBill(b Bill) (*Bill, error) {
	err := b.validate()
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

// GetBill retrieves a single bill with the given ID.
// An error will be returned if the bill is not found, or
// if the HTTP request fails.
func (c *Client) GetBill(id string) (*Bill, error) {
	req, err := c.newRequest(http.MethodGet, "/bills/"+id, nil)
	if err != nil {
		return nil, err
	}

	var result Bill
	res, err := c.do(req, &result)
	if res.StatusCode == 404 {
		return nil, ErrBillNotFound
	}
	return &result, err
}

// DeleteBill deletes a bill with the given ID.
// An error will be returned if the bill is not found, or
// if the HTTP request fails.
func (c *Client) DeleteBill(id string) error {
	req, err := c.newRequest(http.MethodDelete, "/bills/"+id, nil)
	if err != nil {
		return err
	}

	res, err := c.do(req, &struct{}{})
	if res.StatusCode == 404 {
		return ErrBillNotFound
	}
	return err
}

// CheckRegistration checks for a bank account's registration status based on
// the given account number. Returns true if the bank account is registered
// and verified, and false otherwise.
// An error will be returned if the bank account is not found or if the
// HTTP request fails.
func (c *Client) CheckRegistration(accountNumber string) (bool, error) {
	req, err := c.newRequest(http.MethodGet, "/check/bank_account_number/"+accountNumber, nil)
	if err != nil {
		return false, err
	}

	var result BankAccountCheckResponse
	_, err = c.do(req, &result)
	if err != nil {
		return false, err
	}

	switch result.Name {
	case "verified":
		return true, nil
	case "unverified":
		return false, nil
	}
	return false, ErrBankAccountNotFound
}

// GetBillTransactions retrieves a set of transactions related to a bill with the given ID.
// Up to 15 transactions will be returned at a time.
// The page parameter determines the page of the open collection set to retrieve,
// and defaults to 1.
// The status parameter determines whether to retrieve all transactions or
// only transactions that are pending, completed or failed. This parameter can
// take the values "", "pending", "completed", or "failed", and will default to "".
// An error will be returned if the HTTP request fails.
func (c *Client) GetBillTransactions(id string, page int, status string) (*BillTransactions, error) {
	if page <= 0 {
		page = 1
	}
	if status != "pending" && status != "completed" && status != "failed" {
		status = ""
	}

	req, err := c.newRequest(http.MethodGet, "/bills/"+id+"/transactions", nil)
	if err != nil {
		return nil, err
	}

	var q = req.URL.Query()
	q.Set("page", strconv.Itoa(page))
	if status != "" {
		q.Set("status", status)
	}
	req.URL.RawQuery = q.Encode()

	var result BillTransactions
	_, err = c.do(req, &result)
	return &result, err
}

// GetPaymentMethodIndex retrieves all available payment methods that can be
// enabled or disabled on a collection with the given ID.
// An error will be returned if the HTTP request fails.
func (c *Client) GetPaymentMethodIndex(id string) (*[]PaymentMethod, error) {
	req, err := c.newRequest(http.MethodGet, "/collections/"+id+"/payment_methods", nil)
	if err != nil {
		return nil, err
	}

	var result PaymentMethodList
	_, err = c.do(req, &result)
	return result.PaymentMethods, err
}

// UpdatePaymentMethods enables a set of payment methods on a collection with the
// given ID.
// The payment method codes are passed as a slice of strings.
// An error will be returned if the HTTP request fails.
func (c *Client) UpdatePaymentMethods(id string, codes []string) (*[]PaymentMethod, error) {
	methods := []PaymentMethod{}
	for _, element := range codes {
		methods = append(methods, PaymentMethod{
			Code: element,
		})
	}

	body := PaymentMethodList{
		PaymentMethods: &methods,
	}

	req, err := c.newRequest(http.MethodPut, "/collections/"+id+"/payment_methods", body)
	if err != nil {
		return nil, err
	}

	var result PaymentMethodList
	_, err = c.do(req, &result)
	return result.PaymentMethods, err
}

// GetBankAccountIndex gets a set of bank accounts with the given account numbers.
// This function requires the Billplz 'ADMIN' setting to be turned on, and will return
// an error if this condition is not met.
// An error will also be returned if the HTTP request fails.
func (c *Client) GetBankAccountIndex(accountNumbers []string) (*BankAccountList, error) {
	req, err := c.newRequest(http.MethodGet, "/bank_verification_services", nil)
	if err != nil {
		return nil, err
	}

	var q = req.URL.Query()
	for index, element := range accountNumbers {
		if index > 9 {
			break
		}
		q.Add("account_numbers[]", element)
	}
	req.URL.RawQuery = q.Encode()

	var result BankAccountList
	res, err := c.do(req, &result)
	if res.StatusCode == 422 || res.StatusCode == 401 {
		return nil, ErrAdminPrivilegeRequired
	}
	return &result, err
}

// GetBankAccount gets a bank account with the given account number.
// This function requires the Billplz 'ADMIN' setting to be turned on, and will return
// an error if this condition is not met.
// An error will also be returned if the HTTP request fails.
func (c *Client) GetBankAccount(accountNumber string) (*BankAccount, error) {
	req, err := c.newRequest(http.MethodGet, "/bank_verification_services/"+accountNumber, nil)
	if err != nil {
		return nil, err
	}

	var result BankAccount
	res, err := c.do(req, &result)
	if res.StatusCode == 422 {
		return nil, ErrAdminPrivilegeRequired
	}

	return &result, err
}

// CreateBankAccount creates a new bank account through the API's Bank Account
// Direct Verification Service.
// This function requires the Billplz 'ADMIN' setting to be turned on, and will return
// an error if this condition is not met.
// An error will also be returned if the supplied bank account fails validation,
// or if the HTTP request fails.
func (c *Client) CreateBankAccount(b BankAccount) (*BankAccount, error) {
	err := b.validate()
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest(http.MethodPost, "/bank_verification_services", b)
	if err != nil {
		return nil, err
	}

	var result BankAccount
	res, err := c.do(req, &result)
	if res.StatusCode == 422 || res.StatusCode == 401 {
		return nil, ErrAdminPrivilegeRequired
	}
	return &result, err
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
