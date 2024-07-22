# What I'd like to do here

Create intentionally overcomplicated image compressor service, 
which I'll use to learn a bit about gRPC, protobuf and some message queues.

# Design

This is a correct view of what is happening:

```👨‍💻💩```

If we are talking technology-wise:

```
         ┌──────┐                                     
         │CLIENT│                                     
         └──┬───┘                                     
            │                                         
         ┌──▼─────────────┐                           
         │BE FACING CLIENT│                           
         └─┬─────┬───────┬┘                           
           │     │       │                            
           │     │       │                            
┌──────────▼─┐   │    ┌──▼───────────┐                
│BE WITH gRPC│   │    │MESSAGE BROKER│                
└──────────┬─┘   │    └──┬───────────┘                
           │     │       │                            
           │     │    ┌──▼───────────────────────────┐
           │     │    │BE LISTENING TO MESSAGE BROKER│
           │     │    └──┬───────────────────────────┘
           │     │       │                            
           │   ┌─▼┐      │                            
           └───►DB◄──────┘                            
               └──┘                                   
```

This is what I'd like to achieve, but we all know how it goes :-)

# Important addition

This readme is created before the code, so I'll have some reference point, 
so something might change before it's shown in readme

Yes, the .idea files are left intentional