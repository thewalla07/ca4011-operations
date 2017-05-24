set CARS; # car types
param rate {CARS} > 0; # cars per hour
param profit {CARS}; # profit per car
param commit {CARS} >= 0; # lower limit on car demands
param fuel {CARS} >= 0; # fuel efficiency of car
param max_hours > 0; # max factory hours
param min_effic > 0; # min ave fuel efficiency

var Make {c in CARS} >= commit[c]; # cars produced
var Eff == (sum {c in CARS} Make[c] * fuel[c])/(sum {cx in CARS} Make[cx]);
# objective function
maximize Total_Profit: sum {c in CARS} profit[c] * Make[c];

# constraints
subject to Time:
    max_hours >= sum {c in CARS} (rate[c]) * Make[c];

subject to Effic:
    min_effic * sum {cx in CARS} Make[cx] <= sum {c in CARS} (fuel[c]) * Make[c];
