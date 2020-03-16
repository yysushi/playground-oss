FROM python:3.8.2-alpine3.11 as builder

RUN apk add --no-cache git

WORKDIR /src
COPY . .

COPY --from=mypackage /mypackage/lib/ /usr/local/lib
# RUN pip install -r requirements.txt
RUN python setup.py install
