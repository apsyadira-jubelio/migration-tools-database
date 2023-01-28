# jb-chat-migarte

> SQL Schema migration tool for [Jubelio Chat](https://github.com/apsyadira-jubelio/jubelio-chat-api).

## Features

- Usable as a CLI tool or as a library
- Supports PostgreSQL
- Migrations are defined with SQL for full flexibility
- Atomic migrations
- Create Schema & run migration for all Tenant or specific tenant
- Run migration for database system

## Installation

To build library and command line program, use the following:

```bash
make build-migration-tools
```

## Usage

### As a standalone tool

```
$ jb-chat-migration --help
usage: jb-chat-migration [--version] [--help] <command> [<args>]

Available commands are:
    all-tenant    Run create schema and migration for all tenants
    tenant        Run create schema and migration for specific tenant
    db-system     Run all migration for database system
```

Each command requires a configuration file `.env`
