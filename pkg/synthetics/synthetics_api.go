// Code generated by tutone: DO NOT EDIT
package synthetics

import "context"

// Create a Synthetic Broken Links monitor
func (a *Synthetics) SyntheticsCreateBrokenLinksMonitor(
	accountID int,
	monitor SyntheticsCreateBrokenLinksMonitorInput,
) (*SyntheticsBrokenLinksMonitorCreateMutationResult, error) {
	return a.SyntheticsCreateBrokenLinksMonitorWithContext(context.Background(),
		accountID,
		monitor,
	)
}

// Create a Synthetic Broken Links monitor
func (a *Synthetics) SyntheticsCreateBrokenLinksMonitorWithContext(
	ctx context.Context,
	accountID int,
	monitor SyntheticsCreateBrokenLinksMonitorInput,
) (*SyntheticsBrokenLinksMonitorCreateMutationResult, error) {

	resp := SyntheticsCreateBrokenLinksMonitorQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"monitor":   monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsCreateBrokenLinksMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsBrokenLinksMonitorCreateMutationResult, nil
}

type SyntheticsCreateBrokenLinksMonitorQueryResponse struct {
	SyntheticsBrokenLinksMonitorCreateMutationResult SyntheticsBrokenLinksMonitorCreateMutationResult `json:"SyntheticsCreateBrokenLinksMonitor"`
}

const SyntheticsCreateBrokenLinksMonitorMutation = `mutation(
	$accountId: Int!,
	$monitor: SyntheticsCreateBrokenLinksMonitorInput!,
) { syntheticsCreateBrokenLinksMonitor(
	accountId: $accountId,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		createdAt
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		period
		status
		uri
	}
} }`

// Create a Synthetic Cert Check (Certificate check) monitor
func (a *Synthetics) SyntheticsCreateCertCheckMonitor(
	accountID int,
	monitor SyntheticsCreateCertCheckMonitorInput,
) (*SyntheticsCertCheckMonitorCreateMutationResult, error) {
	return a.SyntheticsCreateCertCheckMonitorWithContext(context.Background(),
		accountID,
		monitor,
	)
}

// Create a Synthetic Cert Check (Certificate check) monitor
func (a *Synthetics) SyntheticsCreateCertCheckMonitorWithContext(
	ctx context.Context,
	accountID int,
	monitor SyntheticsCreateCertCheckMonitorInput,
) (*SyntheticsCertCheckMonitorCreateMutationResult, error) {

	resp := SyntheticsCreateCertCheckMonitorQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"monitor":   monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsCreateCertCheckMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsCertCheckMonitorCreateMutationResult, nil
}

type SyntheticsCreateCertCheckMonitorQueryResponse struct {
	SyntheticsCertCheckMonitorCreateMutationResult SyntheticsCertCheckMonitorCreateMutationResult `json:"SyntheticsCreateCertCheckMonitor"`
}

const SyntheticsCreateCertCheckMonitorMutation = `mutation(
	$accountId: Int!,
	$monitor: SyntheticsCreateCertCheckMonitorInput!,
) { syntheticsCreateCertCheckMonitor(
	accountId: $accountId,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		createdAt
		domain
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		numberDaysToFailBeforeCertExpires
		period
		status
	}
} }`

// Create a Synthetics Private Location
func (a *Synthetics) SyntheticsCreatePrivateLocation(
	accountID int,
	description string,
	name string,
	verifiedScriptExecution bool,
) (*SyntheticsPrivateLocationMutationResult, error) {
	return a.SyntheticsCreatePrivateLocationWithContext(context.Background(),
		accountID,
		description,
		name,
		verifiedScriptExecution,
	)
}

