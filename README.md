![bitemyshinymetalGPT2](https://user-images.githubusercontent.com/58280110/112588630-b63a0580-8dcd-11eb-8899-2e461298ca06.png)


## BiteMyShinyMetalGPT-2


Scripts for our dataset were scraped from imdbs.com using the golang package Colly.

After cloning the repository, you can construct the docker image by navigating into the gpt-2 folder and running 

```docker build -f Dockerfile.cpu . -t 'your_cool_title:latest'```

With this, you can access and run gpt-2 by running 

```docker run -it --rm your_cool_title```


