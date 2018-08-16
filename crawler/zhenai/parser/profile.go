package citylist

import (
	"learn-go/crawler/engine"
	"learn-go/crawler/model"
	"regexp"
	"strconv"
)

// \d 表示数字
var ageReg = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var heightReg = regexp.MustCompile(`<td><span class="label">身高：</span><span field="">([\d]+)CM</span></td>`)
var weightReg = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([\d]+)KG</span></td>`)

var marriageReg = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var genderReg = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var incomeReg = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var educationReg = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var occupationReg = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hukouReg = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<]+)</td>`)
var xinzuoReg = regexp.MustCompile(`<td><span class="label">星座：</span><span field="">([^<]+)</span></td>`)
var carReg = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)

func extractString(contents []byte, reg *regexp.Regexp) string {
	match := reg.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}

	return ""
}

func ParseProfile(contents []byte, name string) engine.ParserResult {

	profile := model.Profile{}

	profile.Marriage = extractString(contents, marriageReg)
	profile.Gender = extractString(contents, genderReg)
	profile.Income = extractString(contents, incomeReg)
	profile.Education = extractString(contents, educationReg)
	profile.Occupation = extractString(contents, occupationReg)
	profile.Hukou = extractString(contents, hukouReg)
	profile.Xinzuo = extractString(contents, xinzuoReg)
	profile.Car = extractString(contents, carReg)

	profile.Name = name

	if age, e := strconv.Atoi(extractString(contents, ageReg)); e == nil {
		profile.Age = age
	}
	if height, e := strconv.Atoi(extractString(contents, heightReg)); e == nil {
		profile.Height = height
	}
	if weight, e := strconv.Atoi(extractString(contents, weightReg)); e == nil {
		profile.Weight = weight
	}

	result := engine.ParserResult{
		Items: []interface{}{profile},
	}

	return result
}
