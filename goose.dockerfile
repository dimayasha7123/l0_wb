FROM gomicro/goose:3.7.0

WORKDIR /migrations/
ADD ./migrations/*.sql ./
ADD ./goose_script.sh ./

RUN chmod +x ./goose_script.sh

ENTRYPOINT "./goose_script.sh"