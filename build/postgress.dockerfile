
FROM postgres:latest

ENV POSTGRES_DB=users
ENV POSTGRES_USER=gorm
ENV POSTGRES_PASSWORD=gorm

EXPOSE 5432
#  psql -U gorm -d users