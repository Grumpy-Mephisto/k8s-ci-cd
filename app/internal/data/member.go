package data

import (
	"context"

	"github.com/olivere/elastic/v7"
)

func SetDefaultData(client *elastic.Client) error {
	defaultData := []map[string]interface{}{
		{
			"id":   "65050295",
			"name": "นายณัฐพงศ์ พงศ์จารุมณี",
		},
		{
			"id":   "65050427",
			"name": "นายธีรวัจน์ ลือสัตย์",
		},
		{
			"id":   "65050431",
			"name": "นายนนท์ปวิธ บัวผุย",
		},
		{
			"id":   "65050437",
			"name": "นายนพกร แก้วสลับนิล",
		},
		{
			"id":   "65050492",
			"name": "นายบริพัตร จริยาทัศน์กร",
		},
		{
			"id":   "65050579",
			"name": "นายพงวิชญ์ สมตา",
		},
	}

	for _, data := range defaultData {
		exists, err := client.Exists().Index("members").Id(data["id"].(string)).Do(context.Background())
		if err != nil {
			return err
		}

		if !exists {
			_, err := client.Index().
				Index("members").
				Id(data["id"].(string)).
				BodyJson(data).
				Do(context.Background())
			if err != nil {
				return err
			}
		}
	}

	return nil
}
