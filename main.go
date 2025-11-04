package main

import (
	"fmt"
	"path/filepath"

	"gorm.io/driver/mysql"
	"gorm.io/gen/field"

	"gorm.io/gen"

	"gorm.io/gorm"
)

func main() {
	fmt.Println("vim-go")

	genCfg := gen.Config{
		OutPath:           filepath.Join(".", "queries"),
		ModelPkgPath:      filepath.Join(".", "models"),
		Mode:              gen.WithDefaultQuery,
		FieldWithTypeTag:  true,
		FieldSignable:     true,
		FieldNullable:     false,
		FieldWithIndexTag: true,
		WithUnitTest:      false,
	}

	g := gen.NewGenerator(genCfg)

	g.WithImportPkgPath("gorm.io/gorm", "gorm.io/plugin/optimisticlock", "gorm.io/plugin/soft_delete")

	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN: "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local",
	}))
	g.UseDB(db)

	tables := g.GenerateAllTable()

	stationWaterLevel := g.GenerateModel("station_river_water_level", gen.FieldRelate(field.BelongsTo,
		"Station", g.GenerateModel("station"),
		&field.RelateConfig{
			RelatePointer: true,
			GORMTag: field.GormTag{
				"joinForeignKey": []string{"station_id"},
			},
		}))
	g.ApplyBasic(stationWaterLevel)

	stationRain := g.GenerateModel("station_rain", gen.FieldRelate(field.BelongsTo,
		"Station", g.GenerateModel("station"),
		&field.RelateConfig{
			RelatePointer: true,
			GORMTag: field.GormTag{
				"joinForeignKey": []string{"station_id"},
			},
		}))
	g.ApplyBasic(stationRain)

	reservoir := g.GenerateModel("reservoir",
		gen.FieldRelate(field.HasOne, "WlStation", g.GenerateModel("station"),
			&field.RelateConfig{
				RelatePointer: true,
				GORMTag: field.GormTag{
					"joinForeignKey": []string{"wl_station_id"},
				},
			}),
		gen.FieldRelate(field.HasOne, "RainStation", g.GenerateModel("station_rain"),
			&field.RelateConfig{
				RelatePointer: true,
				GORMTag: field.GormTag{
					"joinForeignKey": []string{"rain_station_id"},
				},
			}))
	g.ApplyBasic(reservoir)
	g.ApplyBasic(tables...)

	g.Execute()
}
