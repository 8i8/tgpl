package itr

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"tgpl/ex_04.14-webserver/github"
)

func (c *Cache) Update() (*Cache, error) {

	var err error
	var edited bool
	page := 0

	edited = true
	for edited {
		page++
		data, err := github.GetAllIssues(NAME, REPO, TOKEN, page)
		if err != nil {
			return c, fmt.Errorf("GetAllIssues: %v", err)
		}
		c, edited = deserialise(c, data)
	}
	err = writeCache(NAME, REPO, c)
	if err != nil {
		return c, fmt.Errorf("writeCache: %v", err)
	}
	return c, nil
}

func (c *Cache) GoUpdate() (*Cache, error) {

	rand.Seed(time.Now().UnixNano())
	ch := make(chan github.Bucket)
	queue := []github.Bucket{}
	var page, N int
	var edited bool

	tUnit := time.Duration(100)
	seconds := time.Duration(10)
	N = 10
	edited = true

	// TODO this code needs improvement, there is not yet any error
	// checking, the solution should perhaps be as follows:
	// Token buckets.
	// Detect errors, put them back into the queue.
	// Use exponential backoff; (After n collisions, a random number of
	// slot times between 0 and 2n − 1 is chosen).
	for edited {
		for i := 0; i < N; i++ {
			time.Sleep(tUnit * time.Millisecond)
			page++

			var result github.Bucket
			result.Page = page
			result.Count++

			go github.GoGetAllIssues(NAME, REPO, TOKEN, result, ch)
		}
		for i := 0; i < N; i++ {
			data := <-ch
			if data.Err == nil {
				c, edited = deserialise(c, data.Issues)
			} else {
				queue = append(queue, data)
			}
		}
		for len(queue) > 0 {
			l := len(queue)
			for i := 0; i < l; i++ {
				data := queue[0]
				queue = queue[1:len(queue)]
				data.Err = nil
				data.Count++
				time.Sleep(tUnit * time.Duration(rand.Intn(data.Count*2-1)+1) * time.Millisecond)
				go github.GoGetAllIssues(NAME, REPO, TOKEN, data, ch)
			}
			for i := 0; i < l; i++ {
				data := <-ch
				if data.Err == nil {
					c, edited = deserialise(c, data.Issues)
				} else {
					queue = append(queue, data)
				}
			}
		}
		err := writeCache(NAME, REPO, c)
		if err != nil {
			return c, fmt.Errorf("writeCache: %v", err)
		}
		time.Sleep(seconds * time.Second)
	}
	return c, nil
}

func (c *Cache) Report() {
	if VERBOSE {
		fmt.Println("PrintIssues: printing")
	}
	for i, id := range c.IssuesIndex {
		fmt.Printf("%.3d: id: %v id:%v created: %v\n", i+1, id, c.Issues[id].Id, c.Issues[id].CreatedAt)
	}
}

