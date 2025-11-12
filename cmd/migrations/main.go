package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"github.com/ricardo-ronchini/budget-flow-app-go/db"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Passe o comando 'up' ou 'down' para seguir...")
	}

	if err := godotenv.Load(".env"); err != nil {
		log.Panic("⚠️  Não foi possível carregar .env, usando variáveis do sistema")
	}

	m, err := migrate.New(
		"file://db/migrations",
		db.ContextDBURL(),
	)
	if err != nil {
		log.Fatal("Erro ao iniciar migrate: ", err)
	}

	command := os.Args[1]
	param := ""

	if len(os.Args) > 2 {
		param = os.Args[2]
	}

	currentVersion, dirty, err := m.Version()

	if err == migrate.ErrNilVersion {
		// force deafult version
		if err := m.Force(-1); err != nil {
			log.Fatal("Erro ao forçar versão. Erro: ", err)
		}
	}

	if dirty {
		if err := m.Force(int(currentVersion)); err != nil {
			log.Fatal("Erro ao forçar versão. Erro: ", err)
		}
	}

	switch command {
	case "up":
		log.Printf("Migration '%s', go to version '%d'", command, currentVersion)

		err := m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal("Erro ao aplicar migrations: ", err)
		}
		if err == migrate.ErrNoChange {
			log.Fatal("❌ Nenhuma alteração feita...")
		}

		log.Println("✅ Migrations aplicadas com sucesso!")

	case "down":
		if param == "all" {
			if err := m.Down(); err != nil && err != migrate.ErrNoChange {
				log.Fatal("Erro ao reverter migrations: ", err)
			}

			log.Println("Migration down all completa!")
		} else {
			if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
				log.Fatal("Erro ao reverter migrations: ", err)
			}

			log.Println("Migrations retonada para a versão: ", currentVersion-1)
		}

	case "version":
		log.Printf("Current version: %d, dirty: %v", currentVersion, dirty)
	default:
		log.Fatal("Comando não reconhecido")
	}
}
