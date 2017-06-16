# backplane-poc
Proof of concept app that uses Backplane.io Core (load balancing) and Identity (zero trust identity proxy)





# Build poc binary

It assumes you build the go binary outside constructing the container.  Do
that using the following command:

    GOOS=linux CGO_ENABLED=0 go build -a -o backplane-poc

That will create a binary in the root of the repo.

Now we want to build the docker image and tag it.

    docker build -t mickeyyawn/backplane-poc:build1 .

To test the image and golang binary, crank up a container with this:


    docker run -p 127.0.0.1:8080:8080  --env PORT=8080  --env BACKPLANE_TOKEN=??? -i -d  mickeyyawn/backplane-poc:build1 


If all goes well, you should see a page proclaiming: "hello from backplane poc !"


