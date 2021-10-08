package schema

import (
	"github.com/ipld/go-ipldtool/errors"
)

const (
	ErrCode_SchemaDSLParseFailed = "schema-dsl-parse-failed"
	ErrCode_SchemaParseFailed    = "scheam-parse-failed"
	ErrCode_SchemaCompileFailed  = "schema-compile-failed"
)

// ErrSchemaDSLParseFailed is an error constructor function.
//
// Errors:
//
//   - schema-dsl-parse-failed -- just that.
func ErrSchemaDSLParseFailed(cause error) error {
	return &errors.Error{
		ErrCode_SchemaDSLParseFailed,
		cause.Error(),
		nil,
		cause,
	}
}

// ErrSchemaDSLParseFailed is an error constructor function.
//
// Errors:
//
//   - schema-compile-failed -- just that.
func ErrSchemaCompileFailed(cause error) error {
	return &errors.Error{
		ErrCode_SchemaCompileFailed,
		cause.Error(),
		nil,
		cause,
	}
}
