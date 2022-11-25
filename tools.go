//go:build tools
// +build tools

package tools

// Necessário para baixar as dependencia para build
import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
