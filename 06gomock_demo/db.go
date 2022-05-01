package gomock_demo

/*
@author RandySun
@create 2022-05-01-15:48
*/

//
// DB
//  @Description: 数据接口
//
type DB interface {
	Get(key string) (int, error)
	Add(key string, value int) error
}

func GetFromDB(db DB, key string) int {
	if v, err := db.Get(key); err == nil {
		return v
	}
	return -1

}
