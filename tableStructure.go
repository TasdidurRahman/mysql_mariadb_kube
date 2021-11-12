package main
//import "C"
//
//type Table struct {
//	Name string `yaml:"name"`
//	Columns []Column `yaml:"columns"`
//}
//
//func (t *Table) create() (ret string){
//	ret = cre + tabl + t.Name + insideParenthesis(t.getColumnconfigs()) + end
//	return ret
//}
//
//func (t *Table) getColumnconfigs()(ret string){
//	ret = ""
//	for i , obj := range t.Columns{
//		if i != 0 {
//			ret += comma
//		}
//		ret += obj.config()
//	}
//	return ret
//}
//
//func (t *Table) drop() (ret string){
//	ret = drop + tabl + t.Name + end
//	return  ret
//}