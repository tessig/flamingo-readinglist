package application_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	"github.com/tessig/flamingo-readinglist/src/app/domain"
	"github.com/tessig/flamingo-readinglist/src/inmemorylist/application"
)

func TestListRepository_Load(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		lists   map[string]*domain.ReadingList
		args    args
		want    *domain.ReadingList
		wantErr bool
	}{
		{
			name: "successful load",
			lists: map[string]*domain.ReadingList{
				"id1": {
					ID: "id1",
				},
			},
			args: args{
				id: "id1",
			},
			want: &domain.ReadingList{
				ID: "id1",
			},
			wantErr: false,
		},
		{
			name: "id not in repo",
			lists: map[string]*domain.ReadingList{
				"id1": {
					ID: "id1",
				},
			},
			args: args{
				id: "id2",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := new(application.ListRepository)
			for _, list := range tt.lists {
				require.NoError(t, l.Save(context.Background(), list))
			}

			got, err := l.Load(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error("Load() -got, +want ", diff)
			}
		})
	}
}
