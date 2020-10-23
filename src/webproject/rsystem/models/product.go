package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

var DefaultProductList *ProductManager

type Product struct {
	Id             int64  // Unique identifier
	Name           string //
	Des            string
	ProductVersion []*ProductVersion `orm:"reverse(many)"`
	JenkinsNode    *JenkinsNode      `json:"jenkinsnode" orm:"rel(fk)"`
}

// NewTask creates a new task given a title, that can't be empty.
func NewProduct(name string, version string) (*Product, error) {
	if name == "" {
		return nil, fmt.Errorf("empty name")
	}
	if version == "" {
		return nil, fmt.Errorf("empty version")
	}
	o := orm.NewOrm()

	// 获取表句柄
	node := JenkinsNode{}
	var maps []orm.Params
	num, err := o.Raw("select * FROM jenkins_node WHERE ID=? ", 1).Values(&maps)
	if err != nil || num <= 0 {
		return nil, err
	}

	tmpId, _ := strconv.ParseInt(maps[0]["id"].(string), 10, 64)
	node.Id = tmpId
	node.Name = maps[0]["name"].(string)
	node.Des = maps[0]["des"].(string)
	node.Domain = maps[0]["domain"].(string)
	node.UserName = maps[0]["user_name"].(string)
	node.Password = maps[0]["password"].(string)

	fmt.Printf("---node:%v\n", node)
	t := Product{Name: name, JenkinsNode: &node}

	fmt.Printf("---Product:%v\n", t)
	num, err = o.InsertOrUpdate(&t)

	if err == nil && num > 0 {
		return &t, nil
	}
	return nil, err
}

// TaskManager manages a list of tasks in memory.
type ProductManager struct {
	Products []*Product
	lastID   int64
}

// NewTaskManager returns an empty TaskManager.
func NewProductManager() *ProductManager {
	return &ProductManager{}
}

func (m *ProductManager) Insert(p *Product) error {
	o := orm.NewOrm()
	o.Begin()
	num, err := o.Insert(p)

	if err == nil && num > 0 {
		o.Commit()
		return nil
	}
	o.Rollback()
	return err
}

func (m *ProductManager) Update(p Product) (maps []orm.Params, err error) {

	o := orm.NewOrm()
	o.Begin()
	num, err := o.Update(&p)

	if err == nil && num > 0 {
		o.Commit()
		return maps, nil
	}
	o.Rollback()
	return nil, err
}

func (m *ProductManager) Save(p *Product) error {
	o := orm.NewOrm()
	num, err := o.InsertOrUpdate(p)

	if err == nil && num > 0 {
		return nil
	}
	return err
}

func cloneProduct(p *Product) *Product {
	c := *p
	return &c
}

func (m *ProductManager) SelectAll() []*Product {
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("select * FROM product; ").Values(&maps)

	if err != nil && num > 0 {
		return nil
	}
	var products []*Product
	for _, mm := range maps {
		tmpId, _ := strconv.ParseInt(mm["id"].(string), 10, 64)
		nodeId, _ := strconv.ParseInt(mm["jenkins_node_id"].(string), 10, 64)
		node, _ := DefaultJenkinsNodeList.Find(nodeId)
		products = append(products, &Product{Id: tmpId, Name: mm["name"].(string),
			Des: mm["des"].(string), JenkinsNode: node})
	}
	return products
}

func (m *ProductManager) Find(ID int64) (maps []orm.Params, err error) {

	o := orm.NewOrm()
	num, err := o.Raw("select * FROM product WHERE ID=? ", ID).Values(&maps)

	if err == nil && num > 0 {

		return maps, nil
	}
	return nil, err
}
func (m *ProductManager) FindInCache(ID int64) (p *Product, err error) {
	for _, p := range m.Products {
		if p.Id == ID {
			return p, nil
		}
	}
	return nil, nil
}
