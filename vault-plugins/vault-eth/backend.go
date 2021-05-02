package vault_eth

import (
	"context"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
)

type backed struct {
	*framework.Backend
}

func Backend(c *logical.BackendConfig) *backend {
	var b backend

	b.Backend = &framework.Backend{
		BackendType: logical.TypeLogical,
		Secrets: []*framework.Secret{},
		Paths: framework.PathAppend(
			// Add paths here
		),
		PathsSpecial: &logical.Paths{
			// add path access req herre
		},
	}
	return b
}

func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	b := Backend(conf)

	if err := b.Setup(ctx, conf); err != nil {
		return nil, err
	}
	return b, nil
}
