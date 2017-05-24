set EXER; # set of exercises
param cals {EXER} > 0; # calories burnt per hour exercise
param tol {EXER} > 0; # tolerance for each exercise
param cal_target > 0; # target calories to burn

var Work {e in EXER} <= tol[e], >= 0; # perform work up to tolerance of each exercise
var Burnt {e in EXER} == cals[e] * Work[e];
# objective function
minimize Total_Time: sum {e in EXER} Work[e];

# constraints
subject to CalorieTarget:
    cal_target <= sum {e in EXER} cals[e] * Work[e];
