package utils

import (
	"os-container-project/internal/model"

	"gorm.io/gorm"
)

func SetDefaultData(db *gorm.DB) error {
	defaultData := []model.Member{
		{
			StudentID: "65050295",
			Name:      "นายณัฐพงศ์ พงศ์จารุมณี",
		},
		{
			StudentID: "65050427",
			Name:      "นายธีรวัจน์ ลือสัตย์",
		},
		{
			StudentID: "65050431",
			Name:      "นายนนท์ปวิธ บัวผุย",
		},
		{
			StudentID: "65050437",
			Name:      "นายนพกร แก้วสลับนิล",
		},
		{
			StudentID: "65050492",
			Name:      "นายบริพัตร จริยาทัศน์กร",
		},
		{
			StudentID: "65050579",
			Name:      "นายพงวิชญ์ สมตา",
		},
	}

	for index := range defaultData {
		var count int64
		db.Model(&model.Member{}).Where("student_id = ?", defaultData[index].StudentID).Count(&count)

		if count == 0 {
			if err := db.Create(&defaultData[index]).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
