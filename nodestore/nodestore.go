package nodestore

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type Node struct {
	Id         int
	Name       string
	Timestamps []time.Time
	Schedule   *Schedule
	Misses     []int
}

type Schedule struct {
	minutes int
	hours   int
	days    int
	weeks   int
	months  int
}

type NodeStore struct {
	Nodes []*Node
}

func (s *Schedule) Check(time time.Duration) bool {
	if s == nil {
		return true
	}
	return true
}

// Open and parse existing nodestore or return empty if file doesn't exist.
func Open() (*NodeStore, error) {
	ns := new(NodeStore)
	if !exists() {
		return ns, nil
	}
	file, err := ioutil.ReadFile("nodestore.json")
	if err != nil {
		return nil, err
	}
	json.Unmarshal(file, &ns)
	return ns, nil
}

// Check if nodestore file exists.
func exists() bool {
	_, err := ioutil.ReadFile("nodestore.json")
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Returns the next avaliable node id
func (ns *NodeStore) nextId() int {
	if len(ns.Nodes) == 0 {
		return 0
	}
	return ns.Nodes[len(ns.Nodes)-1].Id + 1
}

func (ns *NodeStore) InsertOrUpdate(name string) (*Node, error) {
	_, n := ns.FindByName(name)
	// New node
	if n == nil {
		n = new(Node)
		n.Id = ns.nextId()
		n.Name = name
		ns.Nodes = append(ns.Nodes, n)
	}
	n.Timestamps = append(n.Timestamps, time.Now())
	if err := ns.save(); err != nil {
		return nil, err
	}
	return n, nil
}

// Parse Datastore and writes it to file.
func (ns *NodeStore) save() error {
	b, err := json.Marshal(ns)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("nodestore.json", b, 0660)
	if err != nil {
		return err
	}
	return nil
}

// Find a node by name from the nodestore.
func (ns *NodeStore) FindById(id int) (int, *Node) {
	for index, node := range ns.Nodes {
		if node.Id == id {
			return index, node
		}
	}
	return -1, nil
}

// Find a node by id from the nodestore.
func (ns *NodeStore) FindByName(name string) (int, *Node) {
	for index, node := range ns.Nodes {
		if node.Name == name {
			return index, node
		}
	}
	return -1, nil
}

func must(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func (ns *NodeStore) Delete(id int) error {
	index, _ := ns.FindById(id)
	if index == -1 {
		return errors.New("node not found")
	}
	ns.Nodes = append(ns.Nodes[:index], ns.Nodes[index+1:]...)
	return must(ns.save())
}
