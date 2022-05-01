package sqlmock_demo

import "database/sql"

/*
@author RandySun
@create 2022-05-01-10:43
*/

//
//  recordStats
//  @Description: 记录用户浏览产品信息 在`products`表中将当前商品的浏览次数+1 - 在`product_viewers`表中记录浏览当前商品的用户id
//  @param db 数据库连接对象
//  @param userId 用户id
//  @param productId 产品id
//  @return err
//
func recordStats(db *sql.DB, userId, productId int64) (err error) {
	// 开启事务
	// 操作views和product_viewers两张表
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:

			tx.Rollback()
		}
	}()

	// 更新products 表
	if _, err = tx.Exec("update products set views=views+1"); err != nil {
		return
	}

	// product_viewers 表中插入一条数据
	if _, err = tx.Exec("insert into product_viewers(user_id, product_id) values(?, ?)", userId, productId); err != nil {
		return
	}
	return

}

func main() {
	// 注意测试过程中并不需要真正的连接
	db, err := sql.Open("mysql", "root/randy.")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// userId 为1的用户浏览了productId为5的产品
	if err = recordStats(db, 1 /*some user id*/, 5 /*some product id*/); err != nil {
		panic(err)
	}

}
