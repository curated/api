# Curated

Curated is an open source & data platform consisting of a curated dataset of GitHub comments ranked by user reactions for popular repositories. The goal is to provide easy access to meaningful information for anyone interested in learning best practices and development history of popular open source projects.

## Architecture

Curated has a simple ecosystem where **api** and **scheduler** are two distinct processes sharing **domain** logic while interfacing with different external components: user interfaces and data operations, respectively.

```
api users
         \
mobile ----> api -> domain <- scheduler -> github-ql
         /            |                       |
web                postgres                 github
```

The following repositories compose the project ecosystem:

- **api** - open data endpoints
- [assets](https://github.com/curated/assets) - images and related assets
- [docs](https://github.com/curated/docs) - high level documentation
- [domain](https://github.com/curated/domain) - models and persistence
- **github-ql** - github graph client
- **mobile** - open data mobile app
- **scheduler** - github data workers
- **web** - open data responsive app
