FROM ubuntu

RUN apt-get update && apt-get install -y bash

COPY BBDB/BBDB.go /bbdatabase
RUN chmod +x /bbdatabase

EXPOSE 3333

CMD ["/bbdatabase"]
