package db

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const (
	ErrInvalidConfig string = "Cannot connect to mongo database. Invalid connection parameters"
	ErrClientIsNil   string = "Client is nil"
	ErrDBIsNil       string = "DB is nil"
	ErrIndexIsNil    string = "Index is nil"
	ErrNotFound      string = "Not found"
)

// Config 정의
type Config struct {
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Database string `json:"database" yaml:"database"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"username" yaml:"password"`
}

// MongoDB 관련 기능 포함
type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
	conf   *Config
}

// GetConnectString 연결 문자열 변환
func GetConnectString(conf *Config) string {
	if conf.Username != "" && conf.Password != "" {
		return fmt.Sprintf("mongodb://%s:%s@%s:%d", conf.Username, conf.Password, conf.Host, conf.Port)
	}
	return fmt.Sprintf("mongodb://%s:%d", conf.Host, conf.Port)
}

// NewMongoDB 세션을 생성한다.
func NewMongoDB(config ...interface{}) (*MongoDB, error) {
	if len(config) == 0 {
		return nil, errors.New(ErrInvalidConfig)
	}

	connStr := ""
	conf := new(Config)

	switch v := config[0].(type) {
	case string:
		connStr = v
	case Config:
		conf = &v
		connStr = GetConnectString(conf)
	case *Config:
		conf = v
		connStr = GetConnectString(conf)
	}

	o := new(MongoDB)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, err
	}

	o.client = client
	o.conf = conf
	o.db = o.client.Database(conf.Database)

	return o, nil
}

// ParseFilter 검색 파라미터맵에서 sort option 파싱
func ParseFilter(params map[string]string) (FindOptions, bson.M, error) {
	opts := FindOptions{}
	sortOption := FindOption{}

	for k, v := range params {
		switch k {
		case "sort":
			tokens := strings.Split(v, ",")
			if len(tokens) == 1 {
				sortOption.Key = tokens[0]
				sortOption.Value = -1 // default desc
			} else if len(tokens) == 2 {
				sortOption.Key = tokens[0]
				i, err := strconv.Atoi(tokens[1])
				if err != nil {
					continue
				}
				sortOption.Value = int64(i)
			} else {
				continue
			}

			opts["sort"] = sortOption
		case "limit":
			i, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			opts["limit"] = int64(i)
		}
	}
	delete(params, "sort")
	delete(params, "limit")

	filter := bson.M{}
	parseVF := func(k, v string) (string, bson.M) {
		var toVals []interface{}

		vType := "string" // int, float
		if k[0] == '!' {
			vType = "int"
			k = k[1:]
		} else if v[0] == '@' {
			vType = "float"
			k = k[1:]
		}

		toVF := func(v, t string) interface{} {
			switch t {
			case "int":
				i, _ := strconv.ParseInt(v, 10, 0)
				return i
			case "float":
				f, _ := strconv.ParseFloat(v, 64)
				return f
			default:
				return v
			}
		}

		keyTokens := strings.Split(k, ",")
		exp := ""
		if len(keyTokens) == 2 {
			exp = keyTokens[1]
		}

		// TODO: 검색문자열에 ,가 포함된 경우 예외 처리 필요
		tokens := strings.Split(v, ",")
		if exp == "" && len(tokens) >= 2 {
			exp = tokens[0]
			tokens = tokens[1:]

			v = strings.Join(tokens, ",")
		}

		if exp == "$in" {
			for _, tv := range tokens {
				toVals = append(toVals, toVF(tv, vType))
			}
			return keyTokens[0], bson.M{exp: toVals}
		}

		if exp == "" {
			return keyTokens[0], bson.M{"$eq": toVF(v, vType)}
		}

		if vType == "string" {
			if exp == "$lte" || exp == "$lt" || exp == "$gte" || exp == "$gt" {
				i, _ := strconv.Atoi(v)
				return keyTokens[0], bson.M{exp: int64(i)}
			}
		}

		return keyTokens[0], bson.M{exp: toVF(v, vType)}
	}
	appendVF := func(k string, v bson.M) bson.M {
		if oldV, ok := filter[k]; ok {
			switch tv := oldV.(type) {
			case bson.M:
				for ok, ov := range tv {
					v[ok] = ov
				}
			}
		}

		return v
	}
	for k, v := range params {
		switch k {
		case "offset":
			var qv interface{}
			i, err := strconv.Atoi(v)
			if err != nil {
				qv = v
			} else {
				qv = int64(i)
			}

			if sortOption.Value == int64(-1) {
				// desc
				filter[sortOption.Key] = appendVF(sortOption.Key, bson.M{"$lt": qv})
			} else {
				// asc
				filter[sortOption.Key] = appendVF(sortOption.Key, bson.M{"$gt": qv})
			}
		default:
			newK, newV := parseVF(k, v)
			filter[newK] = appendVF(newK, newV)
		}
	}

	return opts, filter, nil
}

