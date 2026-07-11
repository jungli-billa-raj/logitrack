Okay so here is what I"m understanding. 

1. We created our own Server struct to include GO's default Multiplexer and other things that we may need. Here, the repo interface. 
2. We could also add other fields like configurations, database url,port etc if we wanted but we have intentionally kept it simple and minimal here. 
3. To make our custom Server and actual server in GO's eyes to make it allow using it in ServerAndRun() some function like that, we need to implement the default server's interface. 
4. This interface has only one function ServeHTTP. We write our own implementation for this to satisfy the interface contract where we sneak it the pointer to request body and a writer to GO's default multiplexer. 
5. Lastly, we also need the multiplexer know what paths we are going to serve and what functions to run for each. For this we called the HandleFunc() function and registered the path and associated functions. 