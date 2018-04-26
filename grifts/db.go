package grifts

import (
	"fmt"

	"github.com/markbates/grift/grift"
	"github.com/pkg/errors"
	"mmgitl.mattclark.guru/Anchorspace/dashboard/models"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed:uses", func(c *grift.Context) error {
		if err := models.DB.RawQuery("truncate table uses;").Exec(); err != nil {
			return errors.WithStack(err)
		}
		us := models.Uses{models.Use{
			Name:        "Event",
			Description: "",
		}, models.Use{
			Name:        "Artistic",
			Description: "",
		}, models.Use{
			Name:        "Office",
			Description: "",
		}, models.Use{
			Name:        "Retail",
			Description: "",
		}, models.Use{
			Name:        "Kitchen",
			Description: "",
		}, models.Use{
			Name:        "Industrial",
			Description: "",
		}, models.Use{
			Name:        "Concert",
			Description: "",
		}, models.Use{
			Name:        "Theater",
			Description: "",
		}, models.Use{
			Name:        "Co-Working",
			Description: "",
		}, models.Use{
			Name:        "Pop-Up",
			Description: "",
		}, models.Use{
			Name:        "Commissary",
			Description: "",
		}, models.Use{
			Name:        "Make Space",
			Description: "",
		}, models.Use{
			Name:        "Gala/Wedding",
			Description: "",
		}, models.Use{
			Name:        "Art Gallery",
			Description: "",
		}, models.Use{
			Name:        "Executive Office",
			Description: "",
		}, models.Use{
			Name:        "Event Driven",
			Description: "",
		}, models.Use{
			Name:        "Pop-Up Dining",
			Description: "",
		}, models.Use{
			Name:        "Warehouse",
			Description: "",
		}, models.Use{
			Name:        "Conference",
			Description: "",
		}, models.Use{
			Name:        "Photography",
			Description: "",
		}, models.Use{
			Name:        "Conference Room",
			Description: "",
		}, models.Use{
			Name:        "Brick and Mortar",
			Description: "",
		}, models.Use{
			Name:        "Restaurant",
			Description: "",
		}, models.Use{
			Name:        "Creative/Studio",
			Description: "",
		}, models.Use{
			Name:        "Office",
			Description: "",
		},
		}

		for _, u := range us {
			if err := models.DB.Create(&u); err != nil {
				fmt.Printf("Error creating:%v\n", u)
				return errors.WithStack(err)
			}
		}
		// Add DB seeding stuff here
		//		return models.DB.Eager().Create(&us)
		return nil
	})

	grift.Add("seed", func(c *grift.Context) error {
		err := grift.Run("db:seed:uses", c)
		if err != nil {
			return errors.WithStack(err)
		}
		return nil
	})
})
