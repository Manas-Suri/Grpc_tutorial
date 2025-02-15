# GRPC Client Server on Python - Golang

This is an example of running a python server and golang client for having a grpc connection.

## Server file structure

```bash
grpc_go_client_server 
├── greeter
   ├── users.grpc.pb.go
   ├── users.pb.go
├── proto
   ├── user.proto
├── src
   ├── client.go
├── go.mod
├── go.sum 
```

## Commands to run Go code

Go grpc - 

1. To make the file from .proto 
```bash
protoc --plugin=/home/msuri/go/bin/protoc-gen-go --go_out=. --go-grpc_out=. proto/user.proto
```
2.  If this creates the problem, then first add the protoc-gen-go-grpc   plugin.
To install and follow this plugin, follow these commands - 
 
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
echo $PATH
export PATH=$PATH:$HOME/go/bin
source ~/.bashrc
```
After this, again run the proto command.

## Client file structure

```bash
grpc_go_client_server 
├── greeter
   ├── users.grpc.pb.go
   ├── users.pb.go
├── proto
   ├── user.proto
├── server
   ├── server.py
├── README.md
├── requirements.txt
```

## Running the system 

1. Make a virtual environment for both client and server folders

```console
python3 -m venv server/venv
source server/venv/bin/activate
pip install -r requirements.txt
python3 server/server.py

```
Adding the bash script in another terminal

```console
python3 -m venv client/venv
source client/venv/bin/activate
pip install -r client/requirements.txt
python3 client/client.py
```




Go grpc - 

1. To make the file from .proto 
```bash
protoc --plugin=/home/msuri/go/bin/protoc-gen-go --go_out=. --go-grpc_out=. proto/user.proto
```
2.  If this creates the problem, then first add the protoc-gen-go-grpc   plugin.
To install and follow this plugin, follow these commands - 
 
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
echo $PATH
export PATH=$PATH:$HOME/go/bin
source ~/.bashrc
```
After this, again run the proto command. This will provide the python files 


2. To make the file from .proto for python 
   1. Activate the virtual environment
   2. then install the libraries from requirements.txt
   ```bash
   pip install -r client/requirements.txt
   ```

3. To make the file from .proto

```bash
python -m grpc_tools.protoc -Iprotos --python_out=greeter/ --grpc_python_out=greeter protos/user.proto
``` 

4. Use the command to run the file as the module -  
   ```bash
   python -m server.server
   ``` 
If you receive this error 
'''bash
    import user_pb2 as user__pb2
ModuleNotFoundError: No module named 'user_pb2'
'''

