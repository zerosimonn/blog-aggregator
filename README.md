# Gator
gator is a blog-aggregator, specifically a multi-user CLI blog aggregator written in Go.

## Prerequisites

- Go 1.22+
- PostgreSQL

## Setup

1. Create a config file in your home dir as `~/.gatorconfig.json` with:

   ```json
   {
     "db_url": "postgres://user:password@localhost:5432/gator?sslmode=disable"
   }
2. go install ...

## Usage

1. Run the program in dev environement using `go run . `
2. Run the program in production environment using `gator <command1> <comnmand2>`
Possible commands: regoster, login, reset, users, agg, feeds, addfeed, follow, unfollow, following, browse.
Example creating a new user:
`gator register <name>`
Example adding a feed:
`gator addfeed <url>`
Start the aggregator:

```bash
gator agg 30s
```

View the posts:

```bash
gator browse [limit]
```

## Information

1. Explanations in NOTES.md file
2. References: boot.dev 