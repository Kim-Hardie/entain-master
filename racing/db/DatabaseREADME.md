# Database README

This document provides information about the database setup and changes made to the scripts for the racing service.

## Changes Made

The following changes have been made to the database scripts:

1. Added a `status` field to the `races` table: A new column named `status` has been added to the `races` table. This field represents the status of a race and can have two values: "OPEN" if the `advertised_start_time` is in the future, and "CLOSED" if it is in the past.

2. Modified the `seed` function in `db.go`: The method for adding races to the database has been updated to include the `status` field. Now, when races are inserted into the `races` table, the `status` field is automatically set based on the `advertised_start_time`.

3. Included an SQL script to update existing data: A standalone SQL script named `update_status.sql` has been provided. This script can be executed to update the `status` field for all existing races in the database. It checks the `advertised_start_time` and sets the `status` to "OPEN" or "CLOSED" accordingly.

## Running the SQL Script

To update the existing races in the database and add the `status` field, follow these steps:

1. Make sure you have a backup of your database in case of any issues during the update process.

2. Open your preferred SQLite client or command-line tool.

3. Execute the `update_status.sql` script against your database. This script will update the `status` field for races where the status is currently `NULL`, based on the `advertised_start_time`.

   Example using the SQLite command-line tool:
   ```bash
   $ sqlite3 path/to/your/database.db < update_status.sql


# Sports Database
This database is created to store and manage sports matches data. It provides the functionality to create a new match, get a match by its ID, and filter matches based on specific criteria.

## Database Schema
Database Schema
The database consists of a single table named matches. The table schema is as follows:

```sql
CREATE TABLE IF NOT EXISTS matches (
id INTEGER PRIMARY KEY AUTOINCREMENT,
name TEXT,
stadium TEXT,
sport TEXT,
team1 TEXT,
team2 TEXT,
time TIMESTAMP
)
```
### Features
- Init(): Initializes the database by creating the matches table if it doesn't exist.
- CreateMatch(match *Match): Inserts a new match into the database.
- GetMatchByID(id int): Fetches a match from the database using its ID.