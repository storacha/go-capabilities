package pdp

import (
	// for go:embed
	_ "embed"
	"fmt"

	"github.com/filecoin-project/go-data-segment/merkletree"
	ipldprime "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/datamodel"
	ipldschema "github.com/ipld/go-ipld-prime/schema"
	"github.com/storacha/go-capabilities/pkg/types"
	"github.com/storacha/go-piece/pkg/piece"
	"github.com/storacha/go-ucanto/core/ipld"
	"github.com/storacha/go-ucanto/core/schema"
	"github.com/storacha/go-ucanto/validator"
)

const PDPAcceptAbility = "pdp/accept"

//go:embed pdp.ipldsch
var pdpSchema []byte

var pdpTS *ipldschema.TypeSystem

func init() {
	ts, err := ipldprime.LoadSchemaBytes(pdpSchema)
	if err != nil {
		panic(fmt.Errorf("loading blob schema: %w", err))
	}
	pdpTS = ts
}

func PDPAcceptCaveatsType() ipldschema.Type {
	return pdpTS.TypeByName("PDPAcceptCaveats")
}

func PDPAcceptOkType() ipldschema.Type {
	return pdpTS.TypeByName("PDPAcceptOk")
}

func PDPInfoCaveatsType() ipldschema.Type {
	return pdpTS.TypeByName("PDPInfoCaveats")
}

func PDPInfoOkType() ipldschema.Type {
	return pdpTS.TypeByName("PDPInfoOk")
}

type PDPAcceptCaveats struct {
	Piece piece.PieceLink
}

func (pc PDPAcceptCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&pc, PDPAcceptCaveatsType(), types.Converters...)
}

var PDPAcceptCaveatsReader = schema.Struct[PDPAcceptCaveats](PDPAcceptCaveatsType(), nil, types.Converters...)

var PDPAccept = validator.NewCapability(
	PDPAcceptAbility,
	schema.DIDString(),
	PDPAcceptCaveatsReader,
	validator.DefaultDerives,
)

type PDPAcceptOk struct {
	Aggregate      piece.PieceLink
	InclusionProof merkletree.ProofData
	Piece          piece.PieceLink
}

func (po PDPAcceptOk) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&po, PDPAcceptOkType(), types.Converters...)
}

const PDPInfoAbility = "pdp/info"

type PDPInfoCaveats struct {
	Piece piece.PieceLink
}

func (pi PDPInfoCaveats) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&pi, PDPInfoCaveatsType(), types.Converters...)
}

var PDPInfoCaveatsReader = schema.Struct[PDPInfoCaveats](PDPInfoCaveatsType(), nil, types.Converters...)

var PDPInfo = validator.NewCapability(
	PDPInfoAbility,
	schema.DIDString(),
	PDPInfoCaveatsReader,
	validator.DefaultDerives,
)

type PDPInfoAcceptedAggregate struct {
	Aggregate      piece.PieceLink
	InclusionProof merkletree.ProofData
}

type PDPInfoOk struct {
	Piece      piece.PieceLink
	Aggregates []PDPInfoAcceptedAggregate
}

func (po PDPInfoOk) ToIPLD() (datamodel.Node, error) {
	return ipld.WrapWithRecovery(&po, PDPInfoOkType(), types.Converters...)
}
