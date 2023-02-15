package object

type TableConfig struct {
	StructNameToHump bool // 结构体名称是否转为驼峰式，默认为false
	RmTagIfUcFirst   bool // 如果字段首字母本来就是大写, 就不添加tag, 默认false添加, true不添加
	TagToLower       bool // tag的字段名字是否转换为小写, 如果本身有大写字母的话, 默认false不转
	JsonTagToHump    bool // json tag是否转为驼峰，默认为false，不转换
	UcFirstOnly      bool // 字段首字母大写的同时, 是否要把其他字母转换为小写,默认false不转换
	SeparateFile     bool // 每个struct放入单独的文件,默认false,放入同一个文件
}
