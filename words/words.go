package words

type Api interface {
	GetRandom(amount int) []string
}
