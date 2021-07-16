package issueancerpc

import "github.com/pkt-cash/pktd/lnd/macaroons"

// Config is the main configuration file for the router RPC server. It contains
// all the items required for the router RPC server to carry out its duties.
// The fields with struct tags are meant to be parsed as normal configuration
// options, while if able to be populated, the latter fields MUST also be
// specified.
type Config struct {

	// IssuanceMacPath is the path for the issuance macaroon. If unspecified
	// then we assume that the macaroon will be found under the network
	// directory, named DefaultIssuanceMacFilename.
	IssuanceMacPath string `long:"issuancemacaroonpath" description:"Path to the issuance macaroon"`

	// NetworkDir is the main network directory wherein the router rpc
	// server will find the macaroon named DefaultRouterMacFilename.
	NetworkDir string

	// MacService is the main macaroon service that we'll use to handle
	// authentication for the Router rpc server.
	MacService *macaroons.Service
}
