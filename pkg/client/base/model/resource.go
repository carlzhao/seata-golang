package model

import (
	"github.com/carlzhao/seata-golang/v2/pkg/apis"
)

// Resource used to manage transaction resource
type Resource interface {
	GetResourceID() string

	GetBranchType() apis.BranchSession_BranchType
}
