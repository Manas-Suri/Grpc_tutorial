# GRPC Client Server on Python - Golang

## Running the system 

1. Make a virtual environment for both client and server folders

```console
python3 -m venv server/venv
source server/venv/bin/activate
pip install -r server/requirements.txt
python3 server/server.py

```
Adding the bash script in another terminal

```console
python3 -m venv client/venv
source client/venv/bin/activate
pip install -r client/requirements.txt
python3 client/client.py
```
1. After running these 2 commands in seperate terminal client and server communication can be seen

Refer to  -  [Link](https://youtu.be/E0CaocyNYKg?si=Ph_nPMz_f7b30Nfq)




Go grpc - 

1. To make the file from .proto 
```bash
protoc --plugin=/home/msuri/go/bin/protoc-gen-go --go_out=. --go-grpc_out=/greeter proto/greeter.proto
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
