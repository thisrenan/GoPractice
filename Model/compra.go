package Model

import (
	"errors"
	"time"
)

type Buy struct {
	Market string
	Items  []Item
	Date   time.Time
}

func Start(market string, listItem []Item) (*Buy, error) {

	if market == "" {
		return nil, errors.New("market is necessary")
	}

	if len(listItem) == 0 {
		return nil, errors.New("item list is mandatory")
	}

	return &Buy{market, listItem, time.Now()}, nil
}

func StartString(market string, listItem []string) (*Buy, error) {

	if market == "" {
		return nil, errors.New("market is necessary")
	}

	if len(listItem) == 0 {
		return nil, errors.New("item list is mandatory")
	}

	var items []Item
	for _, name := range listItem {
		items = append(items, Item{Name: name})
	}

	return &Buy{market, items, time.Now()}, nil
}
