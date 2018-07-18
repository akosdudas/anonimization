package anondb

type TypeConversionfunc func(string)(interface{}, error)

//Conversionfunc map[string] typeConversionfunc

/*
func (table TypeConversionTable)Convert(fieldName string, value string)interface{}, err error{
	if table[fieldName] != nil
		return table[fieldName](value), nil
	return value, nil
}
*/

func MakeTypeConversionTable(datasetName string) (table map[string]typeConversionfunc, err error){
	dataset, e := GetDataset(datasetName)
	log.Println(e)
	for _, field := range dataset.Fields {
		if(field.Type) == "coords"
			table[field.Name], err := anonbll.PreprocessCoord
		if(err!= nil)
			return
	}
	return
}
