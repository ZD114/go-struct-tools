package repository

import "time"

type FormHang struct {
	FormHangId        string    `xorm:"varchar(100) 'form_hang_id' notnull default('') comment('')" json:"formHangId"`
	ProcessInstanceId string    `xorm:"varchar(100) 'process_instance_id' notnull default('') comment('流程实例ID')" json:"processInstanceId"` // 流程实例ID
	FormDataId        int64     `xorm:"bigint 'form_data_id' notnull default(0) comment('表单数据ID')" json:"formDataId"`                      // 表单数据ID
	UserName          string    `xorm:"varchar(100) 'user_name' notnull default('') comment('用户名')" json:"userName"`                       // 用户名
	HangName          string    `xorm:"varchar(100) 'hang_name' notnull default('') comment('挂靠人')" json:"hangName"`                       // 挂靠人
	CreateTime        time.Time `xorm:"datetime 'create_time' notnull comment('创建时间')" json:"createTime"`                                  // 创建时间
	UpdateTime        time.Time `xorm:"datetime 'update_time' notnull comment('更新时间')" json:"updateTime"`                                  // 更新时间
}
