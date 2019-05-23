package itr

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"sync"
	"time"

	"tgpl/ex_04.14-webserver/github"
)

type Issues map[uint64]Issue
type Users map[uint64]User
type Labels map[uint64]Label
type Milestones map[uint64]Milestone

// Cache is a struct that contains all the local cache data from a github
// repos list.
type Cache struct {
	EnvIssues
	EnvUsers
	EnvLabels
	EnvMilestones
}

type EnvIssues struct {
	Issues     Issues
	listIssues []uint64
}

type EnvUsers struct {
	Users     Users
	listUsers []uint64
}

type EnvLabels struct {
	Labels     Labels
	listLabels []uint64
}

type EnvMilestones struct {
	Milestones     Milestones
	listMilestones []uint64
}

func (c *Cache) Init() {
	c.Issues = make(Issues)
	c.Users = make(Users)
	c.Labels = make(Labels)
	c.Milestones = make(Milestones)
}

// Data contains the data used to run the site.
var Data Cache

// GenerateLists makes lists of all issue id's.
func (c *Cache) GenerateLists() {

	// Make issue index
	c.listIssues = make([]uint64, 0, len(c.Issues))
	for _, issue := range c.Issues {
		c.listIssues = append(c.listIssues, issue.Id)
	}
	// Make user index
	c.listUsers = make([]uint64, 0, len(c.users))
	for _, issue := range c.users {
		c.listUsers = append(c.listUsers, issue.Id)
	}
	// Make label index
	c.listLabels = make([]uint64, 0, len(c.labels))
	for _, issue := range c.labels {
		c.listLabels = append(c.listLabels, issue.Id)
	}
	// Make milestone index
	c.listMilestones = make([]uint64, 0, len(c.milestones))
	for _, issue := range c.milestones {
		c.listMilestones = append(c.listMilestones, issue.Id)
	}
}

// SortByIdDes makes lists of all issue id's.
func (c *Cache) SortByIdDes() {

	// Sort issues by id
	sort.Slice(c.listIssues, func(i, j int) bool {
		return c.listIssues[i] > c.listIssues[j]
	})
	// Sort users by id
	sort.Slice(c.listUsers, func(i, j int) bool {
		return c.listUsers[i] > c.listUsers[j]
	})
	// Sort labels by id
	sort.Slice(c.listLabels, func(i, j int) bool {
		return c.listLabels[i] > c.listLabels[j]
	})
	// Sort milestones by id
	sort.Slice(c.listMilestones, func(i, j int) bool {
		return c.listMilestones[i] > c.listMilestones[j]
	})
}

func (c *Cache) SortByIdAsc() {

	// Sort issues by id
	sort.Slice(c.listIssues, func(i, j int) bool {
		return c.listIssues[i] < c.listIssues[j]
	})
	// Sort users by id
	sort.Slice(c.listUsers, func(i, j int) bool {
		return c.listUsers[i] < c.listUsers[j]
	})
	// Sort labels by id
	sort.Slice(c.listLabels, func(i, j int) bool {
		return c.listLabels[i] < c.listLabels[j]
	})
	// Sort milestones by id
	sort.Slice(c.listMilestones, func(i, j int) bool {
		return c.listMilestones[i] < c.listMilestones[j]
	})
}

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

	var err error
	var edited bool
	var modified bool
	var N int
	ch := make(chan []github.Issue)
	page := 0

	N = 1000
	count := 0

	edited = true

	// TODO this code needs improvement, there is not yet any error
	// checking, the solution should be as follows:
	// token buckets
	// detect errors, put them back into the queue, use exponential
	// backoff.
	for edited {
		for i := 0; i < N; i++ {
			time.Sleep(100 * time.Millisecond)
			page++
			go github.GoGetAllIssues(NAME, REPO, TOKEN, page, ch)
		}
		for i := 0; i < N; i++ {
			data := <-ch
			c, modified = deserialise(c, data)
			if modified {
				count++
			}
		}
		err = writeCache(NAME, REPO, c)
		if err != nil {
			return c, fmt.Errorf("writeCache: %v", err)
		}
		if count < N {
			edited = false
		}
		count = 0
	}
	close(ch)
	return c, nil
}

func (c *Cache) Report() {
	if VERBOSE {
		fmt.Println("PrintIssues: printing")
	}
	for i, id := range c.listIssues {
		fmt.Printf("%.3d: id: %v id:%v created: %v\n", i+1, id, c.issues[id].Id, c.issues[id].CreatedAt)
	}
}

