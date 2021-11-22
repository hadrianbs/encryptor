package provider

import (
	"errors"
)

type LocalKey struct {
	ID  string
	Key string
}

type LocalKeyStore struct {
	keys map[string]LocalKey
}

func NewLocalKeyStore(keys []LocalKey) *LocalKeyStore {
	keyMap := map[string]LocalKey{}
	for _, key := range keys {
		keyMap[key.ID] = key
	}
	return &LocalKeyStore{
		keys: keyMap,
	}
}

func (c *LocalKeyStore) FetchKey(keyID string) ([]byte, error) {
	if _, ok := c.keys[keyID]; !ok {
		return nil, errors.New("key not found")
	}
	return []byte(c.keys[keyID].Key), nil
}
