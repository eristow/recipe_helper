# Use the official PostgreSQL image from the Docker Hub
FROM postgres:16.4

# Set default environment variables for PostgreSQL
# To set env vars when running Docker container, use the -e flag
# Example: docker run --name my_postgres_container -e POSTGRES_USER=myuser -e POSTGRES_PASSWORD=mypassword -e POSTGRES_DB=mydatabase -p 5432:5432 recipe_helper_database
ENV POSTGRES_USER=defaultuser
ENV POSTGRES_PASSWORD=defaultpassword
ENV POSTGRES_DB=defaultdatabase

# Expose the default PostgreSQL port
EXPOSE 5432