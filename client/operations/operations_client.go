// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteApplicationsID Delete the application specified
*/
func (a *Client) DeleteApplicationsID(params *DeleteApplicationsIDParams) (*DeleteApplicationsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteApplicationsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteApplicationsID",
		Method:             "DELETE",
		PathPattern:        "/applications/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteApplicationsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteApplicationsIDNoContent), nil

}

/*
DeleteDeploymentsName Delete a deployment request
*/
func (a *Client) DeleteDeploymentsName(params *DeleteDeploymentsNameParams) (*DeleteDeploymentsNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDeploymentsNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteDeploymentsName",
		Method:             "DELETE",
		PathPattern:        "/deployments/{name}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDeploymentsNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDeploymentsNameNoContent), nil

}

/*
DeleteDynamicnodesID Delete the dynamic node specified
*/
func (a *Client) DeleteDynamicnodesID(params *DeleteDynamicnodesIDParams) (*DeleteDynamicnodesIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDynamicnodesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteDynamicnodesID",
		Method:             "DELETE",
		PathPattern:        "/dynamicnodes/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDynamicnodesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDynamicnodesIDNoContent), nil

}

/*
DeleteExternalendpointsID Delete the external endpoint specified
*/
func (a *Client) DeleteExternalendpointsID(params *DeleteExternalendpointsIDParams) (*DeleteExternalendpointsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteExternalendpointsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteExternalendpointsID",
		Method:             "DELETE",
		PathPattern:        "/externalendpoints/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteExternalendpointsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteExternalendpointsIDNoContent), nil

}

/*
DeleteMicroservicesID Delete a microservice
*/
func (a *Client) DeleteMicroservicesID(params *DeleteMicroservicesIDParams) (*DeleteMicroservicesIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteMicroservicesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteMicroservicesID",
		Method:             "DELETE",
		PathPattern:        "/microservices/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteMicroservicesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMicroservicesIDNoContent), nil

}

/*
DeleteNodesID Delete a node
*/
func (a *Client) DeleteNodesID(params *DeleteNodesIDParams) (*DeleteNodesIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNodesID",
		Method:             "DELETE",
		PathPattern:        "/nodes/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNodesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodesIDNoContent), nil

}

/*
DeleteRegionsID Delete the region
*/
func (a *Client) DeleteRegionsID(params *DeleteRegionsIDParams) (*DeleteRegionsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRegionsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRegionsID",
		Method:             "DELETE",
		PathPattern:        "/regions/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRegionsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRegionsIDNoContent), nil

}

/*
DeleteRelationshipsID Delete the relationship specified
*/
func (a *Client) DeleteRelationshipsID(params *DeleteRelationshipsIDParams) (*DeleteRelationshipsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRelationshipsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRelationshipsID",
		Method:             "DELETE",
		PathPattern:        "/relationships/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRelationshipsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRelationshipsIDNoContent), nil

}

/*
GetApplications Returns the list of all application deployed
*/
func (a *Client) GetApplications(params *GetApplicationsParams) (*GetApplicationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetApplications",
		Method:             "GET",
		PathPattern:        "/applications",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetApplicationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetApplicationsOK), nil

}

/*
GetApplicationsID Returns the given application
*/
func (a *Client) GetApplicationsID(params *GetApplicationsIDParams) (*GetApplicationsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetApplicationsID",
		Method:             "GET",
		PathPattern:        "/applications/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetApplicationsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetApplicationsIDOK), nil

}

/*
GetDataflows Returns the list of all dataflows
*/
func (a *Client) GetDataflows(params *GetDataflowsParams) (*GetDataflowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDataflowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDataflows",
		Method:             "GET",
		PathPattern:        "/dataflows",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDataflowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDataflowsOK), nil

}

/*
GetDataflowsID Returns the given dataflow
*/
func (a *Client) GetDataflowsID(params *GetDataflowsIDParams) (*GetDataflowsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDataflowsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDataflowsID",
		Method:             "GET",
		PathPattern:        "/dataflows/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDataflowsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDataflowsIDOK), nil

}

