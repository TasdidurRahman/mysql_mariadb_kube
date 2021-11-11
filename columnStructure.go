package main

type Column struct {
	Name string `yaml:"name"`
	DataType columnDataType `yaml:"data_type"`
	AllowNull allowNull `yaml:"allow_null"`
}

func (c *Column) config() (ret string){
	ret = c.Name + string(c.DataType) + string(c.AllowNull)

	return ret
}