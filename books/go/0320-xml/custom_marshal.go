package main

import (
	"encoding/xml"
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
	d, err := xml.Marshal(&e)
	if err != nil {
		log.Fatalf("json.MarshalIndent failed with '%s'\n", err)
	}
	fmt.Printf("Standard time JSON: %s\n", string(d))
	// :show end
}

// :show start
type customTime time.Time

const customTimeFormat = `2006-02-01`

func (ct customTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := time.Time(ct)
	v := t.Format(customTimeFormat)
	return e.EncodeElement(v, start)
}

func (ct *customTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	t, err := time.Parse(customTimeFormat, s)
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
	d, err := xml.Marshal(&e)
	if err != nil {
		log.Fatalf("xml.Marshal failed with '%s'\n", err)
	}
	fmt.Printf("\nCustom time XML: %s\n", string(d))
	var decoded Event2
	err = xml.Unmarshal(d, &decoded)
	if err != nil {
		log.Fatalf("xml.Unmarshal failed with '%s'\n", err)
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
