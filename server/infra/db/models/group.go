// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Group is an object representing the database table.
type Group struct { // ID
	ID int64 `boil:"id" json:"id" toml:"id" yaml:"id"`
	// ユーザーID
	UserID string `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	// グループ名
	GroupName string `boil:"group_name" json:"group_name" toml:"group_name" yaml:"group_name"`
	// 作成日時
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	// 更新日時
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *groupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L groupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var GroupColumns = struct {
	ID        string
	UserID    string
	GroupName string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	GroupName: "group_name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var GroupTableColumns = struct {
	ID        string
	UserID    string
	GroupName string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "group.id",
	UserID:    "group.user_id",
	GroupName: "group.group_name",
	CreatedAt: "group.created_at",
	UpdatedAt: "group.updated_at",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var GroupWhere = struct {
	ID        whereHelperint64
	UserID    whereHelperstring
	GroupName whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`group`.`id`"},
	UserID:    whereHelperstring{field: "`group`.`user_id`"},
	GroupName: whereHelperstring{field: "`group`.`group_name`"},
	CreatedAt: whereHelpertime_Time{field: "`group`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`group`.`updated_at`"},
}

// GroupRels is where relationship names are stored.
var GroupRels = struct {
	User     string
	UserTags string
}{
	User:     "User",
	UserTags: "UserTags",
}

// groupR is where relationships are stored.
type groupR struct {
	User     *User        `boil:"User" json:"User" toml:"User" yaml:"User"`
	UserTags UserTagSlice `boil:"UserTags" json:"UserTags" toml:"UserTags" yaml:"UserTags"`
}

// NewStruct creates a new relationship struct
func (*groupR) NewStruct() *groupR {
	return &groupR{}
}

func (r *groupR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

func (r *groupR) GetUserTags() UserTagSlice {
	if r == nil {
		return nil
	}
	return r.UserTags
}

// groupL is where Load methods for each relationship are stored.
type groupL struct{}

var (
	groupAllColumns            = []string{"id", "user_id", "group_name", "created_at", "updated_at"}
	groupColumnsWithoutDefault = []string{"user_id", "group_name", "created_at", "updated_at"}
	groupColumnsWithDefault    = []string{"id"}
	groupPrimaryKeyColumns     = []string{"id"}
	groupGeneratedColumns      = []string{}
)

type (
	// GroupSlice is an alias for a slice of pointers to Group.
	// This should almost always be used instead of []Group.
	GroupSlice []*Group
	// GroupHook is the signature for custom Group hook methods
	GroupHook func(context.Context, boil.ContextExecutor, *Group) error

	groupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	groupType                 = reflect.TypeOf(&Group{})
	groupMapping              = queries.MakeStructMapping(groupType)
	groupPrimaryKeyMapping, _ = queries.BindMapping(groupType, groupMapping, groupPrimaryKeyColumns)
	groupInsertCacheMut       sync.RWMutex
	groupInsertCache          = make(map[string]insertCache)
	groupUpdateCacheMut       sync.RWMutex
	groupUpdateCache          = make(map[string]updateCache)
	groupUpsertCacheMut       sync.RWMutex
	groupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var groupAfterSelectHooks []GroupHook

var groupBeforeInsertHooks []GroupHook
var groupAfterInsertHooks []GroupHook

var groupBeforeUpdateHooks []GroupHook
var groupAfterUpdateHooks []GroupHook

var groupBeforeDeleteHooks []GroupHook
var groupAfterDeleteHooks []GroupHook

var groupBeforeUpsertHooks []GroupHook
var groupAfterUpsertHooks []GroupHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Group) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Group) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Group) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Group) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Group) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Group) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Group) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Group) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Group) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range groupAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddGroupHook registers your hook function for all future operations.
func AddGroupHook(hookPoint boil.HookPoint, groupHook GroupHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		groupAfterSelectHooks = append(groupAfterSelectHooks, groupHook)
	case boil.BeforeInsertHook:
		groupBeforeInsertHooks = append(groupBeforeInsertHooks, groupHook)
	case boil.AfterInsertHook:
		groupAfterInsertHooks = append(groupAfterInsertHooks, groupHook)
	case boil.BeforeUpdateHook:
		groupBeforeUpdateHooks = append(groupBeforeUpdateHooks, groupHook)
	case boil.AfterUpdateHook:
		groupAfterUpdateHooks = append(groupAfterUpdateHooks, groupHook)
	case boil.BeforeDeleteHook:
		groupBeforeDeleteHooks = append(groupBeforeDeleteHooks, groupHook)
	case boil.AfterDeleteHook:
		groupAfterDeleteHooks = append(groupAfterDeleteHooks, groupHook)
	case boil.BeforeUpsertHook:
		groupBeforeUpsertHooks = append(groupBeforeUpsertHooks, groupHook)
	case boil.AfterUpsertHook:
		groupAfterUpsertHooks = append(groupAfterUpsertHooks, groupHook)
	}
}

// One returns a single group record from the query.
func (q groupQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Group, error) {
	o := &Group{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for group")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Group records from the query.
func (q groupQuery) All(ctx context.Context, exec boil.ContextExecutor) (GroupSlice, error) {
	var o []*Group

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Group slice")
	}

	if len(groupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Group records in the query.
func (q groupQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count group rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q groupQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if group exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *Group) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`user_id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// UserTags retrieves all the user_tag's UserTags with an executor.
func (o *Group) UserTags(mods ...qm.QueryMod) userTagQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`user_tag`.`group_id`=?", o.ID),
	)

	return UserTags(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (groupL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGroup interface{}, mods queries.Applicator) error {
	var slice []*Group
	var object *Group

	if singular {
		var ok bool
		object, ok = maybeGroup.(*Group)
		if !ok {
			object = new(Group)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeGroup)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeGroup))
			}
		}
	} else {
		s, ok := maybeGroup.(*[]*Group)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeGroup)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeGroup))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &groupR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &groupR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user`),
		qm.WhereIn(`user.user_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user")
	}

	if len(userAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Groups = append(foreign.R.Groups, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.UserID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Groups = append(foreign.R.Groups, local)
				break
			}
		}
	}

	return nil
}

