package entity

import (
	"context"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel/trace"
)

type IFile interface {
	CreateForJoke(ctx context.Context, jokeID string, req FileCreateRequest) (FileCreateResponse, error)
	CreateForRoom(ctx context.Context, roomID string, req FileCreateRequest) (FileCreateResponse, error)
	Update(ctx context.Context, id string, req FileUpdateRequest) (model.FileInfo, error)
	Delete(ctx context.Context, id string) error

	InfoQuery() query.IFileInfo
	ContentQuery() query.IFileContent
}

type file struct {
	IDependency
}

func newFile(dep IDependency) IFile {
	return file{
		IDependency: dep,
	}
}

func (e file) CreateForJoke(ctx context.Context, jokeID string, req FileCreateRequest) (FileCreateResponse, error) {
	var resp FileCreateResponse
	txFunc := func(
		ctx context.Context,
		txC command.Commands,
		txQ query.Queries,
		txE Entities,
	) error {
		var err error
		resp, err = e.create(ctx, req, txC)
		if err != nil {
			return err
		}

		return txC.FileInfo.LinkWithJoke(ctx, resp.FileInfo.ID, jokeID)
	}

	if err := e.RunInTX(ctx, txFunc); err != nil {
		return FileCreateResponse{}, err
	}

	return resp, nil
}

func (e file) CreateForRoom(ctx context.Context, roomID string, req FileCreateRequest) (FileCreateResponse, error) {
	var resp FileCreateResponse
	txFunc := func(
		ctx context.Context,
		txC command.Commands,
		txQ query.Queries,
		txE Entities,
	) error {
		var err error
		resp, err = e.create(ctx, req, txC)
		if err != nil {
			return err
		}

		return txC.FileInfo.LinkWithRoom(ctx, resp.FileInfo.ID, roomID)
	}

	if err := e.RunInTX(ctx, txFunc); err != nil {
		return FileCreateResponse{}, err
	}

	return resp, nil
}

type FileCreateRequest struct {
	FileName    string         `validate:"required,max=128"`
	FileType    model.FileType `validate:"enum"`
	FileContent []byte         `validate:"required,max=32000000"`
}

type FileCreateResponse struct {
	FileInfo    model.FileInfo
	FileContent model.FileContent
}

// create creates a FileInfo and FileContent
func (file) create(ctx context.Context, req FileCreateRequest, commands command.Commands) (FileCreateResponse, error) {
	if err := global.Validate().Struct(req); err != nil {
		return FileCreateResponse{}, err
	}

	fileInfo := model.FileInfo{
		ID:        uuid.NewString(),
		Name:      req.FileName,
		Type:      req.FileType,
		Size:      len(req.FileContent),
		CreatedAt: time.Now().Truncate(time.Millisecond).Local(),
	}

	fileContent := model.FileContent{
		ID:      fileInfo.ID,
		Content: req.FileContent,
	}

	if err := commands.FileInfo.Create(ctx, fileInfo); err != nil {
		return FileCreateResponse{}, err
	}

	if err := commands.FileContent.Create(ctx, fileContent); err != nil {
		return FileCreateResponse{}, err
	}

	return FileCreateResponse{
		FileInfo:    fileInfo,
		FileContent: fileContent,
	}, nil
}

type FileUpdateRequest struct {
	Name *string `validate:"omitempty,min=1,max=256"`
}

func (e file) Update(ctx context.Context, id string, req FileUpdateRequest) (model.FileInfo, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	if err := global.Validate().Struct(req); err != nil {
		return model.FileInfo{}, err
	}

	fileInfo, err := e.Q().FileInfo.Get(ctx, model.FileInfo{ID: id})
	if err != nil {
		return model.FileInfo{}, err
	}

	fileInfo.UpdatedAt = time.Now().Truncate(time.Millisecond).Local()

	if req.Name != nil {
		fileInfo.Name = *req.Name
	}

	if err := e.C().FileInfo.Update(ctx, fileInfo); err != nil {
		return model.FileInfo{}, err
	}

	return fileInfo, nil
}

func (e file) Delete(ctx context.Context, id string) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	txFunc := func(
		ctx context.Context,
		txC command.Commands,
		txQ query.Queries,
		txE Entities,
	) error {
		if err := txC.FileInfo.Delete(ctx, model.FileInfo{ID: id}); err != nil {
			return err
		}

		return txC.FileContent.Delete(ctx, model.FileContent{ID: id})
	}

	return e.RunInTX(ctx, txFunc)
}

func (e file) InfoQuery() query.IFileInfo {
	return e.Q().FileInfo
}

func (e file) ContentQuery() query.IFileContent {
	return e.Q().FileContent
}
