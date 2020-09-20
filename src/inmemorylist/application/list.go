package application

import (
	"context"
	"fmt"
	"sync"

	"go.opencensus.io/trace"

	"github.com/tessig/flamingo-readinglist/src/app/domain"
)

type (
	// ListRepository stores ReadingLists in memory
	ListRepository struct {
		lists map[string]*domain.ReadingList
		mutex sync.Mutex
	}
)

var _ domain.ReadingListRepository = new(ListRepository)

// Load a list from the storage
func (l *ListRepository) Load(ctx context.Context, id string) (*domain.ReadingList, error) {
	ctx, span := trace.StartSpan(ctx, "inmemorylist/repository/load")
	defer span.End()

	l.mutex.Lock()
	defer l.mutex.Unlock()
	list, ok := l.lists[id]
	if !ok {
		return nil, fmt.Errorf("list with id %q not found", id)
	}

	return list, nil
}

// Save a list to the storage, existing lists with the same ID will be overwritten
func (l *ListRepository) Save(ctx context.Context, list *domain.ReadingList) error {
	ctx, span := trace.StartSpan(ctx, "inmemorylist/repository/save")
	defer span.End()

	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.lists == nil {
		l.lists = make(map[string]*domain.ReadingList)
	}
	l.lists[list.ID] = list

	return nil
}
