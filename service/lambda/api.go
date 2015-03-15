package lambda

// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

import (
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
)

// AddEventSourceRequest generates a request for the AddEventSource operation.
func (c *Lambda) AddEventSourceRequest(input *AddEventSourceInput) (req *aws.Request, output *EventSourceConfiguration) {
	if opAddEventSource == nil {
		opAddEventSource = &aws.Operation{
			Name:       "AddEventSource",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-11-13/event-source-mappings/",
		}
	}

	req = aws.NewRequest(c.Service, opAddEventSource, input, output)
	output = &EventSourceConfiguration{}
	req.Data = output
	return
}

func (c *Lambda) AddEventSource(input *AddEventSourceInput) (output *EventSourceConfiguration, err error) {
	req, out := c.AddEventSourceRequest(input)
	output = out
	err = req.Send()
	return
}

var opAddEventSource *aws.Operation

// DeleteFunctionRequest generates a request for the DeleteFunction operation.
func (c *Lambda) DeleteFunctionRequest(input *DeleteFunctionInput) (req *aws.Request, output *DeleteFunctionOutput) {
	if opDeleteFunction == nil {
		opDeleteFunction = &aws.Operation{
			Name:       "DeleteFunction",
			HTTPMethod: "DELETE",
			HTTPPath:   "/2014-11-13/functions/{FunctionName}",
		}
	}

	req = aws.NewRequest(c.Service, opDeleteFunction, input, output)
	output = &DeleteFunctionOutput{}
	req.Data = output
	return
}

func (c *Lambda) DeleteFunction(input *DeleteFunctionInput) (output *DeleteFunctionOutput, err error) {
	req, out := c.DeleteFunctionRequest(input)
	output = out
	err = req.Send()
	return
}

var opDeleteFunction *aws.Operation

// GetEventSourceRequest generates a request for the GetEventSource operation.
func (c *Lambda) GetEventSourceRequest(input *GetEventSourceInput) (req *aws.Request, output *EventSourceConfiguration) {
	if opGetEventSource == nil {
		opGetEventSource = &aws.Operation{
			Name:       "GetEventSource",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-11-13/event-source-mappings/{UUID}",
		}
	}

	req = aws.NewRequest(c.Service, opGetEventSource, input, output)
	output = &EventSourceConfiguration{}
	req.Data = output
	return
}

