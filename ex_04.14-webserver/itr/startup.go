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
	cache := &Data
	cache.Init()
	var err error
	var ok bool

	NAME, REPO, TOKEN, err = github.ConnectDetails("connect")
	if err != nil {
		return cache, fmt.Errorf("connectionDetails: %v", err)
	}

	// Load cache if available.
	cache, ok, err = loadCache(cache, NAME, REPO)
	if err != nil {
		return cache, fmt.Errorf("loadCache: %v", err)
	}

	if VERBOSE && ok {
		fmt.Printf("issues: %.10d\nusers: %.10d\nlabels: %.10d\nmilestones: %.10d\n",
			len(cache.Issues), len(cache.Users), len(cache.Labels), len(cache.Milestones))
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
			len(cache.Issues), len(cache.Users), len(cache.Labels), len(cache.Milestones))
	}
	return cache, err
}
