package word2vec

const (
	MaxString = 100
)

type VocabWord struct {
	Word        string
	Count       uint64
	Point, Code []int
	CodeLen     int
}
