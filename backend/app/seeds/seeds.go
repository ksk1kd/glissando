package seeds

import (
	"log/slog"

	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}

func All() []Seed {
	return []Seed{
		// Items
		{
			Name: "CreateItem - xxx",
			Run: func(db *gorm.DB) error {
				return CreateItem(db, "xxx", "/images/xxx.png", 4, 700, 1)
			},
		},
	}
}

func RunAll(db *gorm.DB) {
	for _, seed := range All() {
		if err := seed.Run(db); err != nil {
			slog.Error("Seeding error",
				slog.String("seed name", seed.Name),
				slog.String("error message", err.Error()))
		}
	}
}
