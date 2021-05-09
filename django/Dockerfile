ARG PYTHON_IMAGE_VERSION
FROM python:${PYTHON_IMAGE_VERSION}

RUN apt-get update \
  && apt-get install -y mariadb-client

ENV PYTHONUNBUFFERED 1
RUN mkdir /work
WORKDIR /work

# COPY ./requirements.txt /work/
ARG DJANGO_VERSION
RUN echo psycopg2 >> ./requirements.txt \
  && echo uwsgi > ./requirements.txt \
  && echo mysqlclient > ./requirements.txt \
  && echo django==${DJANGO_VERSION} > ./requirements.txt
RUN pip install -r requirements.txt
RUN mv requirements.txt /tmp/requirements.txt

# ARG DJANGO_GIT_BRANCH
# ARG DJANGO_GIT_REPOS
# RUN pwd >> /work/pwd.log
# RUN git clone -b $DJANGO_GIT_BRANCH $DJANGO_GIT_REPOS