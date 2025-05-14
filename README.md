# 🕒 go-cron - Pomodoro Timer

A simple Pomodoro timer implemented in Go.

---

## 📝 Description

**go-cron** is a Pomodoro timer application that alternates between work and short break periods. The default work time is **25 minutes**, and the short break time is **5 minutes**.

This program is perfect for improving productivity by using the Pomodoro Technique, a time management method that encourages working in focused intervals followed by short breaks.

---

## ▶️ Usage

### Steps to Run

1. Clone the repository:

   ```bash
   git clone git@github.com:indraprasetya154/go-cron-pomodoro.git
   ```

2. Navigate to the project directory:

   ```bash
   cd go-cron
   ```

3. Run the timer:

   ```bash
   go run main.go
   ```

4. The timer will start, and you'll see messages indicating the start of work and break periods.

5. Press `Ctrl+C` to stop the timer.

---

## 🧩 Dependencies

This project depends on the following library for scheduling tasks:

* [github.com/go-co-op/gocron/v2](https://github.com/go-co-op/gocron) — A Go package for cron-style job scheduling.

---

## 📂 Project Structure
```
.
├── main.go
├── go.mod
└── go.sum
```
