## Requirement 1

- A fintech company Bandit want to build an internal tool for their employees to easily schedule meeting among themselves.
- [done] Employees can create a meeting with detail of participants, start time, and end time.
- [done] One employee can only participate in one meeting at one particular time.
- [done] If one of participants can't join the meeting, the meeting can't be created.
- [done] Employee can see events between a certain time range.
- Suggest available slots for given list of users for a given day (Optional)

## Requirement 2

- Employee can make group based on their division. One employee can only have one group.
- Meeting participants can be employee or group or both.
- When creating an event, user has an option to specify the number of required representatives (lets call it N) from group (assume that the number of required representatives are same for all group in an event).
- When a group is added, block the slot against the N available members of the group. If the N employees are not available for any group in the event, the event can't be created.
- Suggest available slots for given list of groups for a given day (Optional)

## Requirement 3

- Update meeting time - if it is not possible return the status accordingly
- Support recurring meeting (start_time, end_time, date, frequency (weekly, monthly, yearly))