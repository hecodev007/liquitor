package db

import (
	"fmt"
	"testing"
)

func TestSqlite(t *testing.T)  {
	//const UINT_MAX = ^uint64(0)
   //for i:=0;i<20;i++{
	//   AddData("11111","1112222","932356024711512064","111",fmt.Sprintf("%d",i))
   //}


   row:=GetRows(1)
   for _,r:=range row{
   	fmt.Println(r)
   }

   fmt.Println("finished!")
}
