package controllers

import (
	"anydevelop.cn/recruit-server/common"
	"anydevelop.cn/recruit-server/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

//  AdminController operations for Admin
type AdminController struct {
	AdminBase
}

// URLMapping ...
func (c *AdminController) URLMapping() {
	c.Mapping("AddAdmin", c.AddAdmin)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Modify", c.Modify)
	c.Mapping("Delete", c.Delete)
}

// AddAdmin ...
// @Title AddAdmin
// @Description create Admin
// @Param	body		body 	models.Admin	true		"body for Admin content"
// @Success 201 {int} models.Admin
// @Failure 403 body is empty
// @router /AddAdmin [post]
func (c *AdminController) AddAdmin() {
	var v models.Admin
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddAdmin(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = common.Success(v)
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Admin by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Admin
// @Failure 403 :id is empty
// @router /GetAdmin/:id [get]
func (c *AdminController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetAdminById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = common.Success(v)
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Admin
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Admin
// @Failure 403
// @router /GetAll [get]
func (c *AdminController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllAdmin(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = common.Success(l)
	}
	c.ServeJSON()
}

// Modify ...
// @Title Modify
// @Description update the Admin
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Admin	true		"body for Admin content"
// @Success 200 {object} models.Admin
// @Failure 403 :id is not int
// @router /ModifyAdmin/:id [put]
func (c *AdminController) Modify() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Admin{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.ModifyAdminById(&v); err == nil {
		c.Data["json"] = common.Ok()
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Admin
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /delete/:id [delete]
func (c *AdminController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteAdmin(id); err == nil {
		c.Data["json"] = common.Ok()
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
