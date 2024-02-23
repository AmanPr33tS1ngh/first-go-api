package tools

import ("time")

type mockDB struct{}

var mockLoginDetails map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC"
		Username: "Alex"
	}
}

var mockCoinDetails map[string]CoinDetails{
	"alex": {
		Coins: 12
		Username: "Alex"
	}
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails{
	time.Sleep(time.Second * 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok{
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails{
	time.Sleep(time.Second * 1)
	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok{
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error{
	return nil
}