// Package generated provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get account information.
	// (GET /v2/accounts/{address})
	AccountInformation(ctx echo.Context, address string, params AccountInformationParams) error
	// Get a list of unconfirmed transactions currently in the transaction pool by address.
	// (GET /v2/accounts/{address}/transactions/pending)
	GetPendingTransactionsByAddress(ctx echo.Context, address string, params GetPendingTransactionsByAddressParams) error
	// Get application information.
	// (GET /v2/applications/{application-id})
	GetApplicationByID(ctx echo.Context, applicationId uint64) error
	// Get asset information.
	// (GET /v2/assets/{asset-id})
	GetAssetByID(ctx echo.Context, assetId uint64) error
	// Get the block for the given round.
	// (GET /v2/blocks/{round})
	GetBlock(ctx echo.Context, round uint64, params GetBlockParams) error
	// Get a Merkle proof for a transaction in a block.
	// (GET /v2/blocks/{round}/transactions/{txid}/proof)
	GetProof(ctx echo.Context, round uint64, txid string, params GetProofParams) error
	// Get the current supply reported by the ledger.
	// (GET /v2/ledger/supply)
	GetSupply(ctx echo.Context) error
	// Gets the current node status.
	// (GET /v2/status)
	GetStatus(ctx echo.Context) error
	// Gets the node status after waiting for the given round.
	// (GET /v2/status/wait-for-block-after/{round})
	WaitForBlock(ctx echo.Context, round uint64) error
	// Compile TEAL source code to binary, produce its hash
	// (POST /v2/teal/compile)
	TealCompile(ctx echo.Context) error
	// Provide debugging information for a transaction (or group).
	// (POST /v2/teal/dryrun)
	TealDryrun(ctx echo.Context) error
	// Broadcasts a raw transaction to the network.
	// (POST /v2/transactions)
	RawTransaction(ctx echo.Context) error
	// Get parameters for constructing a new transaction
	// (GET /v2/transactions/params)
	TransactionParams(ctx echo.Context) error
	// Get a list of unconfirmed transactions currently in the transaction pool.
	// (GET /v2/transactions/pending)
	GetPendingTransactions(ctx echo.Context, params GetPendingTransactionsParams) error
	// Get a specific pending transaction.
	// (GET /v2/transactions/pending/{txid})
	PendingTransactionInformation(ctx echo.Context, txid string, params PendingTransactionInformationParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AccountInformation converts echo context to params.
func (w *ServerInterfaceWrapper) AccountInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params AccountInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AccountInformation(ctx, address, params)
	return err
}

// GetPendingTransactionsByAddress converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactionsByAddress(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsByAddressParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactionsByAddress(ctx, address, params)
	return err
}

// GetApplicationByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetApplicationByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "application-id" -------------
	var applicationId uint64

	err = runtime.BindStyledParameter("simple", false, "application-id", ctx.Param("application-id"), &applicationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter application-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetApplicationByID(ctx, applicationId)
	return err
}

// GetAssetByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetAssetByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "asset-id" -------------
	var assetId uint64

	err = runtime.BindStyledParameter("simple", false, "asset-id", ctx.Param("asset-id"), &assetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAssetByID(ctx, assetId)
	return err
}

// GetBlock converts echo context to params.
func (w *ServerInterfaceWrapper) GetBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBlockParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBlock(ctx, round, params)
	return err
}

// GetProof converts echo context to params.
func (w *ServerInterfaceWrapper) GetProof(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProofParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetProof(ctx, round, txid, params)
	return err
}

// GetSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetSupply(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSupply(ctx)
	return err
}

// GetStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetStatus(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStatus(ctx)
	return err
}

// WaitForBlock converts echo context to params.
func (w *ServerInterfaceWrapper) WaitForBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.WaitForBlock(ctx, round)
	return err
}

// TealCompile converts echo context to params.
func (w *ServerInterfaceWrapper) TealCompile(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TealCompile(ctx)
	return err
}

// TealDryrun converts echo context to params.
func (w *ServerInterfaceWrapper) TealDryrun(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TealDryrun(ctx)
	return err
}

// RawTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransaction(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransaction(ctx)
	return err
}

// TransactionParams converts echo context to params.
func (w *ServerInterfaceWrapper) TransactionParams(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TransactionParams(ctx)
	return err
}

// GetPendingTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactions(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactions(ctx, params)
	return err
}

// PendingTransactionInformation converts echo context to params.
func (w *ServerInterfaceWrapper) PendingTransactionInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PendingTransactionInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PendingTransactionInformation(ctx, txid, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/v2/accounts/:address", wrapper.AccountInformation, m...)
	router.GET("/v2/accounts/:address/transactions/pending", wrapper.GetPendingTransactionsByAddress, m...)
	router.GET("/v2/applications/:application-id", wrapper.GetApplicationByID, m...)
	router.GET("/v2/assets/:asset-id", wrapper.GetAssetByID, m...)
	router.GET("/v2/blocks/:round", wrapper.GetBlock, m...)
	router.GET("/v2/blocks/:round/transactions/:txid/proof", wrapper.GetProof, m...)
	router.GET("/v2/ledger/supply", wrapper.GetSupply, m...)
	router.GET("/v2/status", wrapper.GetStatus, m...)
	router.GET("/v2/status/wait-for-block-after/:round", wrapper.WaitForBlock, m...)
	router.POST("/v2/teal/compile", wrapper.TealCompile, m...)
	router.POST("/v2/teal/dryrun", wrapper.TealDryrun, m...)
	router.POST("/v2/transactions", wrapper.RawTransaction, m...)
	router.GET("/v2/transactions/params", wrapper.TransactionParams, m...)
	router.GET("/v2/transactions/pending", wrapper.GetPendingTransactions, m...)
	router.GET("/v2/transactions/pending/:txid", wrapper.PendingTransactionInformation, m...)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"S8kuMdP/++v/HwAA//96ww9J5PoAAA==",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
