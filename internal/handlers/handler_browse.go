package handlers

import (
	"context"
	"fmt"
	"github.com/andycostintoma/blog-aggregator/internal"
	"github.com/andycostintoma/blog-aggregator/internal/database"
	"strconv"
)

func HandlerBrowse(s *internal.State, cmd internal.Command, user database.User) error {
	limit := 2
	var err error
	if len(cmd.Args) >= 1 {
		limit, err = strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}

	}
	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})

	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
}
