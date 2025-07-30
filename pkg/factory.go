package pkg

import (
	"database/sql"

	"github.com/OpsOMI/S.L.A.M/pkg/cronpkg"
	"github.com/OpsOMI/S.L.A.M/pkg/hasherpkg"
	"github.com/OpsOMI/S.L.A.M/pkg/txmanagerpkg"
)

type IPackages interface {
	Hasher() hasherpkg.IHasher
	TXManager() txmanagerpkg.ITXManager
	Cron() *cronpkg.Manager
}

type packages struct {
	hasher    hasherpkg.IHasher
	txManager txmanagerpkg.ITXManager
	cron      *cronpkg.Manager
}

func NewPackages(
	db *sql.DB,
) IPackages {
	return &packages{
		hasher:    hasherpkg.New(),
		txManager: txmanagerpkg.New(db),
		cron:      cronpkg.New(),
	}
}

func (p *packages) Hasher() hasherpkg.IHasher {
	return p.hasher
}

func (p *packages) TXManager() txmanagerpkg.ITXManager {
	return p.txManager
}

func (p *packages) Cron() *cronpkg.Manager {
	return p.cron
}
