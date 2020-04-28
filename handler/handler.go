package handler

import (
	"context"

	proto "github.com/ZeroTechh/VelocityCore/proto/UserMetaService"
	"github.com/ZeroTechh/hades"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"

	"github.com/ZeroTechh/UserMetaService/core/meta"
	"github.com/ZeroTechh/UserMetaService/core/types"
)

var (
	config   = hades.GetConfig("main.yaml", []string{"config", "../config"})
	verified = config.Map("accountStatuses").Str("verified")
	deleted  = config.Map("accountStatuses").Str("deleted")
)

func toProto(d types.Meta) (*proto.Data, error) {
	r := proto.Data{}
	err := copier.Copy(&r, &d)
	err = errors.Wrap(err, "Error while copying")
	return &r, err
}

// New returns a new handler.
func New() *Handler {
	h := Handler{}
	h.init()
	return &h
}

// Handler handles all user meta service functions.
type Handler struct {
	meta *meta.Meta
}

// init is used to initialize.
func (h *Handler) init() {
	h.meta = meta.New()
}

// Add generates and adds new user meta data into db.
func (h Handler) Add(ctx context.Context, request *proto.Identifier) (*proto.Message, error) {
	err := h.meta.Create(ctx, request.UserID)
	err = errors.Wrap(err, "Error while creating and adding new meta data")
	return &proto.Message{}, err
}

// Get is returns meta data of a user.
func (h Handler) Get(ctx context.Context, request *proto.Identifier) (*proto.Data, error) {
	d, err := h.meta.Get(ctx, request.UserID)
	if err != nil {
		return &proto.Data{}, errors.Wrap(err, "Error while getting meta data")
	}
	r, err := toProto(d)
	return r, errors.Wrap(err, "Error while converting to proto")
}

// Activate marks user as verified.
func (h Handler) Activate(ctx context.Context, request *proto.Identifier) (*proto.Message, error) {
	err := h.meta.SetStatus(ctx, request.UserID, verified)
	err = errors.Wrap(err, "Error while changing status")
	return &proto.Message{}, err
}

// Delete marks user as deleted.
func (h Handler) Delete(ctx context.Context, request *proto.Identifier) (*proto.Message, error) {
	err := h.meta.SetStatus(ctx, request.UserID, deleted)
	err = errors.Wrap(err, "Error while changing status")
	return &proto.Message{}, err
}
