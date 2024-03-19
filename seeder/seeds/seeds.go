package seeds

import (
	"github.com/puvadon-artmit/gofiber-template/seeder/builder"
	"github.com/puvadon-artmit/gofiber-template/seeder/entity"
	"gorm.io/gorm"
)

func All() []entity.Seed {
	return []entity.Seed{
		{
			Name: "Create Role",
			Run: func(db *gorm.DB) error {
				err := builder.CreateRole(db, "ce1dbddb-56aa-4610-8476-b785dbcdd49b", "SU", "Super Admin", "ผู้ดูแลระบบ", "4nG8pLx9Wc7q3Sd5Fb2Vh6Kt1Ry0mJzOuPw")
				if err != nil {
					return err
				}
				err = builder.CreateRole(db, "ce2dbddb-56aa-4610-8476-b785dbcdd496", "IT", "IT support", "ไอทีซัพพอต", "2dFp4sHc6qRt8xYv9MwZ5nTbGjX1lV7yU3K")
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name: "Create User",
			Run: func(db *gorm.DB) error {
				err := builder.CreateUsers(db, "ATC00177", "ภูวดล", "อาจมิตร์", "$2a$10$o8.G9E1mgDGcnkAO48//wu23u0abXd9lGpXqON6CduaVKwrFBtkO.", "ce1dbddb-56aa-4610-8476-b785dbcdd49b")
				if err != nil {
					return err
				}
				err = builder.CreateUsers(db, "ATC00052", "จณิสตา", " ตันสิริ", "$2a$10$S.kVrBqzyYjnz5jm2GneSe85m0AyK4jbWrQshfd3hEoftwTES4kIG", "ce1dbddb-56aa-4610-8476-b785dbcdd49b")
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}
