set ORIG; # origins
set DEST; # destinations
param supply {ORIG} >= 0; # amounts available at origins
param demand {DEST} >= 0; # amounts required at destinations
check: sum {i in ORIG} supply[i] >= sum {j in DEST} demand[j];
param time {ORIG,DEST} >= 0; # time to produce a part at a machine
param cost {ORIG,DEST} >= 0; # shipment costs per unit
var Trans {ORIG,DEST} >= 0; # units to be shipped
var Cost {ORIG} >= 0; # time used on each machine
minimize Total_Time:
sum {i in ORIG, j in DEST} time[i,j] * Trans[i,j];
subject to Pricing {i in ORIG}:
sum {j in DEST} cost[i,j] * Trans[i,j] == Cost[i];
subject to Supply {i in ORIG}:
sum {j in DEST} Trans[i,j] * time[i,j] <= supply[i];
subject to Demand {j in DEST}:
sum {i in ORIG} Trans[i,j] >= demand[j];