// loadCache loads a repo's data from local cache files.
func loadCache(cache *Cache, name, repo string) (*Cache, bool, error) {

	pathIssues := PATH + name + "." + repo + ".Issues.json"
	pathUsers := PATH + name + "." + repo + ".Users.json"
	pathMilestones := PATH + name + "." + repo + ".Milestones.json"
	pathLabels := PATH + name + "." + repo + ".Labels.json"

	if VERBOSE {
		fmt.Printf("loadCache: checking for cache file\n")
	}

	// Check that the cache directory is present, if not then make it.
	if _, err := os.Stat("cache"); err != nil {
		err = os.Mkdir("cache", 0755)
		if err != nil {
			return cache, false, fmt.Errorf("os.Mkdir: %v", err)
		}
	}

	// If a file is not present return.
	if _, err := os.Stat(pathIssues); err != nil {
		if VERBOSE {
			fmt.Printf("loadCache: cache not found\n")
		}
		return cache, false, nil
	}
	if _, err := os.Stat(pathUsers); err != nil {
		if VERBOSE {
			fmt.Printf("loadCache: cache not found\n")
		}
		return cache, false, nil
	}
	if _, err := os.Stat(pathLabels); err != nil {
		if VERBOSE {
			fmt.Printf("loadCache: cache not found\n")
		}
		return cache, false, nil
	}
	if _, err := os.Stat(pathMilestones); err != nil {
		if VERBOSE {
			fmt.Printf("loadCache: cache not found\n")
		}
		return cache, false, nil
	}

	if VERBOSE {
		fmt.Printf("loadCache: opening cache file for reading\n")
	}

	// Open issues
	file, err := os.Open(pathIssues)
	if err != nil {
		return cache, false, fmt.Errorf("os.Open: %v", err)
	}
	err = json.NewDecoder(file).Decode(&cache.Issues)
	if err != nil {
		return cache, false, fmt.Errorf("Decode: %v", err)
	}
	file.Close()

	// Open users
	file, err = os.Open(pathUsers)
	if err != nil {
		return cache, false, fmt.Errorf("os.Open: %v", err)
	}
	err = json.NewDecoder(file).Decode(&cache.Users)
	if err != nil {
		return cache, false, fmt.Errorf("Decode: %v", err)
	}
	file.Close()

	// Open labels
	file, err = os.Open(pathLabels)
	if err != nil {
		return cache, false, fmt.Errorf("os.Open: %v", err)
	}
	err = json.NewDecoder(file).Decode(&cache.Labels)
	if err != nil {
		return cache, false, fmt.Errorf("Decode: %v", err)
	}
	file.Close()

	// Open milestones
	file, err = os.Open(pathMilestones)
	if err != nil {
		return cache, false, fmt.Errorf("os.Open: %v", err)
	}
	err = json.NewDecoder(file).Decode(&cache.Milestones)
	if err != nil {
		return cache, false, fmt.Errorf("Decode: %v", err)
	}
	file.Close()

	if VERBOSE {
		fmt.Printf("loadCache: cache loaded\n")
	}

	return cache, true, nil
}

// writeCache writes the current cache data to file.
func writeCache(name, repo string, cache *Cache) error {

	if VERBOSE {
		fmt.Printf("                                                                                \r")
		fmt.Printf(" writeCache: opening cache file for writing\r")
	}

	// Write issues
	issues, err := json.MarshalIndent(cache.Issues, "", "	")
	if err != nil {
		return fmt.Errorf("MarshalIndent: %v", err)
	}
	file, err := os.Create("cache/" + name + "." + repo + ".Issues.json")
	if err != nil {
		return fmt.Errorf("os.Create: %v", err)
	}
	_, err = file.Write(issues)
	if err != nil {
		return fmt.Errorf("file.Write: %v", err)
	}
	file.Close()

	// Write users
	users, err := json.MarshalIndent(cache.Users, "", "	")
	if err != nil {
		return fmt.Errorf("MarshalIndent: %v", err)
	}
	file, err = os.Create("cache/" + name + "." + repo + ".Users.json")
	if err != nil {
		return fmt.Errorf("os.Create: %v", err)
	}
	_, err = file.Write(users)
	if err != nil {
		return fmt.Errorf("file.Write: %v", err)
	}
	file.Close()

	// Write labels
	labels, err := json.MarshalIndent(cache.Labels, "", "	")
	if err != nil {
		return fmt.Errorf("MarshalIndent: %v", err)
	}
	file, err = os.Create("cache/" + name + "." + repo + ".Labels.json")
	if err != nil {
		return fmt.Errorf("os.Create: %v", err)
	}
	_, err = file.Write(labels)
	if err != nil {
		return fmt.Errorf("file.Write: %v", err)
	}
	file.Close()

	// Write milestones
	milestones, err := json.MarshalIndent(cache.Milestones, "", "	")
	if err != nil {
		return fmt.Errorf("MarshalIndent: %v", err)
	}
	file, err = os.Create("cache/" + name + "." + repo + ".Milestones.json")
	if err != nil {
		return fmt.Errorf("os.Create: %v", err)
	}
	_, err = file.Write(milestones)
	if err != nil {
		return fmt.Errorf("file.Write: %v", err)
	}
	file.Close()

	if VERBOSE {
		fmt.Printf("                                                                               \r")
		fmt.Printf(" writeCache: cache written\r")
	}

	return nil
}

