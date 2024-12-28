package types

type BlockEntityRequest struct {
	EntityID   string `json:"entity_id" validate:"required"`
	EntityType string `json:"entity_type" validate:"required"`
}

type BlockEntityResponse struct {
	Success    bool   `json:"success"`
	EntityID   string `json:"entity_id"`
	EntityType string `json:"entity_type"`
	Message    string `json:"message,omitempty"`
}

type UnblockEntityRequest struct {
	EntityID   string `json:"entity_id" validate:"required"`
	EntityType string `json:"entity_type" validate:"required"`
}

type UnblockEntityResponse struct {
	Success    bool   `json:"success"`
	EntityID   string `json:"entity_id"`
	EntityType string `json:"entity_type"`
	Message    string `json:"message,omitempty"`
}

type AssignRoleRequest struct {
	Role string `json:"role" validate:"required"`
}

type CancelRoleRequest struct {
	Role string `json:"role" validate:"required"`
}

type AssignPermissionToRoleRequest struct {
	Role        string   `json:"role" validate:"required"`
	Permissions []string `json:"permissions" validate:"required"`
}

type RevokePermissionFromRoleRequest struct {
	Role        string   `json:"role" validate:"required"`
	Permissions []string `json:"permissions" validate:"required"`
}

type PublishStatementRequest struct {
	UserIds     []int    `json:"user_ids" validate:"required"`
	Action      string   `json:"action" validate:"required"`
	Permissions []string `json:"permissions" validate:"required"`
}

type CancelStatementRequest struct {
	Action      string   `json:"action" validate:"required"`
	Permissions []string `json:"permissions" validate:"required"`
}
