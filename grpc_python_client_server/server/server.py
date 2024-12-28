from concurrent import futures
import grpc
from greeter import user_pb2
from greeter import user_pb2_grpc


class User1(user_pb2_grpc.GreeterServicer):
    def GetUsers(self, request, context):
        users = [
            user_pb2.User(
                id='1',
                name='John Doe',
                email='manassuri98.ms',
                password='manas123',
                phone='123-456-7890',
                address='123 Main St',
                role='user',
                status='active',
                created_at='2024-08-09',
                updated_at='2024-08-09'
            ),
            user_pb2.User(
                id='2',
                name='Manas Suri',
                email='manassuri98.ms',
                password='manas123',
                phone='123-456-7890',
                address='123 Main St',
                role='user',
                status='active',
                created_at='2024-08-09',
                updated_at='2024-08-09'
            ),

        ]
        return user_pb2.GetUsersResponse(users=users)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    user_pb2_grpc.add_GreeterServicer_to_server(User1(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()

