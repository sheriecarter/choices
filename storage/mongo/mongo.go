package mongo

import (
	"log"
	"sync"

	"github.com/foolusion/choices"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type mongo struct {
	namespaces []choices.Namespace
	mu         sync.RWMutex
}

type MongoNamespace struct {
	Name        string
	Segments    [16]byte
	TeamID      []string
	Experiments []MongoExperiments
}

type MongoExperiments struct {
	Name     string
	Segments [16]byte
	Params   []MongoParams
}

type MongoParams struct {
	Name  string
	Type  choices.ValueType
	Value bson.Raw
}

func (m *mongo) Update() {
	sess, err := mgo.Dial("elwin-storage")
	if err != nil {
		log.Println(err)
	}
	c := sess.DB("test").C("namespaces")
	iter := c.Find(bson.M{}).Iter()
	var mongoNamespaces []MongoNamespace
	err = iter.All(&mongoNamespaces)
	if err != nil {
		log.Println(err)
	}

	namespaces := make([]choices.Namespace, len(mongoNamespaces))
	for i, n := range mongoNamespaces {
		namespaces[i] = choices.Namespace{Name: n.Name, Segments: n.Segments, TeamID: n.TeamID, Experiments: make([]choices.Experiment, len(n.Experiments))}
		for j, e := range n.Experiments {
			namespaces[i].Experiments[j] = choices.Experiment{Name: e.Name, Segments: e.Segments, Params: make([]choices.Param, len(e.Params))}
			for k, p := range e.Params {
				namespaces[i].Experiments[j].Params[k] = choices.Param{Name: p.Name}
				switch p.Type {
				case choices.ValueTypeUniform:
					var uniform choices.Uniform
					p.Value.Unmarshal(&uniform)
					namespaces[i].Experiments[j].Params[k].Value = &uniform
				case choices.ValueTypeWeighted:
					var weighted choices.Weighted
					p.Value.Unmarshal(&weighted)
					namespaces[i].Experiments[j].Params[k].Value = &weighted
				}
			}
		}
	}

	m.mu.Lock()
	m.namespaces = namespaces
	m.mu.Unlock()
	return
}

func (m *mongo) Read() []choices.Namespace {
	m.mu.RLock()
	ns := m.namespaces
	m.mu.RUnlock()
	return ns
}