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
	//  at the same time test how the heights of the tree react
	heightTest := func(expectedHeight int) {
		height := tree.trunk.Height
		if height != expectedHeight {
			t.Errorf("bad height, expected %v found %v ", expectedHeight, height)
		}
	}

	printErr(tree.Add([]byte("keyOne"), []byte("valueOne")))
	heightTest(0)
	printErr(tree.Add([]byte("keyTwo"), []byte("valueTwo")))
	heightTest(1)
	printErr(tree.Add([]byte("keyThree"), []byte("valueThree")))
	heightTest(1)
	printErr(tree.Add([]byte("keyFour"), []byte("valueFour")))
	heightTest(2)

	//Test retrieving saved values
	retrieveTest := func(key, expectedVal string) {
		recievedVal, err := tree.Get([]byte(key))
		printErr(err)
		if bytes.Compare(recievedVal, []byte(expectedVal)) != 0 {
			t.Errorf("bad expected %v recieved %v ", expectedVal, string(recievedVal[:]))
		}
	}

	retrieveTest("keyOne", "valueOne")
	retrieveTest("keyTwo", "valueTwo")
	retrieveTest("keyThree", "valueThree")
	retrieveTest("keyFour", "valueFour")

	//Test adding a duplicate value

	//Test updating an existing value

	//Test updating a non existent value

	//Test removing saved values from the tree

	//Test retrieval of saved values
}
