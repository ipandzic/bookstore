#!/usr/bin/env bash

echo "\n########## MIGRATE DB ##########\n"
sql-migrate up -config=config/migrations.yml