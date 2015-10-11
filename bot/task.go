// This file is part of the Plan-B-ot package.
// Copyright (c) 2015 Martin Schenck
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package bot

import (
	"bytes"
	"strconv"
)

var currentTask task

// Task has a name and all the associated votes
type task struct {
	Name  string
	Votes []vote
}

// Vote is one vote on a task
type vote struct {
	Username string
	Vote     float64
}

// Vote adds a vote to the current task
func (t *task) Vote(username string, userVote string) error {
	floatVoat, err := strconv.ParseFloat(userVote, 64)
	if err != nil {
		return err
	}

	existingVote := t.getExistingVoteForUser(username)
	if existingVote != nil {
		existingVote.Vote = floatVoat
	} else {
		newVote := vote{
			Username: username,
			Vote:     floatVoat,
		}

		t.Votes = append(t.Votes, newVote)
	}

	return nil
}

// getExistingVoteForUser returns a pointer to an existing vote, if any
func (t *task) getExistingVoteForUser(userName string) *vote {
	for index, vote := range t.Votes {
		if userName == vote.Username {
			return &t.Votes[index]
		}
	}

	return nil
}

// getVoters returns a string representation of all user's names of all
// currently active votes
func (t *task) getVoters() string {
	var buffer bytes.Buffer

	for index, vote := range currentTask.Votes {
		if index > 0 {
			buffer.WriteString(", ")
		}
		buffer.WriteString("@")
		buffer.WriteString(vote.Username)
	}

	return buffer.String()
}
