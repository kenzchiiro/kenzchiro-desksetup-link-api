package product

import (
	"context"
)

func (r *ProductRepository) Delete(ctx context.Context, id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	result, err := r.db.ExecContext(ctx, `
		DELETE FROM products WHERE id = $1
	`, id)
	if err != nil {
		return false, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return affected > 0, nil
}
