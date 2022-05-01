package varys

/*
@author RandySun
@create 2022-05-01-17:33
*/

type UserInfo struct {
	Name string
}

//
// GetInfoByUID
//  @Description: 获取用户uid
//  @param int64
//  @return *UserInfo
//  @return error
//
func GetInfoByUID(int64) (*UserInfo, error) {
	return &UserInfo{Name: "RandySun"}, nil

}
