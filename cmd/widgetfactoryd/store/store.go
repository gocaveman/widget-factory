package store

import (
	"context"
	"database/sql"
)

func NewStore(DB *sql.DB) *Store {
	return &Store{
		DB: DB,
	}
}

type Store struct {
	*sql.DB
}

func (store *Store) Widget() *WidgetStore {
	return &WidgetStore{DB: store.DB}
}

type WidgetStore struct {
	DB *sql.DB
}

func (s *WidgetStore) Insert(ctx context.Context, o *Widget) error {

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, "INSERT INTO widget(widget_id, name, description) VALUES (?, ?, ?)", o.WidgetID, o.Name, o.Description); err != nil {
		return err
	}

	return tx.Commit()

}

func (s *WidgetStore) Update(ctx context.Context, o *Widget) error {

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "UPDATE widget set name = ?, description = ? where widget_id = ?", o.Name, o.Description, o.WidgetID)
	if err != nil {
		return err
	}

	return tx.Commit()

}

func (s *WidgetStore) Delete(ctx context.Context, id string) error {

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "DELETE FROM widget where widget_id = ?", id)
	if err != nil {
		return err
	}

	return tx.Commit()

}

func (s *WidgetStore) SelectOne(ctx context.Context, id string) (*Widget, error) {

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var ret Widget

	row := tx.QueryRowContext(ctx, "SELECT widget_id, name, description FROM widget WHERE widget_id = ?", id)

	if err := row.Scan(&ret.WidgetID, &ret.Name, &ret.Description); err != nil {
		return nil, err
	}

	return &ret, tx.Commit()

}

func (s *WidgetStore) SelectLimit(ctx context.Context, limit int64) ([]Widget, error) {

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var ret []Widget

	var rows *sql.Rows

	if limit == 0 {
		rows, err = tx.QueryContext(ctx, "SELECT widget_id, name, description FROM widget")
		if err != nil {
			return nil, err
		}
	} else {
		rows, err = tx.QueryContext(ctx, "SELECT widget_id, name, description FROM widget LIMIT ?", limit)
		if err != nil {
			return nil, err
		}
	}

	defer rows.Close()

	for rows.Next() {
		var o Widget
		if err := rows.Scan(&o.WidgetID, &o.Name, &o.Description); err != nil {
			return nil, err
		}
		ret = append(ret, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ret, tx.Commit()

}

func (s *WidgetStore) SelectLimitCount(ctx context.Context, limit int64) ([]Widget, int64, error) {

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, 0, err
	}

	defer tx.Rollback()

	var ret []Widget

	var rows *sql.Rows

	rows, err = tx.QueryContext(ctx, "SELECT widget_id, name, description FROM widget LIMIT ?", limit)
	if err != nil {
		return nil, 0, err
	}

	for rows.Next() {
		var o Widget
		if err := rows.Scan(&o.WidgetID, &o.Name, &o.Description); err != nil {
			return nil, 0, err
		}
		ret = append(ret, o)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	if err := rows.Close(); err != nil {
		return nil, 0, err
	}

	rowsCount := tx.QueryRowContext(ctx, "SELECT COUNT(*) AS count FROM widget")

	var count uint8

	if err := rowsCount.Scan(&count); err != nil {
		return nil, 0, err
	}

	return ret, int64(count), tx.Commit()

}
