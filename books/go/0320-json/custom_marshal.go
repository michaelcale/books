package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func notCustom() {
	// :show start
	type Event struct {
		What string
		When time.Time
	}
	e := Event{
		What: "earthquake",
		When: time.Now(),
	}
	d, err := json.Marshal(&e)
	if err != nil {
		log.Fatalf("json.MarshalIndent failed with '%s'\n", err)
	}
	fmt.Printf("Standard time JSON: %s\n", string(d))
	// :show end
}

// :show start
type customTime time.Time

const customTimeFormat = `"2006-02-01"`

func (ct customTime) MarshalJSON() ([]byte, error) {
	t := time.Time(ct)
	s := t.Format(customTimeFormat)
	return []byte(s), nil
}

func (ct *customTime) UnmarshalJSON(d []byte) error {
	t, err := time.Parse(customTimeFormat, string(d))
	if err != nil {
		return err
	}
	*ct = customTime(t)
	return nil
}

type Event2 struct {
	What string
	When customTime
}

// :show end

func custom() {
	// :show start
	e := Event2{
		What: "earthquake",
		When: customTime(time.Now()),
	}
	d, err := json.Marshal(&e)
	if err != nil {
		log.Fatalf("json.Marshal failed with '%s'\n", err)
	}
	fmt.Printf("\nCustom time JSON: %s\n", string(d))
	var decoded Event2
	err = json.Unmarshal(d, &decoded)
	if err != nil {
		log.Fatalf("json.Unmarshal failed with '%s'\n", err)
	}
	t := time.Time(decoded.When)
	fmt.Printf("Decoded custom time: %s\n", t.Format(customTimeFormat))
	// :show end
}

func main() {
	// :show start
	notCustom()
	custom()
	// :show end
}
