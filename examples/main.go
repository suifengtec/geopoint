package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	geopoints "github.com/suifengtec/go-postgis-points"
)

//Users ...
type Users struct {
	Id              uint64 `xorm:"pk" json:"id"`
	ProvinceAgentId uint64 `json:"province_agent_id"`
	CityAgentId     uint64 `json:"city_agent_id"`
	CountyAgentId   uint64 `json:"county_agent_id"`
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	BizLevel        int8   `json:"biz_level"`
	Role            int8   `json:"role"`
	Grade           int8   `json:"grade"`
	RefType         int8   `json:"ref_type"`
	RefId           uint64 `json:"ref_id"`
	Points          uint64 `json:"points"`
	Geog            string `json:"geog"`
}

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "ylbadmin"
	dbPassword = "ylbadminpwd"
	dbName     = "ylbdb01"
)

func getDBEngine() *xorm.Engine {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	engine, err := xorm.NewEngine("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	engine.ShowSQL(true)

	fmt.Println("connect postgresql success")
	return engine
}

//GetUserCount ...
func GetUserCount() uint64 {
	engine := getDBEngine()
	user := new(Users)
	total, err := engine.Where("id >?", 1).Count(user)
	if err != nil {
		fmt.Println(err.Error())
	}

	return uint64(total)
}

//GetUserByID ...
func GetUserByID(id uint64) Users {
	var user Users
	engine := getDBEngine()
	engine.Id(id).Get(&user)
	return user
}

//PointsInDistanceRange 距离某一点给定距离内的若干个点
func PointsInDistanceRange(p geopoints.GeoPoint, d int64) []Users {

	list := make([]Users, 0)
	qStr := p.GetPointsQueryStringWithIn(d)
	engine := getDBEngine()
	err := engine.Where(qStr).Desc("id").Find(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

// AddNewUser ...
func AddNewUser(user *Users) uint64 {

	engine := getDBEngine()

	userInDb := new(Users)
	engine.Where("phone=?", user.Phone).Get(userInDb)
	//使用该手机号码的用户已存在了
	if userInDb.Id != 0 {

		log.Println("使用该手机号码的用户已存在了?")
		log.Println(userInDb.Id)
		log.Println(user.Phone)
		return 0
	}

	/*
		INSERT INTO "users" ("province_agent_id","city_agent_id","county_agent_id","name","phone","geog") VALUES ($1,$2,$3,$4,$5,$6) []interface {}{41, 371, 2, "cccd", "13800138000", "SRID=4326;POINT(113.538639 34.826563)"}
	*/
	//sqlStr := "INSERT INTO \"users\" (\"province_agent_id\",\"city_agent_id\",\"county_agent_id\",\"name\",\"phone\",\"geog\") VALUES ($1,$2,$3,$4,$5,$6)"
	sqlStr := "INSERT INTO \"users\" (\"province_agent_id\",\"city_agent_id\",\"county_agent_id\",\"name\",\"phone\",\"geog\") VALUES (?,?,?,?,?,?)"
	res, err := engine.Exec(sqlStr, int64(user.ProvinceAgentId), int64(user.CityAgentId), int64(user.CountyAgentId), user.Name, user.Phone, user.Geog)

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
	engine.Where("phone=?", user.Phone).Get(userInDb)
	return userInDb.Id

}

//GetRowIDByPhone ...
func GetRowIDByPhone(phone string) uint64 {

	user := new(Users)
	engine := getDBEngine()
	engine.Where("phone=?", phone).Get(user)
	return user.Id

}

//DeleteUserByPhone ...
func DeleteUserByPhone(phone string) bool {
	user := new(Users)
	engine := getDBEngine()
	engine.Where("phone=?", phone).Get(user)
	id := user.Id
	if id == 0 {
		log.Printf("try to delete not exists user with phone %s", phone)
		return false
	}
	rows, err := engine.Delete(&user)
	if err != nil {
		log.Println(err)
		return false
	}
	if rows == 0 {
		return false
	}
	return true
}

func UpdateUser(user *User) bool {

	engine := getDBEngine()

}

func main() {
	//total row in DT
	{ //用户数量
		total := GetUserCount()
		fmt.Println(total)
	}

	{ //添加新用户
		point := geopoints.GeoPoint{Lng: 113.538639, Lat: 34.826563}
		u1 := &Users{
			//Id:              1102,
			ProvinceAgentId: 41,
			CityAgentId:     371,
			CountyAgentId:   2,
			Name:            "ddddd",
			Phone:           "13800138005",
			BizLevel:        2,
			Role:            100,
			RefType:         1,
			RefId:           0,
			Geog:            point.String(),
		}
		isInsertedInt64 := AddNewUser(u1)
		fmt.Println("isInsertedInt64")
		fmt.Println(isInsertedInt64)
	}

	{ //按照ID获取用户
		var p2 geopoints.GeoPoint
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
		p3 := geopoints.GeoPoint{Lng: 113.739873, Lat: 34.356696}
		ps := PointsInDistanceRange(p3, 5000)
		psLen := len(ps)
		if psLen > 0 {

			for i := 0; i < psLen; i++ {
				var p geopoints.GeoPoint
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
