# Curated

Curated is an open source & data platform consisting of a curated dataset of GitHub comments ranked by user reactions for popular repositories. The goal is to provide easy access to meaningful information for anyone interested in learning best practices and development history of popular open source projects.

## Architecture

Curated has a simple ecosystem where **services** and **workers** are distinct processes sharing **domain** logic while interfacing with different external components: user interfaces and data operations, respectively.

```
api users
         \
mobile ----> services -> domain <- workers -> octograph
         /                 |          |           |
web                     postgres    redis       github
```

The following repositories compose the project ecosystem:

* **core**
  * [domain](https://github.com/curated/domain) - models and persistence
  * [octograph](https://github.com/curated/octograph) - github graphql client
* **services**
  * [issue-service](https://github.com/curated/issue-service) - open data issue endpoint
* **workers**
  * [issue-worker](https://github.com/curated/issue-worker) - github issue data worker
* **interfaces**
  * mobile - react native?
  * web - next.js?

Additional repositories supporting development:

* [coding-standards](https://github.com/curated/coding-standards) - shared configurations across all repositories to help ensure quality and consistency
