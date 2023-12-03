package main

import (
    "log"
    "os"

    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/apis"
    "github.com/pocketbase/pocketbase/core"
		"fmt"
)

// region
// docker





func main() {
    app := pocketbase.New()

    // serves static files from the provided public dir (if exists)
    app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
        e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
        return nil
    })


		app.OnRecordBeforeUpdateRequest("containers").Add(func(e *core.RecordUpdateEvent) error {
			// admin, _ := e.HttpContext.Get(apis.ContextAdminKey).(*models.Admin)
			// if admin != nil {
			// 		return nil // ignore for admins
			// }

			// overwrite the submitted "active" field value to false
			// e.Record.Set("active", false)
			fmt.Println("CREATEDDDDDDD")

			// // or you can also prevent the create event by returning an error, eg.:
			// if (e.Record.GetString("status") != "pending") {
			// 		return apis.NewBadRequestError("status must be pending.", nil)
			// }

			return nil
		})


    if err := app.Start(); err != nil {
        log.Fatal(err)
    }

		
}