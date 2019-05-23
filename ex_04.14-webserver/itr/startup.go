package itr

import (
	"fmt"

	"tgpl/ex_04.14-webserver/github"
)

var (
	VERBOSE = true
	NAME    string
	REPO    string
	TOKEN   string
)

// LoadCache sets the connection details and then loads existing data from
// localy saved files.
func LoadCache() (*Cache, error) {

	if VERBOSE {
		github.VERBOSE = true
	}

	// Retreive a pointer to the data store.
	cache := &data
	cache.Init()
	var err error

	NAME, REPO, TOKEN, err = connectDetails("connect")
	if err != nil {
		return cache, fmt.Errorf("connectionDetails: %v", err)
	}

	// Load cache if available.
	cache, err = loadCache(cache, NAME, REPO)
	if err != nil {
		return cache, fmt.Errorf("loadCache: %v", err)
	}

	if VERBOSE {
		fmt.Printf("LoadCache: cache loaded\n")
		fmt.Printf("issues: %.10d\nusers: %.10d\nlabels: %.10d\nmilestones: %.10d\n",
			len(cache.issues), len(cache.users), len(cache.labels), len(cache.milestones))
	}

	return cache, nil
}

// UpdateCache updates the cache adding new records downloaded from github.
func UpdateCache(cache *Cache) (*Cache, error) {

	cache, err := cache.GoUpdate()
	cache.GenerateLists()
	cache.SortByIdAsc()
	if VERBOSE {
		fmt.Printf("UpdateCache: cache updated\n")
		fmt.Printf("issues: %.10d\nusers: %.10d\nlabels: %.10d\nmilestones: %.10d\n",
			len(cache.issues), len(cache.users), len(cache.labels), len(cache.milestones))
	}
	return cache, err
}
