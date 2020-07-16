package config


var (
	StatusMap=map[int]string{
		0:"新建",
		1:"进行中",
		2:"暂停",
		3:"完成",
		-1:"删除",
	}
	UserStatusMap=map[int]string{
		1:"在职",
		0:"离职",
	}
)