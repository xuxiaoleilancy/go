package models

//	"fmt"
//	"strconv"
//	"time"

//	"github.com/astaxie/beego/orm"

var DefaultCustomizationCenter *AutoCustomizationCenter

type CustomResourcePack struct {
	Id          int64  // Unique identifier
	CodeDir     string // Description
	ResourceDir string // Is this task done?
	Des         string // Is this task done?
	Os          string
	Procuct     string
}

// TaskManager manages a list of tasks in memory.
type AutoCustomizationCenter struct {
	packs []*CustomResourcePack
}

func NewAutoCustomizationCenter() *AutoCustomizationCenter {
	return &AutoCustomizationCenter{}
}

func init() {
	DefaultCustomizationCenter = NewAutoCustomizationCenter()
}
