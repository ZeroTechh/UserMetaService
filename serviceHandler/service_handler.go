package serviceHandler

import (
	"context"

	proto "github.com/ZeroTechh/VelocityCore/proto/UserMetaService"
	"github.com/ZeroTechh/hades"
	"github.com/jinzhu/copier"

	"github.com/ZeroTechh/UserMetaService/core/metaDB"
	"github.com/ZeroTechh/UserMetaService/core/types"
)

var (
	// all the configs
	config = hades.GetConfig(
		"main.yaml",
		[]string{"config", "../config"},
	)
	accountStatuses = config.Map("accountStatuses")
)

func dataToProto(data types.Meta) *proto.Data {
	request := proto.Data{}
	copier.Copy(&request, &data)
	return &request
}

// New returns a new service handler
func New() *Handler {
	handler := Handler{}
	handler.init()
	return &handler
}

// Handler is used to handle all user main service functions
type Handler struct {
	metaDB *metaDB.MetaDB
}

// Init is used to initialize
func (handler *Handler) init() {
	handler.metaDB = metaDB.New()
}

// Add is used to handle Add function
func (handler Handler) Add(ctx context.Context, request *proto.Identifier) (*proto.Message, error) {
	handler.metaDB.Create(request.UserID)
	return &proto.Message{}, nil
}

// Get is used to handle Get function
func (handler Handler) Get(ctx context.Context, request *proto.Identifier) (*proto.Data, error) {
	data := handler.metaDB.Get(request.UserID)
	response := dataToProto(data)
	return response, nil
}

// Activate is used to handle Activate function
func (handler Handler) Activate(ctx context.Context, request *proto.Identifier) (*proto.Message, error) {
	handler.metaDB.ChangeStatus(
		request.UserID, accountStatuses.Str("verified"))
	return &proto.Message{}, nil
}

// Delete is used to handle Delete function
func (handler Handler) Delete(ctx context.Context, request *proto.Identifier) (*proto.Message, error) {
	handler.metaDB.ChangeStatus(
		request.UserID, accountStatuses.Str("deleted"))
	return &proto.Message{}, nil
}
