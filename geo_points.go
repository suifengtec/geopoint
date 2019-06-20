package geopoints

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

//GeoPoint 4326
type GeoPoint struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}

//String ...
// point := GeoPoint{Lng: 113.538639, Lat: 34.826563}
// point.String()
// -->postGIS
func (p *GeoPoint) String() string {
	return fmt.Sprintf("SRID=4326;POINT(%v %v)", p.Lng, p.Lat)
}

//Value ...
func (p GeoPoint) Value() (driver.Value, error) {
	return p.String(), nil
}

//JSONString ...
// 从数据库查询了一条或者多条有GEO点信息的数据ps,那么,可以在循环中使用
/*
	p3 := GeoPoint{Lng: 113.739873, Lat: 34.356696}

	ps := PointsInDistanceRange(p3, 5000)
	psLen := len(ps)
	if psLen > 0 {

		for i := 0; i < psLen; i++ {

			var p GeoPoint
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
*/
func (p *GeoPoint) JSONString() string {
	return fmt.Sprintf("{lng:%v,lat: %v}", p.Lng, p.Lat)
}

//Scan ...interface {} is *string, not []uint8
// val interface{}
// val string
// ---
// var p2 GeoPoint
// user := GetOneRowFromDB(2)
// p2.Scan(user.Geog)
// fmt.Println(p2.String())
//113.739873, 34.356696
func (p *GeoPoint) Scan(val string) error {
	//b, err := hex.DecodeString(string(val.([]uint8)))
	b, err := hex.DecodeString(string([]uint8(val)))
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	var wkbByteOrder uint8
	if err := binary.Read(r, binary.LittleEndian, &wkbByteOrder); err != nil {
		return err
	}
	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case 0:
		byteOrder = binary.BigEndian
	case 1:
		byteOrder = binary.LittleEndian
	default:
		return fmt.Errorf("Invalid byte order %d", wkbByteOrder)
	}
	var wkbGeometryType uint64
	if err := binary.Read(r, byteOrder, &wkbGeometryType); err != nil {
		return err
	}
	if err := binary.Read(r, byteOrder, p); err != nil {
		return err
	}

	return nil
}

//GetPointsQueryStringWithIn ...
/*
//PointsInDistanceRange 距离某一点给定距离内的若干个点
func PointsInDistanceRange(p GeoPoint, d int64) []Users {

	list := make([]Users, 0)
	qStr := p.QueryStringWithIn(d)
	engine := getDBEngine()
	err := engine.Where(qStr).Desc("id").Find(&list)
	if err != nil {
		log.Fatal(err)
	}
	return list
}

*/
func (p *GeoPoint) GetPointsQueryStringWithIn(d int64) string {

	return fmt.Sprintf("ST_DWithin(geog::geography, ST_SetSRID(ST_MakePoint(%f, %f),4326)::geography, %d)", p.Lng, p.Lat, d)
}