// Close the DB Connection
func (o *MongoDB) Close(ctx context.Context) {
	if o.client != nil {
		o.client.Disconnect(ctx)
	}
}

// GetCol 컬렉션 조회
func (o *MongoDB) GetCol(colName string) (interface{}, error) {
	if o.client == nil {
		return nil, errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return nil, errors.New(ErrDBIsNil)
	}

	col := o.db.Collection(colName)
	return col, nil
}

// DropCollection 컬렉션 삭제
func (o *MongoDB) DropCollection(ctx context.Context, colName string) error {
	if o.client == nil {
		return errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return errors.New(ErrDBIsNil)
	}

	col := o.db.Collection(colName)
	return col.Drop(ctx)
}

// GetCollections 전체 컬렉션 조회
func (o *MongoDB) GetCollections(ctx context.Context) ([]string, error) {
	if o.client == nil {
		return nil, errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return nil, errors.New(ErrDBIsNil)
	}

	cur, err := o.db.ListCollections(ctx, bsonx.Doc{})
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	for cur.Next(nil) {
		next := &bsonx.Doc{}
		if err := cur.Decode(next); err != nil {
			return nil, err
		}

		elem, err := next.LookupErr("name")
		if err != nil {
			return nil, err
		}
		if elem.Type() != bson.TypeString {
			return nil, fmt.Errorf("incorrect type for 'name'. got %v, want %v", elem.Type(), bson.TypeString)
		}
		elemName := elem.StringValue()
		result = append(result, elemName)
	}

	return result, nil
}

// DropDatabase 데이터베이스 삭제
func (o *MongoDB) DropDatabase(ctx context.Context) error {
	if o.client == nil {
		return errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return errors.New(ErrDBIsNil)
	}

	return o.db.Drop(ctx)
}

// DeleteOne 도큐먼트 하나 삭제
func (o *MongoDB) DeleteOne(ctx context.Context, colName string, filter interface{}) (int64, error) {
	if o.client == nil {
		return 0, errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return 0, errors.New(ErrDBIsNil)
	}

	col := o.db.Collection(colName)
	ret, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}

	return ret.DeletedCount, nil
}

// DeleteMany 여러 도큐먼트 삭제
func (o *MongoDB) DeleteMany(ctx context.Context, colName string, filter interface{}) (int64, error) {
	if o.client == nil {
		return 0, errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return 0, errors.New(ErrDBIsNil)
	}

	col := o.db.Collection(colName)
	ret, err := col.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	return ret.DeletedCount, nil
}

