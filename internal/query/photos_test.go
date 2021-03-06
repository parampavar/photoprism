package query

import (
	"github.com/photoprism/photoprism/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/photoprism/photoprism/internal/form"
)

func TestPhotos(t *testing.T) {
	t.Run("normal query", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = ""
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		//t.Logf("results: %+v", photos)
		assert.LessOrEqual(t, 4, len(photos))
		for _, r := range photos {
			assert.IsType(t, PhotosResult{}, r)
			assert.NotEmpty(t, r.ID)
			assert.NotEmpty(t, r.CameraID)
			assert.NotEmpty(t, r.LensID)

			if fix, ok := entity.PhotoFixtures[r.PhotoName]; ok {
				assert.Equal(t, fix.PhotoName, r.PhotoName)
			}
		}
	})
	t.Run("label query dog", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "label:dog"
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)

		assert.Equal(t, "label dog not found", err.Error())
		assert.Empty(t, photos)
		//t.Logf("results: %+v", photos)
	})
	t.Run("label query landscape", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "label:landscape Order:relevance"
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)
		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 2, len(photos))
	})
	t.Run("invalid label query", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "label:xxx"
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)

		assert.Error(t, err)
		assert.Empty(t, photos)

		if err != nil {
			assert.Equal(t, err.Error(), "label xxx not found")
		}
	})
	t.Run("form.location true", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = ""
		f.Count = 10
		f.Offset = 0
		f.Location = true

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 3, len(photos))

	})
	t.Run("form.camera", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = ""
		f.Count = 10
		f.Offset = 0
		f.Camera = 1000003

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 4, len(photos))
	})
	t.Run("form.color", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = ""
		f.Count = 3
		f.Offset = 0
		f.Color = "blue"

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 2, len(photos))
	})
	t.Run("form.favorites", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "favorites:true"
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 1, len(photos))
	})
	t.Run("form.country", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "country:zz"
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 3, len(photos))

	})
	t.Run("form.title", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "title:Neckarbrücke"
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		//t.Logf("results: %+v", photos)
		assert.Equal(t, 1, len(photos))

	})
	t.Run("form.hash", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "hash:2cad9168fa6acc5c5c2965ddf6ec465ca42fd818"
		f.Count = 3
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		//t.Logf("results: %+v", photos)
		assert.Equal(t, 1, len(photos))
	})
	t.Run("form.duplicate", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "duplicate:true"
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 1, len(photos))

	})
	t.Run("form.portrait", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "portrait:true"
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 1, len(photos))

	})
	t.Run("form.mono", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "mono:false"
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 4, len(photos))
	})
	t.Run("form.chroma >9 Order:similar", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "chroma:25 Order:similar"
		f.Count = 3
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 2, len(photos))

	})
	t.Run("form.chroma <9", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "chroma:4"
		f.Count = 3
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 1, len(photos))

	})
	t.Run("form.fmin and Order:oldest", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "Fmin:5 Order:oldest"
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 1, len(photos))
	})
	t.Run("form.fmax and Order:newest", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "Fmax:2 Order:newest"
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}

		assert.LessOrEqual(t, 3, len(photos))

	})
	t.Run("form.Lat and form.Lng and Order:imported", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "Lat:33.45343166666667 Lng:25.764711666666667 Dist:2000 Order:imported"
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.LessOrEqual(t, 2, len(photos))

	})
	t.Run("form.Lat and form.Lng and Order:imported Dist:6000", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "Lat:33.45343166666667 Lng:25.764711666666667 Dist:6000 Order:imported"
		f.Count = 10
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.LessOrEqual(t, 2, len(photos))

	})
	t.Run("form.Before and form.After Order:relevance", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "Before:2016-01-01 After:2013-01-01 Order:relevance"
		f.Count = 5000
		f.Offset = 0
		f.Merged = true

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.LessOrEqual(t, 3, len(photos))
	})

	t.Run("search for diff", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = "Diff:800"
		f.Count = 5000
		f.Offset = 0

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.LessOrEqual(t, 1, len(photos))
	})
	t.Run("search for lens, month, year, album", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = ""
		f.Count = 5000
		f.Offset = 0
		f.Lens = 1000000
		f.Month = 2
		f.Year = 2790
		f.Album = "at9lxuqxpogaaba8"

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.LessOrEqual(t, 1, len(photos))
	})
	t.Run("search for private, archived, review", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = ""
		f.Count = 5000
		f.Offset = 0
		f.Private = true
		f.Archived = true
		f.Review = true

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Empty(t, photos)
	})
	t.Run("search for archived and public", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = ""
		f.Count = 5000
		f.Offset = 0
		f.Archived = true
		f.Public = true

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.Empty(t, photos)
		//TODO create test fixture
	})
	t.Run("search for ID", func(t *testing.T) {
		var f form.PhotoSearch
		f.Query = ""
		f.Count = 5000
		f.Offset = 0
		f.ID = "pt9jtdre2lvl0yh7"
		f.Merged = true

		photos, _, err := Photos(f)

		if err != nil {
			t.Fatal(err)
		}
		assert.LessOrEqual(t, 1, len(photos))
	})
}
