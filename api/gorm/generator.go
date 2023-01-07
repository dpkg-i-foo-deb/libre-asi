package gorm

import (
	"gorm.io/gen"
)

func GenerateStructs() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./queries",
		ModelPkgPath: "./models",
		Mode:         gen.WithoutContext,
	})

	g.UseDB(db)

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
