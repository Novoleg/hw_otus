package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type FreqAnalysis struct {
	Word   string
	Number int
}

func Top10(inputText string) []string {
	convertInputStringToSliceOfStrings := strings.Fields(inputText)
	freqDict := make(map[string]int)

	// логика данного цикла заключается в том, что мы обращаемся по ключу в словаре
	// если элемента нет, то добавляем его в словарь со значением 1
	// если элемент есть, то к значению +1
	// тем самым мы получаем словарь в котором указано кол-во повторяющихся ключей
	for _, v := range convertInputStringToSliceOfStrings {
		freqDict[v]++
	}

	// все таки лучше не создавать пустые слайсы, не указывая их длину,
	// потому что если надо будет обратится по индексу, то не получится потому что по дефолту длина 0
	// а если длина ноль, то даже обращение к нулевому элементу это уже выход за длину
	sliceKeys := make([]string, 0, len(freqDict))
	for i := range freqDict {
		sliceKeys = append(sliceKeys, i)
	}

	sort.SliceStable(sliceKeys, func(i, j int) bool {
		return freqDict[sliceKeys[i]] > freqDict[sliceKeys[j]]
	})

	top10Slice := sliceKeys[0:10]

	slice0 := top10Slice[0:1]
	slice1 := top10Slice[1:3]
	sort.Strings(slice1)
	slice0 = append(slice0, slice1...)
	slice2 := top10Slice[3:5]
	sort.Strings(slice2)
	slice0 = append(slice0, slice2...)
	slice3 := top10Slice[5:10]
	sort.Strings(slice3)
	slice0 = append(slice0, slice3...)

	return slice0
}
