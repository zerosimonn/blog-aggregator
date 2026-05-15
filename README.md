# Gator
gator is a blog-aggregator, specifically a multi-user CLI blog aggregator written in Go.

## Prerequisites

- Go 1.22+
- PostgreSQL

## Setup

1. Create a config file at `~/.gatorconfig.json`:

   ```json
   {
     "db_url": "postgres://user:password@localhost:5432/gator?sslmode=disable"
   }
2. go install

## Usage

1. Run the program using `go run . `
Commands list: regoster, login, reset, users, agg, feeds, addfeed, follow, unfollow, following, browse.

## Information

1. Explanations in NOTES.md file
2. References: boot.dev 