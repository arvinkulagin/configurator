package consul

import (
	"strings"

	"github.com/hashicorp/consul/api"
)

const (
	_SEPARATOR = "/"
	_SCHEME    = "http"
)

type Consul struct {
	kv *api.KV
}

func New(addr string) (Consul, error) {
	c := Consul{}
	config := api.DefaultConfig()
	config.Address = addr
	config.Scheme = _SCHEME
	client, err := api.NewClient(config)
	if err != nil {
		return c, err
	}
	c.kv = client.KV()
	return c, nil
}

func (c Consul) Get(prefix string) (map[string]interface{}, error) {
	prefix = strings.TrimPrefix(prefix, "/")
	pairs, _, err := c.kv.List(prefix, nil)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	for _, pair := range pairs {
		path := strings.Split(pair.Key, _SEPARATOR)[len(strings.Split(prefix, "/")):]
		if path[0] == "" {
			continue
		}
		put(data, path, string(pair.Value))
	}
	return data, nil
}

func put(data map[string]interface{}, path []string, value string) {
	if len(path) == 1 && path[0] == "" {
		return
	}
	if len(path) == 1 {
		data[path[0]] = value
		return
	}
	if _, exist := data[path[0]]; !exist {
		data[path[0]] = make(map[string]interface{})
	}
	put(data[path[0]].(map[string]interface{}), path[1:], value)
}
