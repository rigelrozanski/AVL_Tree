package AVL_Tree

type Tree interface {
	Get(key []byte) (value []byte, err error)
	Add(key []byte, value []byte) error
	Update(key []byte, value []byte) error
	Remove(key []byte) error
}
