package model

type Options struct {
	OptionId    int64  `xorm:"not null pk autoincr BIGINT(20)"`
	OptionName  string `xorm:"not null default '' unique VARCHAR(191)"`
	OptionValue string `xorm:"not null LONGTEXT"`
	Autoload    string `xorm:"not null default 'yes' VARCHAR(20)"`
}

// GetOptionsesCount Optionss' Count
func GetOptionsesCount(offset int, limit int) (int64, error) {
	total, err := Engine.Limit(limit, offset).Count(&Options{})
	return total, err
}

// GetOptionsCountViaOptionId Get Options via OptionId
func GetOptionsCountViaOptionId(iOptionId int64) int64 {
	n, _ := Engine.Where("option_id = ?", iOptionId).Count(&Options{OptionId: iOptionId})
	return n
}

// GetOptionsCountViaOptionName Get Options via OptionName
func GetOptionsCountViaOptionName(iOptionName string) int64 {
	n, _ := Engine.Where("option_name = ?", iOptionName).Count(&Options{OptionName: iOptionName})
	return n
}

// GetOptionsCountViaOptionValue Get Options via OptionValue
func GetOptionsCountViaOptionValue(iOptionValue string) int64 {
	n, _ := Engine.Where("option_value = ?", iOptionValue).Count(&Options{OptionValue: iOptionValue})
	return n
}

// GetOptionsCountViaAutoload Get Options via Autoload
func GetOptionsCountViaAutoload(iAutoload string) int64 {
	n, _ := Engine.Where("autoload = ?", iAutoload).Count(&Options{Autoload: iAutoload})
	return n
}

// GetOptionsesByOptionIdAndOptionNameAndOptionValue Get Optionses via OptionIdAndOptionNameAndOptionValue
func GetOptionsesByOptionIdAndOptionNameAndOptionValue(offset int, limit int, OptionId_ int64, OptionName_ string, OptionValue_ string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("option_id = ? and option_name = ? and option_value = ?", OptionId_, OptionName_, OptionValue_).Limit(limit, offset).Find(_Options)
	return _Options, err
}

// GetOptionsesByOptionIdAndOptionNameAndAutoload Get Optionses via OptionIdAndOptionNameAndAutoload
func GetOptionsesByOptionIdAndOptionNameAndAutoload(offset int, limit int, OptionId_ int64, OptionName_ string, Autoload_ string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("option_id = ? and option_name = ? and autoload = ?", OptionId_, OptionName_, Autoload_).Limit(limit, offset).Find(_Options)
	return _Options, err
}

// GetOptionsesByOptionIdAndOptionValueAndAutoload Get Optionses via OptionIdAndOptionValueAndAutoload
func GetOptionsesByOptionIdAndOptionValueAndAutoload(offset int, limit int, OptionId_ int64, OptionValue_ string, Autoload_ string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("option_id = ? and option_value = ? and autoload = ?", OptionId_, OptionValue_, Autoload_).Limit(limit, offset).Find(_Options)
	return _Options, err
}

// GetOptionsesByOptionNameAndOptionValueAndAutoload Get Optionses via OptionNameAndOptionValueAndAutoload
func GetOptionsesByOptionNameAndOptionValueAndAutoload(offset int, limit int, OptionName_ string, OptionValue_ string, Autoload_ string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("option_name = ? and option_value = ? and autoload = ?", OptionName_, OptionValue_, Autoload_).Limit(limit, offset).Find(_Options)
	return _Options, err
}

// GetOptionsesByOptionIdAndOptionName Get Optionses via OptionIdAndOptionName
func GetOptionsesByOptionIdAndOptionName(offset int, limit int, OptionId_ int64, OptionName_ string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("option_id = ? and option_name = ?", OptionId_, OptionName_).Limit(limit, offset).Find(_Options)
	return _Options, err
}

// GetOptionsesByOptionIdAndOptionValue Get Optionses via OptionIdAndOptionValue
func GetOptionsesByOptionIdAndOptionValue(offset int, limit int, OptionId_ int64, OptionValue_ string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("option_id = ? and option_value = ?", OptionId_, OptionValue_).Limit(limit, offset).Find(_Options)
	return _Options, err
}

// GetOptionsesByOptionIdAndAutoload Get Optionses via OptionIdAndAutoload
func GetOptionsesByOptionIdAndAutoload(offset int, limit int, OptionId_ int64, Autoload_ string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("option_id = ? and autoload = ?", OptionId_, Autoload_).Limit(limit, offset).Find(_Options)
	return _Options, err
}

// GetOptionsesByOptionNameAndOptionValue Get Optionses via OptionNameAndOptionValue
func GetOptionsesByOptionNameAndOptionValue(offset int, limit int, OptionName_ string, OptionValue_ string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("option_name = ? and option_value = ?", OptionName_, OptionValue_).Limit(limit, offset).Find(_Options)
	return _Options, err
}

