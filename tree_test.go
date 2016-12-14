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

	//some expected tree forms:
	// a   a   b      b       b       d      d
	//    /   / \    / \     / \     / \    / \
	//   b   a   c  a   c   a   d   b   e  c   e
	//                   \     / \   \
	//                    d   c   e   c

	printErr(tree.Add([]byte("a"), []byte("vA")))
	heightTest(0)
	printErr(tree.Add([]byte("b"), []byte("vB")))
	heightTest(1)
	printErr(tree.Add([]byte("c"), []byte("vC")))
	heightTest(1)
	printErr(tree.Add([]byte("d"), []byte("vD")))
	heightTest(2)
	printErr(tree.Add([]byte("e"), []byte("vE")))
	heightTest(2)

	//Test retrieving saved values
	retrieveTest := func(key, expectedVal string, expectedExists bool) {
		recievedVal, err := tree.Get([]byte(key))
		if expectedExists {
			printErr(err)
			if bytes.Compare(recievedVal, []byte(expectedVal)) != 0 {
				t.Errorf("bad expected %v recieved %v ", expectedVal, string(recievedVal[:]))
			}
		} else {
			if err == nil {
				t.Errorf("expected to receive an error when attempting to retrieve non-existent value for key ", key)
			}
		}
	}

	retrieveTest("a", "vA", true)
	retrieveTest("b", "vB", true)
	retrieveTest("c", "vC", true)
	retrieveTest("d", "vD", true)
	retrieveTest("e", "vE", true)

	//Test adding a duplicate value
	expErr := tree.Add([]byte("a"), []byte("vA"))
	if expErr == nil {
		t.Errorf("expected to receive an error when attempting to add duplicate values")
	}

	//Test updating an existing value
	printErr(tree.Update([]byte("a"), []byte("vAA")))
	retrieveTest("a", "vAA", true)

	//Test updating a non existent value
	expErr = tree.Update([]byte("z"), []byte("vZ"))
	if expErr == nil {
		t.Errorf("expected to receive an error when attempting to update non-existent value")
	}

	//Test removing saved values from the tree
	printErr(tree.Remove([]byte("a")))
	heightTest(2)
	printErr(tree.Remove([]byte("b")))
	heightTest(1)

	//Test bad retrieval of old saved value
	retrieveTest("a", "", false)
	retrieveTest("b", "", false)
	retrieveTest("c", "vC", true)
	retrieveTest("d", "vD", true)
	retrieveTest("e", "vE", true)
}
