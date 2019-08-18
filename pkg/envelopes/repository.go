package envelopes

import (
	"fmt"
	"log"

	"github.com/Lockwarr/Go-Server/common/pkg/cql"
	"github.com/gocql/gocql"
)

type repository struct {
	session *gocql.Session
	kb      *cql.KeyspaceBinder
}

type Repository interface {
	InsertEnvelope(envelope Envelope) error
	GetEnvelope() (Envelope, error)
	Migrations() error
}

// NewRepository - new repository
func NewRepository(session *gocql.Session, kb *cql.KeyspaceBinder) Repository {
	return &repository{session: session, kb: kb}
}

func (r *repository) InsertEnvelope(envelope Envelope) error {
	var err error
	if err != nil {
		return err
	}
	if err = r.session.Query("INSERT INTO envelopes(id, cube, gesmes, sender, subject, text, xmlns) VALUES(uuid(), ?, ?, ?, ?, ?, ?)",
		envelope.Cube, envelope.Gesmes, envelope.Sender, envelope.Subject, envelope.Text, envelope.Xmlns).Exec(); err != nil {
		fmt.Println(err)
		return err
	}

	return err
	//st := []string{"xmlname", "cube", "gesmes", "sender", "subject", "text", "xmlns"}
	//stmt, names := qb.Insert(r.kb.Table("envelopes")).Columns(st...).ToCql()
	//q := gocqlx.Query(r.session.Query(stmt), names).BindStruct(envelope)
	//fmt.Println(q)
	//return q.ExecRelease()
}

func (r *repository) GetEnvelope() (Envelope, error) {
	var envelope Envelope
	err := r.session.Query(`SELECT cube, gesmes, sender, subject, text, xmlns FROM envelopes`).Consistency(gocql.One).Scan(&envelope.Cube, &envelope.Gesmes, &envelope.Sender, &envelope.Subject, &envelope.Text, &envelope.Xmlns)
	if err != nil {
		log.Println(err)
	}
	return envelope, err
}

func (r *repository) Migrations() error {
	//custom types
	err := r.session.Query("CREATE TYPE IF NOT EXISTS sender (text text, name text);").Exec()
	if err != nil {
		return err
	}
	err = r.session.Query("CREATE TYPE IF NOT EXISTS cube3 (text text, currency text, rate text);").Exec()
	if err != nil {
		return err
	}
	err = r.session.Query("CREATE TYPE IF NOT EXISTS cube2 (text text, time text, cube3 list<frozen<cube3>>);").Exec()
	if err != nil {
		return err
	}
	err = r.session.Query("CREATE TYPE IF NOT EXISTS cube (text text, cube2 list<frozen<cube2>>);").Exec()
	if err != nil {
		return err
	}
	//create envelopes table
	err = r.session.Query("CREATE TABLE IF NOT EXISTS envelopes (id uuid, text text, gesmes text, xmlns text, subject text, sender frozen<sender>,	cube frozen<cube>, PRIMARY KEY (id));").Exec()
	if err != nil {
		return err
	}
	return nil
}
