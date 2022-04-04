package database

import (
	"todo/config"
	"todo/models"
)

func GetActivityByName(name_activity string) (int64, error) {
	tx := config.DB.Where("name = ?", name_activity).Find(&models.Activity{})
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected > 0 {
		return tx.RowsAffected, nil
	}
	return 0, nil
}

func PostActivity(activity models.Activity) (models.Activity, error) {
	tx := config.DB.Save(&activity)
	if tx.Error != nil {
		return activity, tx.Error
	}
	return activity, nil
}

func GetAllActivity() (interface{}, error) {
	var activity []models.ActivityResponse
	query := config.DB.Table("activities").Where("activities.deleted_at IS NULL").Find(&activity)
	if query.Error != nil {
		return nil, query.Error
	}
	if query.RowsAffected == 0 {
		return 0, nil
	}
	return activity, nil
}

func DeleteActivity(id int) (*models.Activity, error) {
	tx := config.DB.Where("id=?", id).Delete(&models.Activity{})
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected > 0 {
		return &models.Activity{}, nil
	}
	return nil, nil
}
