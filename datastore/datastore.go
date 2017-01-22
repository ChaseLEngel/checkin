package datastore

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"regexp"
	"time"
)

type Node struct {
	Name       string
	Timestamps []time.Time
}

type DataStore struct {
	Nodes []*Node
}

// Returns duration since last checkin showing only largest time increment
func (n *Node) Since() string {
	last := n.Timestamps[len(n.Timestamps)-1]
	s := time.Since(last).String()
	r, err := regexp.Compile("[0-9]*(\\.[0-9]*)?[a-z]")
	if err != nil {
		return "-1"
	}
	return r.FindString(s)
}

// Check if datastore file exists.
func exists() bool {
	_, err := ioutil.ReadFile("datastore.json")
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (ds *DataStore) InsertOrUpdate(name string) {
	_, node := ds.Find(name)
	if node != nil {
		node.Timestamps = append(node.Timestamps, time.Now())
	} else { // Add new node
		new_node := new(Node)
		new_node.Name = name
		new_node.Timestamps = append(new_node.Timestamps, time.Now())
		ds.Nodes = append(ds.Nodes, new_node)
	}
	ds.save()
}

// Open and parse existing datastore or return empty if file doesn't exist.
func Open() (*DataStore, error) {
	ds := new(DataStore)
	if !exists() {
		return ds, nil
	}
	file, err := ioutil.ReadFile("datastore.json")
	if err != nil {
		return nil, err
	}
	json.Unmarshal(file, &ds)
	return ds, nil
}

// Parse Datastore and writes it to file.
func (ds *DataStore) save() error {
	b, err := json.Marshal(ds)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("datastore.json", b, 0660)
	if err != nil {
		return err
	}
	return nil
}

// Find a node by name from the datastore.
func (ds *DataStore) Find(name string) (int, *Node) {
	for index, node := range ds.Nodes {
		if node.Name == name {
			return index, node
		}
	}
	return -1, nil
}

func (ds *DataStore) Delete(name string) {
	index, _ := ds.Find(name)
	if index == -1 {
		return
	}
	ds.Nodes = append(ds.Nodes[:index], ds.Nodes[index+1:]...)
	ds.save()
}
