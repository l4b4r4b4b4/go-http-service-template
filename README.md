# Go HTTP Service Template

This template is inspired by Mat Ryer's blog post ["How I write HTTP services in Go after 13 years"](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/#map-the-entire-api-surface-in-routesgo). Mat Ryer, a seasoned Go developer and host of the Go Time podcast, shares his experience and best practices in writing HTTP services in Go.

## Overview

This template provides a solid starting point for writing HTTP services in Go. It incorporates the following key principles outlined in Mat Ryer's blog post:

- **Server and Handler Structure**: The template is structured for maximum maintainability. It promotes separation of concerns and makes the code easier to navigate and understand.

- **Optimized Startup and Shutdown**: The template includes tips and tricks for optimizing for a quick startup and graceful shutdown. This ensures that your service is always available when needed and shuts down gracefully to prevent data loss or corruption.

- **Common Work Handling**: The template provides a way to handle common work that applies to many types of requests. This reduces code duplication and makes your service more efficient.

- **Deep Testing**: The template includes a robust testing framework. This ensures that your service works as expected and makes it easier to identify and fix bugs.

## Usage

To use this template, simply clone the repository and start building your HTTP service. The template provides a solid foundation, but feel free to modify it to suit your specific needs.

## Conclusion

Whether you're working on a small project or a large one, this template, inspired by Mat Ryer's 13 years of experience, provides a time-tested starting point. Happy coding!

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.