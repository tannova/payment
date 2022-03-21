package extractor

import (
	"context"
	"errors"
	myid "gitlab.com/greyhole/myid/pkg/api"
	"strconv"

	"google.golang.org/grpc/metadata"

	"gitlab.com/greyhole/myid/pkg/header"
)

type Extractor interface {
	Get(ctx context.Context, name string) []string
	GetTokenID(ctx context.Context) string
	GetUniversalID(ctx context.Context) string
	GetUserID(ctx context.Context) (int64, error)
	GetSafeID(ctx context.Context) (string, bool)
	GetRoleIDs(ctx context.Context) []string
	GetGroupIDs(ctx context.Context) []string
	GetStatus(ctx context.Context) (myid.Status, error)
}

type extractor struct {
}

func New() Extractor {
	return &extractor{}
}

func (t *extractor) Get(ctx context.Context, name string) []string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	return md.Get(name)
}

func (t *extractor) GetTokenID(ctx context.Context) string {
	values := t.Get(ctx, header.TokenID)
	if len(values) == 0 {
		// Return empty in case UniversalID is undefined
		// Empty is default UniversalID
		return ""
	}

	return values[0]
}

func (t *extractor) GetUniversalID(ctx context.Context) string {
	values := t.Get(ctx, header.UniversalID)
	if len(values) == 0 {
		// Return empty in case UniversalID is undefined
		// Empty is default UniversalID
		return ""
	}

	return values[0]
}

func (t *extractor) GetUserID(ctx context.Context) (int64, error) {
	values := t.Get(ctx, header.UserID)
	if len(values) == 0 {
		return 0, errors.New("metadata does not have x-user-id")
	}

	return strconv.ParseInt(values[0], 10, 64)
}

func (t *extractor) GetSafeID(ctx context.Context) (string, bool) {
	values := t.Get(ctx, header.SafeID)
	if len(values) == 0 {
		return "", false
	}

	return values[0], true
}

func (t *extractor) GetRoleIDs(ctx context.Context) []string {
	return t.Get(ctx, header.RoleID)
}

func (t *extractor) GetGroupIDs(ctx context.Context) []string {
	return t.Get(ctx, header.GroupID)
}

func (t *extractor) GetStatus(ctx context.Context) (myid.Status, error) {
	values := t.Get(ctx, header.Status)
	if len(values) == 0 {
		return myid.Status_UNSPECIFIED, errors.New("metadata does not have x-user-status")
	}

	status, err := strconv.ParseInt(values[0], 10, 32)
	if err != nil {
		return myid.Status_UNSPECIFIED, err
	}

	return myid.Status(status), nil
}