// deserialise transfers received github json data into the correct map form
// for the programs cache usage.
func deserialise(cache *Cache, issues []github.Issue) (*Cache, bool) {

	var mu sync.Mutex
	var mod bool
	mu.Lock()
	defer mu.Unlock()

	for _, issue := range issues {

		if cache.Issues[issue.Id].Id == 0 {
			mod = true
			var n Issue

			n.Id = issue.Id
			n.NodeId = issue.NodeId
			n.URL = issue.URL
			n.RepositoryURL = issue.RepositoryURL
			n.LablesURL = issue.LablesURL
			n.CommentsURL = issue.CommentsURL
			n.EventsURL = issue.EventsURL
			n.HtmlURL = issue.HtmlURL
			n.Number = issue.Number
			n.State = issue.State
			n.Title = issue.Title
			n.Body = issue.Body

			if issue.User.Id > 0 {
				n.User = issue.User.Id
				if cache.Users[issue.User.Id].Id == 0 {
					cache.Users[issue.User.Id] = User(issue.User)
				}
			}

			for _, label := range issue.Labels {
				n.Labels = append(n.Labels, label.Id)
				if cache.Labels[label.Id].Id == 0 {
					cache.Labels[label.Id] = Label(label)
				}
			}

			if issue.Assignee.Id > 0 {
				n.Assignee = issue.Assignee.Id
				if cache.Users[issue.Assignee.Id].Id == 0 {
					cache.Users[issue.Assignee.Id] = User(issue.Assignee)
				}
			}

			for _, user := range issue.Assignees {
				if user.Id > 0 {
					n.Assignees = append(n.Assignees, user.Id)
					if cache.Users[user.Id].Id == 0 {
						cache.Users[user.Id] = User(user)
					}
				}
			}

			if issue.Milestone.Id > 0 {
				n.Milestone = issue.Milestone.Id
				var ms Milestone

				ms.URL = issue.Milestone.URL
				ms.HtmlURL = issue.Milestone.HtmlURL
				ms.LabelsURL = issue.Milestone.LabelsURL
				ms.Id = issue.Milestone.Id
				ms.NodeId = issue.Milestone.NodeId
				ms.Number = issue.Milestone.Number
				ms.State = issue.Milestone.State
				ms.Title = issue.Milestone.Title
				ms.Description = issue.Milestone.Description

				if issue.Milestone.Creator.Id > 0 {
					ms.Creator = issue.Milestone.Creator.Id
					if cache.Users[issue.Milestone.Creator.Id].Id == 0 {
						cache.Users[issue.Milestone.Creator.Id] = User(issue.Milestone.Creator)
					}
				}

				ms.OpenIssues = issue.Milestone.OpenIssues
				ms.ClosedIssues = issue.Milestone.ClosedIssues
				ms.CreatedAt = issue.Milestone.CreatedAt
				ms.UpdatedAt = issue.Milestone.UpdatedAt
				ms.ClosedAt = issue.Milestone.ClosedAt
				ms.DueOn = issue.Milestone.DueOn

				if cache.Milestones[issue.Milestone.Id].Id == 0 {
					cache.Milestones[issue.Milestone.Id] = ms
				}
			}

			n.Locked = issue.Locked
			n.Reason = issue.Reason
			n.Comments = issue.Comments
			n.PullRequest = PullRequest(issue.PullRequest)
			n.ClosedAt = issue.ClosedAt
			n.CreatedAt = issue.CreatedAt
			n.UpdatedAt = issue.UpdatedAt
			n.ClosedBy = issue.ClosedBy.Id

			if issue.ClosedBy.Id > 0 {
				n.ClosedBy = issue.ClosedBy.Id
				if cache.Users[issue.ClosedBy.Id].Id == 0 {
					cache.Users[issue.ClosedBy.Id] = User(issue.ClosedBy)
				}
			}

			cache.Issues[issue.Id] = n
		}
	}

	return cache, mod
}

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

	tokens, err := github.ConnectDetails(PATH + "connect")
	if err != nil {
		return cache, fmt.Errorf("connectionDetails: %v", err)
	}
	if len(tokens) != 4 {
		return cache, fmt.Errorf("connectionDetails: insufficient tokens: %v", tokens)
	}
	NAME = tokens[0]
	REPO = tokens[1]
	TOKEN = tokens[2]

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
	if VERBOSE {
		fmt.Printf("UpdateCache: cache updated\n")
		fmt.Printf("issues: %.10d\nusers: %.10d\nlabels: %.10d\nmilestones: %.10d\n",
			len(cache.Issues), len(cache.Users), len(cache.Labels), len(cache.Milestones))
	}
	return cache, err
}
