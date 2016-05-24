![gekko](https://github.com/tonymtz/gekko/blob/master/static/gekko.png)

# Kekko

Sample project using [Echo](https://github.com/tonymtz/gekko).

## Features
- none

## Quick Start
This lib requires:
- virtualbox 5.0.16
- vagrant Vagrant 1.8.1
- internetz

### Installation
```sh
$ vagrant up
```

### Configuration

Create your config file within `config` directory. Its name must match with the pattern `[env].conf`.
Recommended envs: `dev`, `test`.

Use `env.conf.sample` as base since it has the expected format. Just replace the values!

### Vagrant
```sh
$ vagrant ssh
```

#### Test
```sh
$ export GEKKO_ENV=test && go test
```

#### Running
```sh
$ export GEKKO_ENV=[env] && go run main.go
```
