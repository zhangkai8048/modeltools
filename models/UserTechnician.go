package models

import "time"

// 用户-技师基本信息表
type UserTechnician struct {
	Id int64 `json:"id"` 
	Code string `json:"code"` // 技师编码
	MerchantId int64 `json:"merchant_id"` // 商户id
	UserId int64 `json:"user_id"` // 通行证ID
	WorkerId int64 `json:"worker_id"` // 员工ID
	ServiceStatus int `json:"service_status"` // 是否可服务状态：0工作中,1休假中，2已离职
	Mobile string `json:"mobile"` // 联系电话
	ProvinceId int64 `json:"province_id"` // 服务省份ID
	CityId int64 `json:"city_id"` // 服务城市ID
	AreaId int64 `json:"area_id"` // 服务区域ID
	ServiceProductIds string `json:"service_product_ids"` // 服务商品id集合,以","分隔,开头结尾加","
	ContactName string `json:"contact_name"` // 紧急联系人姓名
	ContactMobile string `json:"contact_mobile"` // 紧急联系人手机号
	ReferrerName string `json:"referrer_name"` // 推荐人姓名
	ReferrerMobile string `json:"referrer_mobile"` // 紧急联系人手机号
	Remark string `json:"remark"` // 备注
	IsEnabled int `json:"is_enabled"` // 是否启用,1:启用 0:未启用
	IsDelete int `json:"is_delete"` // 是否删除,0:否 1:是
	Status int `json:"status"` // Incumbency(1, "在职"), Quit(2, "离职"), PartTimeJob(3, "兼职"), Stop(4, "封停");
	RegistrationForm string `json:"registration_form"` // 报名登记表
	HealthCert string `json:"health_cert"` // 健康证
	HealthCertExpiryDate time.Time `json:"health_cert_expiry_date"` // 健康证有效期截止时间(yyyy-MM-dd)
	ExaminationReport string `json:"examination_report"` // 体检报告
	JobSeniorityCard string `json:"job_seniority_card"` // 从业资格证
	PositionalTitles string `json:"positional_titles"` // 职称证书
	CreateTime time.Time `json:"create_time"` // 创建时间
	CreateUserId int64 `json:"create_user_id"` // 创建人id
	UpdateTime time.Time `json:"update_time"` // 更新时间
	UpdateUserId int64 `json:"update_user_id"` // 更新人
	Version int `json:"version"` // 版本号
	Address string `json:"address"` // 技师现地址
	ExamineStatus int `json:"examine_status"` // 审核状态 待审核:0 审核未通过:1 审核已通过:2
	IntentionalSkillIds string `json:"intentional_skill_ids"` // 技师意向技能id集合,以","分隔
	Age int `json:"age"` // 年龄
	WorkAge int `json:"work_age"` // 工龄
	ActivationStatus int `json:"activation_status"` // 激活状态:  未激活:0 待激活:1 已激活:2
	BankName string `json:"bank_name"` // 开户行
	BankNo string `json:"bank_no"` // 卡号
	OtherCertificates string `json:"other_certificates"` // 其他证书  以‘ ，’分隔
	SettlementType int `json:"settlement_type"` // 技师结算方式 0:其他 1:单结 2:周结 3;月结
	SettlementPrompt string `json:"settlement_prompt"` // 结算提示
	SubordinateType int `json:"subordinate_type"` // 从属类型，1：自营；2：加盟
	AcceptCount int64 `json:"accept_count"` // 技师接单量
	IdCard string `json:"id_card"` // 身份证号
	IdCardValidityPeriod string `json:"id_card_validity_period"` // 身份证有效期
	IdCardImg string `json:"id_card_img"` // 身份证照片，多张用逗号分隔
	AccountStatus int `json:"account_status"` // 账户认证状态 待认证:0 认证未通过:1 认证已通过:2
	CentreAddress string `json:"centre_address"` // 工程师服务中心地址
	CentreLng string `json:"centre_lng"` // 工程师服务中心经度
	CentreLat string `json:"centre_lat"` // 工程师服务中心纬度
	OverstepAreaType int `json:"overstep_area_type"` // 跨区范围距离类型，0：不接受；1：接受10公里之内, 2：接受10-20公里之内,3：接受20-40公里之内,4：接受40公里之上,
	PriceStarting float32 `json:"price_starting"` // 起步价
	PriceUnit float32 `json:"price_unit"` // 超出范围的单价
}