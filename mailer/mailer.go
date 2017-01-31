package mailer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Mailer struct {
	Config *MailConfig
	Emails []*Email
}

type Email struct {
	Id      int
	Address string
}

type MailConfig struct {
	Address  string
	Port     string
	From     string
	Username string
	Password string
}

// Check if mail.json file exists.
func exists() bool {
	_, err := ioutil.ReadFile("mail.json")
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (m *Mailer) nextId() int {
	if len(m.Emails) == 0 {
		return 0
	}
	return m.Emails[len(m.Emails)-1].Id + 1
}

func Open() (*Mailer, error) {
	m := new(Mailer)
	if !exists() {
		return m, nil
	}
	file, err := ioutil.ReadFile("mail.json")
	if err != nil {
		return nil, err
	}
	json.Unmarshal(file, &m)
	return m, nil
}

func must(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func (m *Mailer) Configure(address, port, from, username, password string) error {
	config := new(MailConfig)
	config.Address = address
	config.Port = port
	config.From = from
	config.Username = username
	config.Password = password
	m.Config = config
	return must(m.save())
}

func (m *Mailer) FindByAddress(address string) (int, *Email) {
	for index, email := range m.Emails {
		if address == email.Address {
			return index, email
		}
	}
	return -1, nil
}

func (m *Mailer) FindById(id int) (int, *Email) {
	for index, email := range m.Emails {
		if id == email.Id {
			return index, email
		}
	}
	return -1, nil
}

func (m *Mailer) Insert(address string) error {
	if index, _ := m.FindByAddress(address); index != -1 {
		return errors.New("Email already in database")
	}
	mail := new(Email)
	mail.Id = m.nextId()
	mail.Address = address
	m.Emails = append(m.Emails, mail)
	return must(m.save())
}

func (m *Mailer) Delete(id int) error {
	index, _ := m.FindById(id)
	if index == -1 {
		return errors.New("email not nound")
	}
	m.Emails = append(m.Emails[:index], m.Emails[index+1:]...)
	return must(m.save())
}

func (m *Mailer) save() error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("mail.json", b, 0660)
	if err != nil {
		return err
	}
	return nil
}
