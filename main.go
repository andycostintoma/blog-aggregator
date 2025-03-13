package main

import (
	"database/sql"
	"github.com/andycostintoma/blog-aggregator/internal"
	"github.com/andycostintoma/blog-aggregator/internal/database"
	"github.com/andycostintoma/blog-aggregator/internal/handlers"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := internal.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &internal.State{
		Db:  dbQueries,
		Cfg: &cfg,
	}

	cmds := internal.Commands{
		RegisteredCommands: make(map[string]func(*internal.State, internal.Command) error),
	}
	cmds.Register("login", handlers.HandlerLogin)
	cmds.Register("register", handlers.HandlerRegister)
	cmds.Register("reset", handlers.HandlerReset)
	cmds.Register("users", handlers.HandlerListUsers)
	cmds.Register("agg", handlers.HandlerAgg)
	cmds.Register("addfeed", internal.MiddlewareLoggedIn(handlers.HandlerAddFeed))
	cmds.Register("feeds", handlers.HandlerListFeeds)
	cmds.Register("follow", internal.MiddlewareLoggedIn(handlers.HandlerFollow))
	cmds.Register("following", handlers.HandlerListFeedFollows)
	cmds.Register("unfollow", internal.MiddlewareLoggedIn(handlers.HandlerUnfollow))
	cmds.Register("browse", internal.MiddlewareLoggedIn(handlers.HandlerBrowse))

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
		return
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.Run(programState, internal.Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
