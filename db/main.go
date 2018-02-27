package main

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	_ "github.com/aws/aws-sdk-go/service/dynamodb"
	_ "github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"fmt"
)

// Use struct tags much like the standard JSON library,
// you can embed anonymous structs too!
type widget struct {
	UserID int       // Hash key, a.k.a. partition key
	Time   time.Time // Range key, a.k.a. sort key

	Msg       string              `dynamo:"Message"`
	Count     int                 `dynamo:",omitempty"`
	Friends   []string            `dynamo:",set"` // Sets
	Set       map[string]struct{} `dynamo:",set"` // Map sets, too!
	SecretKey string              `dynamo:"-"`    // Ignored
	Children  []struct{}                          // Lists
}

var endpoint = "daxdynamocluster.bogjh7.clustercfg.dax.usw2.cache.amazonaws.com:8111"

func main() {
	//endpoint := "http://localhost:8000"
	//e2 := &endpoint

	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("us-west-2"),
		Endpoint: &endpoint,
		//Endpoint: e2,
	})
	table := db.Table("Widgets")

	describe := table.Describe()
	_ = describe

	// put item
	w := widget{UserID: 613, Time: time.Now(), Msg: "hello 2"}
	err := table.Put(w).Run()
	_ = err
	if err != nil {
		panic(err)
	}

	// get the same item
	var result widget
	err = table.
		Get("UserID", w.UserID).
		Range("Time", dynamo.Equal, w.Time).
		Filter("'Count' = ? AND $ = ?", w.Count, "Message", w.Msg). // placeholders in expressions
		One(&result)

	fmt.Println(result)

	// get all items
	var results []widget
	err = table.Scan().All(&results)

	fmt.Println(results)
}

type DLumiSuperProxy struct {
	ID      string    `dynamo:"ID"`
	Created time.Time `dynamo:"Created"`
	Status  int       `dynamo:"Status"`
}

type DLumiSuperProxyEvent struct {
	ID   string    `dynamo:"ID"`
	Time time.Time `dynamo:"Time"`
}

type LumiSessionStatus int

const (
	LumiSessionPrepare LumiSessionStatus = iota
)

type DSession struct {
	ID        string
	Created   time.Time
	AccountID string
	UserID    string

	SubID  string
	SubID2 string
	SubID3 string

	Started time.Time
	EndAt   time.Time
	Ended   time.Time

	Status int
}

func StartSession() {
	// Retrieve session from Dynamo.
	// If this is first request then try to start it.
}