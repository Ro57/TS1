package replicatorrpc

import (
	"github.com/pkt-cash/pktd/btcutil/er"
	"github.com/pkt-cash/pktd/lnd/lnrpc"
)

type ServerDriver struct {
	Name string
	New  func(subCfgs lnrpc.SubServerConfigDispatcher) (*Server, lnrpc.MacaroonPerms, er.R)
}

// createNewSubServer is a helper method that will create the new router sub
// server given the main config dispatcher method. If we're unable to find the
// config that is meant for us in the config dispatcher, then we'll exit with
// an error.
func createNewSubServer(configRegistry lnrpc.SubServerConfigDispatcher) (
	*Server, lnrpc.MacaroonPerms, er.R) {

	// We'll attempt to look up the config that we expect, according to our
	// subServerName name. If we can't find this, then we'll exit with an
	// error, as we're unable to properly initialize ourselves without this
	// config.
	routeServerConf, ok := configRegistry.FetchConfig(subServerName)
	if !ok {
		return nil, nil, er.Errorf("unable to find config for "+
			"subserver type %s", subServerName)
	}

	// Now that we've found an object mapping to our service name, we'll
	// ensure that it's the type we need.
	config, ok := routeServerConf.(*Config)
	if !ok {
		return nil, nil, er.Errorf("wrong type of config for "+
			"subserver %s, expected %T got %T", subServerName,
			&Config{}, routeServerConf)
	}

	// Before we try to make the new router service instance, we'll perform
	// some sanity checks on the arguments to ensure that they're useable.
	switch {
	// TODO: We need to add validation cases to validate the
	// configuration data and throw an error if it is incorrect.
	}

	return New(config)
}

func init() {
	/*
		subServer := &lnrpc.SubServerDriver{
			subServerName: subServerName,
			New: func(c lnrpc.SubServerConfigDispatcher) (lnrpc.SubServer, lnrpc.MacaroonPerms, er.R) {
				return createNewSubServer(c)
			},
		}

		// If the build tag is active, then we'll register ourselves as a
		// sub-RPC server within the global lnrpc package namespace.
		f err := lnrpc.RegisterSubServer(subServer); err != nil {
			panic(fmt.Sprintf("failed to register sub server driver '%s': %v",
				subServerName, err))
		}*/
}
