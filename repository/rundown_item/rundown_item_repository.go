package rundown_item

import (
	"context"
	"database/sql"

	models "../../models"
	pRepo "../../repository"
)

func InitRundownItemRepository(Connection *sql.DB) pRepo.RundownItemRepository {
	return &RundownItemRepository{
		Connection: Connection,
	}
}

type RundownItemRepository struct {
	Connection *sql.DB
}

func (o *RundownItemRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.RundownItem, error) {
	rows, err := o.Connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.RundownItem, 0)
	for rows.Next() {
		data := new(models.RundownItem)

		err := rows.Scan(
			&data.ID,
			&data.Title,
			&data.Subtitle,
			&data.Text,
			&data.RundownId,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (o *RundownItemRepository) Create(ctx context.Context, rundownItem *models.RundownItem) (*models.RundownItem, error) {
	query := "Insert rundown_items SET title=?, subtitle=?, text=?, rundown_id=?"

	stmt, err := o.Connection.PrepareContext(ctx, query)
	if err != nil {
		return &models.RundownItem{}, err
	}

	res, err := stmt.ExecContext(ctx, rundownItem.Title, rundownItem.Subtitle, rundownItem.Text, rundownItem.RundownId)

	defer stmt.Close()

	if err != nil {
		return &models.RundownItem{}, err
	}

	insertedId, _ := res.LastInsertId()

	rundownItem.ID = insertedId

	return rundownItem, err
}

func (m *RundownItemRepository) GetByRundownId(ctx context.Context, rundownId int64) ([]*models.RundownItem, error) {
	query := "Select id, title, subtitle, text, rundown_id From rundown_items where rundown_id=?"

	return m.fetch(ctx, query, rundownId)
}

func (m *RundownItemRepository) Update(ctx context.Context, p *models.RundownItem) (*models.RundownItem, error) {
	query := "Update rundown_items set title=?, subtitle=?, text=? where id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.Title,
		p.Subtitle,
		p.Text,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return p, nil
}

func (m *RundownItemRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From rundown_items Where id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *RundownItemRepository) DeleteByRundownId(ctx context.Context, rundownId int64) (bool, error) {
	query := "Delete From rundown_items Where rundown_id=?"

	stmt, err := m.Connection.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, rundownId)
	if err != nil {
		return false, err
	}
	return true, nil
}
