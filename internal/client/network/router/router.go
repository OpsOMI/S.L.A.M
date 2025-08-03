package router

import "github.com/OpsOMI/S.L.A.M/internal/client/network/parser"

type Router struct {
}

func NewRouter() Router {
	return Router{}
}

func (r *Router) Route(
	command parser.Command,
) error {

	return nil
}
