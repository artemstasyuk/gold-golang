package db

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type NewsRepo struct {
	db      orm.DB
	filters map[string][]Filter
	sort    map[string][]SortField
	join    map[string][]string
}

// NewNewsRepo returns new repository
func NewNewsRepo(db orm.DB) NewsRepo {
	return NewsRepo{
		db: db,
		filters: map[string][]Filter{
			Tables.Category.Name: {StatusFilter},
			Tables.News.Name:     {StatusFilter},
			Tables.Tag.Name:      {StatusFilter},
		},
		sort: map[string][]SortField{
			Tables.Category.Name: {{Column: Columns.Category.Title, Direction: SortAsc}},
			Tables.News.Name:     {{Column: Columns.News.CreatedAt, Direction: SortDesc}},
			Tables.Tag.Name:      {{Column: Columns.Tag.Title, Direction: SortAsc}},
		},
		join: map[string][]string{
			Tables.Category.Name: {TableColumns},
			Tables.News.Name:     {TableColumns, Columns.News.Category},
			Tables.Tag.Name:      {TableColumns},
		},
	}
}

// WithTransaction is a function that wraps NewsRepo with pg.Tx transaction.
func (nr NewsRepo) WithTransaction(tx *pg.Tx) NewsRepo {
	nr.db = tx
	return nr
}

// WithEnabledOnly is a function that adds "statusId"=1 as base filter.
func (nr NewsRepo) WithEnabledOnly() NewsRepo {
	f := make(map[string][]Filter, len(nr.filters))
	for i := range nr.filters {
		f[i] = make([]Filter, len(nr.filters[i]))
		copy(f[i], nr.filters[i])
		f[i] = append(f[i], StatusEnabledFilter)
	}
	nr.filters = f

	return nr
}

/*** Category ***/

// FullCategory returns full joins with all columns
func (nr NewsRepo) FullCategory() OpFunc {
	return WithColumns(nr.join[Tables.Category.Name]...)
}

// DefaultCategorySort returns default sort.
func (nr NewsRepo) DefaultCategorySort() OpFunc {
	return WithSort(nr.sort[Tables.Category.Name]...)
}

// CategoryByID is a function that returns Category by ID(s) or nil.
func (nr NewsRepo) CategoryByID(ctx context.Context, id int, ops ...OpFunc) (*Category, error) {
	return nr.OneCategory(ctx, &CategorySearch{ID: &id}, ops...)
}

