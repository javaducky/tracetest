/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// ApiApiController binds http requests to an api service and writes the service results to the http response
type ApiApiController struct {
	service      ApiApiServicer
	errorHandler ErrorHandler
}

// ApiApiOption for how the controller is set up.
type ApiApiOption func(*ApiApiController)

// WithApiApiErrorHandler inject ErrorHandler into controller
func WithApiApiErrorHandler(h ErrorHandler) ApiApiOption {
	return func(c *ApiApiController) {
		c.errorHandler = h
	}
}

// NewApiApiController creates a default api controller
func NewApiApiController(s ApiApiServicer, opts ...ApiApiOption) Router {
	controller := &ApiApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the ApiApiController
func (c *ApiApiController) Routes() Routes {
	return Routes{
		{
			"CreateDataStore",
			strings.ToUpper("Post"),
			"/api/datastores",
			c.CreateDataStore,
		},
		{
			"CreateEnvironment",
			strings.ToUpper("Post"),
			"/api/environments",
			c.CreateEnvironment,
		},
		{
			"CreateTest",
			strings.ToUpper("Post"),
			"/api/tests",
			c.CreateTest,
		},
		{
			"CreateTransaction",
			strings.ToUpper("Post"),
			"/api/transactions",
			c.CreateTransaction,
		},
		{
			"DeleteDataStore",
			strings.ToUpper("Delete"),
			"/api/datastores/{dataStoreId}",
			c.DeleteDataStore,
		},
		{
			"DeleteEnvironment",
			strings.ToUpper("Delete"),
			"/api/environments/{environmentId}",
			c.DeleteEnvironment,
		},
		{
			"DeleteTest",
			strings.ToUpper("Delete"),
			"/api/tests/{testId}",
			c.DeleteTest,
		},
		{
			"DeleteTestRun",
			strings.ToUpper("Delete"),
			"/api/tests/{testId}/run/{runId}",
			c.DeleteTestRun,
		},
		{
			"DeleteTransaction",
			strings.ToUpper("Delete"),
			"/api/transactions/{transactionId}",
			c.DeleteTransaction,
		},
		{
			"DeleteTransactionRun",
			strings.ToUpper("Delete"),
			"/api/transactions/{transactionId}/run/{runId}",
			c.DeleteTransactionRun,
		},
		{
			"DryRunAssertion",
			strings.ToUpper("Put"),
			"/api/tests/{testId}/run/{runId}/dry-run",
			c.DryRunAssertion,
		},
		{
			"ExecuteDefinition",
			strings.ToUpper("Post"),
			"/api/definition.yaml",
			c.ExecuteDefinition,
		},
		{
			"ExportTestRun",
			strings.ToUpper("Get"),
			"/api/tests/{testId}/run/{runId}/export",
			c.ExportTestRun,
		},
		{
			"ExpressionResolve",
			strings.ToUpper("Post"),
			"/api/expressions/resolve",
			c.ExpressionResolve,
		},
		{
			"GetDataStore",
			strings.ToUpper("Get"),
			"/api/datastores/{dataStoreId}",
			c.GetDataStore,
		},
		{
			"GetDataStoreDefinitionFile",
			strings.ToUpper("Get"),
			"/api/datastores/{dataStoreId}/definition.yaml",
			c.GetDataStoreDefinitionFile,
		},
		{
			"GetDataStores",
			strings.ToUpper("Get"),
			"/api/datastores",
			c.GetDataStores,
		},
		{
			"GetEnvironment",
			strings.ToUpper("Get"),
			"/api/environments/{environmentId}",
			c.GetEnvironment,
		},
		{
			"GetEnvironmentDefinitionFile",
			strings.ToUpper("Get"),
			"/api/environments/{environmentId}/definition.yaml",
			c.GetEnvironmentDefinitionFile,
		},
		{
			"GetEnvironments",
			strings.ToUpper("Get"),
			"/api/environments",
			c.GetEnvironments,
		},
		{
			"GetResources",
			strings.ToUpper("Get"),
			"/api/resources",
			c.GetResources,
		},
		{
			"GetRunResultJUnit",
			strings.ToUpper("Get"),
			"/api/tests/{testId}/run/{runId}/junit.xml",
			c.GetRunResultJUnit,
		},
		{
			"GetTest",
			strings.ToUpper("Get"),
			"/api/tests/{testId}",
			c.GetTest,
		},
		{
			"GetTestResultSelectedSpans",
			strings.ToUpper("Get"),
			"/api/tests/{testId}/run/{runId}/select",
			c.GetTestResultSelectedSpans,
		},
		{
			"GetTestRun",
			strings.ToUpper("Get"),
			"/api/tests/{testId}/run/{runId}",
			c.GetTestRun,
		},
		{
			"GetTestRuns",
			strings.ToUpper("Get"),
			"/api/tests/{testId}/run",
			c.GetTestRuns,
		},
		{
			"GetTestSpecs",
			strings.ToUpper("Get"),
			"/api/tests/{testId}/definition",
			c.GetTestSpecs,
		},
		{
			"GetTestVersion",
			strings.ToUpper("Get"),
			"/api/tests/{testId}/version/{version}",
			c.GetTestVersion,
		},
		{
			"GetTestVersionDefinitionFile",
			strings.ToUpper("Get"),
			"/api/tests/{testId}/version/{version}/definition.yaml",
			c.GetTestVersionDefinitionFile,
		},
		{
			"GetTests",
			strings.ToUpper("Get"),
			"/api/tests",
			c.GetTests,
		},
		{
			"GetTransaction",
			strings.ToUpper("Get"),
			"/api/transactions/{transactionId}",
			c.GetTransaction,
		},
		{
			"GetTransactionRun",
			strings.ToUpper("Get"),
			"/api/transactions/{transactionId}/run/{runId}",
			c.GetTransactionRun,
		},
		{
			"GetTransactionRuns",
			strings.ToUpper("Get"),
			"/api/transactions/{transactionId}/run",
			c.GetTransactionRuns,
		},
		{
			"GetTransactionVersion",
			strings.ToUpper("Get"),
			"/api/transactions/{transactionId}/version/{version}",
			c.GetTransactionVersion,
		},
		{
			"GetTransactionVersionDefinitionFile",
			strings.ToUpper("Get"),
			"/api/transactions/{transactionId}/version/{version}/definition.yaml",
			c.GetTransactionVersionDefinitionFile,
		},
		{
			"GetTransactions",
			strings.ToUpper("Get"),
			"/api/transactions",
			c.GetTransactions,
		},
		{
			"ImportTestRun",
			strings.ToUpper("Post"),
			"/api/tests/import",
			c.ImportTestRun,
		},
		{
			"RerunTestRun",
			strings.ToUpper("Post"),
			"/api/tests/{testId}/run/{runId}/rerun",
			c.RerunTestRun,
		},
		{
			"RunTest",
			strings.ToUpper("Post"),
			"/api/tests/{testId}/run",
			c.RunTest,
		},
		{
			"RunTransaction",
			strings.ToUpper("Post"),
			"/api/transactions/{transactionId}/run",
			c.RunTransaction,
		},
		{
			"TestConnection",
			strings.ToUpper("Post"),
			"/api/config/connection",
			c.TestConnection,
		},
		{
			"UpdateDataStore",
			strings.ToUpper("Put"),
			"/api/datastores/{dataStoreId}",
			c.UpdateDataStore,
		},
		{
			"UpdateEnvironment",
			strings.ToUpper("Put"),
			"/api/environments/{environmentId}",
			c.UpdateEnvironment,
		},
		{
			"UpdateTest",
			strings.ToUpper("Put"),
			"/api/tests/{testId}",
			c.UpdateTest,
		},
		{
			"UpdateTransaction",
			strings.ToUpper("Put"),
			"/api/transactions/{transactionId}",
			c.UpdateTransaction,
		},
		{
			"UpsertDefinition",
			strings.ToUpper("Put"),
			"/api/definition.yaml",
			c.UpsertDefinition,
		},
	}
}

// CreateDataStore - Create a new Data Store
func (c *ApiApiController) CreateDataStore(w http.ResponseWriter, r *http.Request) {
	dataStoreParam := DataStore{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&dataStoreParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDataStoreRequired(dataStoreParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateDataStore(r.Context(), dataStoreParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// CreateEnvironment - Create new environment
func (c *ApiApiController) CreateEnvironment(w http.ResponseWriter, r *http.Request) {
	environmentParam := Environment{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&environmentParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEnvironmentRequired(environmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateEnvironment(r.Context(), environmentParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// CreateTest - Create new test
func (c *ApiApiController) CreateTest(w http.ResponseWriter, r *http.Request) {
	testParam := Test{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&testParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTestRequired(testParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateTest(r.Context(), testParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// CreateTransaction - Create new transaction
func (c *ApiApiController) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	transactionParam := Transaction{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&transactionParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTransactionRequired(transactionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CreateTransaction(r.Context(), transactionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteDataStore - Delete a Data Store
func (c *ApiApiController) DeleteDataStore(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dataStoreIdParam := params["dataStoreId"]

	result, err := c.service.DeleteDataStore(r.Context(), dataStoreIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteEnvironment - delete a environment
func (c *ApiApiController) DeleteEnvironment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	environmentIdParam := params["environmentId"]

	result, err := c.service.DeleteEnvironment(r.Context(), environmentIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteTest - delete a test
func (c *ApiApiController) DeleteTest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	result, err := c.service.DeleteTest(r.Context(), testIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteTestRun - delete a test run
func (c *ApiApiController) DeleteTestRun(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	runIdParam := params["runId"]

	result, err := c.service.DeleteTestRun(r.Context(), testIdParam, runIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteTransaction - delete a transaction
func (c *ApiApiController) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionIdParam := params["transactionId"]

	result, err := c.service.DeleteTransaction(r.Context(), transactionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeleteTransactionRun - Delete a specific run from a particular transaction
func (c *ApiApiController) DeleteTransactionRun(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionIdParam := params["transactionId"]

	runIdParam, err := parseInt32Parameter(params["runId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.DeleteTransactionRun(r.Context(), transactionIdParam, runIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DryRunAssertion - run given assertions against the traces from the given run without persisting anything
func (c *ApiApiController) DryRunAssertion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	runIdParam := params["runId"]

	testSpecsParam := TestSpecs{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&testSpecsParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTestSpecsRequired(testSpecsParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.DryRunAssertion(r.Context(), testIdParam, runIdParam, testSpecsParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ExecuteDefinition - Execute a definition
func (c *ApiApiController) ExecuteDefinition(w http.ResponseWriter, r *http.Request) {
	textDefinitionParam := TextDefinition{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&textDefinitionParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTextDefinitionRequired(textDefinitionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ExecuteDefinition(r.Context(), textDefinitionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ExportTestRun - export test and test run information
func (c *ApiApiController) ExportTestRun(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	runIdParam := params["runId"]

	result, err := c.service.ExportTestRun(r.Context(), testIdParam, runIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ExpressionResolve - resolves an expression and returns the result string
func (c *ApiApiController) ExpressionResolve(w http.ResponseWriter, r *http.Request) {
	resolveRequestInfoParam := ResolveRequestInfo{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&resolveRequestInfoParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertResolveRequestInfoRequired(resolveRequestInfoParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ExpressionResolve(r.Context(), resolveRequestInfoParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetDataStore - Get a Data Store
func (c *ApiApiController) GetDataStore(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dataStoreIdParam := params["dataStoreId"]

	result, err := c.service.GetDataStore(r.Context(), dataStoreIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetDataStoreDefinitionFile - Get the data store definition as an YAML file
func (c *ApiApiController) GetDataStoreDefinitionFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dataStoreIdParam := params["dataStoreId"]

	result, err := c.service.GetDataStoreDefinitionFile(r.Context(), dataStoreIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetDataStores - Get all Data Stores
func (c *ApiApiController) GetDataStores(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	takeParam, err := parseInt32Parameter(query.Get("take"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skipParam, err := parseInt32Parameter(query.Get("skip"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	queryParam := query.Get("query")
	sortByParam := query.Get("sortBy")
	sortDirectionParam := query.Get("sortDirection")
	result, err := c.service.GetDataStores(r.Context(), takeParam, skipParam, queryParam, sortByParam, sortDirectionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetEnvironment - get environment
func (c *ApiApiController) GetEnvironment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	environmentIdParam := params["environmentId"]

	result, err := c.service.GetEnvironment(r.Context(), environmentIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetEnvironmentDefinitionFile - Get the environment definition as an YAML file
func (c *ApiApiController) GetEnvironmentDefinitionFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	environmentIdParam := params["environmentId"]

	result, err := c.service.GetEnvironmentDefinitionFile(r.Context(), environmentIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetEnvironments - Get Environments
func (c *ApiApiController) GetEnvironments(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	takeParam, err := parseInt32Parameter(query.Get("take"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skipParam, err := parseInt32Parameter(query.Get("skip"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	queryParam := query.Get("query")
	sortByParam := query.Get("sortBy")
	sortDirectionParam := query.Get("sortDirection")
	result, err := c.service.GetEnvironments(r.Context(), takeParam, skipParam, queryParam, sortByParam, sortDirectionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetResources - Get resources
func (c *ApiApiController) GetResources(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	takeParam, err := parseInt32Parameter(query.Get("take"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skipParam, err := parseInt32Parameter(query.Get("skip"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	queryParam := query.Get("query")
	sortByParam := query.Get("sortBy")
	sortDirectionParam := query.Get("sortDirection")
	result, err := c.service.GetResources(r.Context(), takeParam, skipParam, queryParam, sortByParam, sortDirectionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetRunResultJUnit - get test run results in JUnit xml format
func (c *ApiApiController) GetRunResultJUnit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	runIdParam := params["runId"]

	result, err := c.service.GetRunResultJUnit(r.Context(), testIdParam, runIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTest - get test
func (c *ApiApiController) GetTest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	result, err := c.service.GetTest(r.Context(), testIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTestResultSelectedSpans - retrieve spans that will be selected by selector
func (c *ApiApiController) GetTestResultSelectedSpans(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	testIdParam := params["testId"]

	runIdParam := params["runId"]

	queryParam := query.Get("query")
	result, err := c.service.GetTestResultSelectedSpans(r.Context(), testIdParam, runIdParam, queryParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTestRun - get test Run
func (c *ApiApiController) GetTestRun(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	runIdParam := params["runId"]

	result, err := c.service.GetTestRun(r.Context(), testIdParam, runIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTestRuns - get the runs for a test
func (c *ApiApiController) GetTestRuns(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	testIdParam := params["testId"]

	takeParam, err := parseInt32Parameter(query.Get("take"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skipParam, err := parseInt32Parameter(query.Get("skip"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetTestRuns(r.Context(), testIdParam, takeParam, skipParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTestSpecs - Get definition for a test
func (c *ApiApiController) GetTestSpecs(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	result, err := c.service.GetTestSpecs(r.Context(), testIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTestVersion - get a test specific version
func (c *ApiApiController) GetTestVersion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	versionParam, err := parseInt32Parameter(params["version"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetTestVersion(r.Context(), testIdParam, versionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTestVersionDefinitionFile - Get the test definition as an YAML file
func (c *ApiApiController) GetTestVersionDefinitionFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	versionParam, err := parseInt32Parameter(params["version"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetTestVersionDefinitionFile(r.Context(), testIdParam, versionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTests - Get tests
func (c *ApiApiController) GetTests(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	takeParam, err := parseInt32Parameter(query.Get("take"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skipParam, err := parseInt32Parameter(query.Get("skip"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	queryParam := query.Get("query")
	sortByParam := query.Get("sortBy")
	sortDirectionParam := query.Get("sortDirection")
	result, err := c.service.GetTests(r.Context(), takeParam, skipParam, queryParam, sortByParam, sortDirectionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTransaction - get transaction
func (c *ApiApiController) GetTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionIdParam := params["transactionId"]

	result, err := c.service.GetTransaction(r.Context(), transactionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTransactionRun - Get a specific run from a particular transaction
func (c *ApiApiController) GetTransactionRun(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionIdParam := params["transactionId"]

	runIdParam, err := parseInt32Parameter(params["runId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetTransactionRun(r.Context(), transactionIdParam, runIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTransactionRuns - Get all runs from a particular transaction
func (c *ApiApiController) GetTransactionRuns(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	query := r.URL.Query()
	transactionIdParam := params["transactionId"]

	takeParam, err := parseInt32Parameter(query.Get("take"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skipParam, err := parseInt32Parameter(query.Get("skip"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetTransactionRuns(r.Context(), transactionIdParam, takeParam, skipParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTransactionVersion - get a transaction specific version
func (c *ApiApiController) GetTransactionVersion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionIdParam := params["transactionId"]

	versionParam, err := parseInt32Parameter(params["version"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetTransactionVersion(r.Context(), transactionIdParam, versionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTransactionVersionDefinitionFile - Get the transaction definition as an YAML file
func (c *ApiApiController) GetTransactionVersionDefinitionFile(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionIdParam := params["transactionId"]

	versionParam, err := parseInt32Parameter(params["version"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetTransactionVersionDefinitionFile(r.Context(), transactionIdParam, versionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetTransactions - Get transactions
func (c *ApiApiController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	takeParam, err := parseInt32Parameter(query.Get("take"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	skipParam, err := parseInt32Parameter(query.Get("skip"), false)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	queryParam := query.Get("query")
	sortByParam := query.Get("sortBy")
	sortDirectionParam := query.Get("sortDirection")
	result, err := c.service.GetTransactions(r.Context(), takeParam, skipParam, queryParam, sortByParam, sortDirectionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// ImportTestRun - import test and test run information
func (c *ApiApiController) ImportTestRun(w http.ResponseWriter, r *http.Request) {
	exportedTestInformationParam := ExportedTestInformation{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&exportedTestInformationParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertExportedTestInformationRequired(exportedTestInformationParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.ImportTestRun(r.Context(), exportedTestInformationParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// RerunTestRun - rerun a test run
func (c *ApiApiController) RerunTestRun(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	runIdParam := params["runId"]

	result, err := c.service.RerunTestRun(r.Context(), testIdParam, runIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// RunTest - run test
func (c *ApiApiController) RunTest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	runInformationParam := RunInformation{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&runInformationParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertRunInformationRequired(runInformationParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.RunTest(r.Context(), testIdParam, runInformationParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// RunTransaction - run transaction
func (c *ApiApiController) RunTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionIdParam := params["transactionId"]

	runInformationParam := RunInformation{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&runInformationParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertRunInformationRequired(runInformationParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.RunTransaction(r.Context(), transactionIdParam, runInformationParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// TestConnection - Tests the config data store/exporter connection
func (c *ApiApiController) TestConnection(w http.ResponseWriter, r *http.Request) {
	dataStoreParam := DataStore{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&dataStoreParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDataStoreRequired(dataStoreParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.TestConnection(r.Context(), dataStoreParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateDataStore - Update a Data Store
func (c *ApiApiController) UpdateDataStore(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	dataStoreIdParam := params["dataStoreId"]

	dataStoreParam := DataStore{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&dataStoreParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertDataStoreRequired(dataStoreParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateDataStore(r.Context(), dataStoreIdParam, dataStoreParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateEnvironment - update environment
func (c *ApiApiController) UpdateEnvironment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	environmentIdParam := params["environmentId"]

	environmentParam := Environment{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&environmentParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertEnvironmentRequired(environmentParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateEnvironment(r.Context(), environmentIdParam, environmentParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateTest - update test
func (c *ApiApiController) UpdateTest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	testIdParam := params["testId"]

	testParam := Test{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&testParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTestRequired(testParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateTest(r.Context(), testIdParam, testParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdateTransaction - update transaction
func (c *ApiApiController) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transactionIdParam := params["transactionId"]

	transactionParam := Transaction{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&transactionParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTransactionRequired(transactionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateTransaction(r.Context(), transactionIdParam, transactionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpsertDefinition - Upsert a definition
func (c *ApiApiController) UpsertDefinition(w http.ResponseWriter, r *http.Request) {
	textDefinitionParam := TextDefinition{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&textDefinitionParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertTextDefinitionRequired(textDefinitionParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpsertDefinition(r.Context(), textDefinitionParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
