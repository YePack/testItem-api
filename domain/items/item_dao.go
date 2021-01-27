package items

import (
	"errors"
	"github.com/yepack/testItem-api/clients/elasticsearch"
	"github.com/yepack/testUtils-api/rest_errors"
)

var (
	indexItems = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error to save item", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}
