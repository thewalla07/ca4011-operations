set CARS; # car types
param rate {CARS} > 0; # cars per hour
param profit {CARS}; # profit per car
param commit {CARS} >= 0; # lower limit on car demands
param max_hours > 0; # max factory hours

    # param avail {STAGE} >= 0; # hours available/week in each stage
    # param rate {PROD,STAGE} > 0; # tons per hour in each stage
    # set STAGE; # stages


    # param commit {PROD} >= 0; # lower limit on tons sold in week
    # param market {PROD} >= 0; # upper limit on tons sold in week

var Make {c in CARS} >= commit[c]; # cars produced
var Prof {c in CARS} == Make[c] * profit[c];
# objective function
maximize Total_Prod: sum {c in CARS} Make[c];

# constraints
subject to Time:
    max_hours >= sum {c in CARS} (rate[c]) * Make[c];
