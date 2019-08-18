package main

import (
	"log"
	"net/http"

	"github.com/Lockwarr/Go-Server/common/pkg/cql"
	e "github.com/Lockwarr/Go-Server/pkg/envelopes"
	"github.com/gocql/gocql"
)

func main() {
	//rawEnvelope, _ := e.GetXML("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml")
	//var parsedEnvelope e.Envelope
	//xml.Unmarshal([]byte(rawEnvelope), &parsedEnvelope)

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.ProtoVersion = 4
	cluster.Keyspace = "denislav"
	session, _ := cluster.CreateSession()
	defer session.Close()

	envelopesRepository := e.NewRepository(session, cql.NewKeyspaceBinder("denislav"))
	app := &e.App{&e.HandlerLatest{envelopesRepository}, &e.HandlerDate{envelopesRepository, ""}, &e.HandlerAnalyze{envelopesRepository}}
	//err := envelopesRepository.Migrations()
	//if err != nil {
	//	panic(err)
	//}
	//envelopesRepository.InsertEnvelope(parsedEnvelope)
	log.Println("http://localhost:8080/rates/latest")
	if err := http.ListenAndServe(":8080", app); err != nil {
		panic(err)
	}
}
