// This file is generated by SQLBoiler (https://github.com/volatiletech/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

func testTickers(t *testing.T) {
	t.Parallel()

	query := Tickers(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testTickersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ticker.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Tickers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTickersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Tickers(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Tickers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTickersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TickerSlice{ticker}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Tickers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testTickersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := TickerExists(tx, ticker.ID)
	if err != nil {
		t.Errorf("Unable to check if Ticker exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TickerExistsG to return true, but got false.")
	}
}
func testTickersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx); err != nil {
		t.Error(err)
	}

	tickerFound, err := FindTicker(tx, ticker.ID)
	if err != nil {
		t.Error(err)
	}

	if tickerFound == nil {
		t.Error("want a record, got nil")
	}
}
func testTickersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Tickers(tx).Bind(ticker); err != nil {
		t.Error(err)
	}
}

func testTickersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Tickers(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTickersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	tickerOne := &Ticker{}
	tickerTwo := &Ticker{}
	if err = randomize.Struct(seed, tickerOne, tickerDBTypes, false, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}
	if err = randomize.Struct(seed, tickerTwo, tickerDBTypes, false, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tickerOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = tickerTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Tickers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTickersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tickerOne := &Ticker{}
	tickerTwo := &Ticker{}
	if err = randomize.Struct(seed, tickerOne, tickerDBTypes, false, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}
	if err = randomize.Struct(seed, tickerTwo, tickerDBTypes, false, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tickerOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = tickerTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tickers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func tickerBeforeInsertHook(e boil.Executor, o *Ticker) error {
	*o = Ticker{}
	return nil
}

func tickerAfterInsertHook(e boil.Executor, o *Ticker) error {
	*o = Ticker{}
	return nil
}

func tickerAfterSelectHook(e boil.Executor, o *Ticker) error {
	*o = Ticker{}
	return nil
}

func tickerBeforeUpdateHook(e boil.Executor, o *Ticker) error {
	*o = Ticker{}
	return nil
}

func tickerAfterUpdateHook(e boil.Executor, o *Ticker) error {
	*o = Ticker{}
	return nil
}

func tickerBeforeDeleteHook(e boil.Executor, o *Ticker) error {
	*o = Ticker{}
	return nil
}

func tickerAfterDeleteHook(e boil.Executor, o *Ticker) error {
	*o = Ticker{}
	return nil
}

func tickerBeforeUpsertHook(e boil.Executor, o *Ticker) error {
	*o = Ticker{}
	return nil
}

func tickerAfterUpsertHook(e boil.Executor, o *Ticker) error {
	*o = Ticker{}
	return nil
}

func testTickersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Ticker{}
	o := &Ticker{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, tickerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Ticker object: %s", err)
	}

	AddTickerHook(boil.BeforeInsertHook, tickerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	tickerBeforeInsertHooks = []TickerHook{}

	AddTickerHook(boil.AfterInsertHook, tickerAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	tickerAfterInsertHooks = []TickerHook{}

	AddTickerHook(boil.AfterSelectHook, tickerAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	tickerAfterSelectHooks = []TickerHook{}

	AddTickerHook(boil.BeforeUpdateHook, tickerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	tickerBeforeUpdateHooks = []TickerHook{}

	AddTickerHook(boil.AfterUpdateHook, tickerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	tickerAfterUpdateHooks = []TickerHook{}

	AddTickerHook(boil.BeforeDeleteHook, tickerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	tickerBeforeDeleteHooks = []TickerHook{}

	AddTickerHook(boil.AfterDeleteHook, tickerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	tickerAfterDeleteHooks = []TickerHook{}

	AddTickerHook(boil.BeforeUpsertHook, tickerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	tickerBeforeUpsertHooks = []TickerHook{}

	AddTickerHook(boil.AfterUpsertHook, tickerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	tickerAfterUpsertHooks = []TickerHook{}
}
func testTickersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tickers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTickersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx, tickerColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Tickers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTickerToOneExchangeUsingExchange(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var local Ticker
	var foreign Exchange

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, tickerDBTypes, false, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, exchangeDBTypes, false, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.ExchangeID = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.Exchange(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TickerSlice{&local}
	if err = local.L.LoadExchange(tx, false, (*[]*Ticker)(&slice)); err != nil {
		t.Fatal(err)
	}
	if local.R.Exchange == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Exchange = nil
	if err = local.L.LoadExchange(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Exchange == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTickerToOneSetOpExchangeUsingExchange(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Ticker
	var b, c Exchange

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, tickerDBTypes, false, strmangle.SetComplement(tickerPrimaryKeyColumns, tickerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, exchangeDBTypes, false, strmangle.SetComplement(exchangePrimaryKeyColumns, exchangeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, exchangeDBTypes, false, strmangle.SetComplement(exchangePrimaryKeyColumns, exchangeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Exchange{&b, &c} {
		err = a.SetExchange(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Exchange != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Tickers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ExchangeID != x.ID {
			t.Error("foreign key was wrong value", a.ExchangeID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ExchangeID))
		reflect.Indirect(reflect.ValueOf(&a.ExchangeID)).Set(zero)

		if err = a.Reload(tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ExchangeID != x.ID {
			t.Error("foreign key was wrong value", a.ExchangeID, x.ID)
		}
	}
}
func testTickersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = ticker.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testTickersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TickerSlice{ticker}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testTickersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Tickers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	tickerDBTypes = map[string]string{`AdjClose`: `real`, `Close`: `real`, `ExchangeID`: `integer`, `ExecutedOn`: `timestamp without time zone`, `High`: `real`, `ID`: `integer`, `Low`: `real`, `Open`: `real`, `Volume`: `real`}
	_             = bytes.MinRead
)

func testTickersUpdate(t *testing.T) {
	t.Parallel()

	if len(tickerColumns) == len(tickerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tickers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	if err = ticker.Update(tx); err != nil {
		t.Error(err)
	}
}

func testTickersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(tickerColumns) == len(tickerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	ticker := &Ticker{}
	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tickers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, ticker, tickerDBTypes, true, tickerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(tickerColumns, tickerPrimaryKeyColumns) {
		fields = tickerColumns
	} else {
		fields = strmangle.SetComplement(
			tickerColumns,
			tickerPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(ticker))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := TickerSlice{ticker}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testTickersUpsert(t *testing.T) {
	t.Parallel()

	if len(tickerColumns) == len(tickerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	ticker := Ticker{}
	if err = randomize.Struct(seed, &ticker, tickerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = ticker.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Ticker: %s", err)
	}

	count, err := Tickers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &ticker, tickerDBTypes, false, tickerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Ticker struct: %s", err)
	}

	if err = ticker.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Ticker: %s", err)
	}

	count, err = Tickers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
