# Golang Method Set Examples

This project is a collection of **Method Set** examples in Go, designed for learning and interview preparation.

---

## 📂 Project Structure

```
golang-methodset/
│
├── 01_basics/                # Simple struct and method examples
├── 02_value_receiver/        # Methods with value receivers
├── 03_pointer_receiver/      # Methods with pointer receivers
├── 04_methodset_table/       # Comparison table of method sets
├── 05_interface_impl/        # Effect of method sets on interfaces
├── 06_embedding/             # Embedding and method promotion
├── 07_interface_rules/       # Why interfaces cannot have method bodies
├── 08_value_vs_pointer/      # Struct with both value and pointer methods
└── 09_interview_exercises/   # Tricky interview exercises
```

---

## 🚀 Running the Examples

Each example contains a `main.go` file and can be run using:

```bash
cd 01_basics
go run main.go
```

---

## 📖 Key Points

- **Method Set** determines which methods are accessible on a type.
- **Value receiver methods** can be called on both values and pointers.
- **Pointer receiver methods** can only be called on pointers.
- For **interfaces**, the method set determines if a type implements the interface.
- With **embedding**, methods from the embedded struct are promoted to the outer struct.

---

## 💻 Interview Exercises

The `09_interview_exercises` folder contains tricky exercises.  
Try to predict the output before looking at the solutions.

---