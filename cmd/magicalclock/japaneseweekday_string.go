// Code generated by "stringer -type=JapaneseWeekDay"; DO NOT EDIT.

package main

import "strconv"

const _JapaneseWeekDay_name = "日月火水木金土"

var _JapaneseWeekDay_index = [...]uint8{0, 3, 6, 9, 12, 15, 18, 21}

func (i JapaneseWeekDay) String() string {
	if i < 0 || i >= JapaneseWeekDay(len(_JapaneseWeekDay_index)-1) {
		return "JapaneseWeekDay(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _JapaneseWeekDay_name[_JapaneseWeekDay_index[i]:_JapaneseWeekDay_index[i+1]]
}
