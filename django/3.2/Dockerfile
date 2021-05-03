FROM python:3.9-buster

RUN apt-get update \
  && apt-get install -y mariadb-client

ENV PYTHONUNBUFFERED 1
RUN mkdir /work
WORKDIR /work

COPY ./requirements.txt /work/
RUN pip install -r requirements.txt

ARG DJANGO_3_2_GIT
RUN git clone ${DJANGO_3_2_GIT}