// Create a Synthetics Private Location
func (a *Synthetics) SyntheticsCreatePrivateLocationWithContext(
	ctx context.Context,
	accountID int,
	description string,
	name string,
	verifiedScriptExecution bool,
) (*SyntheticsPrivateLocationMutationResult, error) {

	resp := SyntheticsCreatePrivateLocationQueryResponse{}
	vars := map[string]interface{}{
		"accountId":               accountID,
		"description":             description,
		"name":                    name,
		"verifiedScriptExecution": verifiedScriptExecution,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsCreatePrivateLocationMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsPrivateLocationMutationResult, nil
}

type SyntheticsCreatePrivateLocationQueryResponse struct {
	SyntheticsPrivateLocationMutationResult SyntheticsPrivateLocationMutationResult `json:"SyntheticsCreatePrivateLocation"`
}

const SyntheticsCreatePrivateLocationMutation = `mutation(
	$accountId: Int!,
	$description: String,
	$name: String!,
	$verifiedScriptExecution: Boolean!,
) { syntheticsCreatePrivateLocation(
	accountId: $accountId,
	description: $description,
	name: $name,
	verifiedScriptExecution: $verifiedScriptExecution,
) {
	accountId
	description
	domainId
	errors {
		description
		type
	}
	guid
	key
	locationId
	name
	verifiedScriptExecution
} }`

// Create a Synthetic Script Api monitor
func (a *Synthetics) SyntheticsCreateScriptAPIMonitor(
	accountID int,
	monitor SyntheticsCreateScriptAPIMonitorInput,
) (*SyntheticsScriptAPIMonitorCreateMutationResult, error) {
	return a.SyntheticsCreateScriptAPIMonitorWithContext(context.Background(),
		accountID,
		monitor,
	)
}

// Create a Synthetic Script Api monitor
func (a *Synthetics) SyntheticsCreateScriptAPIMonitorWithContext(
	ctx context.Context,
	accountID int,
	monitor SyntheticsCreateScriptAPIMonitorInput,
) (*SyntheticsScriptAPIMonitorCreateMutationResult, error) {

	resp := SyntheticsCreateScriptAPIMonitorQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"monitor":   monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsCreateScriptAPIMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsScriptAPIMonitorCreateMutationResult, nil
}

type SyntheticsCreateScriptAPIMonitorQueryResponse struct {
	SyntheticsScriptAPIMonitorCreateMutationResult SyntheticsScriptAPIMonitorCreateMutationResult `json:"SyntheticsCreateScriptAPIMonitor"`
}

const SyntheticsCreateScriptAPIMonitorMutation = `mutation(
	$accountId: Int!,
	$monitor: SyntheticsCreateScriptApiMonitorInput!,
) { syntheticsCreateScriptApiMonitor(
	accountId: $accountId,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		createdAt
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		period
		runtime {
			runtimeType
			runtimeTypeVersion
			scriptLanguage
		}
		status
	}
} }`

// Create a Synthetic Script Browser Monitor
func (a *Synthetics) SyntheticsCreateScriptBrowserMonitor(
	accountID int,
	monitor SyntheticsCreateScriptBrowserMonitorInput,
) (*SyntheticsScriptBrowserMonitorCreateMutationResult, error) {
	return a.SyntheticsCreateScriptBrowserMonitorWithContext(context.Background(),
		accountID,
		monitor,
	)
}

// Create a Synthetic Script Browser Monitor
func (a *Synthetics) SyntheticsCreateScriptBrowserMonitorWithContext(
	ctx context.Context,
	accountID int,
	monitor SyntheticsCreateScriptBrowserMonitorInput,
) (*SyntheticsScriptBrowserMonitorCreateMutationResult, error) {

	resp := SyntheticsCreateScriptBrowserMonitorQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"monitor":   monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsCreateScriptBrowserMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsScriptBrowserMonitorCreateMutationResult, nil
}

type SyntheticsCreateScriptBrowserMonitorQueryResponse struct {
	SyntheticsScriptBrowserMonitorCreateMutationResult SyntheticsScriptBrowserMonitorCreateMutationResult `json:"SyntheticsCreateScriptBrowserMonitor"`
}

const SyntheticsCreateScriptBrowserMonitorMutation = `mutation(
	$accountId: Int!,
	$monitor: SyntheticsCreateScriptBrowserMonitorInput!,
) { syntheticsCreateScriptBrowserMonitor(
	accountId: $accountId,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		advancedOptions {
			enableScreenshotOnFailureAndScript
		}
		createdAt
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		period
		runtime {
			runtimeType
			runtimeTypeVersion
			scriptLanguage
		}
		status
	}
} }`

// Queues a request to create a secure credential
func (a *Synthetics) SyntheticsCreateSecureCredential(
	accountID int,
	description string,
	key string,
	value SecureValue,
) (*SyntheticsSecureCredentialMutationResult, error) {
	return a.SyntheticsCreateSecureCredentialWithContext(context.Background(),
		accountID,
		description,
		key,
		value,
	)
}

// Queues a request to create a secure credential
func (a *Synthetics) SyntheticsCreateSecureCredentialWithContext(
	ctx context.Context,
	accountID int,
	description string,
	key string,
	value SecureValue,
) (*SyntheticsSecureCredentialMutationResult, error) {

	resp := SyntheticsCreateSecureCredentialQueryResponse{}
	vars := map[string]interface{}{
		"accountId":   accountID,
		"description": description,
		"key":         key,
		"value":       value,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsCreateSecureCredentialMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsSecureCredentialMutationResult, nil
}

type SyntheticsCreateSecureCredentialQueryResponse struct {
	SyntheticsSecureCredentialMutationResult SyntheticsSecureCredentialMutationResult `json:"SyntheticsCreateSecureCredential"`
}

const SyntheticsCreateSecureCredentialMutation = `mutation(
	$accountId: Int!,
	$description: String,
	$key: String!,
	$value: SecureValue!,
) { syntheticsCreateSecureCredential(
	accountId: $accountId,
	description: $description,
	key: $key,
	value: $value,
) {
	createdAt
	description
	errors {
		description
	}
	key
	lastUpdate
} }`

// Create a Synthetic Simple (Ping) monitor
func (a *Synthetics) SyntheticsCreateSimpleBrowserMonitor(
	accountID int,
	monitor SyntheticsCreateSimpleBrowserMonitorInput,
) (*SyntheticsSimpleBrowserMonitorCreateMutationResult, error) {
	return a.SyntheticsCreateSimpleBrowserMonitorWithContext(context.Background(),
		accountID,
		monitor,
	)
}

// Create a Synthetic Simple (Ping) monitor
func (a *Synthetics) SyntheticsCreateSimpleBrowserMonitorWithContext(
	ctx context.Context,
	accountID int,
	monitor SyntheticsCreateSimpleBrowserMonitorInput,
) (*SyntheticsSimpleBrowserMonitorCreateMutationResult, error) {

	resp := SyntheticsCreateSimpleBrowserMonitorQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"monitor":   monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsCreateSimpleBrowserMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsSimpleBrowserMonitorCreateMutationResult, nil
}

type SyntheticsCreateSimpleBrowserMonitorQueryResponse struct {
	SyntheticsSimpleBrowserMonitorCreateMutationResult SyntheticsSimpleBrowserMonitorCreateMutationResult `json:"SyntheticsCreateSimpleBrowserMonitor"`
}

const SyntheticsCreateSimpleBrowserMonitorMutation = `mutation(
	$accountId: Int!,
	$monitor: SyntheticsCreateSimpleBrowserMonitorInput!,
) { syntheticsCreateSimpleBrowserMonitor(
	accountId: $accountId,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		advancedOptions {
			enableScreenshotOnFailureAndScript
			responseValidationText
			useTlsValidation
		}
		createdAt
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		period
		runtime {
			runtimeType
			runtimeTypeVersion
			scriptLanguage
		}
		status
		uri
	}
} }`

// Create a Synthetic Simple (Ping) monitor
func (a *Synthetics) SyntheticsCreateSimpleMonitor(
	accountID int,
	monitor SyntheticsCreateSimpleMonitorInput,
) (*SyntheticsSimpleBrowserMonitorCreateMutationResult, error) {
	return a.SyntheticsCreateSimpleMonitorWithContext(context.Background(),
		accountID,
		monitor,
	)
}

// Create a Synthetic Simple (Ping) monitor
func (a *Synthetics) SyntheticsCreateSimpleMonitorWithContext(
	ctx context.Context,
	accountID int,
	monitor SyntheticsCreateSimpleMonitorInput,
) (*SyntheticsSimpleBrowserMonitorCreateMutationResult, error) {

	resp := SyntheticsCreateSimpleMonitorQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"monitor":   monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsCreateSimpleMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsSimpleBrowserMonitorCreateMutationResult, nil
}

type SyntheticsCreateSimpleMonitorQueryResponse struct {
	SyntheticsSimpleBrowserMonitorCreateMutationResult SyntheticsSimpleBrowserMonitorCreateMutationResult `json:"SyntheticsCreateSimpleMonitor"`
}

const SyntheticsCreateSimpleMonitorMutation = `mutation(
	$accountId: Int!,
	$monitor: SyntheticsCreateSimpleMonitorInput!,
) { syntheticsCreateSimpleMonitor(
	accountId: $accountId,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		advancedOptions {
			enableScreenshotOnFailureAndScript
			responseValidationText
			useTlsValidation
		}
		createdAt
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		period
		runtime {
			runtimeType
			runtimeTypeVersion
			scriptLanguage
		}
		status
		uri
	}
} }`

// Create a Synthetic Step monitor
func (a *Synthetics) SyntheticsCreateStepMonitor(
	accountID int,
	monitor SyntheticsCreateStepMonitorInput,
) (*SyntheticsStepMonitorCreateMutationResult, error) {
	return a.SyntheticsCreateStepMonitorWithContext(context.Background(),
		accountID,
		monitor,
	)
}

// Create a Synthetic Step monitor
func (a *Synthetics) SyntheticsCreateStepMonitorWithContext(
	ctx context.Context,
	accountID int,
	monitor SyntheticsCreateStepMonitorInput,
) (*SyntheticsStepMonitorCreateMutationResult, error) {

	resp := SyntheticsCreateStepMonitorQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"monitor":   monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsCreateStepMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsStepMonitorCreateMutationResult, nil
}

type SyntheticsCreateStepMonitorQueryResponse struct {
	SyntheticsStepMonitorCreateMutationResult SyntheticsStepMonitorCreateMutationResult `json:"SyntheticsCreateStepMonitor"`
}

const SyntheticsCreateStepMonitorMutation = `mutation(
	$accountId: Int!,
	$monitor: SyntheticsCreateStepMonitorInput!,
) { syntheticsCreateStepMonitor(
	accountId: $accountId,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		advancedOptions {
			enableScreenshotOnFailureAndScript
		}
		createdAt
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		period
		status
		steps {
			ordinal
			type
			values
		}
	}
} }`

// Delete a Synthetic Monitor
func (a *Synthetics) SyntheticsDeleteMonitor(
	gUID EntityGUID,
) (*SyntheticsMonitorDeleteMutationResult, error) {
	return a.SyntheticsDeleteMonitorWithContext(context.Background(),
		gUID,
	)
}

// Delete a Synthetic Monitor
func (a *Synthetics) SyntheticsDeleteMonitorWithContext(
	ctx context.Context,
	gUID EntityGUID,
) (*SyntheticsMonitorDeleteMutationResult, error) {

	resp := SyntheticsDeleteMonitorQueryResponse{}
	vars := map[string]interface{}{
		"guid": gUID,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsDeleteMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsMonitorDeleteMutationResult, nil
}

type SyntheticsDeleteMonitorQueryResponse struct {
	SyntheticsMonitorDeleteMutationResult SyntheticsMonitorDeleteMutationResult `json:"SyntheticsDeleteMonitor"`
}

const SyntheticsDeleteMonitorMutation = `mutation(
	$guid: EntityGuid!,
) { syntheticsDeleteMonitor(
	guid: $guid,
) {
	deletedGuid
} }`

// Delete a Synthetics Private Location
func (a *Synthetics) SyntheticsDeletePrivateLocation(
	gUID EntityGUID,
) (*SyntheticsPrivateLocationDeleteResult, error) {
	return a.SyntheticsDeletePrivateLocationWithContext(context.Background(),
		gUID,
	)
}

// Delete a Synthetics Private Location
func (a *Synthetics) SyntheticsDeletePrivateLocationWithContext(
	ctx context.Context,
	gUID EntityGUID,
) (*SyntheticsPrivateLocationDeleteResult, error) {

	resp := SyntheticsDeletePrivateLocationQueryResponse{}
	vars := map[string]interface{}{
		"guid": gUID,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsDeletePrivateLocationMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsPrivateLocationDeleteResult, nil
}

type SyntheticsDeletePrivateLocationQueryResponse struct {
	SyntheticsPrivateLocationDeleteResult SyntheticsPrivateLocationDeleteResult `json:"SyntheticsDeletePrivateLocation"`
}

const SyntheticsDeletePrivateLocationMutation = `mutation(
	$guid: EntityGuid!,
) { syntheticsDeletePrivateLocation(
	guid: $guid,
) {
	errors {
		description
		type
	}
} }`

// Queues a request to delete an existing secure credential
func (a *Synthetics) SyntheticsDeleteSecureCredential(
	accountID int,
	key string,
) (*SyntheticsSecureCredentialMutationResult, error) {
	return a.SyntheticsDeleteSecureCredentialWithContext(context.Background(),
		accountID,
		key,
	)
}

// Queues a request to delete an existing secure credential
func (a *Synthetics) SyntheticsDeleteSecureCredentialWithContext(
	ctx context.Context,
	accountID int,
	key string,
) (*SyntheticsSecureCredentialMutationResult, error) {

	resp := SyntheticsDeleteSecureCredentialQueryResponse{}
	vars := map[string]interface{}{
		"accountId": accountID,
		"key":       key,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsDeleteSecureCredentialMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsSecureCredentialMutationResult, nil
}

type SyntheticsDeleteSecureCredentialQueryResponse struct {
	SyntheticsSecureCredentialMutationResult SyntheticsSecureCredentialMutationResult `json:"SyntheticsDeleteSecureCredential"`
}

const SyntheticsDeleteSecureCredentialMutation = `mutation(
	$accountId: Int!,
	$key: String!,
) { syntheticsDeleteSecureCredential(
	accountId: $accountId,
	key: $key,
) {
	createdAt
	description
	errors {
		description
	}
	key
	lastUpdate
} }`

// Purge the job queue for a specified private location
func (a *Synthetics) SyntheticsPurgePrivateLocationQueue(
	gUID EntityGUID,
) (*SyntheticsPrivateLocationPurgeQueueResult, error) {
	return a.SyntheticsPurgePrivateLocationQueueWithContext(context.Background(),
		gUID,
	)
}

// Purge the job queue for a specified private location
func (a *Synthetics) SyntheticsPurgePrivateLocationQueueWithContext(
	ctx context.Context,
	gUID EntityGUID,
) (*SyntheticsPrivateLocationPurgeQueueResult, error) {

	resp := SyntheticsPurgePrivateLocationQueueQueryResponse{}
	vars := map[string]interface{}{
		"guid": gUID,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsPurgePrivateLocationQueueMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsPrivateLocationPurgeQueueResult, nil
}

type SyntheticsPurgePrivateLocationQueueQueryResponse struct {
	SyntheticsPrivateLocationPurgeQueueResult SyntheticsPrivateLocationPurgeQueueResult `json:"SyntheticsPurgePrivateLocationQueue"`
}

const SyntheticsPurgePrivateLocationQueueMutation = `mutation(
	$guid: EntityGuid!,
) { syntheticsPurgePrivateLocationQueue(
	guid: $guid,
) {
	errors {
		description
		type
	}
} }`

// Update a Synthetic Broken Links monitor
func (a *Synthetics) SyntheticsUpdateBrokenLinksMonitor(
	gUID EntityGUID,
	monitor SyntheticsUpdateBrokenLinksMonitorInput,
) (*SyntheticsBrokenLinksMonitorUpdateMutationResult, error) {
	return a.SyntheticsUpdateBrokenLinksMonitorWithContext(context.Background(),
		gUID,
		monitor,
	)
}

// Update a Synthetic Broken Links monitor
func (a *Synthetics) SyntheticsUpdateBrokenLinksMonitorWithContext(
	ctx context.Context,
	gUID EntityGUID,
	monitor SyntheticsUpdateBrokenLinksMonitorInput,
) (*SyntheticsBrokenLinksMonitorUpdateMutationResult, error) {

	resp := SyntheticsUpdateBrokenLinksMonitorQueryResponse{}
	vars := map[string]interface{}{
		"guid":    gUID,
		"monitor": monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsUpdateBrokenLinksMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsBrokenLinksMonitorUpdateMutationResult, nil
}

type SyntheticsUpdateBrokenLinksMonitorQueryResponse struct {
	SyntheticsBrokenLinksMonitorUpdateMutationResult SyntheticsBrokenLinksMonitorUpdateMutationResult `json:"SyntheticsUpdateBrokenLinksMonitor"`
}

const SyntheticsUpdateBrokenLinksMonitorMutation = `mutation(
	$guid: EntityGuid!,
	$monitor: SyntheticsUpdateBrokenLinksMonitorInput!,
) { syntheticsUpdateBrokenLinksMonitor(
	guid: $guid,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		createdAt
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		period
		status
		uri
	}
} }`

// Update a Synthetic Cert Check (Certificate check) monitor
func (a *Synthetics) SyntheticsUpdateCertCheckMonitor(
	gUID EntityGUID,
	monitor SyntheticsUpdateCertCheckMonitorInput,
) (*SyntheticsCertCheckMonitorUpdateMutationResult, error) {
	return a.SyntheticsUpdateCertCheckMonitorWithContext(context.Background(),
		gUID,
		monitor,
	)
}

// Update a Synthetic Cert Check (Certificate check) monitor
func (a *Synthetics) SyntheticsUpdateCertCheckMonitorWithContext(
	ctx context.Context,
	gUID EntityGUID,
	monitor SyntheticsUpdateCertCheckMonitorInput,
) (*SyntheticsCertCheckMonitorUpdateMutationResult, error) {

	resp := SyntheticsUpdateCertCheckMonitorQueryResponse{}
	vars := map[string]interface{}{
		"guid":    gUID,
		"monitor": monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsUpdateCertCheckMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsCertCheckMonitorUpdateMutationResult, nil
}

type SyntheticsUpdateCertCheckMonitorQueryResponse struct {
	SyntheticsCertCheckMonitorUpdateMutationResult SyntheticsCertCheckMonitorUpdateMutationResult `json:"SyntheticsUpdateCertCheckMonitor"`
}

const SyntheticsUpdateCertCheckMonitorMutation = `mutation(
	$guid: EntityGuid!,
	$monitor: SyntheticsUpdateCertCheckMonitorInput!,
) { syntheticsUpdateCertCheckMonitor(
	guid: $guid,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		createdAt
		domain
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		numberDaysToFailBeforeCertExpires
		period
		status
	}
} }`

// Update a Synthetics Private Location
func (a *Synthetics) SyntheticsUpdatePrivateLocation(
	description string,
	gUID EntityGUID,
	verifiedScriptExecution bool,
) (*SyntheticsPrivateLocationMutationResult, error) {
	return a.SyntheticsUpdatePrivateLocationWithContext(context.Background(),
		description,
		gUID,
		verifiedScriptExecution,
	)
}

// Update a Synthetics Private Location
func (a *Synthetics) SyntheticsUpdatePrivateLocationWithContext(
	ctx context.Context,
	description string,
	gUID EntityGUID,
	verifiedScriptExecution bool,
) (*SyntheticsPrivateLocationMutationResult, error) {

	resp := SyntheticsUpdatePrivateLocationQueryResponse{}
	vars := map[string]interface{}{
		"description":             description,
		"guid":                    gUID,
		"verifiedScriptExecution": verifiedScriptExecution,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsUpdatePrivateLocationMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsPrivateLocationMutationResult, nil
}

type SyntheticsUpdatePrivateLocationQueryResponse struct {
	SyntheticsPrivateLocationMutationResult SyntheticsPrivateLocationMutationResult `json:"SyntheticsUpdatePrivateLocation"`
}

const SyntheticsUpdatePrivateLocationMutation = `mutation(
	$description: String,
	$guid: EntityGuid!,
	$verifiedScriptExecution: Boolean,
) { syntheticsUpdatePrivateLocation(
	description: $description,
	guid: $guid,
	verifiedScriptExecution: $verifiedScriptExecution,
) {
	accountId
	description
	domainId
	errors {
		description
		type
	}
	guid
	key
	locationId
	name
	verifiedScriptExecution
} }`

// Update a Synthetic Script Api monitor
func (a *Synthetics) SyntheticsUpdateScriptAPIMonitor(
	gUID EntityGUID,
	monitor SyntheticsUpdateScriptAPIMonitorInput,
) (*SyntheticsScriptAPIMonitorUpdateMutationResult, error) {
	return a.SyntheticsUpdateScriptAPIMonitorWithContext(context.Background(),
		gUID,
		monitor,
	)
}

// Update a Synthetic Script Api monitor
func (a *Synthetics) SyntheticsUpdateScriptAPIMonitorWithContext(
	ctx context.Context,
	gUID EntityGUID,
	monitor SyntheticsUpdateScriptAPIMonitorInput,
) (*SyntheticsScriptAPIMonitorUpdateMutationResult, error) {

	resp := SyntheticsUpdateScriptAPIMonitorQueryResponse{}
	vars := map[string]interface{}{
		"guid":    gUID,
		"monitor": monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsUpdateScriptAPIMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsScriptAPIMonitorUpdateMutationResult, nil
}

type SyntheticsUpdateScriptAPIMonitorQueryResponse struct {
	SyntheticsScriptAPIMonitorUpdateMutationResult SyntheticsScriptAPIMonitorUpdateMutationResult `json:"SyntheticsUpdateScriptAPIMonitor"`
}

const SyntheticsUpdateScriptAPIMonitorMutation = `mutation(
	$guid: EntityGuid!,
	$monitor: SyntheticsUpdateScriptApiMonitorInput!,
) { syntheticsUpdateScriptApiMonitor(
	guid: $guid,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		createdAt
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		period
		runtime {
			runtimeType
			runtimeTypeVersion
			scriptLanguage
		}
		status
	}
} }`

// Update a Synthetic Script Browser Monitor
func (a *Synthetics) SyntheticsUpdateScriptBrowserMonitor(
	gUID EntityGUID,
	monitor SyntheticsUpdateScriptBrowserMonitorInput,
) (*SyntheticsScriptBrowserMonitorUpdateMutationResult, error) {
	return a.SyntheticsUpdateScriptBrowserMonitorWithContext(context.Background(),
		gUID,
		monitor,
	)
}

// Update a Synthetic Script Browser Monitor
func (a *Synthetics) SyntheticsUpdateScriptBrowserMonitorWithContext(
	ctx context.Context,
	gUID EntityGUID,
	monitor SyntheticsUpdateScriptBrowserMonitorInput,
) (*SyntheticsScriptBrowserMonitorUpdateMutationResult, error) {

	resp := SyntheticsUpdateScriptBrowserMonitorQueryResponse{}
	vars := map[string]interface{}{
		"guid":    gUID,
		"monitor": monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsUpdateScriptBrowserMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsScriptBrowserMonitorUpdateMutationResult, nil
}

type SyntheticsUpdateScriptBrowserMonitorQueryResponse struct {
	SyntheticsScriptBrowserMonitorUpdateMutationResult SyntheticsScriptBrowserMonitorUpdateMutationResult `json:"SyntheticsUpdateScriptBrowserMonitor"`
}

const SyntheticsUpdateScriptBrowserMonitorMutation = `mutation(
	$guid: EntityGuid!,
	$monitor: SyntheticsUpdateScriptBrowserMonitorInput!,
) { syntheticsUpdateScriptBrowserMonitor(
	guid: $guid,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		advancedOptions {
			enableScreenshotOnFailureAndScript
		}
		createdAt
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		period
		runtime {
			runtimeType
			runtimeTypeVersion
			scriptLanguage
		}
		status
	}
} }`

// Queues a request to update an existing secure credential
func (a *Synthetics) SyntheticsUpdateSecureCredential(
	accountID int,
	description string,
	key string,
	value SecureValue,
) (*SyntheticsSecureCredentialMutationResult, error) {
	return a.SyntheticsUpdateSecureCredentialWithContext(context.Background(),
		accountID,
		description,
		key,
		value,
	)
}

// Queues a request to update an existing secure credential
func (a *Synthetics) SyntheticsUpdateSecureCredentialWithContext(
	ctx context.Context,
	accountID int,
	description string,
	key string,
	value SecureValue,
) (*SyntheticsSecureCredentialMutationResult, error) {

	resp := SyntheticsUpdateSecureCredentialQueryResponse{}
	vars := map[string]interface{}{
		"accountId":   accountID,
		"description": description,
		"key":         key,
		"value":       value,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsUpdateSecureCredentialMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsSecureCredentialMutationResult, nil
}

type SyntheticsUpdateSecureCredentialQueryResponse struct {
	SyntheticsSecureCredentialMutationResult SyntheticsSecureCredentialMutationResult `json:"SyntheticsUpdateSecureCredential"`
}

const SyntheticsUpdateSecureCredentialMutation = `mutation(
	$accountId: Int!,
	$description: String,
	$key: String!,
	$value: SecureValue,
) { syntheticsUpdateSecureCredential(
	accountId: $accountId,
	description: $description,
	key: $key,
	value: $value,
) {
	createdAt
	description
	errors {
		description
	}
	key
	lastUpdate
} }`

// Update a Synthetic Simple Browser monitor
func (a *Synthetics) SyntheticsUpdateSimpleBrowserMonitor(
	gUID EntityGUID,
	monitor SyntheticsUpdateSimpleBrowserMonitorInput,
) (*SyntheticsSimpleBrowserMonitorUpdateMutationResult, error) {
	return a.SyntheticsUpdateSimpleBrowserMonitorWithContext(context.Background(),
		gUID,
		monitor,
	)
}

// Update a Synthetic Simple Browser monitor
func (a *Synthetics) SyntheticsUpdateSimpleBrowserMonitorWithContext(
	ctx context.Context,
	gUID EntityGUID,
	monitor SyntheticsUpdateSimpleBrowserMonitorInput,
) (*SyntheticsSimpleBrowserMonitorUpdateMutationResult, error) {

	resp := SyntheticsUpdateSimpleBrowserMonitorQueryResponse{}
	vars := map[string]interface{}{
		"guid":    gUID,
		"monitor": monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsUpdateSimpleBrowserMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsSimpleBrowserMonitorUpdateMutationResult, nil
}

type SyntheticsUpdateSimpleBrowserMonitorQueryResponse struct {
	SyntheticsSimpleBrowserMonitorUpdateMutationResult SyntheticsSimpleBrowserMonitorUpdateMutationResult `json:"SyntheticsUpdateSimpleBrowserMonitor"`
}

const SyntheticsUpdateSimpleBrowserMonitorMutation = `mutation(
	$guid: EntityGuid!,
	$monitor: SyntheticsUpdateSimpleBrowserMonitorInput!,
) { syntheticsUpdateSimpleBrowserMonitor(
	guid: $guid,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		advancedOptions {
			enableScreenshotOnFailureAndScript
			responseValidationText
			useTlsValidation
		}
		createdAt
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		period
		runtime {
			runtimeType
			runtimeTypeVersion
			scriptLanguage
		}
		status
		uri
	}
} }`

// Update a Synthetic Simple (Ping) monitor
func (a *Synthetics) SyntheticsUpdateSimpleMonitor(
	gUID EntityGUID,
	monitor SyntheticsUpdateSimpleMonitorInput,
) (*SyntheticsSimpleMonitorUpdateMutationResult, error) {
	return a.SyntheticsUpdateSimpleMonitorWithContext(context.Background(),
		gUID,
		monitor,
	)
}

// Update a Synthetic Simple (Ping) monitor
func (a *Synthetics) SyntheticsUpdateSimpleMonitorWithContext(
	ctx context.Context,
	gUID EntityGUID,
	monitor SyntheticsUpdateSimpleMonitorInput,
) (*SyntheticsSimpleMonitorUpdateMutationResult, error) {

	resp := SyntheticsUpdateSimpleMonitorQueryResponse{}
	vars := map[string]interface{}{
		"guid":    gUID,
		"monitor": monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsUpdateSimpleMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsSimpleMonitorUpdateMutationResult, nil
}

type SyntheticsUpdateSimpleMonitorQueryResponse struct {
	SyntheticsSimpleMonitorUpdateMutationResult SyntheticsSimpleMonitorUpdateMutationResult `json:"SyntheticsUpdateSimpleMonitor"`
}

const SyntheticsUpdateSimpleMonitorMutation = `mutation(
	$guid: EntityGuid!,
	$monitor: SyntheticsUpdateSimpleMonitorInput!,
) { syntheticsUpdateSimpleMonitor(
	guid: $guid,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		advancedOptions {
			redirectIsFailure
			responseValidationText
			shouldBypassHeadRequest
			useTlsValidation
		}
		createdAt
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		period
		status
		uri
	}
} }`

// Update a Synthetic Step monitor
func (a *Synthetics) SyntheticsUpdateStepMonitor(
	gUID EntityGUID,
	monitor SyntheticsUpdateStepMonitorInput,
) (*SyntheticsStepMonitorUpdateMutationResult, error) {
	return a.SyntheticsUpdateStepMonitorWithContext(context.Background(),
		gUID,
		monitor,
	)
}

// Update a Synthetic Step monitor
func (a *Synthetics) SyntheticsUpdateStepMonitorWithContext(
	ctx context.Context,
	gUID EntityGUID,
	monitor SyntheticsUpdateStepMonitorInput,
) (*SyntheticsStepMonitorUpdateMutationResult, error) {

	resp := SyntheticsUpdateStepMonitorQueryResponse{}
	vars := map[string]interface{}{
		"guid":    gUID,
		"monitor": monitor,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, SyntheticsUpdateStepMonitorMutation, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.SyntheticsStepMonitorUpdateMutationResult, nil
}

type SyntheticsUpdateStepMonitorQueryResponse struct {
	SyntheticsStepMonitorUpdateMutationResult SyntheticsStepMonitorUpdateMutationResult `json:"SyntheticsUpdateStepMonitor"`
}

const SyntheticsUpdateStepMonitorMutation = `mutation(
	$guid: EntityGuid!,
	$monitor: SyntheticsUpdateStepMonitorInput!,
) { syntheticsUpdateStepMonitor(
	guid: $guid,
	monitor: $monitor,
) {
	errors {
		description
		type
	}
	monitor {
		advancedOptions {
			enableScreenshotOnFailureAndScript
		}
		createdAt
		guid
		locations {
			private
			public
		}
		modifiedAt
		name
		period
		status
		steps {
			ordinal
			type
			values
		}
	}
} }`
