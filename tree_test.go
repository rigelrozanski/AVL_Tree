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

func TestAVLTree(t *testing.T) {

	printErr := func(err error) {
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	//First create the AVL tr to be tested with
	tr := NewAVLTree()

	//Test adding several values to the AVL Tree
	//  at the same time test how the heights of the tr react

	heightTest := func(expectedHeight int) {
		height := tr.trunk.Height
		if height != expectedHeight {
			t.Errorf("bad height for %v, expected %v found %v ",
				string(tr.trunk.Key[:]), expectedHeight, height)
			t.Log(tr.trunk.printStructure())
		}
	}

	//Expected tr structures:
	// a   a   b      b       b          d          d
	//    /   / \    / \     / \        / \        /  \
	//   b   a   c  a   c   a   d      b   e      b    f
	//                   \     / \    / \   \    /\    /\
	//                    d   c   e  a   c   f  a  c  e  g

	printErr(tr.Add([]byte("a"), []byte("vA")))
	heightTest(0)
	printErr(tr.Add([]byte("b"), []byte("vB")))
	heightTest(1)
	printErr(tr.Add([]byte("c"), []byte("vC")))
	heightTest(1)
	printErr(tr.Add([]byte("d"), []byte("vD")))
	heightTest(2)
	printErr(tr.Add([]byte("e"), []byte("vE")))
	heightTest(2)
	printErr(tr.Add([]byte("f"), []byte("vF")))
	heightTest(2)
	printErr(tr.Add([]byte("g"), []byte("vG")))
	heightTest(2)

	t.Log(tr.trunk.printStructure())

	//Test retrieving saved values
	retrieveTest := func(key, expectedVal string, expectedExists bool) {
		recievedVal, err := tr.Get([]byte(key))
		if expectedExists {
			printErr(err)
			if bytes.Compare(recievedVal, []byte(expectedVal)) != 0 {
				t.Errorf("bad expected %v recieved %v ", expectedVal, string(recievedVal[:]))
				t.Log(tr.trunk.printStructure())
			}
		} else {
			if err == nil {
				t.Errorf("expected to receive an error when attempting to retrieve non-existent value for key %v", key)
				t.Log(tr.trunk.printStructure())
			}
		}
	}

	retrieveTest("a", "vA", true)
	retrieveTest("b", "vB", true)
	retrieveTest("c", "vC", true)
	retrieveTest("d", "vD", true)
	retrieveTest("e", "vE", true)
	retrieveTest("f", "vF", true)
	retrieveTest("g", "vG", true)

	//Test adding a duplicate value
	expErr := tr.Add([]byte("a"), []byte("vA"))
	if expErr == nil {
		t.Errorf("expected to receive an error when attempting to add duplicate values")
	}

	//Test updating an existing value
	printErr(tr.Update([]byte("a"), []byte("vAA")))
	retrieveTest("a", "vAA", true)

	//Test updating a non existent value
	expErr = tr.Update([]byte("z"), []byte("vZ"))
	if expErr == nil {
		t.Errorf("expected to receive an error when attempting to update non-existent value")
	}

	//Test removing saved values from the tr

	//Expected tr structures:
	//      d         d      e
	//     /  \      / \    / \
	//    b    f    c   f  c   f
	//    \    /\      /
	//     c  e  g    e

	//removal of leafs
	printErr(tr.Remove([]byte("a")))
	printErr(tr.Remove([]byte("g")))
	heightTest(2)
	t.Log(tr.trunk.printStructure())

	//removal of a branch with one child nodes
	printErr(tr.Remove([]byte("b")))
	heightTest(2)
	t.Log(tr.trunk.printStructure())

	//removal of a branch with one child nodes
	printErr(tr.Remove([]byte("d")))
	heightTest(1)
	t.Log(tr.trunk.printStructure())

	//Test retrieval of old saved value
	retrieveTest("a", "cA", false)
	retrieveTest("b", "vB", false)
	retrieveTest("c", "vC", true)
	retrieveTest("d", "vD", false)
	retrieveTest("e", "vE", true)
	retrieveTest("f", "vF", true)
	retrieveTest("g", "vG", false)
}
