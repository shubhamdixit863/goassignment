package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"

	"os"
)

type RestaurantData struct {
	EaterId    string `json:"eater_id"`
	FoodMenuId string `json:"food_menu_id"`
}
type FinalCount struct {
	EaterId string
	Count   int
}

func (rd *RestaurantData) toString() string {

	return fmt.Sprintf("Eater Id = %s , FoodMenuId = %s", rd.EaterId, rd.FoodMenuId)

}

type AllRestaurantData []RestaurantData

func (ard *AllRestaurantData) Find(eaterId string, foodMenuId string) *RestaurantData {
	for _, v := range *ard {
		if v.FoodMenuId == foodMenuId && v.EaterId == eaterId {
			return &v
		}

	}
	return nil
}

func (ard *AllRestaurantData) AddData(rd RestaurantData) {
	*ard = append(*ard, rd)
}

// Utility Method

func SortSlice(slc []FinalCount) {
	sort.Slice(slc, func(i, j int) bool {
		return slc[i].Count > slc[j].Count
	})
}

func (ard *AllRestaurantData) GroupBy() []FinalCount {
	mp := make(map[string]int)
	var countData []FinalCount
	for _, v := range *ard {
		if _, ok := mp[v.FoodMenuId]; ok {
			mp[v.FoodMenuId]++
		} else {
			mp[v.FoodMenuId] = 1
		}
	}

	for k, v := range mp {

		countData = append(countData, FinalCount{
			EaterId: k,
			Count:   v,
		})

	}
	return countData
}

func UnmarshalStringToStruct(string2 string) (*RestaurantData, error) {
	data := &RestaurantData{}
	err := json.Unmarshal([]byte(string2), data)
	if err != nil {
		return data, err
	}

	return data, nil

}

func ReadFile(fileName string) (AllRestaurantData, error) {
	f, err := os.Open(fileName)
	var ar AllRestaurantData

	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		// Read the file line by line and insert the data after parsing inside AllRestaurantData
		toStruct, err := UnmarshalStringToStruct(scanner.Text())
		if err != nil {
			return nil, err
		}

		if ar.Find(toStruct.EaterId, toStruct.FoodMenuId) != nil {

			return nil, errors.New("duplicate Record- " + toStruct.toString())
		}
		// Safe to insert data in the slice
		ar.AddData(*toStruct)

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ar, nil
}

func main() {
	allData, err := ReadFile("data.log")
	if err != nil {
		log.Fatalln("Error Occurred", err)
	}
	// Grouping by the data
	countData := allData.GroupBy()
	// Sorting the countData
	SortSlice(countData)
	fmt.Println(countData[0:3])

}
