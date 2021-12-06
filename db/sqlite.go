package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func UpdateLiquility(liquility string, address, contractAddress string) error {
	db, err := sql.Open("sqlite3", "./lq.db")
	checkErr(err)
	// update
	stmt, err := db.Prepare("update borrow set liquility=? where address=? and contractAddress=?")
	checkErr(err)

	res, err := stmt.Exec(liquility, address, contractAddress)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)

	//fmt.Println(affect)
	return err
}

//select address,contractAddress,liquility from borrow  group by address having liquility+0 >0 and liquility+0=min(liquility+0)  order by liquility+0 asc  limit 0,100
func GetRows(start int) []Borrow {
	db, err := sql.Open("sqlite3", "./lq.db")
	checkErr(err)
	// query
	sql := fmt.Sprintf("select address,contractAddress  from borrow  group by address having liquility+0 >0 and liquility+0=min(liquility+0)  order by liquility+0 asc  limit %d,100", start)
	rows, err := db.Query(sql)
	if err!=nil{
		fmt.Println("db.Query:",err)
	}

	//checkErr(err)

	var address, contractAddress string

	list := make([]Borrow, 0)
	for rows.Next() {
		err = rows.Scan(&address, &contractAddress)
		checkErr(err)
		br := Borrow{
			Address: address,
			//AccountBorrows:  accountBorrows,
			//TotalBorrows:    totalBorrows,
			ContractAddress: contractAddress,
			//Liquility:       liquility,
		}
		list = append(list, br)
	}

	rows.Close()
	return list
}

func AddData(address, contractAddress, txHash string, accountBorrows, totalBorrows, liquity string, bockHeight uint64) error {
	//打开数据库，如果不存在，则创建
	db, err := sql.Open("sqlite3", "./lq.db")
	if checkErr(err) {
		return err
	}
	//	const UINT_MAX = ^uint32(0)

	// insert
	stmt, err := db.Prepare("INSERT INTO borrow(address, accountBorrows, totalBorrows,contractAddress,liquility,blockHeight,txHash, create_at) values(?,?,?,?,?,?,?,?)")
	if checkErr(err) {
		return err
	}

	res, err := stmt.Exec(address, accountBorrows, totalBorrows, contractAddress, liquity, bockHeight, txHash, time.Now())
	if checkErr(err) {
		return err
	}

	_, err = res.LastInsertId()
	if checkErr(err) {
		return err
	}

	//fmt.Println(id)
	return err

}

func checkErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
