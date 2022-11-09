// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// DataPlatformIndustryTextDatum is an object representing the database table.
type DataPlatformIndustryTextDatum struct {
	Industry        string      `boil:"Industry" json:"Industry" toml:"Industry" yaml:"Industry"`
	Language        string      `boil:"Language" json:"Language" toml:"Language" yaml:"Language"`
	IndustryKeyText null.String `boil:"IndustryKeyText" json:"IndustryKeyText,omitempty" toml:"IndustryKeyText" yaml:"IndustryKeyText,omitempty"`

	R *dataPlatformIndustryTextDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L dataPlatformIndustryTextDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DataPlatformIndustryTextDatumColumns = struct {
	Industry        string
	Language        string
	IndustryKeyText string
}{
	Industry:        "Industry",
	Language:        "Language",
	IndustryKeyText: "IndustryKeyText",
}

var DataPlatformIndustryTextDatumTableColumns = struct {
	Industry        string
	Language        string
	IndustryKeyText string
}{
	Industry:        "data_platform_industry_text_data.Industry",
	Language:        "data_platform_industry_text_data.Language",
	IndustryKeyText: "data_platform_industry_text_data.IndustryKeyText",
}

// Generated where

var DataPlatformIndustryTextDatumWhere = struct {
	Industry        whereHelperstring
	Language        whereHelperstring
	IndustryKeyText whereHelpernull_String
}{
	Industry:        whereHelperstring{field: "`data_platform_industry_text_data`.`Industry`"},
	Language:        whereHelperstring{field: "`data_platform_industry_text_data`.`Language`"},
	IndustryKeyText: whereHelpernull_String{field: "`data_platform_industry_text_data`.`IndustryKeyText`"},
}

// DataPlatformIndustryTextDatumRels is where relationship names are stored.
var DataPlatformIndustryTextDatumRels = struct {
	IndustryDataPlatformIndustryIndustryDatum string
}{
	IndustryDataPlatformIndustryIndustryDatum: "IndustryDataPlatformIndustryIndustryDatum",
}

// dataPlatformIndustryTextDatumR is where relationships are stored.
type dataPlatformIndustryTextDatumR struct {
	IndustryDataPlatformIndustryIndustryDatum *DataPlatformIndustryIndustryDatum `boil:"IndustryDataPlatformIndustryIndustryDatum" json:"IndustryDataPlatformIndustryIndustryDatum" toml:"IndustryDataPlatformIndustryIndustryDatum" yaml:"IndustryDataPlatformIndustryIndustryDatum"`
}

// NewStruct creates a new relationship struct
func (*dataPlatformIndustryTextDatumR) NewStruct() *dataPlatformIndustryTextDatumR {
	return &dataPlatformIndustryTextDatumR{}
}

func (r *dataPlatformIndustryTextDatumR) GetIndustryDataPlatformIndustryIndustryDatum() *DataPlatformIndustryIndustryDatum {
	if r == nil {
		return nil
	}
	return r.IndustryDataPlatformIndustryIndustryDatum
}

// dataPlatformIndustryTextDatumL is where Load methods for each relationship are stored.
type dataPlatformIndustryTextDatumL struct{}

var (
	dataPlatformIndustryTextDatumAllColumns            = []string{"Industry", "Language", "IndustryKeyText"}
	dataPlatformIndustryTextDatumColumnsWithoutDefault = []string{"Industry", "Language", "IndustryKeyText"}
	dataPlatformIndustryTextDatumColumnsWithDefault    = []string{}
	dataPlatformIndustryTextDatumPrimaryKeyColumns     = []string{"Industry", "Language"}
	dataPlatformIndustryTextDatumGeneratedColumns      = []string{}
)

type (
	// DataPlatformIndustryTextDatumSlice is an alias for a slice of pointers to DataPlatformIndustryTextDatum.
	// This should almost always be used instead of []DataPlatformIndustryTextDatum.
	DataPlatformIndustryTextDatumSlice []*DataPlatformIndustryTextDatum
	// DataPlatformIndustryTextDatumHook is the signature for custom DataPlatformIndustryTextDatum hook methods
	DataPlatformIndustryTextDatumHook func(context.Context, boil.ContextExecutor, *DataPlatformIndustryTextDatum) error

	dataPlatformIndustryTextDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	dataPlatformIndustryTextDatumType                 = reflect.TypeOf(&DataPlatformIndustryTextDatum{})
	dataPlatformIndustryTextDatumMapping              = queries.MakeStructMapping(dataPlatformIndustryTextDatumType)
	dataPlatformIndustryTextDatumPrimaryKeyMapping, _ = queries.BindMapping(dataPlatformIndustryTextDatumType, dataPlatformIndustryTextDatumMapping, dataPlatformIndustryTextDatumPrimaryKeyColumns)
	dataPlatformIndustryTextDatumInsertCacheMut       sync.RWMutex
	dataPlatformIndustryTextDatumInsertCache          = make(map[string]insertCache)
	dataPlatformIndustryTextDatumUpdateCacheMut       sync.RWMutex
	dataPlatformIndustryTextDatumUpdateCache          = make(map[string]updateCache)
	dataPlatformIndustryTextDatumUpsertCacheMut       sync.RWMutex
	dataPlatformIndustryTextDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var dataPlatformIndustryTextDatumAfterSelectHooks []DataPlatformIndustryTextDatumHook

var dataPlatformIndustryTextDatumBeforeInsertHooks []DataPlatformIndustryTextDatumHook
var dataPlatformIndustryTextDatumAfterInsertHooks []DataPlatformIndustryTextDatumHook

var dataPlatformIndustryTextDatumBeforeUpdateHooks []DataPlatformIndustryTextDatumHook
var dataPlatformIndustryTextDatumAfterUpdateHooks []DataPlatformIndustryTextDatumHook

var dataPlatformIndustryTextDatumBeforeDeleteHooks []DataPlatformIndustryTextDatumHook
var dataPlatformIndustryTextDatumAfterDeleteHooks []DataPlatformIndustryTextDatumHook

var dataPlatformIndustryTextDatumBeforeUpsertHooks []DataPlatformIndustryTextDatumHook
var dataPlatformIndustryTextDatumAfterUpsertHooks []DataPlatformIndustryTextDatumHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DataPlatformIndustryTextDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dataPlatformIndustryTextDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DataPlatformIndustryTextDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dataPlatformIndustryTextDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DataPlatformIndustryTextDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dataPlatformIndustryTextDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DataPlatformIndustryTextDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dataPlatformIndustryTextDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DataPlatformIndustryTextDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dataPlatformIndustryTextDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DataPlatformIndustryTextDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dataPlatformIndustryTextDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DataPlatformIndustryTextDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dataPlatformIndustryTextDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DataPlatformIndustryTextDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dataPlatformIndustryTextDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DataPlatformIndustryTextDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range dataPlatformIndustryTextDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDataPlatformIndustryTextDatumHook registers your hook function for all future operations.
func AddDataPlatformIndustryTextDatumHook(hookPoint boil.HookPoint, dataPlatformIndustryTextDatumHook DataPlatformIndustryTextDatumHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		dataPlatformIndustryTextDatumAfterSelectHooks = append(dataPlatformIndustryTextDatumAfterSelectHooks, dataPlatformIndustryTextDatumHook)
	case boil.BeforeInsertHook:
		dataPlatformIndustryTextDatumBeforeInsertHooks = append(dataPlatformIndustryTextDatumBeforeInsertHooks, dataPlatformIndustryTextDatumHook)
	case boil.AfterInsertHook:
		dataPlatformIndustryTextDatumAfterInsertHooks = append(dataPlatformIndustryTextDatumAfterInsertHooks, dataPlatformIndustryTextDatumHook)
	case boil.BeforeUpdateHook:
		dataPlatformIndustryTextDatumBeforeUpdateHooks = append(dataPlatformIndustryTextDatumBeforeUpdateHooks, dataPlatformIndustryTextDatumHook)
	case boil.AfterUpdateHook:
		dataPlatformIndustryTextDatumAfterUpdateHooks = append(dataPlatformIndustryTextDatumAfterUpdateHooks, dataPlatformIndustryTextDatumHook)
	case boil.BeforeDeleteHook:
		dataPlatformIndustryTextDatumBeforeDeleteHooks = append(dataPlatformIndustryTextDatumBeforeDeleteHooks, dataPlatformIndustryTextDatumHook)
	case boil.AfterDeleteHook:
		dataPlatformIndustryTextDatumAfterDeleteHooks = append(dataPlatformIndustryTextDatumAfterDeleteHooks, dataPlatformIndustryTextDatumHook)
	case boil.BeforeUpsertHook:
		dataPlatformIndustryTextDatumBeforeUpsertHooks = append(dataPlatformIndustryTextDatumBeforeUpsertHooks, dataPlatformIndustryTextDatumHook)
	case boil.AfterUpsertHook:
		dataPlatformIndustryTextDatumAfterUpsertHooks = append(dataPlatformIndustryTextDatumAfterUpsertHooks, dataPlatformIndustryTextDatumHook)
	}
}

// One returns a single dataPlatformIndustryTextDatum record from the query.
func (q dataPlatformIndustryTextDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DataPlatformIndustryTextDatum, error) {
	o := &DataPlatformIndustryTextDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for data_platform_industry_text_data")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DataPlatformIndustryTextDatum records from the query.
func (q dataPlatformIndustryTextDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (DataPlatformIndustryTextDatumSlice, error) {
	var o []*DataPlatformIndustryTextDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DataPlatformIndustryTextDatum slice")
	}

	if len(dataPlatformIndustryTextDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DataPlatformIndustryTextDatum records in the query.
func (q dataPlatformIndustryTextDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count data_platform_industry_text_data rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q dataPlatformIndustryTextDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if data_platform_industry_text_data exists")
	}

	return count > 0, nil
}

// IndustryDataPlatformIndustryIndustryDatum pointed to by the foreign key.
func (o *DataPlatformIndustryTextDatum) IndustryDataPlatformIndustryIndustryDatum(mods ...qm.QueryMod) dataPlatformIndustryIndustryDatumQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`Industry` = ?", o.Industry),
	}

	queryMods = append(queryMods, mods...)

	return DataPlatformIndustryIndustryData(queryMods...)
}

// LoadIndustryDataPlatformIndustryIndustryDatum allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (dataPlatformIndustryTextDatumL) LoadIndustryDataPlatformIndustryIndustryDatum(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDataPlatformIndustryTextDatum interface{}, mods queries.Applicator) error {
	var slice []*DataPlatformIndustryTextDatum
	var object *DataPlatformIndustryTextDatum

	if singular {
		var ok bool
		object, ok = maybeDataPlatformIndustryTextDatum.(*DataPlatformIndustryTextDatum)
		if !ok {
			object = new(DataPlatformIndustryTextDatum)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDataPlatformIndustryTextDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDataPlatformIndustryTextDatum))
			}
		}
	} else {
		s, ok := maybeDataPlatformIndustryTextDatum.(*[]*DataPlatformIndustryTextDatum)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDataPlatformIndustryTextDatum)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDataPlatformIndustryTextDatum))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &dataPlatformIndustryTextDatumR{}
		}
		args = append(args, object.Industry)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &dataPlatformIndustryTextDatumR{}
			}

			for _, a := range args {
				if a == obj.Industry {
					continue Outer
				}
			}

			args = append(args, obj.Industry)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`data_platform_industry_industry_data`),
		qm.WhereIn(`data_platform_industry_industry_data.Industry in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load DataPlatformIndustryIndustryDatum")
	}

	var resultSlice []*DataPlatformIndustryIndustryDatum
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice DataPlatformIndustryIndustryDatum")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for data_platform_industry_industry_data")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for data_platform_industry_industry_data")
	}

	if len(dataPlatformIndustryIndustryDatumAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.IndustryDataPlatformIndustryIndustryDatum = foreign
		if foreign.R == nil {
			foreign.R = &dataPlatformIndustryIndustryDatumR{}
		}
		foreign.R.IndustryDataPlatformIndustryTextData = append(foreign.R.IndustryDataPlatformIndustryTextData, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Industry == foreign.Industry {
				local.R.IndustryDataPlatformIndustryIndustryDatum = foreign
				if foreign.R == nil {
					foreign.R = &dataPlatformIndustryIndustryDatumR{}
				}
				foreign.R.IndustryDataPlatformIndustryTextData = append(foreign.R.IndustryDataPlatformIndustryTextData, local)
				break
			}
		}
	}

	return nil
}

// SetIndustryDataPlatformIndustryIndustryDatum of the dataPlatformIndustryTextDatum to the related item.
// Sets o.R.IndustryDataPlatformIndustryIndustryDatum to related.
// Adds o to related.R.IndustryDataPlatformIndustryTextData.
func (o *DataPlatformIndustryTextDatum) SetIndustryDataPlatformIndustryIndustryDatum(ctx context.Context, exec boil.ContextExecutor, insert bool, related *DataPlatformIndustryIndustryDatum) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `data_platform_industry_text_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"Industry"}),
		strmangle.WhereClause("`", "`", 0, dataPlatformIndustryTextDatumPrimaryKeyColumns),
	)
	values := []interface{}{related.Industry, o.Industry, o.Language}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Industry = related.Industry
	if o.R == nil {
		o.R = &dataPlatformIndustryTextDatumR{
			IndustryDataPlatformIndustryIndustryDatum: related,
		}
	} else {
		o.R.IndustryDataPlatformIndustryIndustryDatum = related
	}

	if related.R == nil {
		related.R = &dataPlatformIndustryIndustryDatumR{
			IndustryDataPlatformIndustryTextData: DataPlatformIndustryTextDatumSlice{o},
		}
	} else {
		related.R.IndustryDataPlatformIndustryTextData = append(related.R.IndustryDataPlatformIndustryTextData, o)
	}

	return nil
}

