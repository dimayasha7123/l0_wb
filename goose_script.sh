CONN_STRING="host=$POSTGRES_HOST user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB sslmode=disable"
echo "$CONN_STRING"
goose postgres "$CONN_STRING" up
exit 0