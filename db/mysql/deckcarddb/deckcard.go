package deckcarddb

type Deckcard struct {
}

type IDeckCard interface {
}

var providers = map[string]IDeckCard{}

func init() {
	providers["mysql"] = Deckcard{}
}

func GetProvider(typeDb string) IDeckCard {
	return providers[typeDb]
}
