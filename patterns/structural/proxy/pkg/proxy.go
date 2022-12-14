package pkg

import "errors"

type ProxyDataBase struct {
	Users map[string]bool
	Db    *DataBase
}

func (pDb ProxyDataBase) GetData(user string) ([]string, error) {
	if !pDb.Users[user] {
		return nil, errors.New("Недостаточно прав доступа к данным!")
	}
	return pDb.Db.GetData(user)
}
