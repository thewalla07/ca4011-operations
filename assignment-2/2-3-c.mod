set SUPP; # set of suppliers
param cane {SUPP} >= 0, <= 100; # percentage of cane
param corn {SUPP} >= 0, <= 100; # percentage of corn
param beet {SUPP} >= 0, <= 100; # percentage of beet
param cost {SUPP} >= 0; # cost per ton of sugar mix

param low_per >= 0;
param high_per <= 100;
param max_per == 100;

var Buy {s in SUPP} >= 0; # amount to buy from each supplier
var Cost {s in SUPP} == cost[s] * Buy[s];

var Cane_per == sum {s in SUPP} (cane[s]/100) * Buy[s];
var Corn_per == sum {s in SUPP} (corn[s]/100) * Buy[s];
var Beet_per == sum {s in SUPP} (beet[s]/100) * Buy[s];

# objective function
minimize Total_Cost: sum {s in SUPP} cost[s] * Buy[s];
# constraints
subject to MaxVolume:
    1 == sum {s in SUPP} Buy[s];
subject to CorrectPercentages:
    1 == Cane_per + Corn_per + Beet_per;
subject to CaneRange:
    low_per/100 <= Cane_per <= high_per/100;
subject to CornRange:
    low_per/100 <= Corn_per <= high_per/100;
subject to BeetRange:
    low_per/100 <= Beet_per <= high_per/100;
