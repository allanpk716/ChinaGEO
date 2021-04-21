package main

import geo "github.com/allanpk716/google-geolocate"

type GEOHelper struct {
	client *geo.GoogleGeo
}

func NewGEOHelper() *GEOHelper {
	c := GEOHelper{}
	c.client = geo.NewGoogleGeo("", "http://127.0.0.1:10809")
	return &c
}

func (c *GEOHelper) queryLocationName(name string) (*geo.Point, error) {
	res, err := c.client.Geocode(name)
	if err != nil {
		return nil, err
	}
	return res, nil
}