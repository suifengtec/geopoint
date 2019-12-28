# geopoint

geopoint is a go package to handle PostGIS point type.

## Init

if you do not use go module function , you need to get it via the following command:

```bash
go get github.com/suifengtec/geopoint

```

import it in your go package or go project:

```go

import(

    "github.com/suifengtec/geopoint"
)

```

## Use it

```go

p1 := geopoint.GeoPoint{Lng: 113.538639, Lat: 34.826563}
//to String via String() or p1.ToString()
p1.String()
//to JSON string
p1.JSONString()
//Scan PostGIS point data in db to Geopiont type.
var p2 geopoint.GeoPoint
user := GetOneRowFromDB(2)
p2.Scan(user.Geog)
fmt.Println(p2.String())


//Query Geopoint relatived instances by distance range helper example:

//GetUsersInDistanceRange ...
func GetUsersInDistanceRange(p geopoint.GeoPoint, d int64) []User {

    list := make([]User, 0)
    // query helper.
    qStr := p.GetPointsQueryStringWithIn(d)
    //X is the xorm.Engine pointer.
    err := X.Where(qStr).Desc("id").Find(&list)
    if err != nil {
        log.Fatal(err)
    }
    return list
}

```

## Licence

MIT.
