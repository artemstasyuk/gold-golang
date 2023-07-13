//nolint:dupl
package vt

import (
	"time"

	"apisrv/pkg/db"
)

type Category struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required,max=256"`
	OrderNumber int    `json:"orderNumber" validate:"required"`
	StatusID    int    `json:"statusId" validate:"required,status"`

	Status *Status `json:"status"`
}

func (c *Category) ToDB() *db.Category {
	if c == nil {
		return nil
	}

	category := &db.Category{
		ID:          c.ID,
		Title:       c.Title,
		OrderNumber: c.OrderNumber,
		StatusID:    c.StatusID,
	}

	return category
}

type CategorySearch struct {
	ID          *int    `json:"id"`
	Title       *string `json:"title"`
	OrderNumber *int    `json:"orderNumber"`
	StatusID    *int    `json:"statusId"`
	IDs         []int   `json:"ids"`
}

func (cs *CategorySearch) ToDB() *db.CategorySearch {
	if cs == nil {
		return nil
	}

	return &db.CategorySearch{
		ID:          cs.ID,
		TitleILike:  cs.Title,
		OrderNumber: cs.OrderNumber,
		StatusID:    cs.StatusID,
		IDs:         cs.IDs,
	}
}

type CategorySummary struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	OrderNumber int    `json:"orderNumber"`

	Status *Status `json:"status"`
}

type News struct {
	ID              int        `json:"id"`
	Title           string     `json:"title" validate:"required,max=256"`
	Alias           string     `json:"alias" validate:"required,alias,max=256"`
	Content         *string    `json:"content"`
	CategoryID      int        `json:"categoryId" validate:"required"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
	PublicationDate time.Time  `json:"publicationDate" validate:"required"`
	TagIDs          []int      `json:"tagIds"`
	StatusID        int        `json:"statusId" validate:"required,status"`

	Category *CategorySummary `json:"category"`
	Status   *Status          `json:"status"`
}

func (n *News) ToDB() *db.News {
	if n == nil {
		return nil
	}

	news := &db.News{
		ID:              n.ID,
		Title:           n.Title,
		Alias:           n.Alias,
		Content:         n.Content,
		CategoryID:      n.CategoryID,
		CreatedAt:       n.CreatedAt,
		UpdatedAt:       n.UpdatedAt,
		PublicationDate: n.PublicationDate,
		TagIDs:          n.TagIDs,
		StatusID:        n.StatusID,
	}

	return news
}

type NewsSearch struct {
	ID              *int       `json:"id"`
	Title           *string    `json:"title"`
	Alias           *string    `json:"alias"`
	Content         *string    `json:"content"`
	CategoryID      *int       `json:"categoryId"`
	CreatedAt       *time.Time `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
	PublicationDate *time.Time `json:"publicationDate"`
	StatusID        *int       `json:"statusId"`
	IDs             []int      `json:"ids"`
	NotID           *int       `json:"notId"`
}

func (ns *NewsSearch) ToDB() *db.NewsSearch {
	if ns == nil {
		return nil
	}

	return &db.NewsSearch{
		ID:              ns.ID,
		TitleILike:      ns.Title,
		Alias:           ns.Alias,
		ContentILike:    ns.Content,
		CategoryID:      ns.CategoryID,
		CreatedAt:       ns.CreatedAt,
		UpdatedAt:       ns.UpdatedAt,
		PublicationDate: ns.PublicationDate,
		StatusID:        ns.StatusID,
		IDs:             ns.IDs,
		NotID:           ns.NotID,
	}
}

type NewsSummary struct {
	ID              int        `json:"id"`
	Title           string     `json:"title"`
	Alias           string     `json:"alias"`
	Content         *string    `json:"content"`
	CategoryID      int        `json:"categoryId"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       *time.Time `json:"updatedAt"`
	PublicationDate time.Time  `json:"publicationDate"`

	Category *CategorySummary `json:"category"`
	Status   *Status          `json:"status"`
}

type Tag struct {
	ID       int    `json:"id"`
	Title    string `json:"title" validate:"required,max=256"`
	StatusID int    `json:"statusId" validate:"required,status"`

	Status *Status `json:"status"`
}

func (t *Tag) ToDB() *db.Tag {
	if t == nil {
		return nil
	}

	tag := &db.Tag{
		ID:       t.ID,
		Title:    t.Title,
		StatusID: t.StatusID,
	}

	return tag
}

type TagSearch struct {
	ID       *int    `json:"id"`
	Title    *string `json:"title"`
	StatusID *int    `json:"statusId"`
	IDs      []int   `json:"ids"`
}

func (ts *TagSearch) ToDB() *db.TagSearch {
	if ts == nil {
		return nil
	}

	return &db.TagSearch{
		ID:         ts.ID,
		TitleILike: ts.Title,
		StatusID:   ts.StatusID,
		IDs:        ts.IDs,
	}
}

type TagSummary struct {
	ID    int    `json:"id"`
	Title string `json:"title"`

	Status *Status `json:"status"`
}
