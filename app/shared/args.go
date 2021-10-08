package shared

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ipld/go-ipld-prime/codec"
	"github.com/ipld/go-ipld-prime/codec/cbor"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
	"github.com/ipld/go-ipld-prime/codec/json"
	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/printer"
)

// StringIsPathish returns true if the string explicitly looks like a filesystem path
// (starts with `./`, `../`, or `/`).
func StringIsPathish(x string) bool {
	return strings.HasPrefix(x, "./") ||
		strings.HasPrefix(x, "../") ||
		strings.HasPrefix(x, "/")
}

// ParseDataSourceArg returns a reader for data based on the argument,
// and a Link if the argument was of that kind.
//
// Errors:
//
//   - ipldtool-error-invalid-args -- if the input arg can't be made into a readable stream.
//
func ParseDataSourceArg(inputArg string) (reader *bufio.Reader, link datamodel.Link, err error) {
	switch {
	case inputArg == "-": // stdin
		reader = bufio.NewReader(os.Stdin) // FIXME does this cli package not have a way to attach a stream so I don't have to use a global for this?
	case StringIsPathish(inputArg): // looks like a filename
		f, err := os.Open(inputArg)
		if err != nil {
			return nil, nil, ErrInvalidArgs("arg looks like a filename but cannot be opened", err)
		}
		reader = bufio.NewReader(f)
	default: // hope this is a CID
		panic("todo")
	}
	return
}

// ParseDataSourceArg returns an IPLD encoder based on the argument string.
// It handles strings of the form "codec:{name}", "codec:0x{code}",
// and the special string "debug".
//
// The argName parameter is used purely for error message formatting purposes.
//
// Errors:
//
//   - ipldtool-error-invalid-args -- if the input arg can't be made into a readable stream.
//
func ParseEncoderArg(arg string, defalt string, argName string) (codec.Encoder, error) {
	if arg == "" {
		arg = defalt
	}
	switch arg {
	case "debug":
		return func(n datamodel.Node, wr io.Writer) error {
			printer.Fprint(wr, n)
			return nil
		}, nil
	default:
		switch {
		case strings.HasPrefix(arg, "codec:0x"):
			panic("todo")
		case strings.HasPrefix(arg, "codec:"):
			// TODO: I thought there'd be a library somewhere with a lookup function from known names to codes.
			// TODO: should probably use the go-ipld-prime multicodec registry, so it's easy for someone to build extended version of the tool?
			switch arg[6:] {
			case "json":
				return json.Encode, nil
			case "dag-json":
				return dagjson.Encode, nil
			case "cbor":
				return cbor.Encode, nil
			case "dag-cbor":
				return dagcbor.Encode, nil
			default:
				return nil, ErrInvalidArgs(fmt.Sprintf("%s argument not recognized: %q is not a supported codec name", argName, arg[6:]), nil)
			}
		default:
			return nil, ErrInvalidArgs(fmt.Sprintf("%s argument format not recognized", argName), nil)
		}
	}
}
