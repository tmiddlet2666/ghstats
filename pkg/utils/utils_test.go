package utils

import (
	. "github.com/onsi/gomega"
	"github.com/tmiddlet2666/ghstats/pkg/config"
	"testing"
)

func TestHttpRequest(t *testing.T) {
	var (
		err    error
		data   []byte
		g      = NewGomegaWithT(t)
		status int
	)

	data, status, err = HttpGETRequest("https://github.com/")
	g.Expect(err).To(BeNil())
	g.Expect(status).To(Equal(200))
	g.Expect(len(data) >= 0).To(BeTrue())
}

func TestGetAPIURL(t *testing.T) {
	var (
		err    error
		data   []byte
		g      = NewGomegaWithT(t)
		status int
		URL    string
	)

	URL = GetAPIURL("oracle", "coherence")

	g.Expect(URL).To(Equal("https://api.github.com/repos/oracle/coherence"))

	data, status, err = HttpGETRequest(URL)
	g.Expect(err).To(BeNil())
	g.Expect(status).To(Equal(200))
	g.Expect(len(data) >= 0).To(BeTrue())
}

func TestGetReleases(t *testing.T) {
	var (
		g        = NewGomegaWithT(t)
		err      error
		releases []config.Release
	)

	releases, err = GetReleases("oracle", "coherence-visualvm")
	g.Expect(err).To(BeNil())
	g.Expect(len(releases) > 0).To(BeTrue())
}

func TestGetRepository(t *testing.T) {
	var (
		g    = NewGomegaWithT(t)
		err  error
		repo config.Repository
	)

	repo, err = GetRepoDetails("oracle", "coherence-visualvm")
	g.Expect(err).To(BeNil())
	g.Expect(repo).To(Not(BeNil()))
	g.Expect(repo.Stars > 0).To(BeTrue())
}
