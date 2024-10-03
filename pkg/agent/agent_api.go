// Code generated by tutone: DO NOT EDIT
package agent

import "context"

// The current release of the provided Agent
func (a *Agent) GetCurrentAgentRelease(
	agentName AgentReleasesFilter,
) (*AgentReleasesAgentRelease, error) {
	return a.GetCurrentAgentReleaseWithContext(context.Background(),
		agentName,
	)
}

// The current release of the provided Agent
func (a *Agent) GetCurrentAgentReleaseWithContext(
	ctx context.Context,
	agentName AgentReleasesFilter,
) (*AgentReleasesAgentRelease, error) {

	resp := currentAgentReleaseResponse{}
	vars := map[string]interface{}{
		"agentName": agentName,
	}

	if err := a.client.NerdGraphQueryWithContext(ctx, getCurrentAgentReleaseQuery, vars, &resp); err != nil {
		return nil, err
	}

	return &resp.DocumentationFields.CurrentAgentRelease, nil
}

const getCurrentAgentReleaseQuery = `query(
	$agentName: AgentReleasesFilter!,
) { docs { currentAgentRelease(
	agentName: $agentName,
) {
	bugs
	date
	downloadLink
	eolDate
	features
	security
	slug
	version
} } }`
