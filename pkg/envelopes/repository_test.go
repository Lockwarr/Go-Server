package envelopes

import (
	"testing"

	"github.com/Lockwarr/Go-Server/common/pkg/cql"
	"github.com/gocql/gocql"
	"github.com/stretchr/testify/assert"
)

var repo Repository
var envelope Envelope

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.ProtoVersion = 4
	cluster.Keyspace = "denislav"
	session, _ := cluster.CreateSession()

	repo = NewRepository(session, cql.NewKeyspaceBinder("denislav"))

	sender := Sender{Text: "testsender", Name: "testName"}
	cube3 := []Cube3{
		{Text: "testc", Currency: "test", Rate: "test"},
		{Text: "testc", Currency: "test", Rate: "test"},
		{Text: "testc", Currency: "test", Rate: "test"},
	}
	cube2 := []Cube2{
		{Text: "test", Time: "Test", Cube: cube3},
		{Text: "test", Time: "Test", Cube: cube3},
		{Text: "test", Time: "Test", Cube: cube3},
	}
	cube := Cube{Text: "test", Cube: cube2}

	envelope = Envelope{ID: "cec76b153cc8a85c5", Text: "test", Gesmes: "STestl", Xmlns: "test", Subject: "test", Sender: sender, Cube: cube}
}

//func TestInsertEnvelope(t *testing.T) {
//	assert := assert.New(t)
//	err := repo.InsertEnvelope(envelope)
//	assert.NoError(err)
//}

func TestGetEnvelope(t *testing.T) {
	assert := assert.New(t)
	envelopeTest := []struct {
		name      string
		wantedErr bool
	}{
		{name: "Get envelopes", wantedErr: false},
	}
	for _, tt := range envelopeTest {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			envelope, _ := repo.GetEnvelope()
			assert.NotEmpty(envelope, "envelope should not be empty")
		})
	}
}

func TestMigrations(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{name: "test1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.Migrations(); (err != nil) != tt.wantErr {
				t.Errorf("repository.Migrations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