// OneCategory is a function that returns one Category by filters. It could return pg.ErrMultiRows.
func (nr NewsRepo) OneCategory(ctx context.Context, search *CategorySearch, ops ...OpFunc) (*Category, error) {
	obj := &Category{}
	err := buildQuery(ctx, nr.db, obj, search, nr.filters[Tables.Category.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// CategoriesByFilters returns Category list.
func (nr NewsRepo) CategoriesByFilters(ctx context.Context, search *CategorySearch, pager Pager, ops ...OpFunc) (categories []Category, err error) {
	err = buildQuery(ctx, nr.db, &categories, search, nr.filters[Tables.Category.Name], pager, ops...).Select()
	return
}

// CountCategories returns count
func (nr NewsRepo) CountCategories(ctx context.Context, search *CategorySearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, nr.db, &Category{}, search, nr.filters[Tables.Category.Name], PagerOne, ops...).Count()
}

// AddCategory adds Category to DB.
func (nr NewsRepo) AddCategory(ctx context.Context, category *Category, ops ...OpFunc) (*Category, error) {
	q := nr.db.ModelContext(ctx, category)
	applyOps(q, ops...)
	_, err := q.Insert()

	return category, err
}

// UpdateCategory updates Category in DB.
func (nr NewsRepo) UpdateCategory(ctx context.Context, category *Category, ops ...OpFunc) (bool, error) {
	q := nr.db.ModelContext(ctx, category).WherePK()
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteCategory set statusId to deleted in DB.
func (nr NewsRepo) DeleteCategory(ctx context.Context, id int) (deleted bool, err error) {
	category := &Category{ID: id, StatusID: StatusDeleted}

	return nr.UpdateCategory(ctx, category, WithColumns(Columns.Category.StatusID))
}

/*** News ***/

// FullNews returns full joins with all columns
func (nr NewsRepo) FullNews() OpFunc {
	return WithColumns(nr.join[Tables.News.Name]...)
}

// DefaultNewsSort returns default sort.
func (nr NewsRepo) DefaultNewsSort() OpFunc {
	return WithSort(nr.sort[Tables.News.Name]...)
}

// NewsByID is a function that returns News by ID(s) or nil.
func (nr NewsRepo) NewsByID(ctx context.Context, id int, ops ...OpFunc) (*News, error) {
	return nr.OneNews(ctx, &NewsSearch{ID: &id}, ops...)
}

// OneNews is a function that returns one News by filters. It could return pg.ErrMultiRows.
func (nr NewsRepo) OneNews(ctx context.Context, search *NewsSearch, ops ...OpFunc) (*News, error) {
	obj := &News{}
	err := buildQuery(ctx, nr.db, obj, search, nr.filters[Tables.News.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// NewsByFilters returns News list.
func (nr NewsRepo) NewsByFilters(ctx context.Context, search *NewsSearch, pager Pager, ops ...OpFunc) (newsList []News, err error) {
	err = buildQuery(ctx, nr.db, &newsList, search, nr.filters[Tables.News.Name], pager, ops...).Select()
	return
}

// CountNews returns count
func (nr NewsRepo) CountNews(ctx context.Context, search *NewsSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, nr.db, &News{}, search, nr.filters[Tables.News.Name], PagerOne, ops...).Count()
}

// AddNews adds News to DB.
func (nr NewsRepo) AddNews(ctx context.Context, news *News, ops ...OpFunc) (*News, error) {
	q := nr.db.ModelContext(ctx, news)
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.News.CreatedAt)
	}
	applyOps(q, ops...)
	_, err := q.Insert()

	return news, err
}

// UpdateNews updates News in DB.
func (nr NewsRepo) UpdateNews(ctx context.Context, news *News, ops ...OpFunc) (bool, error) {
	q := nr.db.ModelContext(ctx, news).WherePK()
	if len(ops) == 0 {
		q = q.ExcludeColumn(Columns.News.CreatedAt)
	}
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteNews set statusId to deleted in DB.
func (nr NewsRepo) DeleteNews(ctx context.Context, id int) (deleted bool, err error) {
	news := &News{ID: id, StatusID: StatusDeleted}

	return nr.UpdateNews(ctx, news, WithColumns(Columns.News.StatusID))
}

/*** Tag ***/

// FullTag returns full joins with all columns
func (nr NewsRepo) FullTag() OpFunc {
	return WithColumns(nr.join[Tables.Tag.Name]...)
}

// DefaultTagSort returns default sort.
func (nr NewsRepo) DefaultTagSort() OpFunc {
	return WithSort(nr.sort[Tables.Tag.Name]...)
}

// TagByID is a function that returns Tag by ID(s) or nil.
func (nr NewsRepo) TagByID(ctx context.Context, id int, ops ...OpFunc) (*Tag, error) {
	return nr.OneTag(ctx, &TagSearch{ID: &id}, ops...)
}

// OneTag is a function that returns one Tag by filters. It could return pg.ErrMultiRows.
func (nr NewsRepo) OneTag(ctx context.Context, search *TagSearch, ops ...OpFunc) (*Tag, error) {
	obj := &Tag{}
	err := buildQuery(ctx, nr.db, obj, search, nr.filters[Tables.Tag.Name], PagerTwo, ops...).Select()

	if errors.Is(err, pg.ErrMultiRows) {
		return nil, err
	} else if errors.Is(err, pg.ErrNoRows) {
		return nil, nil
	}

	return obj, err
}

// TagsByFilters returns Tag list.
func (nr NewsRepo) TagsByFilters(ctx context.Context, search *TagSearch, pager Pager, ops ...OpFunc) (tags []Tag, err error) {
	err = buildQuery(ctx, nr.db, &tags, search, nr.filters[Tables.Tag.Name], pager, ops...).Select()
	return
}

// CountTags returns count
func (nr NewsRepo) CountTags(ctx context.Context, search *TagSearch, ops ...OpFunc) (int, error) {
	return buildQuery(ctx, nr.db, &Tag{}, search, nr.filters[Tables.Tag.Name], PagerOne, ops...).Count()
}

// AddTag adds Tag to DB.
func (nr NewsRepo) AddTag(ctx context.Context, tag *Tag, ops ...OpFunc) (*Tag, error) {
	q := nr.db.ModelContext(ctx, tag)
	applyOps(q, ops...)
	_, err := q.Insert()

	return tag, err
}

// UpdateTag updates Tag in DB.
func (nr NewsRepo) UpdateTag(ctx context.Context, tag *Tag, ops ...OpFunc) (bool, error) {
	q := nr.db.ModelContext(ctx, tag).WherePK()
	applyOps(q, ops...)
	res, err := q.Update()
	if err != nil {
		return false, err
	}

	return res.RowsAffected() > 0, err
}

// DeleteTag set statusId to deleted in DB.
func (nr NewsRepo) DeleteTag(ctx context.Context, id int) (deleted bool, err error) {
	tag := &Tag{ID: id, StatusID: StatusDeleted}

	return nr.UpdateTag(ctx, tag, WithColumns(Columns.Tag.StatusID))
}
