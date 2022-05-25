package contexts

import (
	"github.com/kristiansantos/go-api-rest/api/models"
	"github.com/kristiansantos/go-api-rest/config"
)

func ListPlanets(planets *[]models.Planet) error {
	if err := config.DB.Find(&planets).Error; err != nil {
		return err
	}
	return nil
}

func GetByPlanet(planet *models.Planet, id string) (err error) {
	if err := config.DB.Where("id = ?", id).First(planet).Error; err != nil {
		return err
	}
	return nil
}

func CreatePlanet(planet *models.Planet) (err error) {
	if err = config.DB.Create(planet).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePlanet(planet *models.Planet, id string) (err error) {
	config.DB.Save(planet)
	return nil
}

func DeletePlanet(planet *models.Planet, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(planet)
	return nil
}
