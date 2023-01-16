package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadFile(t *testing.T) {

	data, err := ReadFile("data.log")
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(data), 0)

}

func TestUnmarshalStringToStruct(t *testing.T) {

	str := `{"eater_id":"johnny","food_menu_id":"we122"}`
	toStruct, err := UnmarshalStringToStruct(str)
	assert.Nil(t, err)
	assert.Equal(t, "we122", toStruct.FoodMenuId)
	assert.Equal(t, "johnny", toStruct.EaterId)

}

func TestAllRestaurantData_GroupBy(t *testing.T) {
	data, err := ReadFile("data.log")
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(data), 0)
	countData := data.GroupBy()
	assert.GreaterOrEqual(t, len(countData), 0)

}
