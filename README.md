# Andthensome

Reddit only caches 1000 saved posts+comments. This is a problem if you keep saving posts and want to view your oldest ones.

The name for this project is named "Andthensome", because it has the ability to retrieve all your Reddit posts, and then some, because it can retrieve any updates as well.

Pushes or merges to master branch will initiate github actions.

Docker build command:

```docker build $(pwd) --tag n30w/andthensome:latest --label n30w/andthensome:latest```

Docker run command:

```docker run --env-file=env_vars --publish 4000:4000 andthensome```

## TODO

- [x] Store reddit credentials in .env file
- [x] SQL Comparison function
- [x] Setup webserver
- [x] Use a docker container for API calls
- [ ] Deletion function for SQL db
- [ ] Log planetscale updates to webpage


## Links

- [How to get more JSON Results](https://old.reddit.com/r/redditdev/comments/d7egb/how_to_get_more_json_results_i_get_only_30/)
- [SQL Go Create Table Insert Row](https://golangbot.com/mysql-create-table-insert-row/)
- [Go Docker Tutorial](https://tutorialedge.net/golang/go-docker-tutorial/)