// LoadUserTags allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (groupL) LoadUserTags(ctx context.Context, e boil.ContextExecutor, singular bool, maybeGroup interface{}, mods queries.Applicator) error {
	var slice []*Group
	var object *Group

	if singular {
		var ok bool
		object, ok = maybeGroup.(*Group)
		if !ok {
			object = new(Group)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeGroup)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeGroup))
			}
		}
	} else {
		s, ok := maybeGroup.(*[]*Group)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeGroup)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeGroup))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &groupR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &groupR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`user_tag`),
		qm.WhereIn(`user_tag.group_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load user_tag")
	}

	var resultSlice []*UserTag
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice user_tag")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on user_tag")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user_tag")
	}

	if len(userTagAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.UserTags = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userTagR{}
			}
			foreign.R.Group = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.GroupID) {
				local.R.UserTags = append(local.R.UserTags, foreign)
				if foreign.R == nil {
					foreign.R = &userTagR{}
				}
				foreign.R.Group = local
				break
			}
		}
	}

	return nil
}

// SetUser of the group to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Groups.
func (o *Group) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `group` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, groupPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.UserID
	if o.R == nil {
		o.R = &groupR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Groups: GroupSlice{o},
		}
	} else {
		related.R.Groups = append(related.R.Groups, o)
	}

	return nil
}

// AddUserTags adds the given related objects to the existing relationships
// of the group, optionally inserting them as new records.
// Appends related to o.R.UserTags.
// Sets related.R.Group appropriately.
func (o *Group) AddUserTags(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserTag) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.GroupID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `user_tag` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"group_id"}),
				strmangle.WhereClause("`", "`", 0, userTagPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.GroupID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &groupR{
			UserTags: related,
		}
	} else {
		o.R.UserTags = append(o.R.UserTags, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userTagR{
				Group: o,
			}
		} else {
			rel.R.Group = o
		}
	}
	return nil
}

// SetUserTags removes all previously related items of the
// group replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Group's UserTags accordingly.
// Replaces o.R.UserTags with related.
// Sets related.R.Group's UserTags accordingly.
func (o *Group) SetUserTags(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*UserTag) error {
	query := "update `user_tag` set `group_id` = null where `group_id` = ?"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.UserTags {
			queries.SetScanner(&rel.GroupID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Group = nil
		}
		o.R.UserTags = nil
	}

	return o.AddUserTags(ctx, exec, insert, related...)
}

// RemoveUserTags relationships from objects passed in.
// Removes related items from R.UserTags (uses pointer comparison, removal does not keep order)
// Sets related.R.Group.
func (o *Group) RemoveUserTags(ctx context.Context, exec boil.ContextExecutor, related ...*UserTag) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.GroupID, nil)
		if rel.R != nil {
			rel.R.Group = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("group_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.UserTags {
			if rel != ri {
				continue
			}

			ln := len(o.R.UserTags)
			if ln > 1 && i < ln-1 {
				o.R.UserTags[i] = o.R.UserTags[ln-1]
			}
			o.R.UserTags = o.R.UserTags[:ln-1]
			break
		}
	}

	return nil
}

// Groups retrieves all the records using an executor.
func Groups(mods ...qm.QueryMod) groupQuery {
	mods = append(mods, qm.From("`group`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`group`.*"})
	}

	return groupQuery{q}
}

// FindGroup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindGroup(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Group, error) {
	groupObj := &Group{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `group` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, groupObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from group")
	}

	if err = groupObj.doAfterSelectHooks(ctx, exec); err != nil {
		return groupObj, err
	}

	return groupObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Group) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no group provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(groupColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	groupInsertCacheMut.RLock()
	cache, cached := groupInsertCache[key]
	groupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			groupAllColumns,
			groupColumnsWithDefault,
			groupColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(groupType, groupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(groupType, groupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `group` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `group` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `group` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, groupPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into group")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == groupMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for group")
	}

CacheNoHooks:
	if !cached {
		groupInsertCacheMut.Lock()
		groupInsertCache[key] = cache
		groupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Group.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Group) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	groupUpdateCacheMut.RLock()
	cache, cached := groupUpdateCache[key]
	groupUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			groupAllColumns,
			groupPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update group, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `group` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, groupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(groupType, groupMapping, append(wl, groupPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update group row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for group")
	}

	if !cached {
		groupUpdateCacheMut.Lock()
		groupUpdateCache[key] = cache
		groupUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q groupQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for group")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for group")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o GroupSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `group` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in group slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all group")
	}
	return rowsAff, nil
}

var mySQLGroupUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Group) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no group provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(groupColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLGroupUniqueColumns, o)

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

	groupUpsertCacheMut.RLock()
	cache, cached := groupUpsertCache[key]
	groupUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			groupAllColumns,
			groupColumnsWithDefault,
			groupColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			groupAllColumns,
			groupPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert group, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`group`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `group` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(groupType, groupMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(groupType, groupMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for group")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == groupMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(groupType, groupMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for group")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for group")
	}

CacheNoHooks:
	if !cached {
		groupUpsertCacheMut.Lock()
		groupUpsertCache[key] = cache
		groupUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Group record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Group) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Group provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), groupPrimaryKeyMapping)
	sql := "DELETE FROM `group` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from group")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for group")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q groupQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no groupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from group")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for group")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o GroupSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(groupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `group` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from group slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for group")
	}

	if len(groupAfterDeleteHooks) != 0 {
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
func (o *Group) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindGroup(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *GroupSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := GroupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), groupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `group`.* FROM `group` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, groupPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in GroupSlice")
	}

	*o = slice

	return nil
}

// GroupExists checks if the Group row exists.
func GroupExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `group` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if group exists")
	}

	return exists, nil
}

// Exists checks if the Group row exists.
func (o *Group) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return GroupExists(ctx, exec, o.ID)
}
