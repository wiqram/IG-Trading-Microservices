FROM java:8

# set the working directory in the container
WORKDIR /var/www/java

#COPY --from=builder /usr/local/ig-webapi-java-sample-console/target/ig-webapi-java-sample-console-2.2.0-SNAPSHOT.jar /usr/app/ig-webapi-java-sample-console/target/ig-webapi-java-sample-console-2.2.0-SNAPSHOT.jar
#COPY --from=builder /usr/local/ig-webapi-java-sample-console-ui/target/ig-webapi-java-sample-console-ui-2.2.0-SNAPSHOT.jar /usr/app/ig-webapi-java-sample-console-ui/target/ig-webapi-java-sample-console-ui-2.2.0-SNAPSHOT.jar
#COPY --from=builder /app/ig-webapi-java-client/target/ig-webapi-java-client-2.2.0-SNAPSHOT.jar /ig-webapi-java-client/target/ig-webapi-java-client-2.2.0-SNAPSHOT.jar
#COPY --from=builder /app/ig-webapi-java-client/target/ig-webapi-java-client-2.2.0-SNAPSHOT-jar-with-dependencies.jar /ig-webapi-java-client/target/ig-webapi-java-client-2.2.0-SNAPSHOT-jar-with-dependencies.jar
#COPY --from=builder /app/ig-webapi-java-client/lib/ls-client.6.0.2.20160510.jar /ig-webapi-java-client/lib/ls-client.6.0.2.20160510.jar
#COPY --from=builder /usr/local/ .

# copy the dependencies file to the working directory
#COPY . /var/www/java

# install dependencies
#RUN pip install -r requirements.txt

# copy the content of the local src directory to the working directory
COPY . .

# command to run on container start
EXPOSE 50055
#CMD ["java", "-jar", "-Dspring.profiles.active=default", "ig-webapi-java-client/target/ig-webapi-java-client-2.2.0-SNAPSHOT.jar"]
CMD ["java", "-cp", "ig-webapi-java-client/target/ig-webapi-java-client-2.2.0-SNAPSHOT-jar-with-dependencies.jar", "com.iggroup.webapi.samples.client.grpc.GrpcServer"]
#ENTRYPOINT ["java", "-cp", "ig-webapi-java-client/target/ig-webapi-java-client-2.2.0-SNAPSHOT.jar", "com.iggroup.webapi.samples.client.grpc.App"]
#CMD [ "java","-jar","ig-webapi-java-sample-console-ui/target/ig-webapi-java-sample-console-ui-2.2.0-SNAPSHOT.jar wiqramdemo DemoPWD123 d56147107609d1d2c1e65d8b54a011eff305441b" ]