package orm

import (
	"database/sql"
)

type User struct {
	Id          int64          `column:"id"`
	UserCode    string         `column:"user_code"`
	Password    string         `column:"password"`
	RealName    sql.NullString `column:"real_name"`
	Mobile      sql.NullString `column:"mobile"`
	Email       sql.NullString `column:"email"`
	RoleCodes   sql.NullString `column:"role_codes"`
	WorkCode    sql.NullString `column:"work_code"`
	RegionCode  sql.NullString `column:"region_code"`
	CampusCode  sql.NullString `column:"campus_code"`
	LoginStatus bool           `column:"login_status"`
	Enabled     bool           `column:"enabled"`
	SecretKey   sql.NullString `column:"secret_key"`
	token       sql.NullString `column:"token"`
	CreateTime  sql.NullTime   `column:"create_time"`
	UpdateTime  sql.NullTime   `column:"update_time"`
}

func (User) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
	return "crm_user"
}

//GetUserByUserCode  根据用户编码获取用户名
func (u User) GetUserByUserCode(userCode string) (user *User) {
	db.Where("user_code = ?", userCode).First(&user)
	if user.Id > 0 {
		return user
	} else {
		return nil
	}
}
