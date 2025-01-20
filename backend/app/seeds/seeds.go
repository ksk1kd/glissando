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
		// Members
		{
			Name: "CreateMember - 1",
			Run: func(db *gorm.DB) error {
				return CreateMember(db, "Alice")
			},
		},
		{
			Name: "CreateMember - 2",
			Run: func(db *gorm.DB) error {
				return CreateMember(db, "Bob")
			},
		},
		{
			Name: "CreateMember - 3",
			Run: func(db *gorm.DB) error {
				return CreateMember(db, "Charlie")
			},
		},
		{
			Name: "CreateMember - 4",
			Run: func(db *gorm.DB) error {
				return CreateMember(db, "David")
			},
		},
		{
			Name: "CreateMember - 5",
			Run: func(db *gorm.DB) error {
				return CreateMember(db, "Eve")
			},
		},
		{
			Name: "CreateMember - 6",
			Run: func(db *gorm.DB) error {
				return CreateMember(db, "Frank")
			},
		},
		{
			Name: "CreateMember - 7",
			Run: func(db *gorm.DB) error {
				return CreateMember(db, "Grace")
			},
		},
		{
			Name: "CreateMember - 8",
			Run: func(db *gorm.DB) error {
				return CreateMember(db, "Hank")
			},
		},
		{
			Name: "CreateMember - 9",
			Run: func(db *gorm.DB) error {
				return CreateMember(db, "Ivy")
			},
		},
		{
			Name: "CreateMember - 10",
			Run: func(db *gorm.DB) error {
				return CreateMember(db, "Jack")
			},
		},
		// Projects
		{
			Name: "CreateProject - 1",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Apollo")
			},
		},
		{
			Name: "CreateProject - 2",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Zeus")
			},
		},
		{
			Name: "CreateProject - 3",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Hermes")
			},
		},
		{
			Name: "CreateProject - 4",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Athena")
			},
		},
		{
			Name: "CreateProject - 5",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Poseidon")
			},
		},
		{
			Name: "CreateProject - 6",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Ares")
			},
		},
		{
			Name: "CreateProject - 7",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Hera")
			},
		},
		{
			Name: "CreateProject - 8",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Demeter")
			},
		},
		{
			Name: "CreateProject - 9",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Hestia")
			},
		},
		{
			Name: "CreateProject - 10",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Hephaestus")
			},
		},
		{
			Name: "CreateProject - 11",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Dionysus")
			},
		},
		{
			Name: "CreateProject - 12",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Artemis")
			},
		},
		{
			Name: "CreateProject - 13",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Hades")
			},
		},
		{
			Name: "CreateProject - 14",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Persephone")
			},
		},
		{
			Name: "CreateProject - 15",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Eros")
			},
		},
		{
			Name: "CreateProject - 16",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Helios")
			},
		},
		{
			Name: "CreateProject - 17",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Selene")
			},
		},
		{
			Name: "CreateProject - 18",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Nyx")
			},
		},
		{
			Name: "CreateProject - 19",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Eos")
			},
		},
		{
			Name: "CreateProject - 20",
			Run: func(db *gorm.DB) error {
				return CreateProject(db, "Iris")
			},
		},
		// Resources
		{
			Name: "CreateResource - 1",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 1, 1, "Jan", 10)
			},
		},
		{
			Name: "CreateResource - 2",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 2, 2, "Feb", 20)
			},
		},
		{
			Name: "CreateResource - 3",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 3, 3, "Mar", 30)
			},
		},
		{
			Name: "CreateResource - 4",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 4, 4, "Apr", 40)
			},
		},
		{
			Name: "CreateResource - 5",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 5, 5, "May", 50)
			},
		},
		{
			Name: "CreateResource - 6",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 6, 6, "Jun", 60)
			},
		},
		{
			Name: "CreateResource - 7",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 7, 7, "Jul", 70)
			},
		},
		{
			Name: "CreateResource - 8",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 8, 8, "Aug", 80)
			},
		},
		{
			Name: "CreateResource - 9",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 9, 9, "Sep", 90)
			},
		},
		{
			Name: "CreateResource - 10",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 10, 10, "Oct", 10)
			},
		},
		{
			Name: "CreateResource - 11",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 1, 11, "Nov", 20)
			},
		},
		{
			Name: "CreateResource - 12",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 2, 12, "Dec", 30)
			},
		},
		{
			Name: "CreateResource - 13",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 3, 13, "Jan", 40)
			},
		},
		{
			Name: "CreateResource - 14",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 4, 14, "Feb", 50)
			},
		},
		{
			Name: "CreateResource - 15",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 5, 15, "Mar", 60)
			},
		},
		{
			Name: "CreateResource - 16",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 6, 16, "Apr", 70)
			},
		},
		{
			Name: "CreateResource - 17",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 7, 17, "May", 80)
			},
		},
		{
			Name: "CreateResource - 18",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 8, 18, "Jun", 90)
			},
		},
		{
			Name: "CreateResource - 19",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 9, 19, "Jul", 10)
			},
		},
		{
			Name: "CreateResource - 20",
			Run: func(db *gorm.DB) error {
				return CreateResource(db, 10, 20, "Aug", 20)
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
