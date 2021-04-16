// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testSocialGraphs(t *testing.T) {
	t.Parallel()

	query := SocialGraphs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSocialGraphsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SocialGraphs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSocialGraphsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := SocialGraphs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SocialGraphs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSocialGraphsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SocialGraphSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := SocialGraphs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSocialGraphsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := SocialGraphExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if SocialGraph exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SocialGraphExists to return true, but got false.")
	}
}

func testSocialGraphsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	socialGraphFound, err := FindSocialGraph(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if socialGraphFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSocialGraphsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = SocialGraphs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testSocialGraphsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := SocialGraphs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSocialGraphsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	socialGraphOne := &SocialGraph{}
	socialGraphTwo := &SocialGraph{}
	if err = randomize.Struct(seed, socialGraphOne, socialGraphDBTypes, false, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}
	if err = randomize.Struct(seed, socialGraphTwo, socialGraphDBTypes, false, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = socialGraphOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = socialGraphTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SocialGraphs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSocialGraphsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	socialGraphOne := &SocialGraph{}
	socialGraphTwo := &SocialGraph{}
	if err = randomize.Struct(seed, socialGraphOne, socialGraphDBTypes, false, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}
	if err = randomize.Struct(seed, socialGraphTwo, socialGraphDBTypes, false, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = socialGraphOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = socialGraphTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SocialGraphs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testSocialGraphsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SocialGraphs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSocialGraphsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(socialGraphColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := SocialGraphs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSocialGraphsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSocialGraphsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := SocialGraphSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testSocialGraphsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := SocialGraphs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	socialGraphDBTypes = map[string]string{`ID`: `bigint`, `Amount`: `bigint`, `BuyerAddress`: `text`, `CreatorAddress`: `text`, `Username`: `text`, `ValidatorAddress`: `text`, `CreatedAt`: `timestamp without time zone`, `UpdatedAt`: `timestamp without time zone`, `DeletedAt`: `timestamp without time zone`}
	_                  = bytes.MinRead
)

func testSocialGraphsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(socialGraphPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(socialGraphAllColumns) == len(socialGraphPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SocialGraphs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testSocialGraphsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(socialGraphAllColumns) == len(socialGraphPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &SocialGraph{}
	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := SocialGraphs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, socialGraphDBTypes, true, socialGraphPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(socialGraphAllColumns, socialGraphPrimaryKeyColumns) {
		fields = socialGraphAllColumns
	} else {
		fields = strmangle.SetComplement(
			socialGraphAllColumns,
			socialGraphPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := SocialGraphSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testSocialGraphsUpsert(t *testing.T) {
	t.Parallel()

	if len(socialGraphAllColumns) == len(socialGraphPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := SocialGraph{}
	if err = randomize.Struct(seed, &o, socialGraphDBTypes, true); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SocialGraph: %s", err)
	}

	count, err := SocialGraphs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, socialGraphDBTypes, false, socialGraphPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize SocialGraph struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert SocialGraph: %s", err)
	}

	count, err = SocialGraphs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
