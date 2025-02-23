package bunSqliteErrors

import "github.com/abc-valera/netsly-golang/internal/domain/util/coderr"

// TODO: add error handling here
func HandleQueryResult(err error) error {
	if err != nil {
		// Handler errors here
		return coderr.NewInternalErr(err)
	}

	return nil
}
