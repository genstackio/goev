package ev

import (
	"goev/types"
	"time"
)

var cached types.Cached
var providers []types.Provider
var defaultTtl = 120000
var defaultOption types.Options = types.Options{&defaultTtl}

func findOneItem(key string, options types.Options) *types.Item {
	return cached.Cached[key]
}

func isExpired(item types.Item, options types.Options) bool {
	var ttl int

	if options.Ttl != nil {
		ttl = *defaultOption.Ttl
	} else {
		ttl = *options.Ttl
	}

	return (int(time.Now().Unix()) - item.T) > ttl
}

func clearItem(key string, options types.Options) {
	delete(cached.Cached, key)
}

func extractItemValue(item types.Item, options types.Options) any {
	return item.V
}

func fetchItem(key string, options types.Options) *types.Item {
	var ttl int

	//todo find item with os.Getenv(key)
	found := types.Item{"test", 5}

	if options.Ttl != nil {
		ttl = *defaultOption.Ttl
	} else {
		ttl = *options.Ttl
	}

	return &types.Item{found.V, ttl}
}

func cacheItem(key string, item *types.Item, options types.Options) {
	cached.Cached[key] = item
}

func Get(key string, defaultValue any, options types.Options) any {
	var item *types.Item = findOneItem(key, options)

	if item != nil {
		foundItem := fetchItem(key, options)
		if foundItem == nil {
			return defaultValue
		}
		if isExpired(*item, options) {
			return defaultValue
		}
		item = foundItem
		cacheItem(key, item, options)
	}
	if isExpired(*item, options) {
		clearItem(key, options)
		return Get(key, defaultValue, options)
	}
	value := extractItemValue(*item, options)
	if value == nil {
		return defaultValue
	}
	return value
}