// loadCache loads a repo's data from local cache files.
func loadCache(cache *Cache, name, repo string) (*Cache, error) {

	pathIssues := "cache/" + name + "." + repo + ".issues.json"
	pathUsers := "cache/" + name + "." + repo + ".users.json"
	pathMilestones := "cache/" + name + "." + repo + ".milestones.json"
	pathLabels := "cache/" + name + "." + repo + ".labels.json"

	if VERBOSE {
		fmt.Printf("loadCache: checking for cache file\n")
	}

	// If a file is not present return.
	if _, err := os.Stat(pathIssues); err != nil {
		if VERBOSE {
			fmt.Printf("loadCache: cache not found\n")
		}
		return cache, nil
	}
	if _, err := os.Stat(pathUsers); err != nil {
		if VERBOSE {
			fmt.Printf("loadCache: cache not found\n")
		}
		return cache, nil
	}
	if _, err := os.Stat(pathLabels); err != nil {
		if VERBOSE {
			fmt.Printf("loadCache: cache not found\n")
		}
		return cache, nil
	}
	if _, err := os.Stat(pathMilestones); err != nil {
		if VERBOSE {
			fmt.Printf("loadCache: cache not found\n")
		}
		return cache, nil
	}

	if VERBOSE {
		fmt.Printf("loadCache: opening cache file for reading\n")
	}

	// Open issues
	file, err := os.Open(pathIssues)
	if err != nil {
		return cache, fmt.Errorf("os.Open: %v", err)
	}
	err = json.NewDecoder(file).Decode(&cache.issues)
	if err != nil {
		return cache, fmt.Errorf("Decode: %v", err)
	}
	file.Close()

	// Open users
	file, err = os.Open(pathUsers)
	if err != nil {
		return cache, fmt.Errorf("os.Open: %v", err)
	}
	err = json.NewDecoder(file).Decode(&cache.users)
	if err != nil {
		return cache, fmt.Errorf("Decode: %v", err)
	}
	file.Close()

	// Open labels
	file, err = os.Open(pathLabels)
	if err != nil {
		return cache, fmt.Errorf("os.Open: %v", err)
	}
	err = json.NewDecoder(file).Decode(&cache.labels)
	if err != nil {
		return cache, fmt.Errorf("Decode: %v", err)
	}
	file.Close()

	// Open milestones
	file, err = os.Open(pathMilestones)
	if err != nil {
		return cache, fmt.Errorf("os.Open: %v", err)
	}
	err = json.NewDecoder(file).Decode(&cache.milestones)
	if err != nil {
		return cache, fmt.Errorf("Decode: %v", err)
	}
	file.Close()

	if VERBOSE {
		fmt.Printf("loadCache: cache loaded\n")
	}

	return cache, nil
}

// writeCache writes the current cache data to file.
func writeCache(name, repo string, cache *Cache) error {

	if VERBOSE {
		fmt.Printf("                                                                                \r")
		fmt.Printf(" writeCache: opening cache file for writing\r")
	}

	// Write issues
	issues, err := json.MarshalIndent(cache.issues, "", "	")
	if err != nil {
		return fmt.Errorf("MarshalIndent: %v", err)
	}
	file, err := os.Create("cache/" + name + "." + repo + ".issues.json")
	if err != nil {
		return fmt.Errorf("os.Create: %v", err)
	}
	_, err = file.Write(issues)
	if err != nil {
		return fmt.Errorf("file.Write: %v", err)
	}
	file.Close()

	// Write users
	users, err := json.MarshalIndent(cache.users, "", "	")
	if err != nil {
		return fmt.Errorf("MarshalIndent: %v", err)
	}
	file, err = os.Create("cache/" + name + "." + repo + ".users.json")
	if err != nil {
		return fmt.Errorf("os.Create: %v", err)
	}
	_, err = file.Write(users)
	if err != nil {
		return fmt.Errorf("file.Write: %v", err)
	}
	file.Close()

	// Write labels
	labels, err := json.MarshalIndent(cache.labels, "", "	")
	if err != nil {
		return fmt.Errorf("MarshalIndent: %v", err)
	}
	file, err = os.Create("cache/" + name + "." + repo + ".labels.json")
	if err != nil {
		return fmt.Errorf("os.Create: %v", err)
	}
	_, err = file.Write(labels)
	if err != nil {
		return fmt.Errorf("file.Write: %v", err)
	}
	file.Close()

	// Write milestones
	milestones, err := json.MarshalIndent(cache.milestones, "", "	")
	if err != nil {
		return fmt.Errorf("MarshalIndent: %v", err)
	}
	file, err = os.Create("cache/" + name + "." + repo + ".milestones.json")
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

	var mod bool
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	for _, issue := range issues {

		if cache.issues[issue.Id].Id == 0 {
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
				if cache.users[issue.User.Id].Id == 0 {
					cache.users[issue.User.Id] = User(issue.User)
				}
			}

			for _, label := range issue.Labels {
				n.Labels = append(n.Labels, label.Id)
				if cache.labels[label.Id].Id == 0 {
					cache.labels[label.Id] = Label(label)
				}
			}

			if issue.Assignee.Id > 0 {
				n.Assignee = issue.Assignee.Id
				if cache.users[issue.Assignee.Id].Id == 0 {
					cache.users[issue.Assignee.Id] = User(issue.Assignee)
				}
			}

			for _, user := range issue.Assignees {
				if user.Id > 0 {
					n.Assignees = append(n.Assignees, user.Id)
					if cache.users[user.Id].Id == 0 {
						cache.users[user.Id] = User(user)
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
					if cache.users[issue.Milestone.Creator.Id].Id == 0 {
						cache.users[issue.Milestone.Creator.Id] = User(issue.Milestone.Creator)
					}
				}

				ms.OpenIssues = issue.Milestone.OpenIssues
				ms.ClosedIssues = issue.Milestone.ClosedIssues
				ms.CreatedAt = issue.Milestone.CreatedAt
				ms.UpdatedAt = issue.Milestone.UpdatedAt
				ms.ClosedAt = issue.Milestone.ClosedAt
				ms.DueOn = issue.Milestone.DueOn

				if cache.milestones[issue.Milestone.Id].Id == 0 {
					cache.milestones[issue.Milestone.Id] = ms
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
				if cache.users[issue.ClosedBy.Id].Id == 0 {
					cache.users[issue.ClosedBy.Id] = User(issue.ClosedBy)
				}
			}

			cache.issues[issue.Id] = n
		}
	}

	return cache, mod
}
