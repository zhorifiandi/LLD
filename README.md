# parking-lot-lld
Parking Lot Low Level Design


## Requirements
- Parking lot has multiple floors, multiple entry and exit points.
- Entry is recorded at entry point and payment is done at the time of exit.
- Parking lot contains Multiple floors.
- Parking floor containts multiple parking slots which are of different types.
```
    HANDICAPPED
    CAR
    LARGE
    MOTORBIKE
    ELECTRIC
    EBIKE
```
- Parking is on first come first served basis.
- Time is recorded at the time of issuing the ticket in entry panel.
- Time is recorded at the exit panel and charges are calculated at that point of time.
- User should be able to pay for the ticket.
- Admin should be able to add floor to the lot.
- Admin shoul be able to add parking slots to the floor.
- Admin should be able to remove a slot.
- Admin should be able to add etrance and exit panels.
- User should not be issued a ticket if the slot is not available.
- As the user is allocated a slot the inventory of the slots gets reduced.
- Add user leaves the parking the inventory get updated.

## System Flow
### User Journey
[User Journey System Flow](./diagrams/user-journey-system-flow.png)