package blob

import (
	"fmt"
	// for schema embed
	_ "embed"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/schema"
)

//go:embed blob.ipldsch
var blobSchema []byte

var blobTS *schema.TypeSystem

func init() {
	ts, err := ipld.LoadSchemaBytes(blobSchema)
	if err != nil {
		panic(fmt.Errorf("loading blob schema: %w", err))
	}
	blobTS = ts
}

func AllocateCaveatsType() schema.Type {
	return blobTS.TypeByName("AllocateCaveats")
}

func AllocateOkType() schema.Type {
	return blobTS.TypeByName("AllocateOk")
}

func AcceptCaveatsType() schema.Type {
	return blobTS.TypeByName("AcceptCaveats")
}

func AcceptOkType() schema.Type {
	return blobTS.TypeByName("AcceptOk")
}