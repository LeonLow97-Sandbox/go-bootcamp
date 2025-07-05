# Cron Job

- A cron job is a scheduled task that runs automatically at specified intervals.
- Cron jobs make it easier to manage recurring activities and maintain the health and efficiency of applications.

## Purposes of Cron Jobs

1. Automated Tasks:
    - Cron jobs are commonly used for automating repetitive tasks. This could include tasks such as data cleanup, log rotation, or any other periodic maintenance tasks.
2. Scheduled Jobs:
    - Might have certain jobs that need to run at specific intervals, such as hourly, daily or weekly. Cron jobs provide a way to schedule and run these jobs automatically.
3. Data Processing:
    - Cron jobs can be used for processing data at regular intervals. For example, you might have a cron job that processes data from a database, file, or external API every hour.
4. Notifications:
    - Can use cron jobs to trigger notifications or alerts based on certain conditions. For instance, sending daily summary emails or notifications based on system metrics.
5. Backup Tasks:
    - Regularly scheduled backups are a common use case for cro jobs. Might want to backup databases, files, or any other critical data on a recurring basis.
6. Periodic reports:
    - Generate and send periodic reports automatically. This could be weekly reports, monthly summaries, or any other type of scheduled reporting.
7. Update Operations:
    - Cron jobs can be used to update data or perform certain operations at specific times. For example, updating currency exchange rates daily.
8. Integration with External APIs:
    - If your application relies on data from external APIs, might want to fetch and update that data at regular intervals. Cron jobs are useful for automating this process.
9. Resource Cleanup:
    - Periodically clean up resources that are no longer needed. This might include deleting temporary files, archiving old data, or cleaning up expired sessions.
10. System Maintenance:
    - Perform system maintenance tasks, such as clearing caches, optimizing databases, or checking for updates.

## Golang `robfig.cron` package

```go
// Creating a new cron instance
c := cron.New()

// adding a cron job function to be run on the specified schedule
// AddFunc(spec string, cmd func()) where spec parameter is a cron expression,
// `cmd` is the function to be executed
c.AddFunc("* * * * *", func() {
    // add task here
})

c.AddFunc("@every 5s", func() {
})

// adding a cron job to be run on the specified schedule.
// AddJob(spec string, cmd Job) where `spec` parameter is a cron expression
// `cmd` is an implementation of the `Job` interface
c.AddJob("* * * * *", MyJob{})

// starts the cron scheduler. The cron jobs will start running on the specified
// schedule after this method is called.
c.Start()

// stops the cron scheduler. This method should be called to gracefully stop the
// cron jobs when they are no longer needed
c.Stop()

// manually trigger the execution of all due cron jobs. This can be useful for testing
// or situation where you want to run the jobs outside the normal schedule.
c.Run()

// sets the time zone for the cron instance. By default, cron uses the UTC time zone.
c.Location(time.UTC)
```