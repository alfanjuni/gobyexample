# Defer Keyword in Go

The `defer` keyword postpones the execution of a function until the surrounding function returns. It is commonly used to clean up resources such as closing files or connections. Deferred functions are executed in LIFO (Last In, First Out) order.

In this example, we use `defer` to execute a function after the main function has completed.

---

### Running the Example

Navigate to the directory and run the following command:

```bash
go run examples/golang-25-magic-keywords/15-defer/main.go
```
