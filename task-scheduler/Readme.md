# Low-Level Design Problem Statement: Task Scheduler

## Problem Statement

Design and implement a **Task Scheduler** that can manage and execute tasks efficiently. The system should support scheduling tasks with different execution types, priorities, and delays. It should also allow users to stop tasks and handle concurrent task execution using a worker pool.

---

## Requirements

### Functional Requirements

1. **Task Management**:

   - Users should be able to create tasks with attributes such as:
     - Name
     - Creator
     - Priority (High, Medium, Low)
     - Execution Type (One-Time, Fixed Delay, Fixed Rate)
     - Initial Delay and Recurring Delay (if applicable)
   - Users should be able to schedule tasks for execution.
   - Users should be able to stop tasks before they are executed.

2. **Task Execution**:

   - Tasks should be executed based on their scheduled execution time.
   - Tasks with higher priority should be executed before lower-priority tasks if their execution times are the same.
   - Support for recurring tasks:
     - **Fixed Delay**: Execute the task after a fixed delay from the previous execution's completion.
     - **Fixed Rate**: Execute the task at a fixed interval, regardless of the previous execution's completion.

3. **Concurrency**:

   - Use a worker pool to execute multiple tasks concurrently.
   - Ensure thread safety when managing tasks and scheduling.

4. **Graceful Shutdown**:
   - Provide a mechanism to stop the scheduler and ensure all running tasks are completed before shutting down.

---

### Non-Functional Requirements

1. **Scalability**:

   - The system should handle a large number of tasks efficiently.
   - The worker pool size should be configurable to optimize resource usage.

2. **Fault Tolerance**:

   - Handle task failures gracefully and retry tasks if necessary.

3. **Thread Safety**:

   - Ensure thread-safe operations for task scheduling, execution, and stopping.

4. **Extensibility**:
   - The system should be designed to allow future enhancements, such as adding new task types or execution policies.

---

## Constraints

1. Tasks are represented as objects with attributes like `Name`, `Priority`, `ExecutionTime`, and `Type`.
2. Tasks are stored in a priority queue (min-heap) to ensure they are executed in the correct order.
3. The worker pool size is fixed and configurable.
4. The system should use in-memory data structures (no database).

---

## Example Scenarios

### Scenario 1: One-Time Task

- A user schedules a one-time task to execute after 10 seconds.
- The task is executed once and removed from the scheduler.

### Scenario 2: Fixed Delay Task

- A user schedules a task with a fixed delay of 30 seconds.
- The task is executed repeatedly, with a 30-second delay between the completion of one execution and the start of the next.

### Scenario 3: Fixed Rate Task

- A user schedules a task with a fixed rate of 1 minute.
- The task is executed repeatedly at 1-minute intervals, regardless of the previous execution's completion time.

### Scenario 4: Stopping a Task

- A user schedules multiple tasks and decides to stop one of them before it is executed.
- The task is removed from the scheduler and will not be executed.

---
