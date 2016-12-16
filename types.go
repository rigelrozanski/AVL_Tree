package AVL_Tree

type Tree interface {
	GetHash() (hash []byte, err error)
	Get(key []byte) (value []byte, err error)
	Set(key, value []byte) error
	Add(key, value []byte) error
	Update(key, value []byte) error
	Remove(key []byte) error
	TreeStructure() string
}
