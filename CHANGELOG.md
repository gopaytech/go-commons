# [0.11.0](https://github.com/gopaytech/go-commons/compare/v0.10.2...v0.11.0) (2021-07-27)


### Features

* add seed functionality, add capabilities to scan template with custom delims ([dae9d01](https://github.com/gopaytech/go-commons/commit/dae9d01b56a79e5c31fd95ee93edd27ebbe2cb0c))

## [0.10.2](https://github.com/gopaytech/go-commons/compare/v0.10.1...v0.10.2) (2021-07-21)


### Bug Fixes

* prevent nil error when check file.FileExists ([638da00](https://github.com/gopaytech/go-commons/commit/638da00677c30722638acd4dd1e991ad2eaca9b2))

## [0.10.1](https://github.com/gopaytech/go-commons/compare/v0.10.0...v0.10.1) (2021-07-21)


### Bug Fixes

* add mock and fix test for git clone ([9d4937d](https://github.com/gopaytech/go-commons/commit/9d4937da1de6e7cf062e008698bc4d0273c2af75))

# [0.10.0](https://github.com/gopaytech/go-commons/compare/v0.9.0...v0.10.0) (2021-07-16)


### Features

* complete git pkg features ([5232d78](https://github.com/gopaytech/go-commons/commit/5232d785edc62ab2dd7eec2256c3feae9a62b996))

# [0.9.0](https://github.com/gopaytech/go-commons/compare/v0.8.4...v0.9.0) (2021-07-16)


### Features

* add clone function for git ([399ece3](https://github.com/gopaytech/go-commons/commit/399ece3d89b52cccf1b076960cae386b8b90bd6b))

## [0.8.4](https://github.com/gopaytech/go-commons/compare/v0.8.3...v0.8.4) (2021-07-14)


### Bug Fixes

* refactor templ structure ([a1fd639](https://github.com/gopaytech/go-commons/commit/a1fd639a008ea2d5339a127d0397ff5986277f01))

## [0.8.3](https://github.com/gopaytech/go-commons/compare/v0.8.2...v0.8.3) (2021-07-11)


### Bug Fixes

* add Type alias for LogFields ([3e60c61](https://github.com/gopaytech/go-commons/commit/3e60c616d30a975bf4c08c5b576f74020bfdaf96))

## [0.8.2](https://github.com/gopaytech/go-commons/compare/v0.8.1...v0.8.2) (2021-07-08)


### Bug Fixes

* add fields parameter for zlog loggers ([2c2ae64](https://github.com/gopaytech/go-commons/commit/2c2ae64982b79e5b41b33dc210c9d996f2dc28ff))

## [0.8.1](https://github.com/gopaytech/go-commons/compare/v0.8.0...v0.8.1) (2021-07-08)


### Bug Fixes

* asynq logger ([409c804](https://github.com/gopaytech/go-commons/commit/409c804888c37b0d749417a79f4f106b72ff3f5e))

# [0.8.0](https://github.com/gopaytech/go-commons/compare/v0.7.2...v0.8.0) (2021-07-08)


### Features

* add log for asynq, add missing fatal function ([3396cb3](https://github.com/gopaytech/go-commons/commit/3396cb33a97c7369349a3bcb13414abc9f3739f4))

## [0.7.2](https://github.com/gopaytech/go-commons/compare/v0.7.1...v0.7.2) (2021-06-30)


### Bug Fixes

* modify README to trigger version update ([5dcbd32](https://github.com/gopaytech/go-commons/commit/5dcbd32e2f135e5c5b7f8d540beafe39a8dab199))

## [0.7.1](https://github.com/gopaytech/go-commons/compare/v0.7.0...v0.7.1) (2021-06-30)


### Bug Fixes

* export var instead of struct ([c94a806](https://github.com/gopaytech/go-commons/commit/c94a806a7dd670ea1dbfd1432dcb27c544c295a2))

# [0.7.0](https://github.com/gopaytech/go-commons/compare/v0.6.1...v0.7.0) (2021-06-30)


### Features

* add zlog support for asynq ([25b81a6](https://github.com/gopaytech/go-commons/commit/25b81a61fb89cfb3b138057e348c47e3256b5d8a))

## [0.6.1](https://github.com/gopaytech/go-commons/compare/v0.6.0...v0.6.1) (2021-06-24)


### Bug Fixes

* add zerolog function to return event ([245809f](https://github.com/gopaytech/go-commons/commit/245809f8ed56c4e8b92a0a4f3f299cf6c18a37e6))

# [0.6.0](https://github.com/gopaytech/go-commons/compare/v0.5.0...v0.6.0) (2021-06-23)


### Features

* separate db migrate to mysql and postgresql ([72a667e](https://github.com/gopaytech/go-commons/commit/72a667ebd1c42c3f1c0a8c2c3f0cf24403f6ecaf))

# [0.5.0](https://github.com/gopaytech/go-commons/compare/v0.4.0...v0.5.0) (2021-06-16)


### Features

* add orm client for postgresql ([3f0b1eb](https://github.com/gopaytech/go-commons/commit/3f0b1ebff2010fd806f4f8301b36b62eea41a3be))

# [0.4.0](https://github.com/gopaytech/go-commons/compare/v0.3.0...v0.4.0) (2021-05-07)


### Features

* add job helper for k8s ([45931a9](https://github.com/gopaytech/go-commons/commit/45931a90fecea78db9a3e8562e6f70a11526dfcc))

# [0.3.0](https://github.com/gopaytech/go-commons/compare/v0.2.2...v0.3.0) (2021-05-07)


### Features

* separate docker test to local tag, to exclute those test on gitlab workflow ([4a01a26](https://github.com/gopaytech/go-commons/commit/4a01a261bb6354131f30bb87e24894966ce4eea5))

## [0.2.2](https://github.com/gopaytech/go-commons/compare/v0.2.1...v0.2.2) (2021-05-06)


### Bug Fixes

* add sudo when install apt-get ([d6b633e](https://github.com/gopaytech/go-commons/commit/d6b633e9a6b5acf8278e926ebc4e58eae0e54aee))

## [0.2.1](https://github.com/gopaytech/go-commons/compare/v0.2.0...v0.2.1) (2021-05-06)


### Bug Fixes

* make sure github build run on vm ([483074c](https://github.com/gopaytech/go-commons/commit/483074cf41cd55f1fa9325160d65cf915e77c3a1))
