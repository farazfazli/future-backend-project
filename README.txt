Appointment Scheduling - Future Backend Engineering Project

===============================
CONTENTS OF THIS FILE

* INTRODUCTION
* PREREQUISITES & SETUP
* IMPORTANT INFORMATION
===============================

INTRODUCTION
------------

This is a take-home assignment completed using Go and Postgres.

PREREQUISITES & SETUP
---------------------

Have Go and Postgres installed. Copy the .env.template file to .env and adjust your Postgres info so that the SQL code and program both run properly.

I used a Makefile to simplify the seeding, table creation, and build process. Simply run 'make' and the following will happen *automatically*:

1. sqlc pipeline will run for auto-generating structs based on the SQL schema
2. existing tables/views/functions will be dropped (cascaded drop)
3. sql/schema.sql will be run, which creates a fresh table, view, and function
4. a few appointments (mock data) will be seeded into the database using sql/mock.sql
5. the application will be built (using go build)

After this process is completed, you can run the API using ./cmd/api/api

IMPORTANT INFORMATION
---------------------

1. The OpenAPI specification is available in api-spec/api-spec.yml and can be opened in any compatible REST client (i.e. Insomnia).
2. The database defaults to Los Angeles time (to match with the project specification). This is done while creating the schema, and can be changed as needed.
3. All dependencies are vendored. This is done to ease running a replicable build in a fresh environment without needing much time.
4. We intentionally provide minimal error information to the client (a practice OWASP calls a "generic response"). This is to prevent
reconnaissance by potential attackers. Detailed logging is done on the backend however, so we can see exactly why a particular response code is being thrown.