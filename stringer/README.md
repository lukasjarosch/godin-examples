# [godin] stringer
> Powered by Godin v0.5.0 (caee155)

* **Godin init:** 2019-07-17 13:36:44 +0200 CEST
* **Last godin update :** 2019-07-19 12:50:53 +0200 CEST

## gRPC service: GreeterService
**Hello**

*Hello greets you. This comment is also automatically added to the README.*
*Also make sure that all parameters are named, Godin requires this information in order to work.*
```go
func Hello(ctx context.Context, name string) (greeting string, err error)
```
## AMQP subscriptions

All handlers are located in: `./internal/service/subscriber/`
Each handler has it's own file, named after the subscription topic.

| **Routing key** | **Exchange** | **Queue** | **Handler** |
|-----------------|--------------|-----------|-------------|
| user.created | exchange-name | user-created-queue | UserCreated |
| user.deleted | exchange-name | user-deleted-queue | UserDeleted |

## Transport Options
| **Option**      | **Enabled**                                                                          |
|--------------|----------------------------------------------------------------------------------|
| gRPC Transport layer | ![enabled](https://img.icons8.com/color/24/000000/checked.png) |
| gRPC Server | ![enabled](https://img.icons8.com/color/24/000000/checked.png) |
| gRPC Client | ![disabled](https://img.icons8.com/color/24/000000/close-window.png) |
| AMQP Transport | ![enabled](https://img.icons8.com/color/24/000000/checked.png) |
| AMQP Subscriber | ![disabled](https://img.icons8.com/color/24/000000/close-window.png) |
| AMQP Publisher | ![disabled](https://img.icons8.com/color/24/000000/close-window.png) |

## Endpoint middleware

Endpoint middleware is automatically injected by Godin. It is provided by: [go-godin/middleware](github.com/go-godin/middleware)

| **Middleware**      | **Enabled**                                                               |
|--------------|----------------------------------------------------------------------------------|
| InstrumentGRPC |  ![enabled](https://img.icons8.com/color/24/000000/checked.png)
| Logging |  ![enabled](https://img.icons8.com/color/24/000000/checked.png)
| RequestID |  ![enabled](https://img.icons8.com/color/24/000000/checked.png)

## Service middleware

Service middleware is use-case specific middleware which the developer has to take care of.
Godin only assits in creating and maintaining service middleware, but will never overwrite middleware.

| **Middleware**      | **Enabled**                                                               |
|--------------|----------------------------------------------------------------------------------|
| Authorization |  ![disabled](https://img.icons8.com/color/24/000000/close-window.png)
| Caching |  ![disabled](https://img.icons8.com/color/24/000000/close-window.png)
| Logging |  ![enabled](https://img.icons8.com/color/24/000000/checked.png)
| Recovery |  ![disabled](https://img.icons8.com/color/24/000000/close-window.png)
| Monitoring |  ![disabled](https://img.icons8.com/color/24/000000/close-window.png)
## Subscription middleware

Subscription middleware is automatically injected by Godin. It is provided by: [go-godin/middleware/amqp](github.com/go-godin/middleware/amqp)

| **Middleware**      | **Enabled**                                                               |
|--------------|----------------------------------------------------------------------------------|
| RequestID |  ![enabled](https://img.icons8.com/color/24/000000/checked.png)
| Logging |  ![enabled](https://img.icons8.com/color/24/000000/checked.png)
| PrometheusInstrumentation |  ![enabled](https://img.icons8.com/color/24/000000/checked.png)