// DataPlatformIndustryTextData retrieves all the records using an executor.
func DataPlatformIndustryTextData(mods ...qm.QueryMod) dataPlatformIndustryTextDatumQuery {
	mods = append(mods, qm.From("`data_platform_industry_text_data`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`data_platform_industry_text_data`.*"})
	}

	return dataPlatformIndustryTextDatumQuery{q}
}

// FindDataPlatformIndustryTextDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDataPlatformIndustryTextDatum(ctx context.Context, exec boil.ContextExecutor, industry string, language string, selectCols ...string) (*DataPlatformIndustryTextDatum, error) {
	dataPlatformIndustryTextDatumObj := &DataPlatformIndustryTextDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `data_platform_industry_text_data` where `Industry`=? AND `Language`=?", sel,
	)

	q := queries.Raw(query, industry, language)

	err := q.Bind(ctx, exec, dataPlatformIndustryTextDatumObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from data_platform_industry_text_data")
	}

	if err = dataPlatformIndustryTextDatumObj.doAfterSelectHooks(ctx, exec); err != nil {
		return dataPlatformIndustryTextDatumObj, err
	}

	return dataPlatformIndustryTextDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DataPlatformIndustryTextDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_industry_text_data provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformIndustryTextDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	dataPlatformIndustryTextDatumInsertCacheMut.RLock()
	cache, cached := dataPlatformIndustryTextDatumInsertCache[key]
	dataPlatformIndustryTextDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			dataPlatformIndustryTextDatumAllColumns,
			dataPlatformIndustryTextDatumColumnsWithDefault,
			dataPlatformIndustryTextDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformIndustryTextDatumType, dataPlatformIndustryTextDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(dataPlatformIndustryTextDatumType, dataPlatformIndustryTextDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `data_platform_industry_text_data` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `data_platform_industry_text_data` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `data_platform_industry_text_data` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, dataPlatformIndustryTextDatumPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into data_platform_industry_text_data")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.Industry,
		o.Language,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_industry_text_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformIndustryTextDatumInsertCacheMut.Lock()
		dataPlatformIndustryTextDatumInsertCache[key] = cache
		dataPlatformIndustryTextDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DataPlatformIndustryTextDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DataPlatformIndustryTextDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	dataPlatformIndustryTextDatumUpdateCacheMut.RLock()
	cache, cached := dataPlatformIndustryTextDatumUpdateCache[key]
	dataPlatformIndustryTextDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			dataPlatformIndustryTextDatumAllColumns,
			dataPlatformIndustryTextDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update data_platform_industry_text_data, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `data_platform_industry_text_data` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, dataPlatformIndustryTextDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(dataPlatformIndustryTextDatumType, dataPlatformIndustryTextDatumMapping, append(wl, dataPlatformIndustryTextDatumPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update data_platform_industry_text_data row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for data_platform_industry_text_data")
	}

	if !cached {
		dataPlatformIndustryTextDatumUpdateCacheMut.Lock()
		dataPlatformIndustryTextDatumUpdateCache[key] = cache
		dataPlatformIndustryTextDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q dataPlatformIndustryTextDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for data_platform_industry_text_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for data_platform_industry_text_data")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DataPlatformIndustryTextDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformIndustryTextDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `data_platform_industry_text_data` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformIndustryTextDatumPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in dataPlatformIndustryTextDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all dataPlatformIndustryTextDatum")
	}
	return rowsAff, nil
}

var mySQLDataPlatformIndustryTextDatumUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DataPlatformIndustryTextDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no data_platform_industry_text_data provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(dataPlatformIndustryTextDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDataPlatformIndustryTextDatumUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	dataPlatformIndustryTextDatumUpsertCacheMut.RLock()
	cache, cached := dataPlatformIndustryTextDatumUpsertCache[key]
	dataPlatformIndustryTextDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			dataPlatformIndustryTextDatumAllColumns,
			dataPlatformIndustryTextDatumColumnsWithDefault,
			dataPlatformIndustryTextDatumColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			dataPlatformIndustryTextDatumAllColumns,
			dataPlatformIndustryTextDatumPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert data_platform_industry_text_data, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`data_platform_industry_text_data`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `data_platform_industry_text_data` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(dataPlatformIndustryTextDatumType, dataPlatformIndustryTextDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(dataPlatformIndustryTextDatumType, dataPlatformIndustryTextDatumMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for data_platform_industry_text_data")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(dataPlatformIndustryTextDatumType, dataPlatformIndustryTextDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for data_platform_industry_text_data")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for data_platform_industry_text_data")
	}

CacheNoHooks:
	if !cached {
		dataPlatformIndustryTextDatumUpsertCacheMut.Lock()
		dataPlatformIndustryTextDatumUpsertCache[key] = cache
		dataPlatformIndustryTextDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DataPlatformIndustryTextDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DataPlatformIndustryTextDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DataPlatformIndustryTextDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), dataPlatformIndustryTextDatumPrimaryKeyMapping)
	sql := "DELETE FROM `data_platform_industry_text_data` WHERE `Industry`=? AND `Language`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from data_platform_industry_text_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for data_platform_industry_text_data")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q dataPlatformIndustryTextDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no dataPlatformIndustryTextDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from data_platform_industry_text_data")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for data_platform_industry_text_data")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DataPlatformIndustryTextDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(dataPlatformIndustryTextDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformIndustryTextDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `data_platform_industry_text_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformIndustryTextDatumPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from dataPlatformIndustryTextDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for data_platform_industry_text_data")
	}

	if len(dataPlatformIndustryTextDatumAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DataPlatformIndustryTextDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDataPlatformIndustryTextDatum(ctx, exec, o.Industry, o.Language)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DataPlatformIndustryTextDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DataPlatformIndustryTextDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), dataPlatformIndustryTextDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `data_platform_industry_text_data`.* FROM `data_platform_industry_text_data` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, dataPlatformIndustryTextDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DataPlatformIndustryTextDatumSlice")
	}

	*o = slice

	return nil
}

// DataPlatformIndustryTextDatumExists checks if the DataPlatformIndustryTextDatum row exists.
func DataPlatformIndustryTextDatumExists(ctx context.Context, exec boil.ContextExecutor, industry string, language string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `data_platform_industry_text_data` where `Industry`=? AND `Language`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, industry, language)
	}
	row := exec.QueryRowContext(ctx, sql, industry, language)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if data_platform_industry_text_data exists")
	}

	return exists, nil
}

// Exists checks if the DataPlatformIndustryTextDatum row exists.
func (o *DataPlatformIndustryTextDatum) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DataPlatformIndustryTextDatumExists(ctx, exec, o.Industry, o.Language)
}