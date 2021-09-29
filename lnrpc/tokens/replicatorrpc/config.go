package replicatorrpc

import "github.com/pkt-cash/pktd/lnd/macaroons"

const tokenUrlPattern = "%s/v2/replicator/blocksequence/%s"

// Config is the main configuration file for the router RPC server. It contains
// all the items required for the router RPC server to carry out its duties.
// The fields with struct tags are meant to be parsed as normal configuration
// options, while if able to be populated, the latter fields MUST also be
// specified.
type Config struct {

	// ReplicatorMacPath is the path for the replicator macaroon. If unspecified
	// then we assume that the macaroon will be found under the network
	// directory, named DefaultReplicatorMacFilename.
	ReplicatorMacPath string `long:"replicatormacaroonpath" description:"Path to the replicator macaroon"`

	// NetworkDir is the main network directory wherein the router rpc
	// server will find the macaroon named DefaultRouterMacFilename.
	NetworkDir string

	// MacService is the main macaroon service that we'll use to handle
	// authentication for the Router rpc server.
	MacService *macaroons.Service
}
