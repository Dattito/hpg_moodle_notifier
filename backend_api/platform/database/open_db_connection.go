package database

import "github.com/Dattito/HMN_backend_api/app/queries"

type Queries struct {
	*queries.SignalVerificationQueries
	*queries.AssignmentQueries
}

func OpenDBConnection() (*Queries, error) {
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		SignalVerificationQueries: &queries.SignalVerificationQueries{DB: db},
		AssignmentQueries:         &queries.AssignmentQueries{DB: db},
	}, nil
}
