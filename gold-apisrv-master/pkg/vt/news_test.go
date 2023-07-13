package vt

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"apisrv/pkg/db"
	"apisrv/pkg/embedlog"
)

func TestDB_NewsService(t *testing.T) {
	Convey("Test NewsService", t, func() {
		ctx := context.Background()
		srv := NewNewsService(testDb, embedlog.Logger{})
		So(srv, ShouldNotBeNil)

		Convey("Positive testing", func() {

			Convey("Test CRUD", func() {

				// TestData

				inNews := News{
					Title:           "title",
					Alias:           "alias-1",
					CategoryID:      1,
					TagIDs:          []int{1, 2},
					StatusID:        db.StatusEnabled,
					CreatedAt:       time.Now(),
					PublicationDate: time.Now(),
				}

				// Add
				outNews, err := srv.Add(ctx, inNews)
				So(err, ShouldBeNil)
				So(outNews, ShouldNotBeNil)

				So(outNews.ID, ShouldBeGreaterThan, 0)
				So(outNews.Title, ShouldEqual, inNews.Title)
				So(outNews.Alias, ShouldEqual, inNews.Alias)

				// GetByID  outNews
				news, err := srv.GetByID(ctx, outNews.ID)
				So(err, ShouldBeNil)
				So(news, ShouldNotBeNil)

				So(news.ID, ShouldEqual, outNews.ID)
				So(news.Alias, ShouldEqual, outNews.Alias)
				So(news.Title, ShouldEqual, outNews.Title)
				// Update

				news.Title = "test"

				ok, err := srv.Update(ctx, *news)
				So(err, ShouldBeNil)
				So(ok, ShouldBeTrue)

				updated, err := srv.GetByID(ctx, news.ID)
				So(err, ShouldBeNil)
				So(updated.Title, ShouldEqual, news.Title)

				// Delete
				ok, err = srv.Delete(ctx, news.ID)
				So(err, ShouldBeNil)
				So(ok, ShouldBeTrue)
			})
		})

		Convey("Negative testing", func() {

			Convey("Create news with empty title", func() {
				news := News{
					Title:           "",
					Alias:           "alias-1",
					CategoryID:      1,
					TagIDs:          []int{1, 2},
					StatusID:        db.StatusEnabled,
					CreatedAt:       time.Now(),
					PublicationDate: time.Now(),
				}
				u, err := srv.Add(ctx, news)
				So(err, ShouldNotBeNil)
				So(u, ShouldBeNil)
			})

			Convey("Create news with empty Alias", func() {
				news := News{
					Title:           "title",
					Alias:           "",
					CategoryID:      1,
					TagIDs:          []int{1, 2},
					StatusID:        db.StatusEnabled,
					CreatedAt:       time.Now(),
					PublicationDate: time.Now(),
				}
				u, err := srv.Add(ctx, news)
				So(err, ShouldNotBeNil)
				So(u, ShouldBeNil)
			})

			Convey("Create news with duplicate alias", func() {
				news := News{
					Title:           "unique",
					Alias:           "unique-1",
					CategoryID:      1,
					TagIDs:          []int{1, 2},
					StatusID:        db.StatusEnabled,
					CreatedAt:       time.Now(),
					PublicationDate: time.Now(),
				}
				u, err := srv.Add(ctx, news)
				So(err, ShouldBeNil)
				So(u, ShouldNotBeNil)

				u2, err := srv.Add(ctx, news)
				So(err, ShouldNotBeNil)
				So(u2, ShouldBeNil)
			})
		})

	})
}

func TestDb_TagService(t *testing.T) {
	Convey("Test TagService", t, func() {
		ctx := context.Background()
		srv := NewTagService(testDb, embedlog.Logger{})
		So(srv, ShouldNotBeNil)

		Convey("Positive Testing", func() {
			Convey("Test CRUD", func() {
				inTag := Tag{
					Title:    "unqiue",
					StatusID: db.StatusEnabled,
				}
				// Add
				outTag, err := srv.Add(ctx, inTag)
				So(err, ShouldBeNil)
				So(outTag, ShouldNotBeNil)

				So(outTag.ID, ShouldBeGreaterThan, 0)
				So(outTag.Title, ShouldEqual, inTag.Title)

				// GetByID  outTag
				category, err := srv.GetByID(ctx, outTag.ID)
				So(err, ShouldBeNil)
				So(category, ShouldNotBeNil)

				So(category.ID, ShouldEqual, outTag.ID)
				So(category.Title, ShouldEqual, outTag.ID)

				// Update
				outTag.Title = "test"

				ok, err := srv.Update(ctx, *outTag)
				So(err, ShouldBeNil)
				So(ok, ShouldBeTrue)

				updated, err := srv.GetByID(ctx, outTag.ID)
				So(err, ShouldBeNil)
				So(updated.Title, ShouldEqual, outTag.Title)

				// Delete
				ok, err = srv.Delete(ctx, outTag.ID)
				So(err, ShouldBeNil)
				So(ok, ShouldBeTrue)

			})
		})
		Convey("Negative testing", func() {

			Convey("Create tag with empty title", func() {
				tag := Tag{
					Title:    "",
					StatusID: db.StatusEnabled,
				}
				u, err := srv.Add(ctx, tag)
				So(err, ShouldNotBeNil)
				So(u, ShouldBeNil)
			})
		})

	})
}

func TestDB_CategoryService(t *testing.T) {
	Convey("Test CategoryService", t, func() {
		ctx := context.Background()
		srv := NewCategoryService(testDb, embedlog.Logger{})
		So(srv, ShouldNotBeNil)

		Convey("Positive testing", func() {

			Convey("Test CRUD", func() {

				inCt := Category{
					Title:       "category-title",
					OrderNumber: 3,
					StatusID:    db.StatusEnabled,
				}

				// Add
				outCt, err := srv.Add(ctx, inCt)
				So(err, ShouldBeNil)
				So(outCt, ShouldNotBeNil)

				So(outCt.ID, ShouldBeGreaterThan, 0)
				So(outCt.Title, ShouldEqual, inCt.Title)

				// GetByID  outCategory
				category, err := srv.GetByID(ctx, outCt.ID)
				So(err, ShouldBeNil)
				So(category, ShouldNotBeNil)

				So(category.ID, ShouldEqual, outCt.ID)
				So(category.Title, ShouldEqual, outCt.Title)

				// Update
				outCt.Title = "test"

				ok, err := srv.Update(ctx, *outCt)
				So(err, ShouldBeNil)
				So(ok, ShouldBeTrue)

				updated, err := srv.GetByID(ctx, outCt.ID)
				So(err, ShouldBeNil)
				So(updated.Title, ShouldEqual, outCt.Title)

				// Delete
				ok, err = srv.Delete(ctx, outCt.ID)
				So(err, ShouldBeNil)
				So(ok, ShouldBeTrue)
			})
		})

		Convey("Negative testing", func() {

			Convey("Create category with empty title", func() {
				category := Category{
					Title:       "",
					OrderNumber: 3,
					StatusID:    db.StatusEnabled,
				}
				u, err := srv.Add(ctx, category)
				So(err, ShouldNotBeNil)
				So(u, ShouldBeNil)
			})
		})

	})
}
