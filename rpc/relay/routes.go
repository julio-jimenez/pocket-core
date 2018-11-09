// This package contains files for the Relay API
package relay

import (
	"github.com/julienschmidt/httprouter"
	"github.com/pocket_network/pocket-core/rpc/relay/handlers"
	"github.com/pocket_network/pocket-core/rpc/shared"
)

/*
"relayRoutes" is a function that returns all of the routes of the API.
 */

/*
The "Route" structure defines the generalization of an api route.
*/
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

/*
"Routes" is a slice that holds all of the routes within one structure.
 */
type Routes []Route

func RelayRoutes() shared.Routes {
	routes := shared.Routes{
		shared.Route{"GetRelayAPIVersion", "GET", "/v1", handlers.GetRelayAPIVersion},
		shared.Route{"DispatchOptions", "GET", "/v1/dispatch", handlers.DispatchOptions},
		shared.Route{"DispatchServe", "GET", "/v1/dispatch/serve", handlers.DispatchServe},
		shared.Route{"RelayOptions", "GET", "/v1/relay", handlers.RelayOptions},
		shared.Route{"RelayRead", "GET", "/v1/relay/read", handlers.RelayRead},
		shared.Route{"RelayWrite", "GET", "/v1/relay/write", handlers.RelayWrite},
	}
	return routes
}