package itr

import (
	"encoding/json"
	"fmt"
	"math/rand"
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
	ListIssues []uint64
}

type EnvUsers struct {
	Users     Users
	ListUsers []uint64
}

type EnvLabels struct {
	Labels     Labels
	ListLabels []uint64
}

type EnvMilestones struct {
	Milestones     Milestones
	ListMilestones []uint64
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
	c.listUsers = make([]uint64, 0, len(c.Users))
	for _, issue := range c.Users {
		c.listUsers = append(c.listUsers, issue.Id)
	}
	// Make label index
	c.listLabels = make([]uint64, 0, len(c.Labels))
	for _, issue := range c.Labels {
		c.listLabels = append(c.listLabels, issue.Id)
	}
	// Make milestone index
	c.listMilestones = make([]uint64, 0, len(c.Milestones))
	for _, issue := range c.Milestones {
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
	for i, id := range c.listIssues {
		fmt.Printf("%.3d: id: %v id:%v created: %v\n", i+1, id, c.Issues[id].Id, c.Issues[id].CreatedAt)
	}
}

// loadCache loads a repo's data from local cache files.
func loadCache(cache *Cache, name, repo string) (*Cache, bool, error) {

	pathIssues := "cache/" + name + "." + repo + ".Issues.json"
	pathUsers := "cache/" + name + "." + repo + ".Users.json"
	pathMilestones := "cache/" + name + "." + repo + ".Milestones.json"
	pathLabels := "cache/" + name + "." + repo + ".Labels.json"

	if VERBOSE {
		fmt.Printf("loadCache: checking for cache file\n")
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
