/*
Copyright 2019 Kohl's Department Stores, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package repository

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/miklezzzz/git2consul/config/mock"
	"github.com/miklezzzz/git2consul/repository/mocks"
	"github.com/stretchr/testify/assert"
	git "github.com/go-git/go-git/v5"
)

func TestCheckoutBranch(t *testing.T) {
	_, remotePath := mocks.InitRemote(t)
	defer os.RemoveAll(remotePath)

	repoConfig := mock.RepoConfig(remotePath)

	dstPath, err := ioutil.TempDir("", repoConfig.Name)
	assert.Nil(t, err)
	defer os.RemoveAll(dstPath)

	localRepo, err := git.PlainClone(dstPath, false, &git.CloneOptions{URL: repoConfig.URL})
	assert.Nil(t, err)

	repo := &Repository{
		Repository: localRepo,
		Config:     repoConfig,
	}

	branch := repo.Branch()

	err = repo.CheckoutBranch(branch)
	assert.Nil(t, err)
}
