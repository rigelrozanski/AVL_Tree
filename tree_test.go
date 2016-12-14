//      __      ___        _______ _____  ______ ______
//     /\ \    / / |      |__   __|  __ \|  ____|  ____|
//    /  \ \  / /| |         | |  | |__) | |__  | |__
//   / /\ \ \/ / | |         | |  |  _  /|  __| |  __|
//  / ____ \  /  | |____     | |  | | \ \| |____| |____
// /_/    \_\/   |______|    |_|  |_|  \_\______|______|
//

package AVL_Tree

import (
	"bytes"
	"testing"
)

func TestAVLTREE(t *testing.T) {

	printErr := func(err error) {
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	//First create the AVL tree to be tested with
	tree := NewAVLTree()

	//Test adding several values to the AVL Tree
	printErr(tree.Add([]byte("keyOne"), []byte("valueOne")))
	printErr(tree.Add([]byte("keyTwo"), []byte("valueTwo")))
	//printErr(tree.Add([]byte("keyThree"), []byte("valueThree")))
	//printErr(tree.Add([]byte("keyFour"), []byte("valueFour")))

	trunkNil := (tree.trunk == nil)
	t.Log(trunkNil)

	//Test retrieving saved values
	val1, err1 := tree.Get([]byte("keyOne"))
	printErr(err1)
	if bytes.Compare(val1, []byte("valueOne")) != 0 {
		t.Errorf("bad expected valueOne recieved " + string(val1[:]))
	}

	val2, err2 := tree.Get([]byte("keyTwo"))
	printErr(err2)
	if bytes.Compare(val2, []byte("valueTwo")) != 0 {
		t.Errorf("bad expected valueTwo recieved " + string(val2[:]))
	}

	//Test adding a duplicate value

	//Test updating an existing value

	//Test updating a non existent value

	//Test removing saved values from the tree

	//Test retrieval of saved values
}
