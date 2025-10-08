#!/bin/bash

# Script to apply all database migrations for jobot
# Usage: ./scripts/apply_migrations.sh

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Load environment variables if .env exists
if [ -f .env ]; then
    echo -e "${BLUE}Loading environment variables from .env${NC}"
    export $(cat .env | grep -v '^#' | xargs)
fi

# Database connection settings with defaults
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-postgres}"
DB_PASSWORD="${DB_PASSWORD:-postgres}"
DB_NAME="${DB_NAME:-jobot}"

# Migrations directory
MIGRATIONS_DIR="./migrations"

echo -e "${BLUE}╔════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║   Database Migration Tool - Jobot     ║${NC}"
echo -e "${BLUE}╚════════════════════════════════════════╝${NC}"
echo ""

# Function to print colored messages
print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

print_info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

# Function to check if psql is installed
check_psql() {
    if ! command -v psql &> /dev/null; then
        print_error "psql is not installed. Please install PostgreSQL client."
        exit 1
    fi
    print_success "psql is installed"
}

# Function to test database connection
test_connection() {
    print_info "Testing database connection..."
    print_info "Host: $DB_HOST:$DB_PORT"
    print_info "Database: $DB_NAME"
    print_info "User: $DB_USER"
    
    if PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c "SELECT 1;" > /dev/null 2>&1; then
        print_success "Successfully connected to PostgreSQL server"
    else
        print_error "Cannot connect to PostgreSQL server"
        exit 1
    fi
}

# Function to create database if it doesn't exist
create_database() {
    print_info "Checking if database '$DB_NAME' exists..."
    
    if PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -lqt | cut -d \| -f 1 | grep -qw $DB_NAME; then
        print_success "Database '$DB_NAME' already exists"
    else
        print_warning "Database '$DB_NAME' does not exist. Creating..."
        if PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d postgres -c "CREATE DATABASE $DB_NAME;" > /dev/null 2>&1; then
            print_success "Database '$DB_NAME' created successfully"
        else
            print_error "Failed to create database '$DB_NAME'"
            exit 1
        fi
    fi
}

# Function to apply a single migration
apply_migration() {
    local migration_file=$1
    local migration_name=$(basename $migration_file)
    
    print_info "Applying migration: $migration_name"
    
    if PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$migration_file" > /dev/null 2>&1; then
        print_success "Migration applied: $migration_name"
        return 0
    else
        print_error "Failed to apply migration: $migration_name"
        return 1
    fi
}

# Function to apply all migrations
apply_all_migrations() {
    print_info "Starting migration process..."
    echo ""
    
    # Check if migrations directory exists
    if [ ! -d "$MIGRATIONS_DIR" ]; then
        print_error "Migrations directory not found: $MIGRATIONS_DIR"
        exit 1
    fi
    
    # Count migrations
    migration_count=$(ls -1 $MIGRATIONS_DIR/*.sql 2>/dev/null | wc -l)
    if [ $migration_count -eq 0 ]; then
        print_warning "No migration files found in $MIGRATIONS_DIR"
        exit 0
    fi
    
    print_info "Found $migration_count migration file(s)"
    echo ""
    
    # Apply migrations in order
    local success_count=0
    local fail_count=0
    
    for migration in $(ls $MIGRATIONS_DIR/*.sql | sort); do
        if apply_migration "$migration"; then
            ((success_count++))
        else
            ((fail_count++))
            print_error "Migration process stopped due to error"
            break
        fi
    done
    
    echo ""
    print_info "Migration Summary:"
    print_success "Successfully applied: $success_count migration(s)"
    
    if [ $fail_count -gt 0 ]; then
        print_error "Failed: $fail_count migration(s)"
        exit 1
    fi
}

# Function to show database tables
show_tables() {
    print_info "Database tables:"
    PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "\dt"
}

# Main execution
echo ""
check_psql
echo ""
test_connection
echo ""
create_database
echo ""
apply_all_migrations
echo ""
print_success "All migrations completed successfully!"
echo ""
show_tables
echo ""
print_info "You can now start using the database"

