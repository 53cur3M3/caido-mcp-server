package tools

import (
	caido "github.com/caido-community/sdk-go"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func RegisterAll(server *mcp.Server, client *caido.Client) {
	// HTTP History
	RegisterListRequestsTool(server, client)
	RegisterGetRequestTool(server, client)

	// Automate (Fuzzing)
	RegisterListAutomateSessionsTool(server, client)
	RegisterGetAutomateSessionTool(server, client)
	RegisterGetAutomateEntryTool(server, client)
	RegisterAutomateTaskControlTool(server, client)

	// Replay (Send Requests)
	RegisterSendRequestTool(server, client)
	RegisterBatchSendTool(server, client)
	RegisterEditRequestTool(server, client)
	RegisterExportCurlTool(server, client)
	RegisterCreateReplaySessionTool(server, client)
	RegisterListReplaySessionsTool(server, client)
	RegisterDeleteReplaySessionsTool(server, client)
	RegisterMoveReplaySessionTool(server, client)
	RegisterGetReplayEntryTool(server, client)
	RegisterClearSessionCookiesTool(server, client)
	RegisterGetSessionCookiesTool(server, client)

	// Replay Collections
	RegisterListReplayCollectionsTool(server, client)
	RegisterCreateReplayCollectionTool(server, client)
	RegisterRenameReplayCollectionTool(server, client)
	RegisterDeleteReplayCollectionTool(server, client)

	// Findings
	RegisterListFindingsTool(server, client)
	RegisterCreateFindingTool(server, client)
	RegisterDeleteFindingsTool(server, client)
	RegisterExportFindingsTool(server, client)

	// Sitemap
	RegisterGetSitemapTool(server, client)

	// Scopes
	RegisterListScopesTool(server, client)
	RegisterCreateScopeTool(server, client)
	RegisterRenameScopeTool(server, client)
	RegisterDeleteScopeTool(server, client)

	// Projects
	RegisterListProjectsTool(server, client)
	RegisterSelectProjectTool(server, client)
	RegisterCreateProjectTool(server, client)
	RegisterRenameProjectTool(server, client)
	RegisterDeleteProjectTool(server, client)

	// Workflows
	RegisterListWorkflowsTool(server, client)
	RegisterRunWorkflowTool(server, client)
	RegisterToggleWorkflowTool(server, client)

	// Environments
	RegisterListEnvironmentsTool(server, client)
	RegisterSelectEnvironmentTool(server, client)
	RegisterCreateEnvironmentTool(server, client)
	RegisterDeleteEnvironmentTool(server, client)

	// Instance
	RegisterGetInstanceTool(server, client)

	// Intercept
	RegisterInterceptStatusTool(server, client)
	RegisterInterceptControlTool(server, client)
	RegisterListInterceptEntriesTool(server, client)
	RegisterForwardInterceptTool(server, client)
	RegisterDropInterceptTool(server, client)

	// Tamper (Match & Replace)
	RegisterListTamperRulesTool(server, client)
	RegisterCreateTamperRuleTool(server, client)
	RegisterUpdateTamperRuleTool(server, client)
	RegisterToggleTamperRuleTool(server, client)
	RegisterDeleteTamperRuleTool(server, client)

	// Filters
	RegisterListFiltersTool(server, client)
	RegisterCreateFilterTool(server, client)
	RegisterDeleteFilterTool(server, client)

	// Hosted Files
	RegisterListHostedFilesTool(server, client)

	// Tasks
	RegisterListTasksTool(server, client)
	RegisterCancelTaskTool(server, client)

	// Plugins
	RegisterListPluginsTool(server, client)
}