// GetOptionsesByOptionNameAndAutoload Get Optionses via OptionNameAndAutoload
func GetOptionsesByOptionNameAndAutoload(offset int, limit int, OptionName_ string, Autoload_ string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("option_name = ? and autoload = ?", OptionName_, Autoload_).Limit(limit, offset).Find(_Options)
	return _Options, err
}

// GetOptionsesByOptionValueAndAutoload Get Optionses via OptionValueAndAutoload
func GetOptionsesByOptionValueAndAutoload(offset int, limit int, OptionValue_ string, Autoload_ string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("option_value = ? and autoload = ?", OptionValue_, Autoload_).Limit(limit, offset).Find(_Options)
	return _Options, err
}

// GetOptionses Get Optionses via field
func GetOptionses(offset int, limit int, field string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Limit(limit, offset).Desc(field).Find(_Options)
	return _Options, err
}

// GetOptionsesViaOptionId Get Optionss via OptionId
func GetOptionsesViaOptionId(offset int, limit int, OptionId_ int64, field string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("option_id = ?", OptionId_).Limit(limit, offset).Desc(field).Find(_Options)
	return _Options, err
}

// GetOptionsesViaOptionName Get Optionss via OptionName
func GetOptionsesViaOptionName(offset int, limit int, OptionName_ string, field string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("option_name = ?", OptionName_).Limit(limit, offset).Desc(field).Find(_Options)
	return _Options, err
}

// GetOptionsesViaOptionValue Get Optionss via OptionValue
func GetOptionsesViaOptionValue(offset int, limit int, OptionValue_ string, field string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("option_value = ?", OptionValue_).Limit(limit, offset).Desc(field).Find(_Options)
	return _Options, err
}

// GetOptionsesViaAutoload Get Optionss via Autoload
func GetOptionsesViaAutoload(offset int, limit int, Autoload_ string, field string) (*[]*Options, error) {
	var _Options = new([]*Options)
	err := Engine.Table("options").Where("autoload = ?", Autoload_).Limit(limit, offset).Desc(field).Find(_Options)
	return _Options, err
}

