package handlers

import (
	"context"
	"fmt"
	"github.com/andycostintoma/blog-aggregator/internal"
	"github.com/andycostintoma/blog-aggregator/internal/database"
	"github.com/google/uuid"
	"time"
)

func HandlerFollow(s *internal.State, cmd internal.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
	feed, err := s.Db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}
	feedFollow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Feed followed %v by user %v\n", feedFollow.FeedName, user.Name)
	return nil
}

func HandlerListFeedFollows(s *internal.State, c internal.Command) error {
	following, err := s.Db.GetFeedFollowsForUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't get following: %w", err)
	}
	fmt.Printf("Following: %v\n", following)
	return nil
}

func HandlerUnfollow(s *internal.State, cmd internal.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
	err := s.Db.DeleteFeedFollowsForUser(context.Background(), database.DeleteFeedFollowsForUserParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return err
	}
	return nil
}
