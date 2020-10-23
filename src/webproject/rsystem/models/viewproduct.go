package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

var DefaultViewProductList *ViewProductManager

type ViewProduct struct {
	PId   int64
	PJId  int64
	JId   int64
	PVId  int64
	PVPId int64

	PName     string
	PDes      string
	JName     string
	JDomain   string
	JDes      string
	JUserName string
	JPassword string
	PVName    string
	PVDes     string
	PVJobName string
}

type ViewProductManager struct {
	ViewProducts []*ViewProduct
	lastID       int64
}

// NewTaskManager returns an empty TaskManager.
func NewViewProductManager() *ViewProductManager {
	return &ViewProductManager{}
}

func cloneViewProduct(p *ViewProduct) *ViewProduct {
	c := *p
	return &c
}

func (m *ViewProductManager) SelectAll() []*ViewProduct {
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("select * FROM view_product; ").Values(&maps)

	if err != nil && num > 0 {
		return nil
	}
	var vProducts []*ViewProduct
	for _, mm := range maps {
		pid, _ := strconv.ParseInt(mm["p_id"].(string), 10, 64)
		jid, _ := strconv.ParseInt(mm["j_id"].(string), 10, 64)
		pjid, _ := strconv.ParseInt(mm["p_j_id"].(string), 10, 64)
		pvid, _ := strconv.ParseInt(mm["p_v_id"].(string), 10, 64)
		pvpid, _ := strconv.ParseInt(mm["p_v_p_id"].(string), 10, 64)

		var vp ViewProduct
		vp = ViewProduct{
			PId:       pid,
			JId:       jid,
			PJId:      pjid,
			PVId:      pvid,
			PVPId:     pvpid,
			PName:     mm["p_name"].(string),
			PDes:      mm["p_des"].(string),
			JName:     mm["j_name"].(string),
			JDomain:   mm["j_domain"].(string),
			JDes:      mm["j_des"].(string),
			JUserName: mm["j_user_name"].(string),
			JPassword: mm["j_password"].(string),
			PVName:    mm["p_v_name"].(string),
			PVDes:     mm["p_v_des"].(string),
			PVJobName: mm["p_v_job_name"].(string)}
		vProducts = append(vProducts, &vp)
	}
	return vProducts
}
