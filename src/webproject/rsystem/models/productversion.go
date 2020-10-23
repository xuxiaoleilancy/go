package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

var DefaultProductVersionManageList *ProductVersionManage

type ProductVersion struct {
	Id        int64
	Name      string      //
	Des       string      //
	JobName   string      //
	Product   *Product    `orm:"rel(fk)"`
	AutoTasks []*AutoTask `orm:"reverse(many)"`
}
type ProductVersionManage struct {
	Versions []*ProductVersion
	lastID   int64
}

func NewProductVersionManage() *ProductVersionManage {
	return &ProductVersionManage{}
}
func (m *ProductVersionManage) SelectAll() []*ProductVersion {
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("select * FROM product_version; ").Values(&maps)

	if err != nil && num > 0 {
		return nil
	}
	var versions []*ProductVersion
	for _, mm := range maps {
		version := ProductVersion{}
		versionId, _ := strconv.ParseInt(mm["id"].(string), 10, 64)
		version.Id = versionId
		version.Name = mm["name"].(string)
		version.Des = mm["des"].(string)
		version.JobName = mm["job_name"].(string)
		productId, _ := strconv.ParseInt(mm["product_id"].(string), 10, 64)

		version.Product, _ = DefaultProductList.FindInCache(productId)
		versions = append(versions, &version)
	}

	return versions
}
func (m *ProductVersionManage) Find(p *Product, version string) *ProductVersion {
	return nil
}
