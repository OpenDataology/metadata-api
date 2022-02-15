# portal-backend
The dataset metadata sharing platform backend.

## How to contribute
If you’re interested in contributing code, the best starting point is to have a look at our github issues to see which tasks are the most urgent. 

 accepts PR's (pull requests) from all developers.

Issues can be submitted by anyone - either seasoned developers or newbies. Here is the [code of conduct](https://github.com/dataset-license/community/blob/main/code-of-conduct.md)

## Installation

Deployment Structure

![img](https://github.com/dataset-license/portal-backend/blob/main/docs/images/structure.png?raw=true)

- **Step 1** Setting up the k8s environment, Google GKE or minikube or microk8s are ok for deployment.

- **Step 2** Setting up database and nft storage.

- **Step 3** Setting up `ADDR`, `DSN`, `MAX_IDLE_CONN`environment variables, you can check `/k8s/configMap - template.yaml` for reference.

- **Step 4** Using Dockerfile to build docker image and then upload it to DockerHub.

- **Step 5** Deploy the project by yaml on k8s. And have fun.

## How to use

Portal-backend uses [gin](https://github.com/gin-gonic/gin) and [gorm](https://github.com/go-gorm/gorm) to provide robust webapi, the API Doc is [here](https://github.com/dataset-license/portal-backend/wiki/Portal-beckend-API-Doc). 

Now you can use the webapi to build your own apps.
