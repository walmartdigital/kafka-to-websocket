FROM scratch

ADD /kafka-poc /kafka-poc

ENTRYPOINT [ "/kafka-poc" ]