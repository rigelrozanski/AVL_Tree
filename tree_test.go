package AVL_Tree

import (
	"bytes"
	"testing"
)

func TestAVLTree(t *testing.T) {

	//The AVLTree to be tested with
	tr := NewAVLTree()

	/////////////////////////////////////
	// Core sub-test function variables
	/////////////////////////////////////

	printErr := func(err error) {
		if err != nil {
			t.Errorf(err.Error())
		}
	}

	//Sub-test function used from heightBalanceNodeTest & heightBalanceTrunkTest
	//Test the height and balance against expected values for a node
	heightBalanceSubTest := func(n *node, expdHeight, expdBalance int) {

		//Test balance
		bal := n.getBalance()
		if bal != expdBalance {
			t.Errorf("bad balance for %v, expected %v found %v ",
				string(n.key[:]), expdBalance, bal)
			t.Log(n.outputStructure())
		}

		//Test height
		height := n.height
		if height != expdHeight {
			t.Errorf("bad height for %v, expected %v found %v ",
				string(n.key[:]), expdHeight, height)
			t.Log(n.outputStructure())
		}
	}

	//Test the height and balance against expected values for a node retrieved from key
	heightBalanceNodeTest := func(nodeKey string, expdHeight, expdBalance int) {
		heightBalanceSubTest(tr.trunk.findNode([]byte(nodeKey)), expdHeight, expdBalance)
	}

	//Test the trunk key, height, and balance against expected values
	heightBalanceTrunkTest := func(expdTrunkKey string, expdHeight, expdBalance int) {

		//test trunk node key
		expdTrunkKeyByte := []byte(expdTrunkKey)
		key := tr.trunk.key
		if bytes.Compare(key, expdTrunkKeyByte) != 0 {
			t.Errorf("bad trunk key expected %v found %v ",
				expdTrunkKeyByte, string(key[:]))
			t.Log(tr.trunk.outputStructure())
		}

		heightBalanceSubTest(tr.trunk, expdHeight, expdBalance)
	}

	//Test retrieving a value for a key.
	// The parameter expectedExists specifies whether or not
	// it is expected that the key will be found
	// (used for testing bad retrievals)
	retrieveTest := func(key, expectedVal string, expectedExists bool) {
		recievedVal, err := tr.Get([]byte(key))
		if expectedExists {
			printErr(err)
			if bytes.Compare(recievedVal, []byte(expectedVal)) != 0 {
				t.Errorf("bad expected %v recieved %v ", expectedVal, string(recievedVal[:]))
				t.Log(tr.trunk.outputStructure())
			}
		} else {
			if err == nil {
				t.Errorf("expected to receive an error when attempting to retrieve non-existent value for key %v", key)
				t.Log(tr.trunk.outputStructure())
			}
		}
	}

	/////////////////////////////////////
	// TESTS
	/////////////////////////////////////

	//Test for expected structures when adding key-value pairs to the AVL Tree
	//Expected tr structures:
	// a  a      b      b       b          d          d
	//     \    / \    / \     / \        / \        /  \
	//      b  a   c  a   c   a   d      b   e      b    f
	//                     \     / \    / \   \    /\    /\
	//                      d   c   e  a   c   f  a  c  e  g

	printErr(tr.Add([]byte("a"), []byte("vA")))
	heightBalanceTrunkTest("a", 0, 0)

	printErr(tr.Add([]byte("b"), []byte("vB")))
	heightBalanceTrunkTest("a", 1, 1)
	heightBalanceNodeTest("b", 0, 0)

	printErr(tr.Add([]byte("c"), []byte("vC")))
	heightBalanceTrunkTest("b", 1, 0)
	heightBalanceNodeTest("a", 0, 0)
	heightBalanceNodeTest("c", 0, 0)

	printErr(tr.Add([]byte("d"), []byte("vD")))
	heightBalanceTrunkTest("b", 2, 1)
	heightBalanceNodeTest("a", 0, 0)
	heightBalanceNodeTest("c", 1, 1)
	heightBalanceNodeTest("d", 0, 0)

	printErr(tr.Add([]byte("e"), []byte("vE")))
	heightBalanceTrunkTest("b", 2, 1)
	heightBalanceNodeTest("a", 0, 0)
	heightBalanceNodeTest("d", 1, 0)
	heightBalanceNodeTest("c", 0, 0)
	heightBalanceNodeTest("e", 0, 0)

	printErr(tr.Add([]byte("f"), []byte("vF")))
	heightBalanceTrunkTest("d", 2, 0)
	heightBalanceNodeTest("b", 1, 0)
	heightBalanceNodeTest("e", 1, 1)
	heightBalanceNodeTest("a", 0, 0)
	heightBalanceNodeTest("c", 0, 0)
	heightBalanceNodeTest("f", 0, 0)

	printErr(tr.Add([]byte("g"), []byte("vG")))
	heightBalanceTrunkTest("d", 2, 0)
	heightBalanceNodeTest("b", 1, 0)
	heightBalanceNodeTest("f", 1, 0)
	heightBalanceNodeTest("a", 0, 0)
	heightBalanceNodeTest("c", 0, 0)
	heightBalanceNodeTest("e", 0, 0)
	heightBalanceNodeTest("g", 0, 0)

	t.Log(tr.trunk.outputStructure())

	//Test retrieving saved values
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

	//Test removing saved values from the tr.
	//  Simultaneously test whether the hashes are updated
	//Expected tr structures:
	//      d         d      e
	//     /  \      / \    / \
	//    b    f    c   f  c   f
	//    \    /       /
	//     c  e       e

	//starting hash before removal
	hash1 := tr.trunk.hash

	//Removal of leafs
	printErr(tr.Remove([]byte("a")))
	printErr(tr.Remove([]byte("g")))
	heightBalanceTrunkTest("d", 2, 0)
	heightBalanceNodeTest("b", 1, 1)
	heightBalanceNodeTest("f", 1, -1)
	heightBalanceNodeTest("c", 0, 0)
	heightBalanceNodeTest("e", 0, 0)

	//test to see if the hash has changed
	if bytes.Compare(tr.trunk.hash, hash1) == 0 {
		t.Errorf("expected hash to change after having deleted files")
	}

	//Removal of a branch with one child nodes
	printErr(tr.Remove([]byte("b")))
	heightBalanceTrunkTest("d", 2, 1)
	heightBalanceNodeTest("f", 1, -1)
	heightBalanceNodeTest("c", 0, 0)
	heightBalanceNodeTest("e", 0, 0)

	//Removal of a branch with one child nodes
	printErr(tr.Remove([]byte("d")))
	heightBalanceTrunkTest("e", 1, 0)
	heightBalanceNodeTest("c", 0, 0)
	heightBalanceNodeTest("f", 0, 0)

	//Test retrieval of old saved value
	retrieveTest("a", "cA", false)
	retrieveTest("b", "vB", false)
	retrieveTest("c", "vC", true)
	retrieveTest("d", "vD", false)
	retrieveTest("e", "vE", true)
	retrieveTest("f", "vF", true)
	retrieveTest("g", "vG", false)
}
