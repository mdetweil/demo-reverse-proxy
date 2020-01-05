package design

import . "goa.design/goa/v3/dsl"

var _ = API("calc", func() {
    Title("Calculator Service")
    Description("A service for adding numbers, a goa teaser")

    // Server describes a single process listening for client requests. The DSL
    // defines the set of services that the server hosts as well as hosts details.
    Server("calc", func() {
        Description("calc hosts the Calculator Service.")

        // List the services hosted by this server.
        Services("calc", "world")

        // List the Hosts and their transport URLs.
        Host("development", func() {
            Description("Development hosts.")
            // Transport specific URLs, supported schemes are:
            // 'http', 'https', 'grpc' and 'grpcs' with the respective default
            // ports: 80, 443, 8080, 8443.
            URI("grpc://localhost:8080")
        })

    })
})

var _ = Service("calc", func() {
    Description("The calc service performs operations on numbers")

    // Method describes a service method (endpoint)
    Method("add", func() {
        // Payload describes the method payload.
        // Here the payload is an object that consists of two fields.
        Payload(func() {
            // Field describes an object field given a field index, a field
            // name, a type and a description.
            Field(1, "a", Int, "Left operand")
            Field(2, "b", Int, "Right operand")
            // Required list the names of fields that are required.
            Required("a", "b")
        })

        // Result describes the method result.
        // Here the result is a simple integer value.
        Result(Int)

        // GRPC describes the gRPC transport mapping.
        GRPC(func() {
            // Responses use a "OK" gRPC code.
            // The result is encoded in the response message (default).
            Response(CodeOK)
        })
    })

    Method("multiply", func() {
        Payload(func() {
            Field(1, "a", Int, "Left")
            Field(2, "b", Int, "Right")
            Required("a", "b")
        })
        Result(Int)

        GRPC(func() {
            Response(CodeOK)
        })
    })
})

var _ = Service("world", func() {
    Description("The world service returns hello world")

    // Method describes a service method (endpoint)
    Method("hello", func() {

        // Result describes the method result.
        // Here the result is a simple integer value.
        Result(String)

        // GRPC describes the gRPC transport mapping.
        GRPC(func() {
            // Responses use a "OK" gRPC code.
            // The result is encoded in the response message (default).
            Response(CodeOK)
        })
    })
})