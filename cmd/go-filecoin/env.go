package commands

import (
	"context"

	cmds "github.com/ipfs/go-ipfs-cmds"

	"github.com/sbwtw/go-filecoin/internal/app/go-filecoin/porcelain"
	"github.com/sbwtw/go-filecoin/internal/pkg/protocol/drand"
	"github.com/sbwtw/go-filecoin/internal/pkg/protocol/mining"
	"github.com/sbwtw/go-filecoin/internal/pkg/protocol/retrieval"
	"github.com/sbwtw/go-filecoin/internal/pkg/protocol/storage"
)

// Env is the environment for command API handlers.
type Env struct {
	blockMiningAPI *mining.API
	ctx            context.Context
	drandAPI       *drand.API
	porcelainAPI   *porcelain.API
	retrievalAPI   retrieval.API
	storageAPI     *storage.API
	inspectorAPI   *Inspector
}

var _ cmds.Environment = (*Env)(nil)

// NewClientEnv returns a new environment for command API clients.
// This environment lacks direct access to any internal APIs.
func NewClientEnv(ctx context.Context) *Env {
	return &Env{ctx: ctx}
}

// Context returns the context of the environment.
func (ce *Env) Context() context.Context {
	return ce.ctx
}

// GetPorcelainAPI returns the porcelain.API interface from the environment.
func GetPorcelainAPI(env cmds.Environment) *porcelain.API {
	ce := env.(*Env)
	return ce.porcelainAPI
}

// GetBlockAPI returns the block protocol api from the given environment.
func GetBlockAPI(env cmds.Environment) *mining.API {
	ce := env.(*Env)
	return ce.blockMiningAPI
}

// GetRetrievalAPI returns the retrieval protocol api from the given environment.
func GetRetrievalAPI(env cmds.Environment) retrieval.API {
	ce := env.(*Env)
	return ce.retrievalAPI
}

// GetStorageAPI returns the storage protocol api from the given environment.
func GetStorageAPI(env cmds.Environment) *storage.API {
	ce := env.(*Env)
	return ce.storageAPI
}

// GetInspectorAPI returns the inspector api from the given environment.
func GetInspectorAPI(env cmds.Environment) *Inspector {
	ce := env.(*Env)
	return ce.inspectorAPI
}

// GetDrandAPI returns the drand api from the given environment.
func GetDrandAPI(env cmds.Environment) *drand.API {
	ce := env.(*Env)
	return ce.drandAPI
}
