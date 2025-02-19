#!/bin/bash

# Directory where migrations will be created
MIGRATIONS_DIR="migrations/sqlite"

# Create the migrations directory if it doesn't exist
mkdir -p "$MIGRATIONS_DIR"

# List of migration names (replace with your actual table names)
MIGRATION_NAMES=(
  "create_users_table"
  "create_follows_table"
  "create_posts_table"
  "create_comments_table"
  "create_events_table"
  "create_engagement_table"
  "create_notifications_table"
  "create_sessions_table"
  "create_groups_table"
  "create_group_members_table"
  "create_chats_table"
  "create_messages_table"
  "add_foreign_keys"
  "create_indexes"
)

# Loop through the migration names and create files
for ((i=0; i<${#MIGRATION_NAMES[@]}; i++)); do
  # Generate the migration number (e.g., 000001, 000002, etc.)
  MIGRATION_NUMBER=$(printf "%06d" $((i+1)))

  # Define the file names
  UP_FILE="${MIGRATIONS_DIR}/${MIGRATION_NUMBER}_${MIGRATION_NAMES[$i]}.up.sql"
  DOWN_FILE="${MIGRATIONS_DIR}/${MIGRATION_NUMBER}_${MIGRATION_NAMES[$i]}.down.sql"

  # Create the up.sql file
  echo "-- ${MIGRATION_NAMES[$i]} UP" > "$UP_FILE"
  echo "" >> "$UP_FILE"

  # Create the down.sql file
  echo "-- ${MIGRATION_NAMES[$i]} DOWN" > "$DOWN_FILE"
  echo "" >> "$DOWN_FILE"

  # Print success message
  echo "Created migration files:"
  echo "  - $UP_FILE"
  echo "  - $DOWN_FILE"
done

echo "All migration files created successfully in $MIGRATIONS_DIR."