/*
GetDeployments Returns the list of all deployment requests
*/
func (a *Client) GetDeployments(params *GetDeploymentsParams) (*GetDeploymentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDeployments",
		Method:             "GET",
		PathPattern:        "/deployments",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDeploymentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeploymentsOK), nil

}

/*
GetDeploymentsName Returns the given deployment request
*/
func (a *Client) GetDeploymentsName(params *GetDeploymentsNameParams) (*GetDeploymentsNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentsNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDeploymentsName",
		Method:             "GET",
		PathPattern:        "/deployments/{name}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDeploymentsNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeploymentsNameOK), nil

}

/*
GetDynamicnodes Returns the list of all dynamic nodes
*/
func (a *Client) GetDynamicnodes(params *GetDynamicnodesParams) (*GetDynamicnodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDynamicnodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDynamicnodes",
		Method:             "GET",
		PathPattern:        "/dynamicnodes",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDynamicnodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDynamicnodesOK), nil

}

/*
GetDynamicnodesID Returns the given dynamic node
*/
func (a *Client) GetDynamicnodesID(params *GetDynamicnodesIDParams) (*GetDynamicnodesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDynamicnodesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetDynamicnodesID",
		Method:             "GET",
		PathPattern:        "/dynamicnodes/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDynamicnodesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDynamicnodesIDOK), nil

}

/*
GetExternalendpoints Returns the list of all external endpoints (services, sensors, devices)
*/
func (a *Client) GetExternalendpoints(params *GetExternalendpointsParams) (*GetExternalendpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExternalendpointsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetExternalendpoints",
		Method:             "GET",
		PathPattern:        "/externalendpoints",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExternalendpointsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetExternalendpointsOK), nil

}

/*
GetExternalendpointsID Returns the given external endpoint
*/
func (a *Client) GetExternalendpointsID(params *GetExternalendpointsIDParams) (*GetExternalendpointsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExternalendpointsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetExternalendpointsID",
		Method:             "GET",
		PathPattern:        "/externalendpoints/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetExternalendpointsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetExternalendpointsIDOK), nil

}

/*
GetMicroservices Returns the list of all microservices
*/
func (a *Client) GetMicroservices(params *GetMicroservicesParams) (*GetMicroservicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMicroservicesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMicroservices",
		Method:             "GET",
		PathPattern:        "/microservices",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMicroservicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMicroservicesOK), nil

}

/*
GetMicroservicesID Returns the given microservice
*/
func (a *Client) GetMicroservicesID(params *GetMicroservicesIDParams) (*GetMicroservicesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMicroservicesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMicroservicesID",
		Method:             "GET",
		PathPattern:        "/microservices/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMicroservicesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMicroservicesIDOK), nil

}

/*
GetNodes Returns the list of all physical or virtual nodes
*/
func (a *Client) GetNodes(params *GetNodesParams) (*GetNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodes",
		Method:             "GET",
		PathPattern:        "/nodes",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesOK), nil

}

/*
GetNodesID Returns the given physical or virtual node
*/
func (a *Client) GetNodesID(params *GetNodesIDParams) (*GetNodesIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNodesID",
		Method:             "GET",
		PathPattern:        "/nodes/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIDOK), nil

}

/*
GetRegions Returns the list of the regions
*/
func (a *Client) GetRegions(params *GetRegionsParams) (*GetRegionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRegionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRegions",
		Method:             "GET",
		PathPattern:        "/regions",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRegionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRegionsOK), nil

}

/*
GetRegionsID Returns the given region
*/
func (a *Client) GetRegionsID(params *GetRegionsIDParams) (*GetRegionsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRegionsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRegionsID",
		Method:             "GET",
		PathPattern:        "/regions/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRegionsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRegionsIDOK), nil

}

