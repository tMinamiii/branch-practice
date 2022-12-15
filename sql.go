package branch

import "context"

type Sql struct {
}

func (s *Sql) LoadByID(ctx context.Context, id int64) (any, error) {
	return nil, nil
}