func (c *Lambda) GetEventSource(input *GetEventSourceInput) (output *EventSourceConfiguration, err error) {
	req, out := c.GetEventSourceRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetEventSource *aws.Operation

// GetFunctionRequest generates a request for the GetFunction operation.
func (c *Lambda) GetFunctionRequest(input *GetFunctionInput) (req *aws.Request, output *GetFunctionOutput) {
	if opGetFunction == nil {
		opGetFunction = &aws.Operation{
			Name:       "GetFunction",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-11-13/functions/{FunctionName}",
		}
	}

	req = aws.NewRequest(c.Service, opGetFunction, input, output)
	output = &GetFunctionOutput{}
	req.Data = output
	return
}

func (c *Lambda) GetFunction(input *GetFunctionInput) (output *GetFunctionOutput, err error) {
	req, out := c.GetFunctionRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetFunction *aws.Operation

// GetFunctionConfigurationRequest generates a request for the GetFunctionConfiguration operation.
func (c *Lambda) GetFunctionConfigurationRequest(input *GetFunctionConfigurationInput) (req *aws.Request, output *FunctionConfiguration) {
	if opGetFunctionConfiguration == nil {
		opGetFunctionConfiguration = &aws.Operation{
			Name:       "GetFunctionConfiguration",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-11-13/functions/{FunctionName}/configuration",
		}
	}

	req = aws.NewRequest(c.Service, opGetFunctionConfiguration, input, output)
	output = &FunctionConfiguration{}
	req.Data = output
	return
}

func (c *Lambda) GetFunctionConfiguration(input *GetFunctionConfigurationInput) (output *FunctionConfiguration, err error) {
	req, out := c.GetFunctionConfigurationRequest(input)
	output = out
	err = req.Send()
	return
}

var opGetFunctionConfiguration *aws.Operation

// InvokeAsyncRequest generates a request for the InvokeAsync operation.
func (c *Lambda) InvokeAsyncRequest(input *InvokeAsyncInput) (req *aws.Request, output *InvokeAsyncOutput) {
	if opInvokeAsync == nil {
		opInvokeAsync = &aws.Operation{
			Name:       "InvokeAsync",
			HTTPMethod: "POST",
			HTTPPath:   "/2014-11-13/functions/{FunctionName}/invoke-async/",
		}
	}

	req = aws.NewRequest(c.Service, opInvokeAsync, input, output)
	output = &InvokeAsyncOutput{}
	req.Data = output
	return
}

func (c *Lambda) InvokeAsync(input *InvokeAsyncInput) (output *InvokeAsyncOutput, err error) {
	req, out := c.InvokeAsyncRequest(input)
	output = out
	err = req.Send()
	return
}

var opInvokeAsync *aws.Operation

// ListEventSourcesRequest generates a request for the ListEventSources operation.
func (c *Lambda) ListEventSourcesRequest(input *ListEventSourcesInput) (req *aws.Request, output *ListEventSourcesOutput) {
	if opListEventSources == nil {
		opListEventSources = &aws.Operation{
			Name:       "ListEventSources",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-11-13/event-source-mappings/",
		}
	}

	req = aws.NewRequest(c.Service, opListEventSources, input, output)
	output = &ListEventSourcesOutput{}
	req.Data = output
	return
}

func (c *Lambda) ListEventSources(input *ListEventSourcesInput) (output *ListEventSourcesOutput, err error) {
	req, out := c.ListEventSourcesRequest(input)
	output = out
	err = req.Send()
	return
}

var opListEventSources *aws.Operation

// ListFunctionsRequest generates a request for the ListFunctions operation.
func (c *Lambda) ListFunctionsRequest(input *ListFunctionsInput) (req *aws.Request, output *ListFunctionsOutput) {
	if opListFunctions == nil {
		opListFunctions = &aws.Operation{
			Name:       "ListFunctions",
			HTTPMethod: "GET",
			HTTPPath:   "/2014-11-13/functions/",
		}
	}

	req = aws.NewRequest(c.Service, opListFunctions, input, output)
	output = &ListFunctionsOutput{}
	req.Data = output
	return
}

func (c *Lambda) ListFunctions(input *ListFunctionsInput) (output *ListFunctionsOutput, err error) {
	req, out := c.ListFunctionsRequest(input)
	output = out
	err = req.Send()
	return
}

var opListFunctions *aws.Operation

// RemoveEventSourceRequest generates a request for the RemoveEventSource operation.
func (c *Lambda) RemoveEventSourceRequest(input *RemoveEventSourceInput) (req *aws.Request, output *RemoveEventSourceOutput) {
	if opRemoveEventSource == nil {
		opRemoveEventSource = &aws.Operation{
			Name:       "RemoveEventSource",
			HTTPMethod: "DELETE",
			HTTPPath:   "/2014-11-13/event-source-mappings/{UUID}",
		}
	}

	req = aws.NewRequest(c.Service, opRemoveEventSource, input, output)
	output = &RemoveEventSourceOutput{}
	req.Data = output
	return
}

func (c *Lambda) RemoveEventSource(input *RemoveEventSourceInput) (output *RemoveEventSourceOutput, err error) {
	req, out := c.RemoveEventSourceRequest(input)
	output = out
	err = req.Send()
	return
}

var opRemoveEventSource *aws.Operation

// UpdateFunctionConfigurationRequest generates a request for the UpdateFunctionConfiguration operation.
func (c *Lambda) UpdateFunctionConfigurationRequest(input *UpdateFunctionConfigurationInput) (req *aws.Request, output *FunctionConfiguration) {
	if opUpdateFunctionConfiguration == nil {
		opUpdateFunctionConfiguration = &aws.Operation{
			Name:       "UpdateFunctionConfiguration",
			HTTPMethod: "PUT",
			HTTPPath:   "/2014-11-13/functions/{FunctionName}/configuration",
		}
	}

	req = aws.NewRequest(c.Service, opUpdateFunctionConfiguration, input, output)
	output = &FunctionConfiguration{}
	req.Data = output
	return
}

func (c *Lambda) UpdateFunctionConfiguration(input *UpdateFunctionConfigurationInput) (output *FunctionConfiguration, err error) {
	req, out := c.UpdateFunctionConfigurationRequest(input)
	output = out
	err = req.Send()
	return
}

var opUpdateFunctionConfiguration *aws.Operation

// UploadFunctionRequest generates a request for the UploadFunction operation.
func (c *Lambda) UploadFunctionRequest(input *UploadFunctionInput) (req *aws.Request, output *FunctionConfiguration) {
	if opUploadFunction == nil {
		opUploadFunction = &aws.Operation{
			Name:       "UploadFunction",
			HTTPMethod: "PUT",
			HTTPPath:   "/2014-11-13/functions/{FunctionName}",
		}
	}

	req = aws.NewRequest(c.Service, opUploadFunction, input, output)
	output = &FunctionConfiguration{}
	req.Data = output
	return
}

func (c *Lambda) UploadFunction(input *UploadFunctionInput) (output *FunctionConfiguration, err error) {
	req, out := c.UploadFunctionRequest(input)
	output = out
	err = req.Send()
	return
}

var opUploadFunction *aws.Operation

type AddEventSourceInput struct {
	BatchSize    *int                `type:"integer" json:",omitempty"`
	EventSource  *string             `type:"string" required:"true"json:",omitempty"`
	FunctionName *string             `type:"string" required:"true"json:",omitempty"`
	Parameters   *map[string]*string `type:"map" json:",omitempty"`
	Role         *string             `type:"string" required:"true"json:",omitempty"`

	metadataAddEventSourceInput `json:"-", xml:"-"`
}

type metadataAddEventSourceInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeleteFunctionInput struct {
	FunctionName *string `location:"uri" locationName:"FunctionName" type:"string" required:"true"json:"-" xml:"-"`

	metadataDeleteFunctionInput `json:"-", xml:"-"`
}

type metadataDeleteFunctionInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type DeleteFunctionOutput struct {
	metadataDeleteFunctionOutput `json:"-", xml:"-"`
}

type metadataDeleteFunctionOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type EventSourceConfiguration struct {
	BatchSize    *int                `type:"integer" json:",omitempty"`
	EventSource  *string             `type:"string" json:",omitempty"`
	FunctionName *string             `type:"string" json:",omitempty"`
	IsActive     *bool               `type:"boolean" json:",omitempty"`
	LastModified *time.Time          `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	Parameters   *map[string]*string `type:"map" json:",omitempty"`
	Role         *string             `type:"string" json:",omitempty"`
	Status       *string             `type:"string" json:",omitempty"`
	UUID         *string             `type:"string" json:",omitempty"`

	metadataEventSourceConfiguration `json:"-", xml:"-"`
}

type metadataEventSourceConfiguration struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type FunctionCodeLocation struct {
	Location       *string `type:"string" json:",omitempty"`
	RepositoryType *string `type:"string" json:",omitempty"`

	metadataFunctionCodeLocation `json:"-", xml:"-"`
}

type metadataFunctionCodeLocation struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type FunctionConfiguration struct {
	CodeSize        *int64     `type:"long" json:",omitempty"`
	ConfigurationID *string    `locationName:"ConfigurationId" type:"string" json:"ConfigurationId,omitempty"`
	Description     *string    `type:"string" json:",omitempty"`
	FunctionARN     *string    `type:"string" json:",omitempty"`
	FunctionName    *string    `type:"string" json:",omitempty"`
	Handler         *string    `type:"string" json:",omitempty"`
	LastModified    *time.Time `type:"timestamp" timestampFormat:"unix" json:",omitempty"`
	MemorySize      *int       `type:"integer" json:",omitempty"`
	Mode            *string    `type:"string" json:",omitempty"`
	Role            *string    `type:"string" json:",omitempty"`
	Runtime         *string    `type:"string" json:",omitempty"`
	Timeout         *int       `type:"integer" json:",omitempty"`

	metadataFunctionConfiguration `json:"-", xml:"-"`
}

type metadataFunctionConfiguration struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetEventSourceInput struct {
	UUID *string `location:"uri" locationName:"UUID" type:"string" required:"true"json:"-" xml:"-"`

	metadataGetEventSourceInput `json:"-", xml:"-"`
}

type metadataGetEventSourceInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetFunctionConfigurationInput struct {
	FunctionName *string `location:"uri" locationName:"FunctionName" type:"string" required:"true"json:"-" xml:"-"`

	metadataGetFunctionConfigurationInput `json:"-", xml:"-"`
}

type metadataGetFunctionConfigurationInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetFunctionInput struct {
	FunctionName *string `location:"uri" locationName:"FunctionName" type:"string" required:"true"json:"-" xml:"-"`

	metadataGetFunctionInput `json:"-", xml:"-"`
}

type metadataGetFunctionInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type GetFunctionOutput struct {
	Code          *FunctionCodeLocation  `type:"structure" json:",omitempty"`
	Configuration *FunctionConfiguration `type:"structure" json:",omitempty"`

	metadataGetFunctionOutput `json:"-", xml:"-"`
}

type metadataGetFunctionOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type InvokeAsyncInput struct {
	FunctionName *string `location:"uri" locationName:"FunctionName" type:"string" required:"true"json:"-" xml:"-"`
	InvokeArgs   []byte  `type:"blob" required:"true"json:",omitempty"`

	metadataInvokeAsyncInput `json:"-", xml:"-"`
}

type metadataInvokeAsyncInput struct {
	SDKShapeTraits bool `type:"structure" payload:"InvokeArgs" json:",omitempty"`
}

type InvokeAsyncOutput struct {
	Status *int `location:"statusCode" type:"integer" json:"-" xml:"-"`

	metadataInvokeAsyncOutput `json:"-", xml:"-"`
}

type metadataInvokeAsyncOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListEventSourcesInput struct {
	EventSourceARN *string `location:"querystring" locationName:"EventSource" type:"string" json:"-" xml:"-"`
	FunctionName   *string `location:"querystring" locationName:"FunctionName" type:"string" json:"-" xml:"-"`
	Marker         *string `location:"querystring" locationName:"Marker" type:"string" json:"-" xml:"-"`
	MaxItems       *int    `location:"querystring" locationName:"MaxItems" type:"integer" json:"-" xml:"-"`

	metadataListEventSourcesInput `json:"-", xml:"-"`
}

type metadataListEventSourcesInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListEventSourcesOutput struct {
	EventSources []*EventSourceConfiguration `type:"list" json:",omitempty"`
	NextMarker   *string                     `type:"string" json:",omitempty"`

	metadataListEventSourcesOutput `json:"-", xml:"-"`
}

type metadataListEventSourcesOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListFunctionsInput struct {
	Marker   *string `location:"querystring" locationName:"Marker" type:"string" json:"-" xml:"-"`
	MaxItems *int    `location:"querystring" locationName:"MaxItems" type:"integer" json:"-" xml:"-"`

	metadataListFunctionsInput `json:"-", xml:"-"`
}

type metadataListFunctionsInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type ListFunctionsOutput struct {
	Functions  []*FunctionConfiguration `type:"list" json:",omitempty"`
	NextMarker *string                  `type:"string" json:",omitempty"`

	metadataListFunctionsOutput `json:"-", xml:"-"`
}

type metadataListFunctionsOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type RemoveEventSourceInput struct {
	UUID *string `location:"uri" locationName:"UUID" type:"string" required:"true"json:"-" xml:"-"`

	metadataRemoveEventSourceInput `json:"-", xml:"-"`
}

type metadataRemoveEventSourceInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type RemoveEventSourceOutput struct {
	metadataRemoveEventSourceOutput `json:"-", xml:"-"`
}

type metadataRemoveEventSourceOutput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type UpdateFunctionConfigurationInput struct {
	Description  *string `location:"querystring" locationName:"Description" type:"string" json:"-" xml:"-"`
	FunctionName *string `location:"uri" locationName:"FunctionName" type:"string" required:"true"json:"-" xml:"-"`
	Handler      *string `location:"querystring" locationName:"Handler" type:"string" json:"-" xml:"-"`
	MemorySize   *int    `location:"querystring" locationName:"MemorySize" type:"integer" json:"-" xml:"-"`
	Role         *string `location:"querystring" locationName:"Role" type:"string" json:"-" xml:"-"`
	Timeout      *int    `location:"querystring" locationName:"Timeout" type:"integer" json:"-" xml:"-"`

	metadataUpdateFunctionConfigurationInput `json:"-", xml:"-"`
}

type metadataUpdateFunctionConfigurationInput struct {
	SDKShapeTraits bool `type:"structure" json:",omitempty"`
}

type UploadFunctionInput struct {
	Description  *string `location:"querystring" locationName:"Description" type:"string" json:"-" xml:"-"`
	FunctionName *string `location:"uri" locationName:"FunctionName" type:"string" required:"true"json:"-" xml:"-"`
	FunctionZip  []byte  `type:"blob" required:"true"json:",omitempty"`
	Handler      *string `location:"querystring" locationName:"Handler" type:"string" required:"true"json:"-" xml:"-"`
	MemorySize   *int    `location:"querystring" locationName:"MemorySize" type:"integer" json:"-" xml:"-"`
	Mode         *string `location:"querystring" locationName:"Mode" type:"string" required:"true"json:"-" xml:"-"`
	Role         *string `location:"querystring" locationName:"Role" type:"string" required:"true"json:"-" xml:"-"`
	Runtime      *string `location:"querystring" locationName:"Runtime" type:"string" required:"true"json:"-" xml:"-"`
	Timeout      *int    `location:"querystring" locationName:"Timeout" type:"integer" json:"-" xml:"-"`

	metadataUploadFunctionInput `json:"-", xml:"-"`
}

type metadataUploadFunctionInput struct {
	SDKShapeTraits bool `type:"structure" payload:"FunctionZip" json:",omitempty"`
}