FROM rabbitmq:3.6

ENV RABBITMQ_DEFAULT_USER guest
ENV RABBITMQ_DEFAULT_PASS guest

RUN rabbitmq-plugins enable --offline rabbitmq_management
RUN rabbitmq-plugins enable --offline rabbitmq_mqtt

# COPY rabbitmq-env.conf /etc/rabbitmq/rabbitmq-env.conf
COPY docker/rabbitmq/rabbitmq.config /etc/rabbitmq/rabbitmq.config
RUN chown -R rabbitmq:rabbitmq /var/lib/rabbitmq /etc/rabbitmq &&\
    chmod 777 /var/lib/rabbitmq /etc/rabbitmq

EXPOSE 15672
EXPOSE 5672
EXPOSE 1883