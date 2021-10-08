package schema

import (
	"io"

	"github.com/ipld/go-ipld-prime/schema"
	schemadmt "github.com/ipld/go-ipld-prime/schema/dmt"
	schemadsl "github.com/ipld/go-ipld-prime/schema/dsl"

	"github.com/ipld/go-ipldtool/errors"
)

// DSLParse is just the `schemadsl.Parse` feature, but wrapped in error tagging.
//
// Errors:
//
//   - schema-dsl-parse-failed -- if the DSL document didn't parse.
//
func DSLParse(inputName string, input io.Reader) (*schemadmt.Schema, error) {
	dmt, err := schemadsl.Parse(inputName, input)
	if err != nil {
		return nil, &errors.Error{
			ErrCode_SchemaDSLParseFailed,
			err.Error(),
			nil,
			err,
		}
	}
	return dmt, nil
}

// DSLParse is just the `schemadmt.Compile` feature, but wrapped in error tagging.
//
// Errors:
//
//   - schema-compile-failed -- if the DSL document didn't parse.
//
func SchemaCompile(dmt *schemadmt.Schema) (*schema.TypeSystem, error) {
	var ts schema.TypeSystem
	ts.Init()
	if err := schemadmt.Compile(&ts, dmt); err != nil {
		return nil, &errors.Error{
			ErrCode_SchemaCompileFailed,
			err.Error(),
			nil,
			err,
		}
	}
	return &ts, nil
}
