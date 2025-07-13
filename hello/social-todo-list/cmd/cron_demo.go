package cmd

import (
	_ "embed"
	"log"

	"github.com/spf13/cobra"
	"gorm.io/gorm"

	"main/common"

	sdk_gorm "github.com/200Lab-Education/go-sdk/plugin/storage/sdkgorm"

	goservice "github.com/200Lab-Education/go-sdk"
)

// Vi du cho cronjob update lai so like count tu bang user_like_items bo table todo_items
// UPDATE todo_items ti INNER JOIN (
// SELECT item_id, COUNT(item_id) as `count` FROM user_like_items
// GROUP BY item_id
// ) c ON c.item_id = ti.id SET ti.liked_count = c.count

var cronDemoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Run demo cron job",
	Run: func(cmd *cobra.Command, args []string) {
		service := goservice.New(
			goservice.WithName("social-todo-list"),
			goservice.WithVersion("1.0.0"),
			goservice.WithInitRunnable(sdk_gorm.NewGormDB("main.mysql", common.PluginDBMain)),
		)

		if err := service.Init(); err != nil {
			log.Fatalln(err)
		}

		db := service.MustGet(common.PluginDBMain).(*gorm.DB)

		///SQL Query Example
		// db.Exec()

		log.Println("...i am demo cron with DB connection:", db)
	},
}

func init() {
	rootCmd.AddCommand(cronDemoCmd)
}
