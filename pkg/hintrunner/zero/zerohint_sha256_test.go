package zero

import (
	"testing"

	"github.com/NethermindEth/cairo-vm-go/pkg/hintrunner/hinter"
	"github.com/consensys/gnark-crypto/ecc/stark-curve/fp"
)

func TestZeroHintSha256(t *testing.T) {
	runHinterTests(t, map[string][]hintTestCase{
		"PackedSha256": {
			{
				operanders: []*hintOperander{
					{Name: "sha256_start", Kind: apRelative, Value: addr(6)},
					{Name: "output", Kind: apRelative, Value: addr(22)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
				},
				makeHinter: func(ctx *hintTestContext) hinter.Hinter {
					return newPackedSha256Hint(ctx.operanders["sha256_start"], ctx.operanders["output"])
				},
				check: consecutiveVarAddrResolvedValueEquals(
					"output",
					[]*fp.Element{
						feltString("3663108286"),
						feltString("398046313"),
						feltString("1647531929"),
						feltString("2006957770"),
						feltString("2363872401"),
						feltString("3235013187"),
						feltString("3137272298"),
						feltString("406301144"),
					}),
			},
		},
		"Sha256Chunk": {
			{
				operanders: []*hintOperander{
					{Name: "sha256_start", Kind: apRelative, Value: addr(6)},
					{Name: "output", Kind: apRelative, Value: addr(40)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "input", Kind: apRelative, Value: feltUint64(0)},
					{Name: "state", Kind: apRelative, Value: addr(23)},
					{Name: "-", Kind: apRelative, Value: feltUint64(0)},
					{Name: "-", Kind: apRelative, Value: feltUint64(0)},
					{Name: "-", Kind: apRelative, Value: feltUint64(0)},
					{Name: "-", Kind: apRelative, Value: feltUint64(0)},
					{Name: "-", Kind: apRelative, Value: feltUint64(0)},
					{Name: "-", Kind: apRelative, Value: feltUint64(0)},
					{Name: "-", Kind: apRelative, Value: feltUint64(0)},
					{Name: "-", Kind: apRelative, Value: feltUint64(0)},
				},
				makeHinter: func(ctx *hintTestContext) hinter.Hinter {
					return newSha256ChunkHint(ctx.operanders["sha256_start"], ctx.operanders["state"], ctx.operanders["output"])
				},
				check: consecutiveVarAddrResolvedValueEquals(
					"output",
					[]*fp.Element{
						feltString("2091193876"),
						feltString("1113340840"),
						feltString("3461668143"),
						feltString("3254913767"),
						feltString("3068490961"),
						feltString("2551409935"),
						feltString("2927503052"),
						feltString("3205228454"),
					}),
			},
		},
		"FinalizeSha256": {
			{
				operanders: []*hintOperander{
					{Name: "sha256_ptr_end", Kind: apRelative, Value: addr(5)},
				},
				makeHinter: func(ctx *hintTestContext) hinter.Hinter {
					return newFinalizeSha256Hint(ctx.operanders["sha256_ptr_end"])
				},
				check: consecutiveVarAddrResolvedValueEquals(
					"sha256_ptr_end",
					[]*fp.Element{
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("1779033703"),
						feltString("3144134277"),
						feltString("1013904242"),
						feltString("2773480762"),
						feltString("1359893119"),
						feltString("2600822924"),
						feltString("528734635"),
						feltString("1541459225"),
						feltString("3663108286"),
						feltString("398046313"),
						feltString("1647531929"),
						feltString("2006957770"),
						feltString("2363872401"),
						feltString("3235013187"),
						feltString("3137272298"),
						feltString("406301144"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("1779033703"),
						feltString("3144134277"),
						feltString("1013904242"),
						feltString("2773480762"),
						feltString("1359893119"),
						feltString("2600822924"),
						feltString("528734635"),
						feltString("1541459225"),
						feltString("3663108286"),
						feltString("398046313"),
						feltString("1647531929"),
						feltString("2006957770"),
						feltString("2363872401"),
						feltString("3235013187"),
						feltString("3137272298"),
						feltString("406301144"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("1779033703"),
						feltString("3144134277"),
						feltString("1013904242"),
						feltString("2773480762"),
						feltString("1359893119"),
						feltString("2600822924"),
						feltString("528734635"),
						feltString("1541459225"),
						feltString("3663108286"),
						feltString("398046313"),
						feltString("1647531929"),
						feltString("2006957770"),
						feltString("2363872401"),
						feltString("3235013187"),
						feltString("3137272298"),
						feltString("406301144"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("1779033703"),
						feltString("3144134277"),
						feltString("1013904242"),
						feltString("2773480762"),
						feltString("1359893119"),
						feltString("2600822924"),
						feltString("528734635"),
						feltString("1541459225"),
						feltString("3663108286"),
						feltString("398046313"),
						feltString("1647531929"),
						feltString("2006957770"),
						feltString("2363872401"),
						feltString("3235013187"),
						feltString("3137272298"),
						feltString("406301144"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("1779033703"),
						feltString("3144134277"),
						feltString("1013904242"),
						feltString("2773480762"),
						feltString("1359893119"),
						feltString("2600822924"),
						feltString("528734635"),
						feltString("1541459225"),
						feltString("3663108286"),
						feltString("398046313"),
						feltString("1647531929"),
						feltString("2006957770"),
						feltString("2363872401"),
						feltString("3235013187"),
						feltString("3137272298"),
						feltString("406301144"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("0"),
						feltString("1779033703"),
						feltString("3144134277"),
						feltString("1013904242"),
						feltString("2773480762"),
						feltString("1359893119"),
						feltString("2600822924"),
						feltString("528734635"),
						feltString("1541459225"),
						feltString("3663108286"),
						feltString("398046313"),
						feltString("1647531929"),
						feltString("2006957770"),
						feltString("2363872401"),
						feltString("3235013187"),
						feltString("3137272298"),
						feltString("406301144"),
					}),
			},
		},
	})
}
