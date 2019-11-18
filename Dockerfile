FROM tensorflow/tensorflow:1.15.0-py3
ENV BOT_ENV=production

COPY . /var/www
WORKDIR /var/www

RUN pip3 install keras
COPY model.h5
RUN from keras.models load_model
RUN model=load_model("model.h5")
ENTRYPOINT [ "-p", "8080"]
