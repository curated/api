# Curated

Curated is an open source & data platform consisting of a curated dataset of GitHub comments ranked by user reactions for popular repositories. The goal is to provide easy access to meaningful information for anyone interested in learning best practices and development history of popular open source projects.

## Architecture

Curated has a simple ecosystem where **api** and **workers** are two distinct processes sharing **domain** logic while interfacing with different external components: user interfaces and data operations, respectively.

```
api users
         \
mobile ----> api -> domain <- workers -> octograph
         /            |          |           |
web                postgres    redis       github
```

The following repositories compose the project ecosystem:

* **api** - open data endpoints
* [domain](https://github.com/curated/domain) - models and persistence
* **mobile** - open data mobile app
* [octograph](https://github.com/curated/octograph) - github graphql client
* **web** - open data responsive app
* [workers](https://github.com/curated/workers) - data workers

Additional repositories supporting development:

* [coding-standards](https://github.com/curated/coding-standards) - shared configurations across all repositories to help ensure quality and consistency
