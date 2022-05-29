General Information
-------------------
The server on which the prediction error energy for prediction orders (from 2 to 15) is calculated from the .wav file loaded into the database. The calculation is based on three methods: autocorrelation (Darbin method), covariance (Holetsky method) and staircase (Burg method).\
At the first stage, the decoder module decodes the file - using the api/databases/mysql, the .wav file from the mysql database is requested. Further, in the algorithm module, linear prediction coefficients are calculated on signal segments of 20 ms, by which prediction error energy values. The output is the average value of the array of energies of the prediction error.\
Event logging is performed in .log logger/logs files.\
The <a href="https://swagger.io/">Swagger</a> library was used to write the server. The server logic is located in the gen folder. More information about writing the server can be found <a href="https://github.com/go-swagger/go-swagger">here</a>. Files downloaded from the database are saved to the upload folder.\
Methods of adding and deleting files, viewing a list of files in the database, viability (healthz), calculating the energy of prediction errors (3 methods) are implemented.

-------------------
### Getting Started
- First, you need to run the ConfigMap file, in which you specify the necessary input data in addition to the ip pod mysql (the field will have to be changed later).\
- Before starting calculations, you need to start the mysql database (volume, secret and deployment - in order).\
- In .yaml of volume, you must specify the size of the allocated memory (storage). In secret.yaml, write a password.\
- Next, you must apply a job with <a href="https://github.com/Vetalb60/ErrorsPredictCalculator/tree/main/mig">migrations</a> to organize tables in the database.\
- After implementing the database, you can start the demonstration of server.\
- Communication with the server is carried out using the <a href="https://github.com/Vetalb60/ErrorCalculateClient">client</a>.

Run:
-------------------
### 1.mysql

    kubectl apply -f configmap.yaml
    kubectl apply -f secret.yaml
    kubectl apply -f mysql-volume.yaml
    kubectl apply -f mysql-deployment.yaml

2. Execute the mysql <a href="">migrations</a>.

### 2.server

    docker build -t {YOUR DOCKER HUB PROFILE}/server:tag
    docker push {YOUR DOCKER HUB PROFILE}/server:tag
    kubectl apply -f server-deployment.yaml
    kubectl apply -f server-service.yaml

### 3. Helm

    helm secrets install mysql ./k8s/mysql/mysql-workflow/ -n default -f ./k8s/mysql/mysql-workflow/secrets.yaml
    helm install server ./k8s/server/server-workflow

-------------------
### Modules

- <a href="https://github.com/Vetalb60/ErrorsPredictCalculator">server</a>
- <a href="https://github.com/Vetalb60/ErrorsPredictCalculator/tree/main/mig">migrations</a>
- <a href="https://github.com/Vetalb60/ErrorCalculateClient">client</a>

-------------------
### Get Help
Send questions to gmail: al9xgr99n@gmail.com

Author:Alex Green

-------------------