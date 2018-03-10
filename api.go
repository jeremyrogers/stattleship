package baseballSim

// StattleshipAPI interfaces with the API
type StattleshipAPI struct {
	APIToken string
}

// APIInit initialize the API with your token
func APIInit(token string) *StattleshipAPI {
	rv := new(StattleshipAPI)
	rv.APIToken = token
	return rv
}
