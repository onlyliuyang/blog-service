package _const

//用户相关
const UserStateNormal = 1
const UserStateCancel = 2
const UserStateDisable = 3

var UserStateMap = map[int]string{
	UserStateNormal:  "正常",
	UserStateCancel:  "用户已经注销",
	UserStateDisable: "用户已经禁用",
}
