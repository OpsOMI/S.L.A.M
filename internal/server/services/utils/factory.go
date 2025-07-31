package utils

import "github.com/OpsOMI/S.L.A.M/internal/server/services/utils/parse"

type IUtilServices interface {
	Parse() parse.IParseService
}

type services struct {
	parse parse.IParseService
}

func NewServices() IUtilServices {
	parse := parse.NewService()

	return &services{
		parse: parse,
	}
}

// Parse returns the parse service.
func (u *services) Parse() parse.IParseService {
	return u.parse
}
