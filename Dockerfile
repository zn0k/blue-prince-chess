FROM nginx:alpine

# remove default logs 
RUN rm /var/log/nginx/*

# create a minimal config file
COPY nginx.conf /etc/nginx/nginx.conf

# copy static HTML file into the container
COPY ./index.html /usr/share/nginx/html

EXPOSE 80

# Run nginx in the foreground
CMD ["nginx", "-g", "daemon off;"]
