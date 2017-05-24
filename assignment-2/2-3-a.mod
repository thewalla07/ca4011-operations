set SUPP; # set of suppliers
param cane {SUPP} >= 0, <= 100; # percentage of cane
param corn {SUPP} >= 0, <= 100; # percentage of corn
param beet {SUPP} >= 0, <= 100; # percentage of beet
param cost {SUPP} >= 0; # cost per ton of sugar mix
param cane_target >= 0; # target volume of cane sugar
param corn_target >= 0; # target volume of corn sugar
param beet_target >= 0; # target volume of beet sugar

var Buy {s in SUPP} >= 0; # amount to buy from each supplier
var Cost {s in SUPP} == cost[s] * Buy[s];
# objective function
minimize Total_Cost: sum {s in SUPP} cost[s] * Buy[s];
# constraints
subject to CaneTarget:
    cane_target == sum {s in SUPP} (cane[s]/100) * Buy[s];
subject to CornTarget:
    corn_target == sum {s in SUPP} (corn[s]/100) * Buy[s];
subject to BeetTarget:
    beet_target == sum {s in SUPP} (beet[s]/100) * Buy[s];
