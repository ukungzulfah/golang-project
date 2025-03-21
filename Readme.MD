docker run --name postgres-container -e POSTGRES_PASSWORD=871206 -p 5432:5432 -v postgres_data:/var/lib/postgresql/data -d postgres

docker exec -it postgres-container psql -U postgres -d analytics

docker run -it --name pgadmin -p 8080:80 -e PGADMIN_DEFAULT_EMAIL=admin@admin.com -e PGADMIN_DEFAULT_PASSWORD=root dpage/pgadmin4# golang-project
