module github.com/ipld/go-ipldtool

go 1.16

// replace github.com/ipld/go-ipld-prime => ../go-ipld-prime
// replace github.com/ipld/go-ipld-prime/storage/dsadapter => ../go-ipld-prime/storage/dsadapter

require (
	github.com/frankban/quicktest v1.13.1
	github.com/ipfs/go-cid v0.1.0
	github.com/ipfs/go-ds-flatfs v0.4.5
	github.com/ipld/go-ipld-prime v0.12.4-0.20211010135705-522500cfab8b
	github.com/ipld/go-ipld-prime/storage/dsadapter v0.0.0-20211010135705-522500cfab8b
	github.com/urfave/cli/v2 v2.3.0
	github.com/warpfork/go-testmark v0.6.0
)
