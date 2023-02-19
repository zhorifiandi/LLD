# parking-lot-lld
Parking Lot Low Level Design


## Requirements
### Requirement 1
- A Company has a Parking Building, with entrance and exit gate placed side by side
- The Parking Building has 2 Floors with 3 Vehicle slot for each floors
- They want to have a system to handle ticket entrance and ticket exit
- Customer will be assigned to a nearest available slot from the entrance gate
- Each customer will be identified by its Vehicle ID
- A slot become unavailable for customer if it's assigned
- A slot become available if previous customer has leave with their vehicle
- Admin must be able to see current slot assignment visually

### Requirement 2
- A company has done minimal renovattion on the building by adding 3 new slots on floor level 1
- They want to reflect this change on the systems

### Requirement 3
- A company has done a quite renovation on the building by adding 2 floors with 7 slots and 6 slots each
- They want to reflect this change on the systems

### Requirement 4
The company want to start automate the Parking Fee Calculation
Parking Fee per hour is 2000
At the time customer exit, systems needs to show
- How many time has elapsed (1 sec ~= 1 hour)
- Total Fee

### Requirement #5:
- The company decides to sell the software to other company (Company B)
- Company B building specs:
- Floor 0: 6 slots
- Floor 1: 3 slots
- Floor 2: 4 slots
- Floor 3: 5 slots
- They want to initialize the app with input of list of integer
- Where the index represents the floorLevel, and the value represents the slot nums
