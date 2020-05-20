package query

import (
	"net/url"

	"gopkg.in/mgo.v2/bson"
)

var (
	nameKey      = "name"
	genreKey     = "genres"
	platformKey  = "platforms"
	franchiseKey = "franchise"
	sortKey      = "$natural"
)

//Query parameters
type Query struct {
	Name      string
	Genre     string
	Platform  string
	Franchise string
	Sort      string
}

//New Query
func New(params url.Values) []bson.M {
	q := Query{
		Name:      params.Get(nameKey),
		Genre:     params.Get(genreKey),
		Platform:  params.Get(platformKey),
		Franchise: params.Get(franchiseKey),
		Sort:      params.Get(sortKey),
	}

	return []bson.M{{"$match": q.build()}}
}

func (q *Query) build() bson.M {
	pipe := make([]bson.M, 0)

	if hasValue(q.Name) {
		pipe = append(pipe, bson.M{nameKey: bson.M{"$in": q.Name}})
	}
	if hasValue(q.Genre) {
		pipe = append(pipe, bson.M{genreKey: bson.M{"$in": []string{q.Genre}}})
	}
	if hasValue(q.Platform) {
		pipe = append(pipe, bson.M{platformKey: bson.M{"$in": []string{q.Platform}}})
	}
	if hasValue(q.Franchise) {
		pipe = append(pipe, bson.M{franchiseKey: bson.M{"$in": q.Franchise}})
	}
	if hasValue(q.Sort) {
		pipe = append(pipe, bson.M{"_id": q.Sort})
	}

	if len(pipe) > 0 {
		return bson.M{"$and": pipe}
	}

	return bson.M{}
}

func hasValue(value string) bool {
	if value == "" {
		return false
	}
	return true
}
