# Create an HTTP server to server "wasm micro services"
> This part will depend of the solutions found in the previous exercise(s).


> In Golang, a `sync.Mutex` is a mutual exclusion lock, also known as a mutex, that allows you to synchronize access to shared resources between multiple goroutines. A mutex ensures that **only one goroutine can access a shared resource at a time**, preventing data races and ensuring data integrity.

Here's how a mutex works in Golang:

1. Acquiring the lock: When a goroutine needs to access a shared resource protected by a mutex, it calls the `Lock()` method on the mutex. This effectively "locks" the mutex, preventing other goroutines from accessing the resource until the current goroutine releases the lock.

2. Accessing the resource: Once a goroutine has acquired the lock, it can safely access the shared resource. The goroutine can perform any operations on the resource without worrying about interference from other goroutines.

3. Releasing the lock: After finishing with the shared resource, the goroutine calls the `Unlock()` method to release the lock. This allows other goroutines that are waiting to access the resource to proceed.

Using mutexes ensures that data remains consistent and prevents race conditions, which can lead to unpredictable and buggy behavior in concurrent programs. It's a fundamental tool for writing safe and reliable code in Go.


> In Golang, a goroutine is a lightweight thread of execution that allows you to run multiple operations concurrently within the same program. It's like having multiple threads running simultaneously, but without the overhead of managing threads manually.

**Every request made to the fiber HTTP server will be handled by a goroutine**