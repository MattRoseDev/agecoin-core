package graph

import "github.com/favecode/agecoin-core/service"

//go:generate go run github.com/99designs/gqlgen
type Resolver struct{
  Service *service.Service
}
