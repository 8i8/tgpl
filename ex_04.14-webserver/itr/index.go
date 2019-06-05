package itr

import (
	"sort"
	"strings"
)

// GenerateIndex makes lists of all issue id's.
func (c *Cache) GenerateIndex() {

	// Make issue index
	c.IssuesIndex = make([]uint64, 0, len(c.Issues))
	for _, issue := range c.Issues {
		c.IssuesIndex = append(c.IssuesIndex, issue.Id)
	}
	// Make user index
	c.UsersIndex = make([]uint64, 0, len(c.Users))
	for _, issue := range c.Users {
		c.UsersIndex = append(c.UsersIndex, issue.Id)
	}
	// Make label index
	c.LabelsIndex = make([]uint64, 0, len(c.Labels))
	for _, issue := range c.Labels {
		c.LabelsIndex = append(c.LabelsIndex, issue.Id)
	}
	// Make milestone index
	c.MilestonesIndex = make([]uint64, 0, len(c.Milestones))
	for _, issue := range c.Milestones {
		c.MilestonesIndex = append(c.MilestonesIndex, issue.Id)
	}
}

// SortByIdDes makes lists of all issue id's.
func (c *Cache) SortByIdDes() {

	// Sort issues by id
	sort.Slice(c.IssuesIndex, func(i, j int) bool {
		return c.IssuesIndex[i] > c.IssuesIndex[j]
	})
	// Sort users by Login name
	sort.Slice(c.UsersIndex, func(i, j int) bool {
		if strings.Compare(c.Users[c.UsersIndex[i]].Login, c.Users[c.UsersIndex[j]].Login) > 0 {
			return true
		}
		return false
	})
	// Sort labels by id
	sort.Slice(c.LabelsIndex, func(i, j int) bool {
		return c.LabelsIndex[i] > c.LabelsIndex[j]
	})
	// Sort milestones by id
	sort.Slice(c.MilestonesIndex, func(i, j int) bool {
		return c.MilestonesIndex[i] > c.MilestonesIndex[j]
	})
}

func (c *Cache) SortByIdAsc() {

	// Sort issues by id
	sort.Slice(c.IssuesIndex, func(i, j int) bool {
		return c.IssuesIndex[i] < c.IssuesIndex[j]
	})
	// Sort users by Login name
	sort.Slice(c.UsersIndex, func(i, j int) bool {
		if strings.Compare(c.Users[c.UsersIndex[i]].Login, c.Users[c.UsersIndex[j]].Login) < 0 {
			return true
		}
		return false
	})
	// Sort labels by id
	sort.Slice(c.LabelsIndex, func(i, j int) bool {
		return c.LabelsIndex[i] < c.LabelsIndex[j]
	})
	// Sort milestones by id
	sort.Slice(c.MilestonesIndex, func(i, j int) bool {
		return c.MilestonesIndex[i] < c.MilestonesIndex[j]
	})
}
