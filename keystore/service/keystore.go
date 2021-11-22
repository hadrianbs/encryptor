package service

type KeyService struct {
	keyStore KeyStorer
}

type KeyStorer interface {
	FetchKey(keyID string) ([]byte, error)
}

func NewKeyService(k KeyStorer) *KeyService {
	return &KeyService{
		keyStore: k,
	}
}

func (c *KeyService) FetchKey(keyID string) ([]byte, error) {
	key, err := c.keyStore.FetchKey(keyID)
	if err != nil {
		return nil, err
	}

	return key, nil
}
