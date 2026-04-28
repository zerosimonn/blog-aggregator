package main

import (
	"fmt"
	"context"
	"time"
	"github.com/google/uuid"
	"github.com/zerosimonn/blog-aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <feed_name> <feed_url>", cmd.Name)
	}
	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	
	printFeed(feed)
	
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Println("Feed created successfully:")
	fmt.Printf("ID: %s\n", feed.ID)
	fmt.Printf("Name: %s\n", feed.Name)
	fmt.Printf("URL: %s\n", feed.Url)
	fmt.Printf("Created At: %s\n", feed.CreatedAt.Format(time.RFC3339))
	fmt.Printf("Updated At: %s\n", feed.UpdatedAt.Format(time.RFC3339))
	fmt.Printf("User ID: %s\n", feed.UserID)
}
