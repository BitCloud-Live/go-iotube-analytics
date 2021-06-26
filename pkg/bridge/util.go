package bridge

import (
	"context"
	"math/big"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/prometheus/promql"
	"github.com/prometheus/prometheus/tsdb"
)

// LastCheckedBlockNo returns last checked block number.
func LastCheckedBlockNo(ctx context.Context, engine *promql.Engine, db *tsdb.DB, network string) (*big.Int, error) {
	query, err := engine.NewInstantQuery(
		db,
		`last_block{network=`+network+`}`,
		time.Now(),
	)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	result := query.Exec(ctx)
	if result.Err != nil {
		return nil, errors.Wrapf(result.Err, "error evaluating query:%v", query.Statement())
	}
	last, ok := big.NewInt(0).SetString(result.Value.(promql.Vector).String(), 10)
	if !ok {
		return nil, errors.New("converting last block no to big.Int")
	}
	return last, nil
}
