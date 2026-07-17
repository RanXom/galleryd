package gallery

import (
	"strconv"
	"testing"
)

func makePhotos(n int) []Photo {
	photos := make([]Photo, n)

	for i := range n {
		photos[i] = Photo{
			ID: strconv.Itoa(i),
		}
	}

	return photos
}

func assertIDs(
	t *testing.T,
	photos []Photo,
	want ...string,
) {
	t.Helper()

	if len(photos) != len(want) {
		t.Fatalf(
			"expected %d photos, got %d",
			len(want),
			len(photos),
		)
	}

	for i := range want {
		if photos[i].ID != want[i] {
			t.Fatalf(
				"photo %d: expected %q, got %q",
				i,
				want[i],
				photos[i].ID,
			)
		}
	}
}

func TestQueryPhotos(t *testing.T) {
	photos := makePhotos(10)

	t.Run("default", func(t *testing.T) {
		got := QueryPhotos(photos, Query{})

		assertIDs(
			t,
			got,
			"0", "1", "2", "3", "4",
			"5", "6", "7", "8", "9",
		)
	})

	t.Run("limit", func(t *testing.T) {
		got := QueryPhotos(
			photos,
			Query{
				Limit: 3,
			},
		)

		assertIDs(t, got, "0", "1", "2")
	})

	t.Run("offset", func(t *testing.T) {
		got := QueryPhotos(
			photos,
			Query{
				Offset: 5,
			},
		)

		assertIDs(
			t,
			got,
			"5", "6", "7", "8", "9",
		)
	})

	t.Run("offset and limit", func(t *testing.T) {
		got := QueryPhotos(
			photos,
			Query{
				Offset: 5,
				Limit:  2,
			},
		)

		assertIDs(t, got, "5", "6")
	})

	t.Run("offset beyond end", func(t *testing.T) {
		got := QueryPhotos(
			photos,
			Query{
				Offset: 100,
			},
		)

		assertIDs(t, got)
	})

	t.Run("limit larger than gallery", func(t *testing.T) {
		got := QueryPhotos(
			photos,
			Query{
				Limit: 100,
			},
		)

		assertIDs(
			t,
			got,
			"0", "1", "2", "3", "4",
			"5", "6", "7", "8", "9",
		)
	})

	t.Run("ignores sorting for now", func(t *testing.T) {
		got := QueryPhotos(
			photos,
			Query{
				Sort:  SortByPath,
				Order: SortAsc,
			},
		)

		assertIDs(
			t,
			got,
			"0", "1", "2", "3", "4",
			"5", "6", "7", "8", "9",
		)
	})
}
