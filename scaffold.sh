#!/bin/sh
# Usage:
#   ./scaffold.sh component <name>  # create/remove domain object stack
#   ./scaffold.sh adapter <name>    # create/remove infrastructure adapter
#   ./scaffold.sh remove component <name>
#   ./scaffold.sh remove adapter <name>

ACTION=$1
TYPE=$2
NAME=$3

if [ -z "$ACTION" ] || [ -z "$TYPE" ] || [ -z "$NAME" ]; then
  echo "Usage: $0 create|remove component|adapter <name>"
  exit 1
fi

# Capitalize first letter (portable)
capitalize() {
  echo "$1" | awk '{print toupper(substr($0,1,1)) substr($0,2)}'
}

COMPONENT=$(echo "$NAME" | tr '[:upper:]' '[:lower:]')
COMPONENT_CAP=$(capitalize "$COMPONENT")
MODULE_NAME=$(basename "$(pwd)")

# Directories
DOMAIN_DIR="internal/domain"
PORTS_DIR="internal/application/ports"
SERVICES_DIR="internal/application/services"
REPO_DIR="internal/infrastructure/repository"
HANDLER_DIR="internal/presentation/http"
ADAPTER_DIR="internal/infrastructure"

mkdir -p "$DOMAIN_DIR" "$PORTS_DIR" "$SERVICES_DIR" "$REPO_DIR" "$HANDLER_DIR" "$ADAPTER_DIR"

# Helper functions
create_file() {
  filepath=$1
  content=$2
  if [ ! -f "$filepath" ]; then
    echo "$content" > "$filepath"
    echo "Created $filepath"
  else
    echo "File $filepath already exists, skipping..."
  fi
}

remove_file() {
  filepath=$1
  if [ -f "$filepath" ]; then
    rm "$filepath"
    echo "Removed $filepath"
  fi
}

# -------------------------------
# CREATE
# -------------------------------
if [ "$ACTION" = "create" ]; then

  if [ "$TYPE" = "component" ]; then
    # Domain
    create_file "$DOMAIN_DIR/${COMPONENT}.go" "package domain

type ${COMPONENT_CAP} struct {
    // TODO: Add fields
}
"
    create_file "$DOMAIN_DIR/${COMPONENT}_test.go" "package domain

import \"testing\"

func Test${COMPONENT_CAP}(t *testing.T) {
    // TODO: Add tests
}
"

    # Ports (repository interface)
    create_file "$PORTS_DIR/${COMPONENT}.go" "package ports

// TODO: Define repository interface for ${COMPONENT_CAP}
type ${COMPONENT_CAP}Repository interface {
    // TODO: Add methods
}
"

    # Service
    create_file "$SERVICES_DIR/${COMPONENT}_service.go" "package services

import \"${MODULE_NAME}/internal/application/ports\"

// TODO: Implement service for ${COMPONENT_CAP}
type ${COMPONENT_CAP}Service struct {
    repo ports.${COMPONENT_CAP}Repository
}
"

    # Repository implementation
    create_file "$REPO_DIR/pg_${COMPONENT}_repository.go" "package repository

// TODO: Implement Postgres repository for ${COMPONENT_CAP}
type PG${COMPONENT_CAP}Repository struct {}
"

    # Handler
    create_file "$HANDLER_DIR/${COMPONENT}_handler.go" "package http

import \"${MODULE_NAME}/internal/application/services\"

// TODO: Implement HTTP handler for ${COMPONENT_CAP}
type ${COMPONENT_CAP}Handler struct {
    service *services.${COMPONENT_CAP}Service
}
"

  elif [ "$TYPE" = "adapter" ]; then
    ADAPTER_PATH="$ADAPTER_DIR/${COMPONENT}"
    mkdir -p "$ADAPTER_PATH"

    # Ports interface
    create_file "$PORTS_DIR/${COMPONENT}.go" "package ports

// TODO: Define interface for ${COMPONENT_CAP} adapter
type ${COMPONENT_CAP} interface {
    // TODO: Add methods
}
"

    # Adapter implementation
    create_file "$ADAPTER_PATH/${COMPONENT}.go" "package ${COMPONENT}

// TODO: Implement ${COMPONENT_CAP} adapter
type ${COMPONENT_CAP}Impl struct {}
"

  else
    echo "Unknown type: $TYPE"
    exit 1
  fi

  echo "✅ Scaffold for '$TYPE' '$COMPONENT' created successfully!"

# -------------------------------
# REMOVE
# -------------------------------
elif [ "$ACTION" = "remove" ]; then

  if [ "$TYPE" = "component" ]; then
    remove_file "$DOMAIN_DIR/${COMPONENT}.go"
    remove_file "$DOMAIN_DIR/${COMPONENT}_test.go"
    remove_file "$PORTS_DIR/${COMPONENT}.go"
    remove_file "$SERVICES_DIR/${COMPONENT}_service.go"
    remove_file "$REPO_DIR/pg_${COMPONENT}_repository.go"
    remove_file "$HANDLER_DIR/${COMPONENT}_handler.go"

  elif [ "$TYPE" = "adapter" ]; then
    remove_file "$PORTS_DIR/${COMPONENT}.go"
    rm -rf "$ADAPTER_DIR/${COMPONENT}" && echo "Removed $ADAPTER_DIR/${COMPONENT}"

  else
    echo "Unknown type: $TYPE"
    exit 1
  fi

  echo "✅ Scaffold for '$TYPE' '$COMPONENT' removed successfully!"

else
  echo "Unknown action: $ACTION"
  exit 1
fi

