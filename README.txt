Appointment Scheduling - Future Backend Engineering Project

===============================
CONTENTS OF THIS FILE

* INTRODUCTION
* PREREQUISITES & SETUP
* IMPORTANT INFORMATION
===============================

INTRODUCTION
------------

This is a take-home assignment completed using Go and Postgres. Docs can be viewed by opening docs.html

PREREQUISITES & SETUP
---------------------

Step 1: Ensure you have a recent version of Go and Postgres installed.
Step 2. Copy the .env.template file to .env and adjust your Postgres environment variables so that the SQL code and server both run properly.
Step 3. Run 'make'

The Makefile simplifies the seeding, table creation, and build process. The following steps will happen *automatically*:

1. The sqlc pipeline will be executed for auto-generating structs based on the SQL schema.
2. Existing tables/constraints as well as the view and function will be dropped (cascade drops dependent objects).
3. The file sql/schema.sql will be executed, which creates fresh tables, a view, and a function.
4. A few members, trainers, and appointments (mock data) will be seeded into the database using the sql/mock.sql file.
5. The application will be built (using go build).

After this process is completed, you can run the API using ./cmd/api/api

IMPORTANT INFORMATION
---------------------

1. The OpenAPI specification is available in api-spec/openapi.yml (JSON file is also available which contains the same data) and can be opened in any compatible REST client (e.g. Insomnia).
2. Stoplight Elements is used to view docs from the OpenAPI definition, it includes a "Try it" button for endpoints as well as example code.
2. The database defaults to Los Angeles time (to match with the project specification). This is done in schema.sql before creating the schema, and can be changed as needed.
3. All dependencies are vendored. This is done to ease running a replicable build in a fresh environment without needing much time.
4. I intentionally provided minimal error information to the client (a practice OWASP calls a "generic response"). By leaking minimum info, the thought is to prevent reconnaissance by potential attackers. Detailed logging is done on the backend however, so we can see exactly why a particular response code is being thrown. Also, the API spec has example respones for each cooresponding HTTP status code.