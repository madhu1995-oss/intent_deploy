FROM tensorflow/tensorflow:1.15.0-py3
ENV BOT_ENV=production

COPY . /var/www
WORKDIR /var/www

RUN pip3 install keras
RUN python3 main.py
ENTRYPOINT [ "-p", "8080"]