/*
GetRelationships Returns the list of all relationships
*/
func (a *Client) GetRelationships(params *GetRelationshipsParams) (*GetRelationshipsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRelationshipsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRelationships",
		Method:             "GET",
		PathPattern:        "/relationships",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRelationshipsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRelationshipsOK), nil

}

/*
GetRelationshipsID Returns the given relationship
*/
func (a *Client) GetRelationshipsID(params *GetRelationshipsIDParams) (*GetRelationshipsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRelationshipsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRelationshipsID",
		Method:             "GET",
		PathPattern:        "/relationships/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRelationshipsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRelationshipsIDOK), nil

}

/*
PatchDeploymentsName Update the status of deployment request
*/
func (a *Client) PatchDeploymentsName(params *PatchDeploymentsNameParams) (*PatchDeploymentsNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchDeploymentsNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchDeploymentsName",
		Method:             "PATCH",
		PathPattern:        "/deployments/{name}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PatchDeploymentsNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchDeploymentsNameOK), nil

}

/*
PutApplicationsID Update/create an application
*/
func (a *Client) PutApplicationsID(params *PutApplicationsIDParams) (*PutApplicationsIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutApplicationsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutApplicationsID",
		Method:             "PUT",
		PathPattern:        "/applications/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutApplicationsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutApplicationsIDCreated), nil

}

/*
PutDataflowsID Update/create a dataflow
*/
func (a *Client) PutDataflowsID(params *PutDataflowsIDParams) (*PutDataflowsIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutDataflowsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutDataflowsID",
		Method:             "PUT",
		PathPattern:        "/dataflows/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutDataflowsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutDataflowsIDCreated), nil

}

/*
PutDeploymentsName Update/create a deployment request
*/
func (a *Client) PutDeploymentsName(params *PutDeploymentsNameParams) (*PutDeploymentsNameCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutDeploymentsNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutDeploymentsName",
		Method:             "PUT",
		PathPattern:        "/deployments/{name}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutDeploymentsNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutDeploymentsNameCreated), nil

}

/*
PutDynamicnodesID Update/create a dynamic node
*/
func (a *Client) PutDynamicnodesID(params *PutDynamicnodesIDParams) (*PutDynamicnodesIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutDynamicnodesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutDynamicnodesID",
		Method:             "PUT",
		PathPattern:        "/dynamicnodes/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutDynamicnodesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutDynamicnodesIDCreated), nil

}

/*
PutExternalendpointsID Update/create an external endpoint
*/
func (a *Client) PutExternalendpointsID(params *PutExternalendpointsIDParams) (*PutExternalendpointsIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutExternalendpointsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutExternalendpointsID",
		Method:             "PUT",
		PathPattern:        "/externalendpoints/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutExternalendpointsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutExternalendpointsIDCreated), nil

}

/*
PutMicroservicesID Update/create a microservice
*/
func (a *Client) PutMicroservicesID(params *PutMicroservicesIDParams) (*PutMicroservicesIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutMicroservicesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutMicroservicesID",
		Method:             "PUT",
		PathPattern:        "/microservices/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutMicroservicesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutMicroservicesIDCreated), nil

}

/*
PutNodesID Update/create a node
*/
func (a *Client) PutNodesID(params *PutNodesIDParams) (*PutNodesIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutNodesIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutNodesID",
		Method:             "PUT",
		PathPattern:        "/nodes/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutNodesIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutNodesIDCreated), nil

}

/*
PutRegionsID Update/create a region
*/
func (a *Client) PutRegionsID(params *PutRegionsIDParams) (*PutRegionsIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutRegionsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutRegionsID",
		Method:             "PUT",
		PathPattern:        "/regions/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutRegionsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutRegionsIDCreated), nil

}

/*
PutRelationshipsID Update/create a relationship
*/
func (a *Client) PutRelationshipsID(params *PutRelationshipsIDParams) (*PutRelationshipsIDCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutRelationshipsIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutRelationshipsID",
		Method:             "PUT",
		PathPattern:        "/relationships/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutRelationshipsIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutRelationshipsIDCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}