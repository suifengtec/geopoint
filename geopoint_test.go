package geopoint

import "testing"

/*
* @Author: suifengtec
* @Date:   2019-12-29 02:57:50
* @Last Modified by:   suifengtec
* @Last Modified time: 2019-12-29 03:35:19
 */

func TestString(t *testing.T) {
	want := "SRID=4326;POINT(113.538639 34.826563)"
	p := GeoPoint{Lng: 113.538639, Lat: 34.826563}
	if got := p.String(); got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestToString(t *testing.T) {
	want := "SRID=4326;POINT(113.538639 34.826563)"
	p := GeoPoint{Lng: 113.538639, Lat: 34.826563}
	if got := p.ToString(); got != want {
		t.Errorf("ToString() = %q, want %q", got, want)
	}
}

func TestValue(t *testing.T) {
	want := "SRID=4326;POINT(113.538639 34.826563)"
	p := GeoPoint{Lng: 113.538639, Lat: 34.826563}
	got, err := p.Value()
	if err != nil {
		t.Errorf("Value() return error %s", err.Error())
	}
	if got != want {
		t.Errorf("Value() = %q, want %q", got, want)
	}
}

func TestJSONString(t *testing.T) {
	want := "{lng:113.538639,lat: 34.826563}"
	p := GeoPoint{Lng: 113.538639, Lat: 34.826563}
	if got := p.JSONString(); got != want {
		t.Errorf("JSONString() = %q, want %q", got, want)
	}
}

func TestScan(t *testing.T) {

	p1 := GeoPoint{Lng: 113.739873, Lat: 34.356696}
	want := p1.String()
	p2 := &GeoPoint{}
	if got1 := p2.Scan("0101000020E6100000618C48145A6F5C40984EEB36A82D4140"); got1 != nil {
		t.Errorf("Scan() = %q, want nil", got1)
	}

	got2 := p2.String()

	if got2 != want {
		t.Errorf("Scan()->p2 = %q, want %q", got2, want)
	}
}

func TestGetPointsQueryStringWithIn(t *testing.T) {
	want := "ST_DWithin(geog::geography, ST_SetSRID(ST_MakePoint(113.538639, 34.826563),4326)::geography, 5000)"
	p := GeoPoint{Lng: 113.538639, Lat: 34.826563}
	if got := p.GetPointsQueryStringWithIn(5000); got != want {
		t.Errorf("GetPointsQueryStringWithIn() = %q, want %q", got, want)
	}
}
