package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	"github.com/suifengtec/geopoint"
)

//User Db Model
type User struct {
	ID       uint64 `xorm:"pk 'id'" json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"` //only for test!!
	Geog     string `json:"geog"`     //xorm 以及gorm 目前尚未支持PostGIS
}

/*

psql -U postgres

CREATE USER geopoint_test_user WITH PASSWORD 'geopoint_test_user_passsword';
CREATE DATABASE geopoint_test_db;
GRANT ALL PRIVILEGES ON DATABASE geopoint_test_db to geopoint_test_user;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO geopoint_test_user;

\c geopoint_test_db

CREATE EXTENSION postgis;
psql -d geopoint_test_db < bk.sql

*/

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "geopoint_test_user"
	dbPassword = "geopoint_test_user_passsword"
	dbName     = "geopoint_test_db"
)

//X ...
var X *xorm.Engine

func init() {
	initDBEngine()
}

func initDBEngine() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	e, err := xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	e.ShowSQL(true)

	fmt.Println("connect postgresql success")
	X = e
	//return e
}

//GetUserCount ...
func GetUserCount() uint64 {

	user := new(User)
	total, err := X.Where("id >?", 1).Count(user)
	if err != nil {
		fmt.Println(err.Error())
	}

	return uint64(total)
}

//GetUserByID ...
func GetUserByID(id uint64) User {
	var user User
	X.Id(id).Get(&user)
	return user
}

//PointsInDistanceRange 距离某一点给定距离内的若干个点
func PointsInDistanceRange(p geopoint.GeoPoint, d int64) []User {

	list := make([]User, 0)
	qStr := p.GetPointsQueryStringWithIn(d)
	err := X.Where(qStr).Desc("id").Find(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

// AddNewUser ...
func AddNewUser(user *User) uint64 {

	userInDb := new(User)

	X.Where("phone=?", user.Phone).Get(userInDb)
	//使用该手机号码的用户已存在了
	if userInDb.ID != 0 {

		log.Println("使用该手机号码的用户已存在了?")
		log.Println(userInDb.ID)
		log.Println(user.Phone)
		return 0
	}

	/*
		INSERT INTO "users" ("province_agent_id","city_agent_id","county_agent_id","name","phone","geog") VALUES ($1,$2,$3,$4,$5,$6) []interface {}{41, 371, 2, "cccd", "13800138000", "SRID=4326;POINT(113.538639 34.826563)"}
	*/
	//sqlStr := "INSERT INTO \"users\" (\"province_agent_id\",\"city_agent_id\",\"county_agent_id\",\"name\",\"phone\",\"geog\") VALUES ($1,$2,$3,$4,$5,$6)"
	sqlStr := "INSERT INTO \"user\" (\"name\",\"phone\",\"geog\") VALUES (?,?,?)"
	res, err := X.Exec(sqlStr, user.Name, user.Phone, user.Geog)

	if err != nil {
		log.Println(err)
		return 0
	}
	//isInsertedInt64, err := engine.Insert(user)
	//lastInsertIDInt64, err := res.LastInsertId()
	isInsertedInt64, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0
	}
	if isInsertedInt64 == 0 {
		log.Println("isInsertedInt64==0?")
		return 0
	}
	X.Where("phone=?", user.Phone).Get(userInDb)
	return userInDb.ID

}

//GetRowIDByPhone ...
func GetRowIDByPhone(phone string) uint64 {

	user := new(User)

	X.Where("phone=?", phone).Get(user)
	return user.ID

}

//DeleteUserByPhone ...
func DeleteUserByPhone(phone string) bool {
	user := new(User)

	X.Where("phone=?", phone).Get(user)
	id := user.ID
	if id == 0 {
		log.Printf("try to delete not exists user with phone %s", phone)
		return false
	}
	rows, err := X.Delete(&user)
	if err != nil {
		log.Println(err)
		return false
	}
	if rows == 0 {
		return false
	}
	return true
}

func main() {
	//total row in DT
	{ //用户数量
		total := GetUserCount()
		fmt.Println(total)
	}

	{ //添加新用户
		point := geopoint.GeoPoint{Lng: 113.538639, Lat: 34.826563}
		u1 := &User{
			Name:  "ddddd",
			Phone: "13800138005",

			Geog: point.String(),
		}
		isInsertedInt64 := AddNewUser(u1)
		fmt.Println("isInsertedInt64")
		fmt.Println(isInsertedInt64)
	}

	{ //按照ID获取用户
		var p2 geopoint.GeoPoint
		user := GetUserByID(2)
		fmt.Println(user)

		//p2.Scan([]uint8(user.Geog))
		p2.Scan(user.Geog)
		fmt.Println(p2.String())

	}
	{ //按照phone获取用户
		theID := GetRowIDByPhone("13800138000")

		fmt.Println(theID)
	}
	{ //查找某个点周围指定距离内的用户
		//113.739873, 34.356696
		p3 := geopoint.GeoPoint{Lng: 113.739873, Lat: 34.356696}
		ps := PointsInDistanceRange(p3, 5000)
		psLen := len(ps)
		if psLen > 0 {

			for i := 0; i < psLen; i++ {
				var p geopoint.GeoPoint
				p.Scan(ps[i].Geog)
				ps[i].Geog = p.JSONString()
				m, err := json.Marshal(ps[i])
				if err != nil {
					fmt.Println("error!!")
					continue
				}
				fmt.Println(string(m))
			}
		}

	}
}
