package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
	"github.com/bndr/gojenkins"
)

var DefaultJenkinsNodeList *JenkinsNodeManager

type JenkinsNode struct {
	Id       int64  // Unique identifier
	Name     string //
	Domain   string
	Des      string     //
	UserName string     //
	Password string     //
	Products []*Product `orm:"reverse(many)"`

	Client *gojenkins.Jenkins `orm:"-"'`
}
type JenkinsNodeManager struct {
	nodes  []*JenkinsNode
	lastID int64
}

func NewJenkinsNodeManager() *JenkinsNodeManager {
	return &JenkinsNodeManager{}
}

func (m *JenkinsNodeManager) SelectAll() []*JenkinsNode {
	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("select * FROM jenkins_node; ").Values(&maps)

	if err != nil && num > 0 {
		return nil
	}
	var nodes []*JenkinsNode
	for _, mm := range maps {
		node := JenkinsNode{}
		tmpId, _ := strconv.ParseInt(mm["id"].(string), 10, 64)
		node.Id = tmpId
		node.Name = mm["name"].(string)
		node.Des = mm["des"].(string)
		node.Domain = mm["domain"].(string)
		node.UserName = mm["user_name"].(string)
		node.Password = mm["password"].(string)
		fmt.Printf("node:%v %v\n", node.Id, node)
		nodes = append(nodes, &node)
	}

	return nodes
}
func cloneJenkinsNode(p *JenkinsNode) *JenkinsNode {
	c := *p
	return &c
}

func (n *JenkinsNode) Connect() error {
	n.Client = gojenkins.CreateJenkins(nil, n.Domain, n.UserName, n.Password)

	_, err := n.Client.Init()

	if err != nil {
		fmt.Printf(" Jenkins Init 失败, %v\n", err)
		return err
	}
	fmt.Println(n.Name + " Jenkins连接成功 domain=" + n.Domain)
	return nil
}
func (m *JenkinsNodeManager) Find(ID int64) (node *JenkinsNode, err error) {
	for _, node := range m.nodes {
		if node.Id == ID {
			return node, nil
		}
	}
	return nil, nil
}

func (m *JenkinsNodeManager) ConnectAll() {
	for _, node := range m.nodes {
		node.Connect()
	}
}
