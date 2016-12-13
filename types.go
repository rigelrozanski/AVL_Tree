package tree

import (
	"errors"
	//"path"

	cmn "github.com/rigelrozanski/passwerk/common"

	dbm "github.com/tendermint/go-db"
	"github.com/tendermint/go-merkle"
)

type Tree interface {
	Get(key []byte) (value []byte, err error)
	Add(key []byte, value []byte) error
	Update(key []byte, value []byte) error
	Remove(key []byte) error
	Hash() (hash []byte)
}
