package types

type ProviderInterface interface {
	FindOne(key string, options any) Item
}

type Item struct {
	V any
	T int
}

type Cached struct {
	Cached map[string]*Item
}

type Options struct {
	Ttl *int
}

type Provider struct {
	Name     string
	Priority int
	Provider ProviderInterface
}
