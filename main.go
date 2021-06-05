package main
import(
	"fmt"
	"errors"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
func WerrWrap(n int)(string,error){
	//创建sql.db对象
	db,err:=sql.Open("mysql", "root:yuchunwang@tcp(192.168.203.112:3306)/test_api")
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()
	var (
		name string
	)
	err = db.QueryRow("select name from users where id = ?", n).Scan(&name)
	//fmt.Println(err)
	if err!=nil {
		fmt.Println("query error")
		if err == sql.ErrNoRows {
			//log.Fatal("sql.ErrNoRows")
			return " ", fmt.Errorf("%w,data is nil", err)
		}
		log.Fatal(err)
	}
	return "",err
}
func main() {
	_,err:=WerrWrap(4)
	err=errors.Unwrap(err)
	//fmt.Println(err)
	if errors.Is(err,sql.ErrNoRows){
		//log.Fatal("ErrNoRows")
		fmt.Println("It's ErrNoRows")
	}
}