
package main

import (
    "errors"
    "fmt"
    "log"
    "os"
    "github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"

    "github.com/srikanta0427/listing_app/internal/config"
)

func main() {
    cfg := config.MustLoadConfig()

    if len(os.Args) < 2 {
        log.Fatal("Usage: go run ./cmd/migrate [up|down|version]")
    }

  

    m, err := migrate.New(
        "file://migrations",
        cfg.DatabaseURL,
    )
    if err != nil {
        log.Fatal("Migration initialization failed:", err)
    }

    switch os.Args[1] {

    case "up":
        err = m.Up()
		fmt.Println(os.Args[1])
    case "down":
        err = m.Down()

    case "version":
        version, dirty, versionErr := m.Version()
        if versionErr != nil {
            log.Fatal("Unable to get migration version:", versionErr)
        }

        fmt.Printf("Version: %d, Dirty: %t\n", version, dirty)
        return

    default:
        log.Fatalf("Unknown command: %s", os.Args[1])
    }

    if err != nil && !errors.Is(err, migrate.ErrNoChange) {
        log.Fatal("Migration failed:", err)
    }

    fmt.Println("Migration completed successfully")
}
