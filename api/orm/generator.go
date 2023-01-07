package orm

import (
	"gorm.io/gen"
)

func GenerateStructs() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "./queries",
		ModelPkgPath: "./models",
		Mode:         gen.WithoutContext,
	})

	g.UseDB(DB)

	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}