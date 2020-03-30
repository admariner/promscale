package pgmodel

import (
	"github.com/prometheus/prometheus/prompb"
)

// Reader reads the data based on the provided read request.
type Reader interface {
	Read(*prompb.ReadRequest) (*prompb.ReadResponse, error)
}

// Querier queries the data using the provided query data and returns the
// matching timeseries.
type Querier interface {
	Query(*prompb.Query) ([]*prompb.TimeSeries, error)
}

//HealthChecker allows checking for proper operations
type HealthChecker interface {
	HealthCheck() error
}

// QueryHealthChecker can query and check its own health
type QueryHealthChecker interface {
	Querier
	HealthChecker
}

// DBReader reads data from the database.
type DBReader struct {
	db QueryHealthChecker
}

func (r *DBReader) Read(req *prompb.ReadRequest) (*prompb.ReadResponse, error) {
	if req == nil {
		return nil, nil
	}

	resp := prompb.ReadResponse{
		Results: make([]*prompb.QueryResult, len(req.Queries)),
	}

	for i, q := range req.Queries {
		tts, err := r.db.Query(q)
		if err != nil {
			return nil, err
		}
		resp.Results[i] = &prompb.QueryResult{
			Timeseries: tts,
		}
	}

	return &resp, nil
}

// HealthCheck checks that the reader is properly connected
func (r *DBReader) HealthCheck() error {
	return r.HealthCheck()
}
