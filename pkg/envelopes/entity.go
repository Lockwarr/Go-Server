package envelopes

import (
	"encoding/xml"
	"fmt"

	"github.com/gocql/gocql"
)

// Envelope struct
type Envelope struct {
	ID      string   `xml:"id" cql:"id"`
	XMLName xml.Name `xml:"Envelope" cql:"xmlname"`
	Text    string   `xml:",chardata" cql:"text"`
	Gesmes  string   `xml:"gesmes,attr" cql:"gesmes"`
	Xmlns   string   `xml:"xmlns,attr" cql:"xmlns"`
	Subject string   `xml:"subject" cql:"subject"`
	Sender  Sender   `xml:"Sender" cql:"sender"`
	Cube    Cube     `xml:"Cube" cql:"cube"`
}

// DB envelopes
type DB struct {
	Host string `xml:"db_host" validate:"nonzero"` // DB host address
	Port string `xml:"db_port" validate:"nonzero"` // DB port
}

// MarshalUDT - Marshal the DB{} type to CQL User Defined Type 'db' type
func (p *DB) MarshalUDT(name string, info gocql.TypeInfo) ([]byte, error) {
	switch name {
	case "host":
		return gocql.Marshal(info, p.Host)
	case "port":
		return gocql.Marshal(info, p.Port)
	default:
		return nil, fmt.Errorf("unknown column for position: %q", name)
	}
}

// UnmarshalUDT - Unmarshal the CQL User Defined Type 'db' to the DB{} type
func (p *DB) UnmarshalUDT(name string, info gocql.TypeInfo, data []byte) error {
	switch name {
	case "host":
		return gocql.Unmarshal(info, data, &p.Host)
	case "port":
		return gocql.Unmarshal(info, data, &p.Port)
	default:
		return fmt.Errorf("unknown column for position: %q", name)
	}
}

// ResponseLatest -is response struct for the first endpoint from the task
type ResponseLatest struct {
	Base  string
	Rates map[string]string
}

type ResponseAnalyze struct {
	Base  string
	Rates map[string]AnalyzedRates
}

type AnalyzedRates struct {
	Min float64
	Max float64
	Avg float64
}

// Sender struct
type Sender struct {
	Text string `xml:",chardata" cql:"text"`
	Name string `xml:"name" cql:"name"`
}

// Cube3 level 3
type Cube3 struct {
	Text     string `xml:",chardata" cql:"text"`
	Currency string `xml:"currency,attr" cql:"currency"`
	Rate     string `xml:"rate,attr" cql:"rate"`
}

// Cube2 level 2
type Cube2 struct {
	Text string  `xml:",chardata" cql:"text"`
	Time string  `xml:"time,attr" cql:"time"`
	Cube []Cube3 `xml:"Cube" cql:"cube3"`
}

// Cube level 1
type Cube struct {
	Text string  `xml:",chardata" cql:"text"`
	Cube []Cube2 `xml:"Cube" cql:"cube2"`
}

// MarshalUDT - Marshal the Sender{} type to CQL User Defined Type 'sender' type
func (p Sender) MarshalUDT(name string, info gocql.TypeInfo) ([]byte, error) {
	switch name {
	case "text":
		return gocql.Marshal(info, p.Text)
	case "name":
		return gocql.Marshal(info, p.Name)
	default:
		return nil, fmt.Errorf("unknown column for position: %q", name)
	}
}

// UnmarshalUDT - Unmarshal the CQL User Defined Type 'sender' to the Sender{} type
func (p *Sender) UnmarshalUDT(name string, info gocql.TypeInfo, data []byte) error {
	switch name {
	case "text":
		return gocql.Unmarshal(info, data, &p.Text)
	case "name":
		return gocql.Unmarshal(info, data, &p.Name)
	default:
		return fmt.Errorf("unknown column for position: %q", name)
	}
}

// MarshalUDT - Marshal the Cube3{} type to CQL User Defined Type 'cube3' type
func (p Cube3) MarshalUDT(name string, info gocql.TypeInfo) ([]byte, error) {
	switch name {
	case "text":
		return gocql.Marshal(info, p.Text)
	case "currency":
		return gocql.Marshal(info, p.Currency)
	case "rate":
		return gocql.Marshal(info, p.Rate)
	default:
		return nil, fmt.Errorf("unknown column for position: %q", name)
	}
}

// UnmarshalUDT - Unmarshal the CQL User Defined Type 'cube3' to the Cube3{} type
func (p *Cube3) UnmarshalUDT(name string, info gocql.TypeInfo, data []byte) error {
	switch name {
	case "text":
		return gocql.Unmarshal(info, data, &p.Text)
	case "currency":
		return gocql.Unmarshal(info, data, &p.Currency)
	case "rate":
		return gocql.Unmarshal(info, data, &p.Rate)
	default:
		fmt.Println("5")
		return fmt.Errorf("unknown column for position: %q", name)
	}
}

// MarshalUDT - Marshal the Cube2{} type to CQL User Defined Type 'cube2' type
func (p Cube2) MarshalUDT(name string, info gocql.TypeInfo) ([]byte, error) {
	switch name {
	case "text":
		return gocql.Marshal(info, p.Text)
	case "time":
		return gocql.Marshal(info, p.Time)
	case "cube3":
		return gocql.Marshal(info, p.Cube)
	default:
		return nil, fmt.Errorf("unknown column for position: %q", name)
	}
}

// UnmarshalUDT - Unmarshal the CQL User Defined Type 'cube2' to the Cube2{} type
func (p *Cube2) UnmarshalUDT(name string, info gocql.TypeInfo, data []byte) error {
	switch name {
	case "text":
		return gocql.Unmarshal(info, data, &p.Text)
	case "time":
		return gocql.Unmarshal(info, data, &p.Time)
	case "cube3":
		return gocql.Unmarshal(info, data, &p.Cube)
	default:
		return fmt.Errorf("unknown column for position: %q", name)
	}
}

// MarshalUDT - Marshal the Cube{} type to CQL User Defined Type 'cube' type
func (p Cube) MarshalUDT(name string, info gocql.TypeInfo) ([]byte, error) {
	switch name {
	case "text":
		return gocql.Marshal(info, p.Text)
	case "cube2":
		return gocql.Marshal(info, p.Cube)
	default:
		return nil, fmt.Errorf("unknown column for position: %q", name)
	}
}

// UnmarshalUDT - Unmarshal the CQL User Defined Type 'cube' to the Cube{} type
func (p *Cube) UnmarshalUDT(name string, info gocql.TypeInfo, data []byte) error {
	switch name {
	case "text":
		return gocql.Unmarshal(info, data, &p.Text)
	case "cube2":
		return gocql.Unmarshal(info, data, &p.Cube)
	default:
		return fmt.Errorf("unknown column for position: %q", name)
	}
}
