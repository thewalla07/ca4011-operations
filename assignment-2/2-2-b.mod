set EXER; # set of exercises
param cals {EXER} > 0; # calories burnt per hour exercise
param tol {EXER} > 0; # tolerance for each exercise
param min_time {EXER} >= 0; # minimum time for each exercise
param cal_target > 0; # target calories to burn

var Work {e in EXER} <= tol[e], >= min_time[e]; # perform work up to tolerance of each exercise
var Burnt {e in EXER} == cals[e] * Work[e];
# objective function
minimize Total_Time: sum {e in EXER} Work[e];

# constraints
subject to CalorieTarget:
    cal_target <= sum {e in EXER} cals[e] * Work[e];

subject to Variety:
    4 >= Work['walking'] + Work['jogging'] + Work['machine']
