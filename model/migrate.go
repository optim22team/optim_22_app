// https://gorm.io/docs/connecting_to_the_database.html#Clickhouse に詳細が記載されている。

package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"typefile"
)

// 外部でdb操作をするためのパッケージ変数
var Db *gorm.DB

func InitDB() {
	var err error
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name に詳細が記載されている。
	// DSN(データソース名)の作成。
	// 開発用のデータベース名はoptim_dev,テスト用のデータベース名はotpim_testである。
	dsn := "root:rootpass@tcp(mysql_container:3306)/optim_dev?parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("database successfully configure")
	}

	// 接続したdbをパッケージ変数Dbに代入している。
	Db = db
}

// Insert
// db.Create(&request)

// Select
// db.Find(&request, "id = ?", 10)

// Batch Insert
// var requests = []User{request1, request2, request3}
// db.Create(&requests)


// テストを実行するために前もって必要なデータを作成する。
func CreateTestData() {
	var users = []typefile.User{
		{Name: "user1"},
		{Name: "user2"},
		{Name: "user3"}}
	Db.Create(&users)

	// userが作成された直後にengineerも作成する。
	var engineers = []typefile.Engineer{
		{User: typefile.User{ID: 1,Name: "user1"}},
		{User: typefile.User{ID: 2,Name: "user2"}},
		{User: typefile.User{ID: 3,Name: "user3"}}}
	Db.Create(&engineers)

	//userが作成された直後にclientも作成する。
	var clients = []typefile.Client{
		{User: typefile.User{ID: 1,Name: "user1"}},
		{User: typefile.User{ID: 2,Name: "user2"}},
		{User: typefile.User{ID: 3,Name: "user3"}}}
	Db.Create(&clients)

	var requests = []typefile.Request{
		{ClientID: 1,RequestName: "request1 from clientID 1",Content: "request1 content",Finish: false},
		{ClientID: 1,RequestName: "request2 from clientID 1",Content: "request2 content",Finish: false},
		{ClientID: 2,RequestName: "request3 from clientID 2",Content: "request3 content",Finish: false}}
	Db.Create(&requests)
	
	var winners = []typefile.Winner{
		{EngineerID: 1,RequestID: 1},
		{EngineerID: 2,RequestID: 2}}
	Db.Create(&winners)

	var submissions = []typefile.Submission{
		{RequestID: 3,EngineerID: 1,Content: "submission1 of engineerID 1"},
		{RequestID: 3,EngineerID: 2,Content: "submission1 of engineerID 2"},
		{RequestID: 3,EngineerID: 3,Content: "submission1 of engineerID 3"}}
	Db.Create(&submissions)

	// id=1のClient構造体データを格納するためのインスタンスを生成
	client1 := typefile.Client{}
	// id=1を持つclientを抽出する。
	Db.Find(&client1,"id = ?",1)
	// SELECT * FROM `clients` WHERE id = 1

	// id=2のClient構造体データを格納するためのインスタンスを生成
	client2 := typefile.Client{}
	// id=2を持つclientを抽出する。
	Db.Find(&client2,"id = ?",2)
	// SELECT * FROM `clients` WHERE id = 2

	var clients_association = []typefile.Client{
		client1,
		client1,
		client2}
	Db.Model(&requests).Association("Client").Append(&clients_association)

	// id=3のRequest構造体データを格納するためのインスタンスを生成
	request3 := typefile.Request{}
	// id=3を持つrequestを抽出する。
	Db.Find(&request3,"id = ?",3)
	// SELECT * FROM `requests` WHERE id = 3

	var requests_association = []typefile.Request{
		request3,
		request3,
		request3}

	Db.Model(&submissions).Association("Request").Append(&requests_association)

	// id=1のEngineer構造体データを格納するためのインスタンスを生成
	engineer1 := typefile.Engineer{}
	// id=1を持つengineerを抽出する。
	Db.Find(&engineer1,"id = ?",1)
	// SELECT * FROM `engineers` WHERE id = 1

	// id=2のEngineer構造体データを格納するためのインスタンスを生成
	engineer2 := typefile.Engineer{}
	// id=2を持つengineerを抽出する。
	Db.Find(&engineer2,"id = ?",2)
	// SELECT * FROM `engineers` WHERE id = 2

	// id=3のEngineer構造体データを格納するためのインスタンスを生成
	engineer3 := typefile.Engineer{}
	// id=3を持つengineerを抽出する。
	Db.Find(&engineer3,"id = ?",3)
	// SELECT * FROM `engineers` WHERE id = 3

	var engineers_association = []typefile.Engineer{
		engineer1,
		engineer2,
		engineer3}

	Db.Model(&engineers).Association("Engineer").Append(&engineers_association)
}