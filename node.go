//      __      ___        _______ _____  ______ ______
//     /\ \    / / |      |__   __|  __ \|  ____|  ____|
//    /  \ \  / /| |         | |  | |__) | |__  | |__
//   / /\ \ \/ / | |         | |  |  _  /|  __| |  __|
//  / ____ \  /  | |____     | |  | | \ \| |____| |____
// /_/    \_\/   |______|    |_|  |_|  \_\______|______|
//

package AvlTree

import (
	"github.com/tendermint/go-db"
)

type AVLNode struct {
	key []byte
	value []byte
	height int
	balance int
}

func NewAVLTree(
	cacheSize int,
	db dbm.DB,
	dBName string) AVLTree {

	return AVLTree{
		trunk:     NewAVLNode,
		cacheSize: cacheSize,
		db:        db,
		dBName:    dBName,
	}
}

func Rotate(
