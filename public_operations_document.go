package mosir_sdk_go

import _ "embed"

//go:embed public.operations.graphql
var publicOperationsDocumentLiteral string

var publicOperationsDocument = publicOperationsDocumentLiteral
