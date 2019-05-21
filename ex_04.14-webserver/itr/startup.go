package itr

import (
	"fmt"

	"tgpl/ex_04.14-webserver/github"
)

var VERBOSE = true

func ServerStartup() error {

	if VERBOSE {
		github.VERBOSE = true
	}

	NAME, REPO, TOKEN, err := connectDetails("connect")
	if err != nil {
		return fmt.Errorf("connectionDetails: %v", err)
	}

	// Load cache if available.
	cache, isCache, err := loadCache(NAME, REPO)
	if err != nil {
		return fmt.Errorf("loadCache: %v", err)
	}

	// If there is a cache, check for updates.
	if isCache {
		return nil
	} else {
		data, err = github.GetAllIssues(NAME, REPO, TOKEN)
		if err != nil {
			return fmt.Errorf("GetAllIssues: %v", err)
		}
		err = writeCache(NAME, REPO, data)
	}
	return nil
}
