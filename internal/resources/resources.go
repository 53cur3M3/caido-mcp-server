package resources

import (
	caido "github.com/caido-community/sdk-go"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RegisterAll registers every resource on the server and returns the
// number of resources registered.
func RegisterAll(server *mcp.Server, client *caido.Client) int {
	registers := []func(*mcp.Server, *caido.Client){
		registerRequestResource,
		registerReplaySessionResource,
		registerSitemapResource,
		registerFindingsResource,
	}
	for _, register := range registers {
		register(server, client)
	}
	return len(registers)
}
