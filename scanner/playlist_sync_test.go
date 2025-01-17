package scanner

import (
	"context"

	"github.com/navidrome/navidrome/model"
	"github.com/navidrome/navidrome/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("playlistSync", func() {
	Describe("parsePlaylist", func() {
		var ds model.DataStore
		var ps *playlistSync
		ctx := context.TODO()
		BeforeEach(func() {
			ds = &tests.MockDataStore{
				MockedMediaFile: &mockedMediaFile{},
			}
			ps = newPlaylistSync(ds)
		})

		It("parses well-formed playlists", func() {
			pls, err := ps.parsePlaylist(ctx, "playlists/pls1.m3u", "tests/fixtures")
			Expect(err).To(BeNil())
			Expect(pls.Tracks).To(HaveLen(3))
			Expect(pls.Tracks[0].Path).To(Equal("tests/fixtures/test.mp3"))
			Expect(pls.Tracks[1].Path).To(Equal("tests/fixtures/test.ogg"))
			Expect(pls.Tracks[2].Path).To(Equal("/tests/fixtures/01 Invisible (RED) Edit Version.mp3"))
		})

		It("parses playlists using LF ending", func() {
			pls, err := ps.parsePlaylist(ctx, "lf-ended.m3u", "tests/fixtures/playlists")
			Expect(err).To(BeNil())
			Expect(pls.Tracks).To(HaveLen(2))
		})

		It("parses playlists using CR ending (old Mac format)", func() {
			pls, err := ps.parsePlaylist(ctx, "cr-ended.m3u", "tests/fixtures/playlists")
			Expect(err).To(BeNil())
			Expect(pls.Tracks).To(HaveLen(2))
		})
	})
})

type mockedMediaFile struct {
	model.MediaFileRepository
}

func (r *mockedMediaFile) FindByPath(s string) (*model.MediaFile, error) {
	return &model.MediaFile{
		ID:   "123",
		Path: s,
	}, nil
}