// HasOptionsViaOptionId Has Options via OptionId
func HasOptionsViaOptionId(iOptionId int64) bool {
	if has, err := Engine.Where("option_id = ?", iOptionId).Get(new(Options)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasOptionsViaOptionName Has Options via OptionName
func HasOptionsViaOptionName(iOptionName string) bool {
	if has, err := Engine.Where("option_name = ?", iOptionName).Get(new(Options)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasOptionsViaOptionValue Has Options via OptionValue
func HasOptionsViaOptionValue(iOptionValue string) bool {
	if has, err := Engine.Where("option_value = ?", iOptionValue).Get(new(Options)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// HasOptionsViaAutoload Has Options via Autoload
func HasOptionsViaAutoload(iAutoload string) bool {
	if has, err := Engine.Where("autoload = ?", iAutoload).Get(new(Options)); err != nil {
		return false
	} else {
		if has {
			return true
		}
		return false
	}
}

// GetOptionsViaOptionId Get Options via OptionId
func GetOptionsViaOptionId(iOptionId int64) (*Options, error) {
	var _Options = &Options{OptionId: iOptionId}
	has, err := Engine.Get(_Options)
	if has {
		return _Options, err
	} else {
		return nil, err
	}
}

// GetOptionsViaOptionName Get Options via OptionName
func GetOptionsViaOptionName(iOptionName string) (*Options, error) {
	var _Options = &Options{OptionName: iOptionName}
	has, err := Engine.Get(_Options)
	if has {
		return _Options, err
	} else {
		return nil, err
	}
}

// GetOptionsViaOptionValue Get Options via OptionValue
func GetOptionsViaOptionValue(iOptionValue string) (*Options, error) {
	var _Options = &Options{OptionValue: iOptionValue}
	has, err := Engine.Get(_Options)
	if has {
		return _Options, err
	} else {
		return nil, err
	}
}

// GetOptionsViaAutoload Get Options via Autoload
func GetOptionsViaAutoload(iAutoload string) (*Options, error) {
	var _Options = &Options{Autoload: iAutoload}
	has, err := Engine.Get(_Options)
	if has {
		return _Options, err
	} else {
		return nil, err
	}
}

// SetOptionsViaOptionId Set Options via OptionId
func SetOptionsViaOptionId(iOptionId int64, options *Options) (int64, error) {
	options.OptionId = iOptionId
	return Engine.Insert(options)
}

// SetOptionsViaOptionName Set Options via OptionName
func SetOptionsViaOptionName(iOptionName string, options *Options) (int64, error) {
	options.OptionName = iOptionName
	return Engine.Insert(options)
}

// SetOptionsViaOptionValue Set Options via OptionValue
func SetOptionsViaOptionValue(iOptionValue string, options *Options) (int64, error) {
	options.OptionValue = iOptionValue
	return Engine.Insert(options)
}

// SetOptionsViaAutoload Set Options via Autoload
func SetOptionsViaAutoload(iAutoload string, options *Options) (int64, error) {
	options.Autoload = iAutoload
	return Engine.Insert(options)
}

// AddOptions Add Options via all columns
func AddOptions(iOptionId int64, iOptionName string, iOptionValue string, iAutoload string) error {
	_Options := &Options{OptionId: iOptionId, OptionName: iOptionName, OptionValue: iOptionValue, Autoload: iAutoload}
	if _, err := Engine.Insert(_Options); err != nil {
		return err
	} else {
		return nil
	}
}

// PostOptions Post Options via iOptions
func PostOptions(iOptions *Options) (int64, error) {
	_, err := Engine.Insert(iOptions)
	return iOptions.OptionId, err
}

// PutOptions Put Options
func PutOptions(iOptions *Options) (int64, error) {
	_, err := Engine.Id(iOptions.OptionId).Update(iOptions)
	return iOptions.OptionId, err
}

// PutOptionsViaOptionId Put Options via OptionId
func PutOptionsViaOptionId(OptionId_ int64, iOptions *Options) (int64, error) {
	row, err := Engine.Update(iOptions, &Options{OptionId: OptionId_})
	return row, err
}

// PutOptionsViaOptionName Put Options via OptionName
func PutOptionsViaOptionName(OptionName_ string, iOptions *Options) (int64, error) {
	row, err := Engine.Update(iOptions, &Options{OptionName: OptionName_})
	return row, err
}

// PutOptionsViaOptionValue Put Options via OptionValue
func PutOptionsViaOptionValue(OptionValue_ string, iOptions *Options) (int64, error) {
	row, err := Engine.Update(iOptions, &Options{OptionValue: OptionValue_})
	return row, err
}

// PutOptionsViaAutoload Put Options via Autoload
func PutOptionsViaAutoload(Autoload_ string, iOptions *Options) (int64, error) {
	row, err := Engine.Update(iOptions, &Options{Autoload: Autoload_})
	return row, err
}

// UpdateOptionsViaOptionId via map[string]interface{}{}
func UpdateOptionsViaOptionId(iOptionId int64, iOptionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Options)).Where("option_id = ?", iOptionId).Update(iOptionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateOptionsViaOptionName via map[string]interface{}{}
func UpdateOptionsViaOptionName(iOptionName string, iOptionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Options)).Where("option_name = ?", iOptionName).Update(iOptionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateOptionsViaOptionValue via map[string]interface{}{}
func UpdateOptionsViaOptionValue(iOptionValue string, iOptionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Options)).Where("option_value = ?", iOptionValue).Update(iOptionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// UpdateOptionsViaAutoload via map[string]interface{}{}
func UpdateOptionsViaAutoload(iAutoload string, iOptionsMap *map[string]interface{}) error {
	if row, err := Engine.Table(new(Options)).Where("autoload = ?", iAutoload).Update(iOptionsMap); (err != nil) || (row <= 0) {
		return err
	} else {
		return nil
	}
}

// DeleteOptions Delete Options
func DeleteOptions(iOptionId int64) (int64, error) {
	row, err := Engine.Id(iOptionId).Delete(new(Options))
	return row, err
}

// DeleteOptionsViaOptionId Delete Options via OptionId
func DeleteOptionsViaOptionId(iOptionId int64) (err error) {
	var has bool
	var _Options = &Options{OptionId: iOptionId}
	if has, err = Engine.Get(_Options); (has == true) && (err == nil) {
		if row, err := Engine.Where("option_id = ?", iOptionId).Delete(new(Options)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteOptionsViaOptionName Delete Options via OptionName
func DeleteOptionsViaOptionName(iOptionName string) (err error) {
	var has bool
	var _Options = &Options{OptionName: iOptionName}
	if has, err = Engine.Get(_Options); (has == true) && (err == nil) {
		if row, err := Engine.Where("option_name = ?", iOptionName).Delete(new(Options)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteOptionsViaOptionValue Delete Options via OptionValue
func DeleteOptionsViaOptionValue(iOptionValue string) (err error) {
	var has bool
	var _Options = &Options{OptionValue: iOptionValue}
	if has, err = Engine.Get(_Options); (has == true) && (err == nil) {
		if row, err := Engine.Where("option_value = ?", iOptionValue).Delete(new(Options)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}

// DeleteOptionsViaAutoload Delete Options via Autoload
func DeleteOptionsViaAutoload(iAutoload string) (err error) {
	var has bool
	var _Options = &Options{Autoload: iAutoload}
	if has, err = Engine.Get(_Options); (has == true) && (err == nil) {
		if row, err := Engine.Where("autoload = ?", iAutoload).Delete(new(Options)); (err != nil) || (row <= 0) {
			return err
		} else {
			return nil
		}
	}
	return
}