// InsertOne 도큐먼트 하나 추가
func (o *MongoDB) InsertOne(ctx context.Context, colName string, d interface{}) (interface{}, error) {
	if o.client == nil {
		return nil, errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return nil, errors.New(ErrDBIsNil)
	}

	col := o.db.Collection(colName)
	result, err := col.InsertOne(ctx, d)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

// InsertMany 여러 도큐먼트 추가
func (o *MongoDB) InsertMany(ctx context.Context, colName string, d ...interface{}) ([]interface{}, error) {
	if o.client == nil {
		return nil, errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return nil, errors.New(ErrDBIsNil)
	}

	col := o.db.Collection(colName)
	result, err := col.InsertMany(ctx, d)
	if err != nil {
		return nil, err
	}

	return result.InsertedIDs, nil
}

// UpdateOne 도큐먼트 하나 수정
func (o *MongoDB) UpdateOne(ctx context.Context, colName string, filter interface{}, d interface{}) (int64, error) {
	if o.client == nil {
		return 0, errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return 0, errors.New(ErrDBIsNil)
	}

	col := o.db.Collection(colName)
	result, err := col.UpdateOne(ctx, filter, d)
	if err != nil {
		return 0, err
	}

	return result.MatchedCount, nil
}

// UpdateMany 여러 도큐먼트 수정
func (o *MongoDB) UpdateMany(ctx context.Context, colName string, filter interface{}, d interface{}) (int64, error) {
	if o.client == nil {
		return 0, errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return 0, errors.New(ErrDBIsNil)
	}

	col := o.db.Collection(colName)
	result, err := col.UpdateMany(ctx, filter, d)
	if err != nil {
		return 0, err
	}

	return result.MatchedCount, nil
}

// FindOptions 정의
type FindOptions map[string]interface{}

// FindOption 정의
type FindOption struct {
	Key   string
	Value interface{}
}

// FindMany 여러 데이터 찾기
func (o *MongoDB) FindMany(ctx context.Context, colName string, filter interface{}, opts FindOptions, unmarshalF func(bson.Raw) (interface{}, error)) ([]interface{}, error) {
	if o.client == nil {
		return nil, errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return nil, errors.New(ErrDBIsNil)
	}

	col := o.db.Collection(colName)
	findopts := &options.FindOptions{}
	for k, v := range opts {
		if k == "sort" {
			opt, ok := v.(FindOption)
			if !ok {
				continue
			}
			i, ok := opt.Value.(int64)
			if !ok {
				continue
			}
			findopts.SetSort(bson.M{opt.Key: i})
		}
		if k == "limit" {
			i, ok := v.(int64)
			if !ok {
				continue
			}
			findopts.SetLimit(i)
		}
	}

	cur, err := col.Find(ctx, filter, findopts)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, 0)

	for cur.Next(ctx) {
		obj, err := unmarshalF(cur.Current)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}

	return result, err
}

// FindOne 데이터 하나 찾기
func (o *MongoDB) FindOne(ctx context.Context, colName string, filter interface{}, unmarshalF func(bson.Raw) (interface{}, error)) (interface{}, error) {
	if o.client == nil {
		return nil, errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return nil, errors.New(ErrDBIsNil)
	}

	col := o.db.Collection(colName)
	cur, err := col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		obj, err := unmarshalF(cur.Current)
		if err != nil {
			return nil, err
		}
		return obj, nil
	}

	return nil, errors.New(ErrNotFound)
}

// Count 필터와 일치하는 도큐먼트 개수 반환
func (o *MongoDB) Count(ctx context.Context, colName string, filter interface{}) (int64, error) {
	if o.client == nil {
		return 0, errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return 0, errors.New(ErrDBIsNil)
	}

	col := o.db.Collection(colName)
	return col.CountDocuments(ctx, filter)
}

// Index  정의
type Index struct {
	ColName   string
	FieldName string
	Unique    bool
	Asc       bool
	MaxTime   time.Duration // 쿼리 최대 허용 실행 시간
}

// CreateIndexOne 인덱스 생성
func (o *MongoDB) CreateIndexOne(ctx context.Context, index *Index) error {
	if o.client == nil {
		return errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return errors.New(ErrDBIsNil)
	}
	if index == nil {
		return errors.New(ErrIndexIsNil)
	}

	col := o.db.Collection(index.ColName)

	opts := options.CreateIndexes()
	opts.SetMaxTime(index.MaxTime)

	orderby := -1 // asc : 1, desc : -1
	if index.Asc {
		orderby = 1
	}
	keys := bsonx.Doc{{Key: index.FieldName, Value: bsonx.Int32(int32(orderby))}}
	im := mongo.IndexModel{Keys: keys}

	im.Options = &options.IndexOptions{}
	im.Options.SetUnique(index.Unique)

	_, err := col.Indexes().CreateOne(ctx, im, opts)
	if err != nil {
		return err
	}

	return nil
}

// DropIndexOne 인덱스 삭제
func (o *MongoDB) DropIndexOne(ctx context.Context, colName, name string) error {
	if o.client == nil {
		return errors.New(ErrClientIsNil)
	}
	if o.db == nil {
		return errors.New(ErrDBIsNil)
	}

	col := o.db.Collection(colName)
	_, err := col.Indexes().DropOne(ctx, name)
	if err != nil {
		return err
	}

	return nil
}
