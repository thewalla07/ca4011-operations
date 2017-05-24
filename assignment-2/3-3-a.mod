set ORIG; # origins
set DEST; # destinations
param supply {ORIG} >= 0; # amounts available at origins
param demand {DEST} >= 0; # amounts required at destinations
param supply_pct >= 0, <= 100;
param demand_pct >= 0, <= 100;
check: sum {i in ORIG} supply[i] == sum {j in DEST} demand[j];
param cost {ORIG,DEST} >= 0; # shipment costs per unit
var Trans {ORIG,DEST} >= 0; # units to be shipped
minimize Total_Cost:
sum {i in ORIG, j in DEST} cost[i,j] * Trans[i,j];

subject to Supply {i in ORIG}:
sum {j in DEST} Trans[i,j] == supply[i];
subject to Demand {j in DEST}:
sum {i in ORIG} Trans[i,j] == demand[j];

subject to SupplyPct {i in ORIG, j in DEST}:
Trans[i,j] <= (supply_pct/100) * supply[i];
subject to DemandPct {j in DEST, i in ORIG}:
Trans[i,j] <= (demand_pct/100) * demand[j];
