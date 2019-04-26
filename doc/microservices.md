# Micro Services Architecture

[Microsoft doc](https://docs.microsoft.com/en-us/azure/architecture/guide/architecture-styles/microservices)
[Microservices.io](https://microservices.io/patterns/microservices.html)

## What is Micro Services

    A microservices architecture consists of a collection of small, autonomous services. Each services is self-contained and should implement a single business capability.

    Micro Services:
    - Services are small, independent and loosly coupled
    - Each service is a seperate code base
    - Services can be deployed independently
    - Services are responsible for its own data or external state
    - Services communicate with each other by well-defined APIs
    - Services don't need to share same stack, libraries, or framework

    Besides services, some additional components appear in typical microservices architecture.
    - Management
      - Placing services on nodes, identifying failures, rebalancing services across nodes, and so forth
    
    - Service discovery
      - Maintains a list of services and which nodes they are located on.
 
    - API Gateway
      - The entry point for clients. Clients don't call services directly.
      - It decouples clients from services.
      - Services can use message protocol that are not web friendly
      - It can perform other cross-cutting functions such as authentication, logging, SSL termination, and load balancing.

如图所示
<image src="microservices.svg"><image>

## When to use Micro Services Architecture
