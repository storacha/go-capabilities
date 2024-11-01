package blob

import (
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/storacha/go-capabilities/pkg/types"
	"github.com/storacha/go-ucanto/core/ipld"
	"github.com/storacha/go-ucanto/core/schema"
	"github.com/storacha/go-ucanto/did"
	"github.com/storacha/go-ucanto/ucan"
	"github.com/storacha/go-ucanto/validator"
)

const AcceptAbility = "blob/accept"

type Await struct {
	Selector string
	Link     ipld.Link
}

type Promise struct {
	UcanAwait Await
}

type AcceptCaveats struct {
	Space   did.DID
	Blob    Blob
	Expires uint64
	Put     Promise
}

func (ac AcceptCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ac, AcceptCaveatsType(), types.Converters...)
}

type AcceptOk struct {
	Site ucan.Link
}

func (ao AcceptOk) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&ao, AcceptOkType(), types.Converters...)
}

var AcceptCaveatsReader = schema.Struct[AcceptCaveats](AcceptCaveatsType(), nil, types.Converters...)
var Accept = validator.NewCapability(
	AcceptAbility,
	schema.DIDString(),
	AcceptCaveatsReader,
	validator.DefaultDerives,
)
