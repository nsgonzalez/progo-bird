package main

import (
	"github.com/faiface/pixel"
	"github.com/nsgonzalez/progo-bird/models"
	"golang.org/x/image/colornames"
)

const (
	PLAT_W        = 20
	PLAT_MIN_DIST = 100
	PLAT_MIN_H    = 135
	PLAT_MAX_H    = 175
)

func genManualPlatforms() []models.Platform {
	platforms := []models.Platform{
		{Rect: pixel.R(141, -230, 161, -73), Color: colornames.Silver, Type: models.PLATFORM_BOTTOM},
		{Rect: pixel.R(243, 27, 263, 230), Color: colornames.Silver, Type: models.PLATFORM_TOP},
		{Rect: pixel.R(387, -230, 407, -40), Color: colornames.Silver, Type: models.PLATFORM_BOTTOM},
		{Rect: pixel.R(528, 19, 548, 230), Color: colornames.Silver, Type: models.PLATFORM_TOP},
		{Rect: pixel.R(636, -230, 656, -37), Color: colornames.Silver, Type: models.PLATFORM_BOTTOM},
		{Rect: pixel.R(790, -230, 810, 59), Color: colornames.Silver, Type: models.PLATFORM_BOTTOM},
		{Rect: pixel.R(917, 93, 937, 230), Color: colornames.Silver, Type: models.PLATFORM_TOP},
		{Rect: pixel.R(922, -230, 942, -87), Color: colornames.Silver, Type: models.PLATFORM_BOTTOM},
		{Rect: pixel.R(1050, -30, 1070, 230), Color: colornames.Silver, Type: models.PLATFORM_TOP},
	}

	return platforms
}

func loadScenario() (*models.Character, []models.Platform, *models.Goal) {
	character := loadMainCharacter()

	basePlatform := models.Platform{
		Rect:  pixel.R(0, -SCENARIO_H*2, SCENARIO_MAP_W, -SCENARIO_H),
		Color: colornames.Silver,
		Type:  models.PLATFORM_BOTTOM,
	}
	topPlatform := models.Platform{
		Rect:  pixel.R(0, SCENARIO_H, SCENARIO_MAP_W, SCENARIO_H*2),
		Color: colornames.Silver,
		Type:  models.PLATFORM_TOP,
	}
	platforms := append([]models.Platform{basePlatform, topPlatform}, genManualPlatforms()...)

	goal := &models.Goal{
		Pos:    pixel.V(SCENARIO_MAP_W, 0),
		Radius: 30,
		Step:   1.0 / 7,
	}

	return character, platforms, goal
}
