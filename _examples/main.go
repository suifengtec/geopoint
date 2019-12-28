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
//xorm and gorm all not support PostGIS types now,so we prefer not sync tables used PostGIS types.
//so we do not use their datatable Sync feature.
type User struct {
	ID    uint64 `xorm:"bigint pk autoincr 'id'" json:"id"`
	Label string `xorm:"varchar(64)" json:"label"`
	Geog  string `xorm:"varchar(254)" json:"geog"`
}

/*

psql -U postgres

CREATE USER geopoint_test_user WITH PASSWORD 'geopoint_test_user_passsword';
CREATE DATABASE geopoint_test_db;
GRANT ALL PRIVILEGES ON DATABASE geopoint_test_db to geopoint_test_user;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO geopoint_test_user;

\c geopoint_test_db

CREATE EXTENSION postgis;
\q


psql -d geopoint_test_db < bk.sql

psql -U geopoint_test_user -d geopoint_test_db

\dt

//after tested,drop the test table.

psql -U postgres -d geopoint_test_db

DROP TABLE IF EXISTS public.user;
\q

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

func initDBEngine() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	e, err := xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)

	}
	e.ShowSQL(false)

	fmt.Println("Connected to postgresql !")
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

//GetUserPtrByID ...
func GetUserPtrByID(id uint64) *User {
	user := new(User)
	X.Id(id).Get(user)
	return user
}

//GetIDByLabel ...
func GetIDByLabel(label string) uint64 {
	user := User{Label: label}
	X.Where("label=?", label).Get(&user)

	log.Println("==GetIDByLabel==")
	log.Println(user)
	return user.ID

}

//GetUsersInDistanceRange ...
func GetUsersInDistanceRange(p geopoint.GeoPoint, d int64) []User {

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

	idIfExist := GetIDByLabel(user.Label)

	if idIfExist > 0 {
		log.Println("AddNewUser failed,caused by the label has be used.")
		return uint64(0)
	}
	isInsertedInt64, err := X.Insert(user)
	if err != nil {
		log.Println(err)
		return uint64(0)
	}
	if isInsertedInt64 == 0 {
		return uint64(0)
	}
	return user.ID

}

//DeleteUserByLabel ...
func DeleteUserByLabel(label string) bool {
	user := new(User)

	X.Where("label=?", label).Get(user)
	id := user.ID
	if id == 0 {
		log.Printf("try to delete not exists user with label %s", label)
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

	initDBEngine()
	//total row in DT
	{ //User Count
		fmt.Println("==User Count==")
		total := GetUserCount()
		fmt.Println(total)
	}

	{ //add a new User if not exist.
		fmt.Println("The ID of the new user be added or adding user failed info.")
		point := geopoint.GeoPoint{Lng: 113.538639, Lat: 34.826563}
		u1 := &User{
			Label: "13800138005",
			Geog:  point.String(),
		}
		isInsertedInt64 := AddNewUser(u1)

		fmt.Println(isInsertedInt64)
	}

	{ //fetch a user by ID
		fmt.Println("==Scan Geopoint data in db to raw string==")
		var p geopoint.GeoPoint
		user := GetUserByID(2)
		fmt.Println(user)
		p.Scan(user.Geog)
		fmt.Println(p.String())

	}
	{ //fetch user ID by label
		theID := GetIDByLabel("13800138000")
		fmt.Println("The ID of fetched user by label", theID)
	}
	{ //find some users in a given distance range.
		fmt.Println("==Users in 5000 meters by geopoint.GeoPoint{Lng: 113.739873, Lat: 34.356696}==")
		//113.739873, 34.356696
		p := geopoint.GeoPoint{Lng: 113.739873, Lat: 34.356696}
		ps := GetUsersInDistanceRange(p, 5000)
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
